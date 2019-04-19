package project

import (
	"time"
)

type Admin struct {
	AdminId    int64     `sorm:"size(36); primary" json:"id"` //主键
	Name       string    `sorm:"index; size(16)" json:"name"`
	CreateTime time.Time `sorm:"index" json:"createTime"`
	Status     int64     `sorm:"index" json:"status"`
	Password   string    `sorm:"index; size(16)" json:"Password"` //管理员密码
	NameCity   string    `sorm:"index; size(16)" json:"nameCity"` //管理员所在城市名称
	CityId     int64     `sorm:"index" json:"cityId"`
	// City       *City     `sorm:"- <- ->"` //所对应的城市结构体（基础表结构体）
}
