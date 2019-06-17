package main

import (
	"fmt"
	"git.yichui.net/tudy/go-rest/contrib/story"
	"log"
)

var (
	TokenNowAdmin     string
	ConstEndpoints    = []string{"https://api.yichui.net/api/three/graphql"}
	testAdminName     = "jane"
	testAdminPassword = "111111"
)

func main() {
	var (
		limit int
		skip  int
	)
	limit = 50
	for i := 0; i < 20; i++ {
		arr := getGoodList(limit, skip)
		for _, v := range arr {
			if getGood(v) == 4 {
				continue
			}
			editGqlGood(v)
		}
		skip += limit
	}
}

// 获取商品列表
func getGoodList(limit, skip int) (ret []string) {
	str := fmt.Sprintf(`query mm($t:String){
          me(token:$t){
					goodList(query:{limit:%d,skip:%d,sort:["7295fff194d56240c180606e142d7b220bd8abc8.jpg"],
            keyJson:"{\"icon\":\"7295fff194d56240c180606e142d7b220bd8abc8.jpg\"}"}){
      data{
        gid
      }
    }
          }
        }`, limit, skip)
	token := getTokenAdmin()
	if res, err := getGqlRoot().Request(str, nil, &story.GraphqlReq{
		Server: story.PubPointString("mall-admin"),
		Token:  &token,
	}); err == nil {
		var d []*struct {
			Gid string `json:"gid"`
		}
		_ = res.Unmarshal(&d, "me", "goodList", "data")
		for _, v := range d {
			ret = append(ret, v.Gid)
		}
	} else {
		log.Fatal(err)
	}
	return
}

// 编辑good
func editGqlGood(gid string) {
	str := fmt.Sprintf(`mutation setGoods($t:String){
  meEdit(token:$t){
    setGood(form:{gid:"%s",status:4},content:{}){
      status
    }
  }
}`, gid)
	token := getTokenAdmin()
	if res, err := getGqlRoot().Request(str, nil, &story.GraphqlReq{
		Server: story.PubPointString("mall-admin"),
		Token:  &token,
	}); err == nil {
		log.Printf("商品状态 status %v", res.GetDataString("meEdit", "setGood", "status"))
	} else {
		log.Fatal(err)
		return
	}
}

// 取管理员token
func getTokenAdmin() (ret string) {
	if len(TokenNowAdmin) > 0 {
		return TokenNowAdmin
	}
	if res, err := getGqlRoot().Request(`
mutation la($n: String!, $p: String!) {
  login(form: {adminName: $n, password: $p}) {
    token
  }
}
`, map[string]interface{}{
		"n": testAdminName,
		"p": testAdminPassword,
	}, &story.GraphqlReq{
		Server: story.PubPointString("user-admin"),
	}); err != nil {
		log.Fatal(err)
	} else {
		TokenNowAdmin = res.GetDataString("login", "token")
	}
	return TokenNowAdmin
}
func getGqlRoot() (ret *story.Graphql) {
	return &story.Graphql{
		Endpoints: ConstEndpoints,
	}
}

func getGood(gid string) (ret int) {
	str := fmt.Sprintf(`query mm($t:String){
          me(token:$t){
    good(gid:"%s"){
      status
    }
          }
        }`, gid)
	if res, err := getGqlRoot().Request(
		str, nil, &story.GraphqlReq{
			Server: story.PubPointString("mall-admin"),
			Token:  &TokenNowAdmin,
		}); err == nil {
		ret = res.GetDataInt("me", "good", "status")
	} else {
		log.Fatal(err)
	}
	return
}

