package main

import (
	"fmt"
	_ "github.com/olongfen/note/test/ctrl"
	"os"
	"path"
)


func main()  {
	var(
		dir = "./ddddd/aade/dq/data.log"
		err error
		f *os.File
	)
	dirHead, fl :=path.Split(dir)
	fmt.Println(dirHead, fl)
	if _,err = os.Stat(dirHead);err!=nil && os.IsNotExist(err){
		if err =os.MkdirAll(dirHead,os.ModePerm);err!=nil{
			panic(err)
		}
	}
	if f,err = os.Create(dir);err!=nil{
		panic(err)
	}
	fmt.Println(f.Name())

}
