package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `json:"dictValue" gorm:"default:'';comment('字典键值') ;type:varchar(100);uniqueIndex:dict_value_"`
	if strings.Contains(str, "uniqueIndex:") {
		fmt.Println(strings.Contains(strings.Split(str, "uniqueIndex:")[1], "dict_value_"))
	}
}

func d() (err error) {
	s := func() bool {
		err = fmt.Errorf("%s", "dfadasfasd")
		return false
	}
	s()
	return
}
