package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

var (
	NodeString = `// %s
func(n *Node%s) %s () (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.%s
	}
	return
}
`
	NodeInts = `// %s
func(n *Node%s) %s () (ret int32, err error) {
	if n.Data != nil {
		ret = int32(n.Data.%s)
	}
	return
}
`
	NodeBool = `// %s
func(n *Node%s) %s () (ret bool, err error) {
	if n.Data != nil {
		ret = n.Data.%s
	}
	return
}
`
	NodeTime = `// %s
func (n *Node%s) %s() (ret *string) {
	if n.Data != nil && n.Data.%s.Unix() > 0 {
		_t := rest.PubTimeToStr(n.Data.%s)
		ret = &_t
	}
	return
}`
)

// gql_profile和gql_router的自动生成
func main() {
	PubProfileInit(Business{})
	return
}

func PubProfileInit(typeList ...interface{}) {
	var (
		text       []string
		textRouter []string
	)
	file, _ := os.Create(fmt.Sprintf(`./profile/profile_gql.go`))
	text = append(text, "package ctrl \n")
	textRouter = append(textRouter, fmt.Sprintf("const standSchema=`"))

	for _, ty := range typeList {

		v := reflect.ValueOf(ty)
		t := reflect.TypeOf(ty)
		count := v.NumField()

		// profile
		for i := 0; i < count; i++ {
			f := v.Field(i)
			_type := v.Field(i).Kind()
			_name := t.Field(i).Name
			switch _type {
			case reflect.String:
				con := fmt.Sprintf(`// %s
func(n *Node%s) %s () (ret string, err error) {
	if n.Data != nil {
		ret = n.Data.%s
	}
	return
}
`, _name, t.Name(), _name, _name)
				text = append(text, con)
			case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32:
				con := fmt.Sprintf(`// %s
func(n *Node%s) %s () (ret int32, err error) {
	if n.Data != nil {
		ret = n.Data.%s
	}
	return
}
`, _name, t.Name(), _name, _name)
				text = append(text, con)
			case reflect.Bool:
				con := fmt.Sprintf(`// %s
func(n *Node%s) %s () (ret bool, err error) {
	if n.Data != nil {
		ret = n.Data.%s
	}
	return
}
`, _name, t.Name(), _name, _name)
				text = append(text, con)
			case reflect.Struct:
				switch f.Interface().(type) {
				case time.Time:
					con := fmt.Sprintf(`// %s
func (n *Node%s) %s() (ret *string) {
	if n.Data != nil && n.Data.%s.Unix() > 0 {
		_t := rest.PubTimeToStr(n.Data.%s)
		ret = &_t
	}
	return
}`, _name, t.Name(), _name, _name, _name)
					text = append(text, con)
				default:
					PubProfileInit(f.Interface())
				}

			default:

			}
			fmt.Println(text)
		}

		var ()
		textRouter = append(textRouter, "\n")
		textRouter = append(textRouter, fmt.Sprintf("	type %s {\n", t.Name()))
		// router
		for i := 0; i < count; i++ {
			f := v.Field(i)
			_t := v.Field(i).Kind()
			_name := t.Field(i).Name
			fmt.Println("acccaaa", _t)
			switch _t {
			case reflect.String:
				var con string
				if _name == "UID" {
					con = fmt.Sprintf(`		#
		%s: String!
`, strings.ToLower(_name))
				} else {
					con = fmt.Sprintf(`		#
		%s: String!
`, CapitalLower(_name))
				}

				textRouter = append(textRouter, con)
			case reflect.Int, reflect.Int64, reflect.Int8, reflect.Int16, reflect.Int32:
				con := fmt.Sprintf(`		#
		%s: Int!
`, CapitalLower(_name))
				textRouter = append(textRouter, con)
			case reflect.Bool:
				con := fmt.Sprintf(`		#
		%s: Boolean!
`, CapitalLower(_name))
				textRouter = append(textRouter, con)
			case reflect.Struct:
				switch f.Interface().(type) {
				case time.Time:
					con := fmt.Sprintf(`		# 
		%s: String
`, CapitalLower(_name))
					textRouter = append(textRouter, con)
				}

			}
		}
		textRouter = append(textRouter, "	}")
	}
	textRouter = append(textRouter, fmt.Sprintf("\n`"))

	ret := strings.Join(text, "\n")
	ret2 := strings.Join(textRouter, "")
	io.Copy(file, bytes.NewBuffer([]byte(ret+"\n"+ret2)))

	return
}

// CapitalLower 字符首字母小写
func CapitalLower(str string) string {
	var upperStr string

	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 90 { // 后文有介绍
				vv[i] += 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// 商户表
type Business struct {
	UID       string  `sorm:"size(36);primary"json:"uid"`
	Category  string  `sorm:"size(16);index" json:"category"`    // 商家分类
	Name      string  `sorm:"size(64);index;unique" json:"name"` // 商家名称
	Icon      string  `sorm:"index" json:"icon"`                 // 商家头像
	Address   string  `sorm:"size(256);index" json:"address"`
	Contact   string  `sorm:"size(16);index" json:"contact"`
	Phone     string  `sorm:"size(36),index" json:"phone"`
	Longitude float64 `sorm:"index" json:"longitude"` // 商家经度
	Latitude  float64 `sorm:"index" json:"latitude"`  // 商家纬度
	// TODO 商家类型待处理
	// TypeName  string    `sorm:"index" json:"typeName"`           // 商家类型
	CreateTime time.Time    `sorm:"index" json:"createTime,omitempty"` // 创建时间
	UpdateTime time.Time    `sorm:"index" json:"updateTime,omitempty"` // 更新时间
	Meta       BusinessMeta `sorm:"json" json:"meta,omitempty"`
}

// meta
type BusinessMeta struct {
	Introduce string   ` json:"Introduce"` // 商家介绍
	ImageShow []string ` json:"imageShow"` // 图片展示
}
