package main

import (
	"database/sql"
	"errors"
	"fmt"
	lcasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/olongfen/note/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

func main() {

	a ,_:= NewAdapter(postgres.Open("dbname=business user=business password=business host=127.0.0.1 port=5432 sslmode=disable"),nil)

	e,_err:=NewCasbinMiddleware("./model.conf",a, func(c *gin.Context) string {
		return "user"
	})
	if _err!=nil{
		panic(_err)
	}
	e.LoadPolicy()
	e.AddPolicy("admin","/v1/admin/userList","*")
	// e.AddRoleForUser("user","admin")
	// e.AddRoleForUser("user1","admin")
	g:=e.RequiresPermissions([]string{"/alice_data/resource1:admin"})
	r:=gin.Default()
	r.GET("/alice_data/resource1",g, func(c *gin.Context) {
		c.JSON(200,gin.H{"sdasd":2341231})
	})
	fmt.Println(e.Enforce("4543534534","/v1/admin/userList","admin"))
	r.Run(":3344")
}


type CasbinMiddleware struct {
	 *lcasbin.Enforcer
	subFn    SubjectFn
}

// SubjectFn is used to look up current subject in runtime.
// If it can not find anything, just return an empty string.
type SubjectFn func(c *gin.Context) string

// Logic is the logical operation (AND/OR) used in permission checks
// in case multiple permissions or roles are specified.
type Logic int

const (
	AND Logic = iota
	OR
)

var (
	SubFnNilErr = errors.New("subFn is nil")
)

// NewCasbinMiddleware returns a new CasbinMiddleware using Casbin's Enforcer internally.
// modelFile is the file path to Casbin model file e.g. path/to/rbac_model.conf.
// policyAdapter can be a file or a DB adapter.
// File: path/to/basic_policy.csv
// MySQL DB: mysqladapter.NewDBAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/")
// subFn is a function that looks up the current subject in runtime and returns an empty string if nothing found.
func NewCasbinMiddleware(modelFile string, policyAdapter interface{}, subFn SubjectFn) (*CasbinMiddleware, error) {
	if subFn == nil {
		return nil, SubFnNilErr
	}
	en,err:=lcasbin.NewEnforcer(modelFile, policyAdapter)
	if err!=nil{
		return nil, err
	}
	return &CasbinMiddleware{
		 en,
		   subFn,
	}, nil
}

// Option is used to change some default behaviors.
type Option interface {
	apply(*options)
}

type options struct {
	logic Logic
}

type logicOption Logic

func (lo logicOption) apply(opts *options) {
	opts.logic = Logic(lo)
}

// WithLogic sets the logical operator used in permission or role checks.
func WithLogic(logic Logic) Option {
	return logicOption(logic)
}

// RequiresPermissions tries to find the current subject by calling SubjectFn
// and determine if the subject has the required permissions according to predefined Casbin policies.
// permissions are formatted strings. For example, "file:read" represents the permission to read a file.
// opts is some optional configurations such as the logical operator (default is AND) in case multiple permissions are specified.
func (am *CasbinMiddleware) RequiresPermissions(permissions []string, opts ...Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(permissions) == 0 {
			c.Next()
			return
		}

		// Here we provide default options.
		actualOptions := options{
			logic: AND,
		}
		// Apply actual options.
		for _, opt := range opts {
			opt.apply(&actualOptions)
		}

		// Look up current subject.
		sub := am.subFn(c)
		if sub == "" {
			c.AbortWithStatus(401)
			return
		}

		// Enforce Casbin policies.
		if actualOptions.logic == AND {
			// Must pass all tests.
			for _, permission := range permissions {
				obj, act := parsePermissionStrings(permission)
				if obj == "" || act == "" {
					// Can not handle any illegal permission strings.
					log.Println("illegal permission string: ", permission)
					c.AbortWithStatus(500)
					return
				}
				fmt.Println(sub,obj,act)
				if ok,err := am.Enforce(sub, obj, act); !ok || err!=nil {
					fmt.Println(ok,"aaa",err)
					c.AbortWithStatus(401)
					return
				}
			}
			c.Next()
		} else {
			// Need to pass at least one test.
			for _, permission := range permissions {
				obj, act := parsePermissionStrings(permission)
				if obj == "" || act == "" {
					log.Println("illegal permission string: ", permission)
					c.AbortWithStatus(500)
					continue
				}

				if ok,err := am.Enforce(sub, obj, act); ok && err==nil{
					c.Next()
					return
				}
			}
			c.AbortWithStatus(401)
		}
	}
}

func parsePermissionStrings(str string) (string, string) {
	if !strings.Contains(str, ":") {
		return "", ""
	}
	vals := strings.Split(str, ":")
	return vals[0], vals[1]
}

// RequiresPermissions tries to find the current subject by calling SubjectFn
// and determine if the subject has the required roles according to predefined Casbin policies.
// opts is some optional configurations such as the logical operator (default is AND) in case multiple roles are specified.
func (am *CasbinMiddleware) RequiresRoles(requiredRoles []string, opts ...Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(requiredRoles) == 0 {
			c.Next()
			return
		}

		// Look up current subject.
		sub := am.subFn(c)
		if sub == "" {
			c.AbortWithStatus(401)
			return
		}

		// Here we provide default options.
		actualOptions := options{
			logic: AND,
		}
		// Apply actual options.
		for _, opt := range opts {
			opt.apply(&actualOptions)
		}

		actualRoles, err := am.GetRolesForUser(sub)
		if err != nil {
			log.Println("couldn't get roles of subject: ", err)
			c.AbortWithStatus(500)
			return
		}

		// Enforce Casbin policies.
		sort.Strings(requiredRoles)
		sort.Strings(actualRoles)
		if actualOptions.logic == AND {
			// Must have all required roles.
			if !reflect.DeepEqual(requiredRoles, actualRoles) {
				c.AbortWithStatus(401)
			} else {
				c.Next()
			}
		} else {
			// Need to have at least one of required roles.
			for _, requiredRole := range requiredRoles {
				if i := sort.SearchStrings(actualRoles, requiredRole); i >= 0 &&
					i < len(actualRoles) &&
					actualRoles[i] == requiredRole {
					c.Next()
					return
				}
			}
			c.AbortWithStatus(401)
		}
	}
}



type CasbinRule struct {
	TablePrefix string `gorm:"-"`
	ID        uint `gorm:"primarykey"`
	PType       string `gorm:"size:100"`
	V0          string `gorm:"size:100"`
	V1          string `gorm:"size:100"`
	V2          string `gorm:"size:100"`
	V3          string `gorm:"size:100"`
	V4          string `gorm:"size:100"`
	V5          string `gorm:"size:100"`
}

type Filter struct {
	PType []string
	V0    []string
	V1    []string
	V2    []string
	V3    []string
	V4    []string
	V5    []string
}

func (c *CasbinRule) TableName() string {
	return c.TablePrefix + "casbin_rule" //as Gorm keeps table names are plural, and we love consistency
}

// Adapter represents the Gorm adapter for policy storage.
type Adapter struct {
	tablePrefix    string
	// dbSpecified    bool
	db             *gorm.DB
	isFiltered     bool
}

// finalizer is the destructor for Adapter.
func finalizer(a *Adapter) {
	err := a.db.ConnPool.(*sql.DB).Close()
	if err != nil {
		panic(err)
	}
}

// NewAdapter is the constructor for Adapter.
// dbSpecified is an optional bool parameter. The default value is false.
// It's up to whether you have specified an existing DB in dataSourceName.
// If dbSpecified == true, you need to make sure the DB in dataSourceName exists.
// If dbSpecified == false, the adapter will automatically create a DB named "casbin".
func NewAdapter(migrator gorm.Dialector,c *gorm.Config) (*Adapter, error) {
	a := &Adapter{}


	// Open the DB, create it if not existed.
	err := a.open(migrator,c)
	if err != nil {
		return nil, err
	}

	// Call the destructor when the object is released.
	runtime.SetFinalizer(a, finalizer)

	return a, nil
}

// NewAdapterByDB obtained through an existing Gorm instance get  a adapter, specify the table prefix
// Example: gormadapter.NewAdapterByDBUsePrefix(&db, "cms_") Automatically generate table name like this "cms_casbin_rule"
func NewAdapterByDBUsePrefix(db *gorm.DB, prefix string) (*Adapter, error) {
	a := &Adapter{
		tablePrefix: prefix,
		db:          db,
	}

	err := a.createTable()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func NewAdapterByDB(db *gorm.DB) (*Adapter, error) {
	a := &Adapter{
		db: db,
	}

	err := a.createTable()
	if err != nil {
		return nil, err
	}

	return a, nil
}


func (a *Adapter) open(migrator gorm.Dialector,c *gorm.Config) error {
	var err error
	var db *gorm.DB
	db, err = gorm.Open(migrator,c)
	if err != nil {
		return err
	}
	a.db = db

	return a.createTable()
}

func (a *Adapter) close() error {
	err := a.db.ConnPool.(*sql.DB).Close()
	if err != nil {
		return err
	}

	a.db = nil
	return nil
}

// getTableInstance return the dynamic table name
func (a *Adapter) getTableInstance() *CasbinRule {
	return &CasbinRule{TablePrefix: a.tablePrefix}
}

func (a *Adapter) createTable() error {
	return a.db.AutoMigrate(a.getTableInstance())
}


func loadPolicyLine(line CasbinRule, model model.Model) {
	var p = []string{line.PType,
		line.V0, line.V1, line.V2, line.V3, line.V4, line.V5}

	var lineText string
	if line.V5 != "" {
		lineText = strings.Join(p, ", ")
	} else if line.V4 != "" {
		lineText = strings.Join(p[:6], ", ")
	} else if line.V3 != "" {
		lineText = strings.Join(p[:5], ", ")
	} else if line.V2 != "" {
		lineText = strings.Join(p[:4], ", ")
	} else if line.V1 != "" {
		lineText = strings.Join(p[:3], ", ")
	} else if line.V0 != "" {
		lineText = strings.Join(p[:2], ", ")
	}

	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	var lines []CasbinRule
	if err := a.db.Table(a.tablePrefix + "casbin_rule").Find(&lines).Error; err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}

	return nil
}

// LoadFilteredPolicy loads only policy rules that match the filter.
func (a *Adapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	var lines []CasbinRule

	filterValue, ok := filter.(Filter)
	if !ok {
		return errors.New("invalid filter type")
	}

	if err := a.db.Scopes(a.filterQuery(a.db, filterValue)).Find(&lines).Error; err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	a.isFiltered = true

	return nil
}

// IsFiltered returns true if the loaded policy has been filtered.
func (a *Adapter) IsFiltered() bool {
	return a.isFiltered
}

// filterQuery builds the gorm query to match the rule filter to use within a scope.
func (a *Adapter) filterQuery(db *gorm.DB, filter Filter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(filter.PType) > 0 {
			db = db.Where("p_type in (?)", filter.PType)
		}
		if len(filter.V0) > 0 {
			db = db.Where("v0 in (?)", filter.V0)
		}
		if len(filter.V1) > 0 {
			db = db.Where("v1 in (?)", filter.V1)
		}
		if len(filter.V2) > 0 {
			db = db.Where("v2 in (?)", filter.V2)
		}
		if len(filter.V3) > 0 {
			db = db.Where("v3 in (?)", filter.V3)
		}
		if len(filter.V4) > 0 {
			db = db.Where("v4 in (?)", filter.V4)
		}
		if len(filter.V5) > 0 {
			db = db.Where("v5 in (?)", filter.V5)
		}
		return db
	}
}

func (a *Adapter) savePolicyLine(ptype string, rule []string) CasbinRule {
	line := a.getTableInstance()

	line.PType = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}

	return *line
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	err := a.createTable()
	if err != nil {
		return err
	}

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := a.savePolicyLine(ptype, rule)
			err := a.db.Create(&line).Error
			if err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := a.savePolicyLine(ptype, rule)
			err := a.db.Create(&line).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := a.savePolicyLine(ptype, rule)
	err := a.db.Create(&line).Error
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := a.savePolicyLine(ptype, rule)
	err := a.rawDelete(a.db, line) //can't use db.Delete as we're not using primary key http://jinzhu.me/gorm/crud.html#delete
	return err
}

// AddPolicies adds multiple policy rules to the storage.
func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	return a.db.Transaction(func(tx *gorm.DB) error {
		for _, rule := range rules {
			line := a.savePolicyLine(ptype, rule)
			if err := tx.Create(&line).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// RemovePolicy removes multiple policy rules from the storage.
func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return a.db.Transaction(func(tx *gorm.DB) error {
		for _, rule := range rules {
			line := a.savePolicyLine(ptype, rule)
			if err := a.rawDelete(tx, line); err != nil { //can't use db.Delete as we're not using primary key http://jinzhu.me/gorm/crud.html#delete
				return err
			}
		}
		return nil
	})
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := a.getTableInstance()

	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err := a.rawDelete(a.db, *line)
	return err
}

func (a *Adapter) rawDelete(db *gorm.DB, line CasbinRule) error {
	queryArgs := []interface{}{line.PType}

	queryStr := "p_type = ?"
	if line.V0 != "" {
		queryStr += " and v0 = ?"
		queryArgs = append(queryArgs, line.V0)
	}
	if line.V1 != "" {
		queryStr += " and v1 = ?"
		queryArgs = append(queryArgs, line.V1)
	}
	if line.V2 != "" {
		queryStr += " and v2 = ?"
		queryArgs = append(queryArgs, line.V2)
	}
	if line.V3 != "" {
		queryStr += " and v3 = ?"
		queryArgs = append(queryArgs, line.V3)
	}
	if line.V4 != "" {
		queryStr += " and v4 = ?"
		queryArgs = append(queryArgs, line.V4)
	}
	if line.V5 != "" {
		queryStr += " and v5 = ?"
		queryArgs = append(queryArgs, line.V5)
	}
	args := append([]interface{}{queryStr}, queryArgs...)
	err := db.Delete(a.getTableInstance(), args...).Error
	return err
}

