package main

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"regexp"
)

type exchange struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Result struct {
		From       string  `json:"from"`
		To         string  `json:"to"`
		Fromname   string  `json:"fromname"`
		Toname     string  `json:"toname"`
		Updatetime string  `json:"updatetime"`
		Rate       string  `json:"rate"`
		Camount    float64 `json:"camount"`
	} `json:"result"`
}

func main() {
	r := regexp.MustCompile("([0-9]+.?[0-9]+)")
	fmt.Println(r.FindAllString("8.98亿马币、25亿人民币", -1))
	fmt.Println(138 * 0.180400)
	//params := url.Values{}
	//params.Add("from", "CNY")
	//params.Add("to", "USD")
	//params.Add("amount", "1")
	//params.Add("appkey", "b2beb246120c15ec")
	//fmt.Println(params.Encode())
	//if resp, err := http.Get("https://api.jisuapi.com/exchange/convert?" + params.Encode()); err != nil {
	//	log.Fatal(err)
	//	return
	//} else {
	//	defer resp.Body.Close()
	//	var (
	//		m = new(exchange)
	//	)
	//	ff, _ := ioutil.ReadAll(resp.Body)
	//	json.Unmarshal(ff, &m)
	//	fmt.Println(m)
	//}
	f, err := excelize.OpenFile("/home/olongfen/Documents/ccb.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(parsePort(f))
}
func parsePark(f *excelize.File) (err error) {
	var (
		key     []string
		rows    [][]string
		errHead = errors.New("table head data error")
	)
	if rows, err = f.GetRows("园区"); err != nil {
		return
	}
	for i, row := range rows {
		if i == 0 {
			key = append(key, row...)
			continue
		}
		if len(key) == 0 {
			err = errHead
			break
		}
		// 忽略空的数据
		if len(row) == 0 || (len(row) != 0 && len(row[0]) == 0) {
			continue
		}
		var (
			d = new(ProjectReview)
		)
		d.Type = "park"
		for _i, c := range row {
			if _i >= len(key) {
				break
			}
			switch key[_i] {
			case "序号":
				d.SerialNum = c
			case "亚投行成员否":
				d.IsMemberAsiaBank = c
			case "园区所在国加入亚投行时间":
				d.JoinAsiaBankTime = c
			case "园区所在洲":
				d.Continent = c
			case "地区":
				d.Region = c
			case "国家":
				d.Country = c
			case "国家发展水平":
				d.NationalDevelopmentLevel = c
			case "园区名称":
				d.Name = c
			case "中国实施企业":
				d.Enterprise = c
			case "企业性质":
				d.EnterpriseNature = c
			case "企业识别编码":
				d.EnterpriseCode = c
			case "园区所属类别":
				d.ParkType = c
			case "园区建设起始年份":
				d.StartTime = c
			default:
				err = fmt.Errorf(`dose not support this table head %v`, key[_i])
				return
			}
		}

	}
	return
}

func parseHydropowerStation(f *excelize.File) (err error) {
	var (
		key     []string
		rows    [][]string
		errHead = errors.New("table head data error")
	)
	if rows, err = f.GetRows("水电站"); err != nil {
		return
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			key = append(key, row...)
			continue
		}
		if len(key) == 0 {
			err = errHead
			break
		}
		// 忽略空的数据
		if len(row) == 0 || (len(row) != 0 && len(row[0]) == 0) {
			continue
		}
		var (
			d = new(ProjectReview)
		)
		d.Type = "hydropower_station"
		for _i, c := range row {
			if _i >= len(key) {
				break
			}
			switch key[_i] {
			case "所在地区":
				d.Region = c
			case "Region":
				d.RegionES = c
			case "所在国家":
				d.Country = c
			case "Country":
				d.CountryES = c
			case "投资金额":
				d.InvestmentAmount = c
			case "Investment amount":
				d.InvestmentAmountES = c
			case "水电站名称":
				d.HydroelectricPowerStation = c
			case "Hydroelectric power station":
				d.HydroelectricPowerStationES = c
			case "项目名称":
				d.Name = c
			case "Project":
				d.NameES = c
			case "央企集团":
				d.Enterprise = c
			case "Central Enterprise Group":
				d.EnterpriseES = c
			case "二级单位":
				d.SecondaryUnit = c
			case "Secondary unit":
				d.SecondaryUnitES = c
			case "中国建设的起始年份":
				d.StartTime = c
			case "The starting year of China's construction":
				d.StartTimeES = c
			case "水电站类型":
				d.HydropowerType = c
			case "Hydropower station type":
				d.HydropowerTypeES = c
			case "终止建设年份":
				d.TerminationTime = c
			case "Termination of construction year":
				d.TerminationTimeES = c
			case "装机容量(MW)":
				d.InstalledCapacityMW = c
			case "Installed capacity (MW)":
				d.InstalledCapacityMWES = c
			case "合作方式（投资、承建、收购）":
				d.CooperationMethod = c
			case "cooperation method":
				d.CooperationMethodES = c
			case "在91位图上定位准确性（1级、2级、3级、4级、5级）":
				d.PositionLevel = c
			case "Positioning accuracy on 91 bitmaps":
				d.PositionLevelES = c
			default:
				err = fmt.Errorf(`dose not support this table head %v`, key[_i])
				return
			}
		}

	}
	return
}

func parseOverseasPower(f *excelize.File) (err error) {
	var (
		key     []string
		rows    [][]string
		errHead = errors.New("table head data error")
	)
	if rows, err = f.GetRows("海外电力"); err != nil {
		return
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			for _, _v := range row {
				if len(_v) == 0 {
					_v = "项目状态"
				}
				key = append(key, _v)
			}
			continue
		}
		if len(key) == 0 {
			err = errHead
			break
		}
		// 忽略空的数据
		if len(row) == 0 || (len(row) != 0 && len(row[0]) == 0) {
			continue
		}
		var (
			d = new(ProjectReview)
		)
		d.Type = "overseas_power"
		for _i, c := range row {
			if _i >= len(key) {
				break
			}
			switch key[_i] {
			case "编号":
				d.SerialNum = c
			case "项目状态":
				if _i == 1 {
					d.StartTime = c
				}
				if _i == 2 {
					d.Status = c
				}
			case "集团公司":
				d.Enterprise = c
			case "二级单位":
				d.SecondaryUnit = c
			case "项目名称":
				d.Name = c
			case "规模":
				d.Scale = c
			case "所属地区":
				d.Region = c
			case "所在国家":
				d.Country = c
			case "项目类型":
				d.OverseasPowerType = c
			case "中标信息":
				d.SuccessBidInformation = c

			default:
				err = fmt.Errorf(`dose not support this table head %v`, key[_i])
				return
			}
		}
		fmt.Println(d)
	}
	return
}

func parseRailway(f *excelize.File) (err error) {
	var (
		key     []string
		rows    [][]string
		errHead = errors.New("table head data error")
	)
	if rows, err = f.GetRows("铁路"); err != nil {
		return
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			key = append(key, row...)
			continue
		}
		if len(key) == 0 {
			err = errHead
			break
		}
		// 忽略空的数据
		if len(row) == 0 || (len(row) != 0 && len(row[0]) == 0) {
			continue
		}
		var (
			d = new(ProjectReview)
		)
		d.Type = "railWay"
		for _i, c := range row {
			if _i >= len(key) {
				break
			}
			switch key[_i] {
			case "序号":
				d.SerialNum = c
			case "大洲":
				d.Continent = c
			case "国家":
				d.Country = c
			case "国家类型":
				d.NationalDevelopmentLevel = c
			case "地区":
				d.Region = c
			case "铁路类型":
				d.RailWayType = c
			case "铁路项目名称":
				d.Name = c
			case "建设企业":
				d.Enterprise = c
			case "线路长度（km）":
				d.RoadLength = c
			case "设计时速（km/h）":
				d.Speed = c
			case "预计签约/开工时间":
				d.StartTime = c
			case "预计完成时间":
				d.CompleteTime = c
			case "项目金额":
				d.InvestmentAmount = c
			case "合作方式（投资、承建、收购）":
				d.CooperationMethod = c
			default:
				err = fmt.Errorf(`dose not support this table head %v`, key[_i])
				return
			}
		}
		fmt.Println(d)
	}
	return
}

func parseHighway(f *excelize.File) (err error) {
	var (
		key     []string
		rows    [][]string
		errHead = errors.New("table head data error")
	)
	if rows, err = f.GetRows("高速公路"); err != nil {
		return
	}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if i == 1 {
			for _, r := range row {
				if len(r) == 0 {
					continue
				}
				key = append(key, r)
			}

			continue
		}
		if len(key) == 0 {
			err = errHead
			break
		}
		// 忽略空的数据
		if len(row) == 0 || (len(row) != 0 && len(row[0]) == 0) {
			continue
		}
		var (
			d = new(ProjectReview)
		)
		d.Type = "highway"
		for _i, c := range row {
			if _i >= len(key) {
				break
			}
			switch key[_i] {
			case "大洲":
				d.Continent = c
			case "详细大洲":
				d.Region = c
			case "公路（国家+公路名）":
				d.RoadName = c
			case "项目名称":
				d.Name = c
			case "中国公司":
				d.Enterprise = c
			case "线路长度":
				d.RoadLength = c
			case "建设的起始年份":
				d.StartTime = c
			case "建设期限":
				d.ConstructionPeriod = c
			case "合作方式（投资、承建、收购）":
				d.CooperationMethod = c
			case "定位等级":
				d.PositionLevel = c
			case "修建情况":
				d.RoadStatus = c
			case "公路等级":
				d.RoadLevel = c
			case "项目的战略性地位分级":
				d.StrategicPositionClass = c
			case "金额大小":
				d.InvestmentAmount = c
			default:
				err = fmt.Errorf(`dose not support this table head %v`, key[_i])
				return
			}
		}
		fmt.Println(d)
	}
	return
}

func parsePort(f *excelize.File) (err error) {
	var (
		key     []string
		rows    [][]string
		errHead = errors.New("table head data error")
	)
	if rows, err = f.GetRows("港口"); err != nil {
		return
	}
	for i, row := range rows {
		if i == 0 {
			key = append(key, row...)
			continue
		}
		if len(key) == 0 {
			err = errHead
			break
		}
		// 忽略空的数据
		if len(row) == 0 || (len(row) != 0 && len(row[0]) == 0) {
			continue
		}
		var (
			d = new(ProjectReview)
		)
		d.Type = "port"
		for _i, c := range row {
			if _i >= len(key) {
				break
			}
			switch key[_i] {
			case "序号":
				d.SerialNum = c
			case "港口所在洲":
				d.Continent = c
			case "港口项目名称":
				d.Name = c
			default:
				err = fmt.Errorf(`dose not support this table head %v`, key[_i])
				return
			}
		}
	}
	return
}

type ProjectReview struct {
	Name              string `json:"name" gorm:"type:varchar(64);index"`   // 项目名称
	NameES            string `json:"nameEs" gorm:"type:varchar(64);index"` // 项目名称
	Address           string `json:"address" gorm:"type:varchar(256);index"`
	Longitude         string `json:"longitude" gorm:"type:varchar(12);index"`          // 经度
	Latitude          string `json:"latitude" gorm:"type:varchar(12);index"`           // 纬度
	ProjectTime       string `json:"projectTime" gorm:"type:varchar(12);index"`        // 项目立项时间
	StartTime         string `json:"startTime" gorm:"type:varchar(12);index"`          // 项目启动时间
	StartTimeES       string `json:"startTimeEs" gorm:"type:varchar(12);index"`        // 项目启动时间
	OperationTime     string `json:"operationTime" gorm:"type:varchar(12);index"`      // 项目运行时间
	TerminationTime   string `json:"terminationTime" gorm:"type:varchar(12);index"`    // 终止项目时间
	TerminationTimeES string `json:"terminationTimeEs" gorm:"type:varchar(12);index"`  // 终止项目时间
	Type              string `json:"type" gorm:"type:varchar(20)"`                     //　类型
	Status            string `json:"status" gorm:"type:varchar(64);index;default:'u'"` // 立项（e: established），建设（c: construction），运营（o: operating），
	// 终止（t: terminated）和未知（u: unknown）】
	Introduction        string `json:"introduction" `                             // 项目简介
	Code                string `json:"code" gorm:"type:varchar(64);unique_index"` // 项目代码
	ImagePath           string `json:"imagePath" gorm:"type:varchar(128)"`
	Conclusion          string `json:"conclusion" gorm:"type:varchar(300)"`
	Country             string `json:"country" gorm:"type:varchar(128)"`
	CountryES           string `json:"countryEs" gorm:"type:varchar(128)"` // 国家
	Continent           string `json:"continent" gorm:"type:varchar(300)"`
	RegionES            string `json:"regionEs" gorm:"type:varchar(128)"` // 区域
	Region              string `json:"region" gorm:"type:varchar(128)"`
	Enterprise          string `json:"enterprise" gorm:"type:varchar(64)"` // 企业
	EnterpriseES        string `json:"enterpriseEs" gorm:"type:varchar(64)"`
	EnterpriseNature    string `json:"enterpriseNature" gorm:"type:varchar(12)"`   // 企业性质
	EnterpriseCode      string `json:"enterpriseCode" gorm:"type:varchar(12)"`     // 企业识别编码
	SecondaryUnit       string `json:"secondaryUnit" gorm:"type:varchar(256)"`     // 二级单位
	SecondaryUnitES     string `json:"secondaryUnitEs" gorm:"type:varchar(256)"`   // 二级单位
	CooperationMethod   string `json:"cooperationMethod" gorm:"type:varchar(128)"` // 合作方式（投资、承建、收购）
	CooperationMethodES string `json:"cooperationMethodEs" gorm:"type:varchar(128)"`
	PositionLevel       string `json:"positionLevel"  gorm:"type:varchar(36)"` // 定位等级（1级、2级、3级、4级、5级）
	PositionLevelES     string `json:"positionLevelEs"  gorm:"type:varchar(36)"`
	SerialNum           string `json:"serialNum"  gorm:"type:varchar(12)"`
	InvestmentAmount    string `json:"investmentAmount" gorm:"type:varchar(36)"` // 投资金额
	InvestmentAmountES  string `json:"investmentAmountEs" gorm:"type:varchar(36)"`
	// 水电站
	CurrentStatus               int64  `json:"currentStatus" gorm:"type:int"`
	HydroelectricPowerStation   string `json:"hydroelectricPowerStation" gorm:"type:varchar(128)"` // 水电站名称
	HydroelectricPowerStationES string `json:"hydroelectricPowerStationEs" gorm:"type:varchar(128)"`
	InstalledCapacityMW         string `json:"installedCapacityMw" gorm:"type:varchar(128)"` // 装机容量(MW)
	InstalledCapacityMWES       string `json:"installedCapacityMwEs" gorm:"type:varchar(128)"`
	HydropowerType              string `json:"hydropowerType" gorm:"type:varchar(24)"`   // 水电站类型
	HydropowerTypeES            string `json:"hydropowerTypeEs" gorm:"type:varchar(24)"` // 水电站类型
	// 高速路
	RoadName               string `json:"roadName" gorm:"type:varchar(64)"`               // 路名
	RoadLength             string `json:"roadLength"  gorm:"type:varchar(64)"`            // 路长度
	ConstructionPeriod     string `json:"constructionPeriod"  gorm:"type:varchar(12)"`    // 建设期限
	StrategicPositionClass string `json:"strategicPositionClass" gorm:"type:varchar(12)"` // 战略性地位分级
	RoadLevel              string `json:"roadLevel" gorm:"type:varchar(12)"`              // 公路等级
	RoadStatus             string `json:"roadStatus" gorm:"type:varchar(12)"`             // 公路修建情况
	// 铁路
	Speed              string `json:"speed" gorm:"type:varchar(12)"`
	EstimateFinishTime string `json:"estimateFinishTime"  gorm:"type:varchar(64)"` // 预计完成时间
	RailWayType        string `json:"railWayType" gorm:"type:varchar(12)"`         // 铁路类型
	CompleteTime       string `json:"completeTime" gorm:"type:varchar(32)"`        // 完成时间
	// 电力项目
	Scale                 string `json:"scale" gorm:"type:varchar(12)"`
	OverseasPowerType     string `json:"overseasPowerType" gorm:"type:varchar(12)"` // 海外电力项目类型
	SuccessBidInformation string `json:"successBidInformation"`                     // 中标信息
	// 园区
	IsMemberAsiaBank         string `json:"isMemberAsiaBank"` // 是否是亚投行成员
	JoinAsiaBankTime         string `json:"joinAsiaBankTime" gorm:"type:varchar(12)"`
	NationalDevelopmentLevel string `json:"nationalDevelopmentLevel" gorm:"type:varchar(32)"` // 国家发展水平
	ParkType                 string `json:"parkType" gorm:"type:varchar(12)"`                 // 园区类别
}
