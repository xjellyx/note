package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//f, err := os.Open("/data/allen/csv/1995/广西.csv")
	//if err != nil {
	//	panic(err)
	//}
	//
	//c := csv.NewReader(f)
	//s, _ := c.ReadAll()
	//fmt.Println(s[1])
	patchEnterprise()

}

func isExisItem(val interface{}, arr interface{}) bool {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

// patchEnterprise 广西企业信息录入补丁
func patchEnterprise() (err error) {
	var (
		dirPath  string
		dirList  []string
		name     = "广西.csv"
		filePath []string
		total, n int32
	)
	dirPath = "/data/allen/csv/"
	if err = filepath.Walk(dirPath,
		func(p string, f os.FileInfo, _err error) error {
			if f == nil {
				return _err
			}

			// 获取dir
			if f.IsDir() {
				dirList = append(dirList, p)
				return nil

			}
			return nil
		}); err != nil {
		return
	}

	for _, v := range dirList {
		// 不要根目录
		if v == dirPath {
			continue
		}
		if err = filepath.Walk(v,
			func(p string, f os.FileInfo, _err error) error {
				if f == nil {
					return _err
				}

				// 获取dir
				if !f.IsDir() {
					filePath = append(filePath, p)
					return nil

				}
				return nil
			}); err != nil {
			return
		}

		//for _, _ := range filePath {
		//
		//}
		if isExisItem(v+"/"+name, filePath) {
			if n, err = priParseCsvAndSave(v + "/" + name); err != nil {
				continue
			} else {
				total += n
			}
		}

	}
	println(total)
	return
}

func priParseCsvAndSave(feliPath string) (ret int32, err error) {
	var (
		file    *os.File
		dataArr [][]string
		data    *EnterpriseGX
		n       int32
	)
	if file, err = os.Open(feliPath); err != nil {
		println("[patch-enterprise] err: %v", err)
		err = nil
		return
	}
	c := csv.NewReader(file)
	if dataArr, err = c.ReadAll(); err != nil {
		println("[patch-enterprise] err: %v", err)
		err = nil
		return
	}
	for i := 1; i < len(dataArr); i++ {
		data = NewEnterpriseGX(nil)
		arr := dataArr[i]
		if len(arr) == 10 {
			data.Name = arr[0]
			data.CreditCodeID = arr[1]
			data.RegisterTime = arr[2]
			fmt.Println(time.Parse("2006-01-02", arr[2]))
			data.EnterpriseType = arr[3]
			data.LegalRepresentative = arr[4]
			if arr[5] != "N/A" {
				reg := regexp.MustCompile(`[\d\.]`)
				d := reg.FindAllString(arr[5], -1)
				str := strings.Join(d, "")
				da, _ := strconv.ParseFloat(str, 64)
				println(da)
				data.RegisterCapital = da
			}
			data.BusinessScope = arr[6]
			data.Province = arr[7]
			data.Area = arr[8]
			data.RegisterAddr = arr[9]
			n++
		}
	}

	ret = n
	return
}

//func getFilelist(path string) (ret []string, err error) {
//	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
//		if f == nil {
//			return err
//		}
//		if f.IsDir() {
//			return nil
//		}
//		ret = append(ret, path)
//		return nil
//	})
//	if err != nil {
//		return
//	}
//	return
//}
//
//func getDirList(dirpath string) ([]string, error) {
//	var dir_list []string
//	dir_err := filepath.Walk(dirpath,
//		func(path string, f os.FileInfo, err error) error {
//			if f == nil {
//				return err
//			}
//			if f.IsDir() {
//				dir_list = append(dir_list, path)
//				return nil
//			}
//
//			return nil
//		})
//	return dir_list, dir_err
//}

// EnterpriseGX p2.广西企业
type EnterpriseGX struct {
	// 基础逻辑
	CreditCodeID        string    `sorm:"primary;size(18)" json:"creditCodeID"` // 信用代码，唯一
	CreateTime          time.Time `sorm:"index" json:"createTime"`              // 创建时间
	UpdateTime          time.Time `sorm:"index" json:"updateTime"`              // 更换新时间
	Name                string    `sorm:"" json:"name"`                         // 企业名称
	RegisterTime        string    `sorm:"index" json:"registerTime"`            // 企业注册时间
	EnterpriseType      string    `sorm:"" json:"enterpriseType"`               // 企业类型
	LegalRepresentative string    `sorm:"" json:"legalRepresentative"`          // 法人代表
	RegisterCapital     float64   `sorm:"" json:"registerCapital"`              // 注册资金
	BusinessScope       string    `sorm:"" json:"businessScope"`                // 经营范围
	Province            string    `sorm:"" json:"province"`                     // 省份
	Area                string    `sorm:"" json:"area"`                         // 地区
	RegisterAddr        string    `sorm:"" json:"registerAddr"`                 // 注册地址
}

// NewUser 新建User并初始化
func NewEnterpriseGX(s *EnterpriseGX) (d *EnterpriseGX) {
	if s != nil {
		d = s
	} else {
		d = new(EnterpriseGX)
	}
	if d.CreateTime.Unix() < 0 {
		d.CreateTime = time.Now()
	}
	return
}
