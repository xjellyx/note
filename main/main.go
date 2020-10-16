package main

import (
	"gorm.io/gorm"
	"text/template"
)

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"unique_index"`
	Age  int
}

func main() {
	//g := gormgen.NewGenerator("s.go", "demo")
	//err := g.ParserStruct([]interface{}{&User{}}).Generate().Format().Flush()
	//fmt.Println(err)
}

func parseTemplateOrPanic(t string) *template.Template {
	tpl, err := template.New("output_template").Parse(t)
	if err != nil {
		panic(err)
	}
	return tpl
}

var outputTemplate = parseTemplateOrPanic(`
package {{.PkgName}}
import "gorm.io/gorm"

{{range .Structs}}

	func (t *{{.StructName}}) Add(db *gorm.DB)(err error) {
		return db.Create(t).Error
	}

	func (t *{{.StructName}}) Updates(db *gorm.DB, m map[string]interface{}) error {
		return db.Where("id = ?",t.ID).Updates(m).Error
	}

	func Get{{.StructName}}All(db *gorm.DB)(ret []*{{.StructName}},err error){
		if err = db.Find(&ret).Error;err!=nil{
			return
		}
		return
	}
	
	func {{.StructName}}Count(db *gorm.DB)(ret int64,err error){
		db.Model(&{{.StructName}}{}).Count(&ret)
		return
	}

	{{$StructName := .StructName}}
	{{range .Fields}}
		func (t *{{$StructName}})GetBy{{.FieldName}}(db *gorm.DB)(err error){
			if err = db.First(t,"{{.ColumnName}} = ?",t.{{.FieldName}}).Error;err!=nil{
				return
			}
			return
		}
		
		// DeleteBy{{.FieldName}} delete record by {{.FieldName}}
		func (t *{{.StructName}}) DeleteBy{{.FieldName}}(db *gorm.DB)(err error) {
			return db.Delete(t).Error
		}
	{{end}}

{{end}}
`)
