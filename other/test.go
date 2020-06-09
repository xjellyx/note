package main

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"io/ioutil"
	"strings"
)

type Info struct {
	Date   string `csv:"Date"`
	Tank string `csv:"Tank#"`
	Capacity string `csv:"Capacity"`
	VolumeStored string `csv:"Volume Stored"`
	WeeklyAverageVolumeStored  float64 `csv:"Weekly Average Volume Stored"`
	VolumeSummary  float64 `csv:"Volume Summary"`
	Company string `csv:"Company"`
}
func main()  {
	dirRoot :="/home/olongfen/Documents/resource_root"
	d,_:=ioutil.ReadDir(dirRoot)
	datas := []*Info{}
	for _,v:=range d{
		if v.Name() == "2018"{
			dd,_:=ioutil.ReadDir(dirRoot+"/"+v.Name())
			arr1 := strings.Split(dd[0].Name(), ".")
			arr2 := strings.Split(arr1[0], "_")
			arr3 := strings.Split(arr2[1], "-")
			fmt.Println(arr1)
			fmt.Println(arr2)
			fmt.Println(arr3[0])
			_f,_:=ioutil.ReadFile(dirRoot+"/"+v.Name()+"/"+dd[0].Name())
			if err:=csvutil.Unmarshal(_f,&datas);err!=nil{
				fmt.Println(err)
			}
		}
	}

	fmt.Println(datas[366])
}