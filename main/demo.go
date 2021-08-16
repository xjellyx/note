package main

import (
	"crypto/rand"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

var (
	db *gorm.DB
)

var (
	reds = [33]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26",
		"27", "28", "29", "30", "31", "32", "33"}
	blues = [16]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16"}
)

func init() {
	var (
		err error
	)
	if db, err = gorm.Open(postgres.Open(fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "business",
		"business", "127.0.0.1", "5432", "business"))); err != nil {
		logrus.Panic(err)
	}
}

type Reds []string

func (r Reds) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Reds) Less(i, j int) bool {
	iv, _ := strconv.ParseInt(r[i], 10, 64)
	jv, _ := strconv.ParseInt(r[j], 10, 64)
	return iv < jv
}

func (r Reds) Len() int {
	return len(r)
}

func main() {
	db.AutoMigrate(&CaiPiao{})
	//bd, _ := ioutil.ReadFile("./ssq.json")
	//var data CaiPiaos
	//json.Unmarshal(bd, &data)
	//sort.Sort(data)
	//for _, v := range data {
	//	if v.Code != "2021078" {
	//		continue
	//	}
	//	arr := strings.Split(v.Red, ",")
	//	db.Create(v)
	//}
	var (
		count int64
	)
	for i := 0; i < 5; i++ {
		r, b := gen()

		var (
			res = new(CaiPiao)
		)
		for {
			count++
			d, _ := rand.Int(rand.Reader, big.NewInt(1500000000))
			zhongiang, _ := rand.Int(rand.Reader, big.NewInt(1500000000))
			if d.Int64() == zhongiang.Int64() {
				break
			}
		}

		if err := db.Model(&CaiPiao{}).Where("red = ?", r).Find(res).Error; err != nil {
			logrus.Errorln(err)
			i--
			continue
		}
		fmt.Println(r, b)
	}
	// var data CaiPiaos
	fmt.Println(count)

}

func gen() (string, string) {
	var (
		redKey  = make(map[int64]bool)
		redData []string
	)

	for len(redData) < 6 {
		indexRed, _ := rand.Int(rand.Reader, big.NewInt(int64(len(reds))))
		val := reds[indexRed.Int64()]
		if _, ok := redKey[indexRed.Int64()]; !ok {
			redKey[indexRed.Int64()] = true
			redData = append(redData, val)
		} else {
			continue
		}
	}
	sort.Sort(Reds(redData))
	r := strings.Join(redData, ",")
	index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(blues))))
	b := blues[index.Int64()]
	return r, b
}

type CaiPiao struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name" gorm:"type:varchar(36)"`
	Code        string      `json:"code" gorm:"type:varchar(12);uniqueIndex"`
	Date        string      `json:"date" gorm:"type:varchar(36)"`
	Week        string      `json:"week" gorm:"type:varchar(36)"`
	Red         string      `json:"red" gorm:"type:varchar(36)"`
	Blue        string      `json:"blue" gorm:"type:varchar(36)"`
	Blue2       string      `json:"blue2" gorm:"type:varchar(36)"`
	Sales       string      `json:"sales" gorm:"type:varchar(36)"`
	PoolMoney   string      `json:"poolmoney" gorm:"type:varchar(36)"`
	Content     string      `json:"content" `
	AddMoney    string      `json:"addmoney" gorm:"type:varchar(36)"`
	AddMoney2   string      `json:"addmoney2" gorm:"type:varchar(36)"`
	Msg         string      `json:"msg" `
	Z2Add       string      `json:"z2add" gorm:"type:varchar(36)"`
	M2Add       string      `json:"m2add" gorm:"type:varchar(36)"`
	PrizeGrades PrizeGrades `json:"prizegrades" gorm:"type: json"`
	Zj1         string      `json:"zj1,omitempty"`
	Mj1         string      `json:"mj1,omitempty"`
	Zj6         string      `json:"zj6,omitempty"`
	Mj6         string      `json:"mj6,omitempty"`
}

type PrizeGrades []*PrizeGrade

func (p *PrizeGrades) Scan(in interface{}) error {
	return json.Unmarshal(in.([]byte), p)
}

func (p PrizeGrades) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type PrizeGrade struct {
	Type      int    `json:"type"`
	TypeNum   string `json:"type_num"`
	TypeMoney string `json:"type_money"`
}

type CaiPiaos []*CaiPiao

func (r CaiPiaos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r CaiPiaos) Less(i, j int) bool {
	iv, _ := strconv.ParseInt(r[i].Code, 10, 64)
	jv, _ := strconv.ParseInt(r[j].Code, 10, 64)
	return iv < jv
}

type KV struct {
	Key   string
	Value float64
}

type SortKV []KV

func (r SortKV) Len() int {
	return len(r)
}

func (r SortKV) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r SortKV) Less(i, j int) bool {

	return r[i].Value > r[j].Value
}

func (r CaiPiaos) Len() int {
	return len(r)
}

func count() {
	var (
		data  SortKV
		total int64
	)
	db.Model(&CaiPiao{}).Count(&total)
	for _, v := range reds {
		var (
			count int64
		)
		db.Model(&CaiPiao{}).Select("count(*) as count ").Where("red like ?", "%"+v+"%").Count(&count)
		d := KV{Key: v, Value: float64(count) / float64(total)}
		data = append(data, d)
	}
	sort.Sort(data)
	fmt.Println(data)
}
