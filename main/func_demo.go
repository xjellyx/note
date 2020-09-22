package main

import (
	"fmt"
	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	"strings"
)

func main() {
	e, _ := casbin.NewEnforcer("model.conf", "policy.csv")
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	fmt.Println(strings.Split("/v1/admin/addMenu", "?")[0])
	fmt.Println(e.Enforce("user1", strings.Split("/v1/admin/addMenu?a=10", "?")[0], "POST"))
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return ParamsMatch(name1, name2), nil
}
