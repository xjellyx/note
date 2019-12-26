package game

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User
type User struct {
	gorm.Model
	Uid         string `json:"uid" gorm:"primary_key;column=uid;type:varchar(36)"`
	Nickname    string `gorm:"column:nickname;" json:"nickname" form:"nickname"`
	HeadIcon    string `gorm:"column:head_icon" json:"head_icon" form:"head_icon"`
	AllHeadIcon string `gorm:"column:all_head_icon;type:varchar(255)" json:"all_head_icon" form:"all_head_icon"`
	// Imheadurl              string    `gorm:"column:imheadurl" json:"imheadurl" form:"imheadurl"`
	Email                  string    `gorm:"column:email;unique_index" json:"email" form:"email"`
	Level                  int64     `gorm:"column:level" json:"level" form:"level"`
	CurTitle               int64     `gorm:"column:cur_title" json:"cur_title" form:"cur_title"`             // 当前称号
	Exp                    float64   `gorm:"column:exp" json:"exp" form:"exp"`                               // 经验值
	MoneyConsume           int64     `gorm:"column:money_consume" json:"money_consume" form:"money_consume"` // 玩家累计消费金额
	Chips                  int64     `gorm:"column:chips;type:bigint(20)" json:"chips" form:"chips"`
	LoseChips              int64     `gorm:"column:lose_chips;type:bigint(20)" json:"lose_chips" form:"lose_chips"`
	FreeChipCount          int64     `gorm:"column:free_chip_count;type:int(11)" json:"free_chip_count" form:"free_chip_count"`
	FreeChipStartTime      int64     `gorm:"column:free_chip_start_time;type:bigint(20)" json:"free_chip_start_time" form:"free_chip_start_time"` // 在线领奖开始时间
	FreeChipLeftTime       int64     `gorm:"column:free_chip_left_time;type:int(11)" json:"free_chip_left_time" form:"free_chip_left_time"`
	VipRewardTime          time.Time `gorm:"column:vip_reward_time" json:"vip_reward_time" form:"vip_reward_time"`
	SavingPotChips         int64     `gorm:"column:saving_pot_chips" json:"saving_pot_chips" form:"saving_pot_chips"`
	SavingPotLevel         int64     `gorm:"column:saving_pot_level" json:"saving_pot_level" form:"saving_pot_level"`
	SavingPotHalfPush      int64     `gorm:"column:saving_pot_half_push" json:"saving_pot_half_push" form:"saving_pot_half_push"`
	SavingPotFullPush      int64     `gorm:"column:saving_pot_full_push" json:"saving_pot_full_push" form:"saving_pot_full_push"`
	PayFlag                int64     `gorm:"column:pay_flag" json:"pay_flag" form:"pay_flag"`
	FirstPayFlag           int64     `gorm:"column:first_pay_flag" json:"first_pay_flag" form:"first_pay_flag"`
	BigPayFlag             int64     `gorm:"column:big_pay_flag" json:"big_pay_flag" form:"big_pay_flag"`
	SevenDayFlag           int64     `gorm:"column:seven_day_flag" json:"seven_day_flag" form:"seven_day_flag"`
	PushFlag               int64     `gorm:"column:push_flag" json:"push_flag" form:"push_flag"`
	Openid                 string    `gorm:"column:openid" json:"openid" form:"openid"`
	Token                  string    `gorm:"column:token" json:"token" form:"token"`
	Devicetoken            string    `gorm:"column:devicetoken" json:"devicetoken" form:"devicetoken"`
	Passporttype           string    `gorm:"column:passporttype" json:"passporttype" form:"passporttype"`
	Regip                  string    `gorm:"column:regip" json:"regip" form:"regip"`
	Regtime                time.Time `gorm:"column:regtime" json:"regtime" form:"regtime"`
	Regdevice              int64     `gorm:"column:regdevice" json:"regdevice" form:"regdevice"`
	Regversion             string    `gorm:"column:regversion" json:"regversion" form:"regversion"`
	Banker                 int64     `gorm:"column:banker" json:"banker" form:"banker"`
	Versionid              string    `gorm:"column:versionid" json:"versionid" form:"versionid"`
	Forbiddentime          int64     `gorm:"column:forbiddentime" json:"forbiddentime" form:"forbiddentime"`
	Forbiddendesc          string    `gorm:"column:forbiddendesc" json:"forbiddendesc" form:"forbiddendesc"`
	Ai                     int64     `gorm:"column:ai" json:"ai" form:"ai"`
	Disable                int64     `gorm:"column:disable" json:"disable" form:"disable"`
	Lotto                  float64   `gorm:"column:lotto" json:"lotto" form:"lotto"`
	Vipchipflag            int64     `gorm:"column:vipchipflag" json:"vipchipflag" form:"vipchipflag"`
	Diamond                int64     `gorm:"column:diamond" json:"diamond" form:"diamond"`
	Changenamenum          int64     `gorm:"column:changenamenum" json:"changenamenum" form:"changenamenum"`
	Sign                   string    `gorm:"column:sign" json:"sign" form:"sign"`
	Reliefchipcount        int64     `gorm:"column:reliefchipcount" json:"reliefchipcount" form:"reliefchipcount"`
	Reliefchipstarttime    int64     `gorm:"column:reliefchipstarttime" json:"reliefchipstarttime" form:"reliefchipstarttime"`
	Bankchips              int64     `gorm:"column:bankchips" json:"bankchips" form:"bankchips"`
	Bankpassword           string    `gorm:"column:bankpassword" json:"bankpassword" form:"bankpassword"`
	Online                 int64     `gorm:"column:online" json:"online" form:"online"`
	Pointcontrol           int64     `gorm:"column:pointcontrol" json:"pointcontrol" form:"pointcontrol"`
	Activationcodelock     int64     `gorm:"column:activationcodelock" json:"activationcodelock" form:"activationcodelock"`
	Activationcodelocktime int64     `gorm:"column:activationcodelocktime" json:"activationcodelocktime" form:"activationcodelocktime"`
	LoginIp                string    `gorm:"column:login_ip" json:"login_ip" form:"login_ip"`
	Phone                  string    `gorm:"column:phone;unique_index;size(11)" json:"phone" form:"phone"`
	Bankcard               string    `gorm:"column:bankcard" json:"bankcard" form:"bankcard"`
	AlipayAccount          string    `gorm:"column:alipay_account" json:"alipay_account" form:"alipay_account"`
	RealName               string    `gorm:"column:real_name;unique_index" json:"real_name" form:"real_name"`
}
