package main

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"time"
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
	a := ConfigWeapon{
		ID:           0,
		FireInterval: 0,
		Hit:          "",
		BulletRadius: "",
		BulletSpeed:  "",
		BulletCost:   "",
		Rate:         "",
		CatchRadius:  "",
	}
	json.Marshal(a)
}

func Shuffle(vals []int) []int {
	newVals := make([]int, 0, len(vals))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(vals)) {
		newVals = append(newVals, vals[i])
	}
	return newVals
}

func randSeed() (seed int64) {
	var (
		s   = uuid.NewV4().String()
		i   = 0
		n   = ""
		err error
	)
	for len(n) < 18 {
		n += fmt.Sprintf("%d", s[i])
		i += 1
	}
	n = n[0:18]
	if seed, err = strconv.ParseInt(n, 10, 64); err != nil {
		panic(err)
	}
	return
}
