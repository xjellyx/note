package main

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"strconv"
)

type Hw2WeaponData struct {
	ID           string `json:"id"`
	FireInterval string `json:"射击间隔"`
	Hit          string `json:"命中值"`
	// BulletRadius string `json:"子弹碰撞范围"`
	// BulletSpeed string `json:"子弹速度"`
	// BulletCost  string `json:"子弹消耗"`
	// CatchRadius string `json:"子弹爆炸范围"`
}

type ConfigWeapon struct {
	ID           int    `yaml:"id"`
	FireInterval int64  `yaml:"fireInterval"` // 射击间隔
	Hit          string `yaml:"hit"`          // 命中值
	BulletRadius string `yaml:"bulletRadius"` // 子弹碰撞范围
	BulletSpeed  string `yaml:"bulletSpeed"`  // 子弹速度
	BulletCost   string `yaml:"bulletCost"`   // 子弹消耗
	Rate         string `yaml:"rate"`
	CatchRadius  string `yaml:"catchRadius"` // 子弹爆炸范围
}

func main() {
	var (
		data  = []*Hw2WeaponData{}
		datas = []*ConfigWeapon{}
	)
	d, _ := ioutil.ReadFile("/data/fedora-data/gocode/src/jinguoyule/conf/海王2炮塔表.json")
	_ = json.Unmarshal(d, &data)

	for _, v := range data {
		var (
			d = &ConfigWeapon{}
		)
		d.ID, _ = strconv.Atoi(v.ID)
		d.FireInterval, _ = strconv.ParseInt(v.FireInterval, 10, 64)
		datas = append(datas, d)
	}
	_d, _ := json.Marshal(datas)

	dd, _ := yaml.JSONToYAML(_d)
	ioutil.WriteFile("test_hw2.yaml", dd, os.ModePerm)
}
