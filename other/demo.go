package main

import (
	"fmt"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"time"
)

type ss struct {
	Viper *viper.Viper `yaml:"-"`
	Conf  *Conf
}

type Conf struct {
	DBHost string `json:"db_host"`
	DBPort string `json:"db_port"`
	Data   Data
}
type Data struct {
	Name          string `json:"name"`
	TransportFlow string `json:"transportFlow"`
}

func main() {
	a := time.Now().Unix()
	time.Sleep(time.Second * 5)
	b := time.Now().Unix()
	fmt.Println(b - a)
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	body, _ := proto.Marshal(&protobuf.SAHHDZJoinRoom{
	//		IRet: proto.Int32(54 + 564564654),
	//	})
	//	w.Write(body)
	//
	//})
	//
	//http.ListenAndServe("192.168.31.184:12345", nil)
}

// 前端传入数据时，将float转为int
func PubFloatToInt(input float64) (ret int64) {
	// 初始化成相应n次方类型
	power := decimal.NewFromFloat(5)
	// 保留相应的小数位数然后再乘以相应的10的n次方，最后再转换为int类型
	ret = decimal.NewFromFloat(input).Truncate(5).Mul(power).IntPart()
	return
}

// 前端取数据时，将int转为float
func PubIntToFloat(input int64) (ret float64) {
	// 保留相应的小数
	ret = float64(input) / 5
	return
}

// 保留有效小数
func PubKeepDecimal(input float64) (ret float64) {
	// 保留相应的小数并转换为float
	ret, _ = decimal.NewFromFloat(input).Truncate(5).Float64()
	return
}

// 浮点数加
func PubFloatAdd(a float64, b float64, vals ...float64) (c float64) {
	n := decimal.NewFromFloat(a)
	n = n.Add(decimal.NewFromFloat(b))
	for _, v := range vals {
		n = n.Add(decimal.NewFromFloat(v))
	}
	c, _ = n.Float64()

	return
}

// 浮点数加,并保留有效位数
func PubFloatAddRound(a float64, b float64, vals ...float64) (c float64) {
	c = PubFloatRoundAuto(PubFloatAdd(a, b, vals...))
	return
}

// 浮点数乘
func PubFloatMul(a float64, b float64) (c float64) {
	n := decimal.NewFromFloat(a)
	n = n.Mul(decimal.NewFromFloat(b))
	c, _ = n.Float64()
	return
}

// 浮点数乘,并保留有效位数
func PubFloatMulRound(a float64, b float64) (c float64) {
	c = PubFloatRoundAuto(PubFloatMul(a, b))
	return
}

// 浮点数除
func PubFloatDiv(a float64, b float64) (c float64) {
	n := decimal.NewFromFloat(a)
	n = n.Div(decimal.NewFromFloat(b))
	c, _ = n.Float64()
	return
}

// 整形转浮点数除
func PubIntDivToFloat(a, b int64) (c float64) {
	c = PubFloatDiv(PubIntToFloat(a), PubIntToFloat(b))
	return
}

// 浮点数除,并保留有效位数
func PubFloatDivRound(a float64, b float64) (c float64) {
	c = PubFloatRoundAuto(PubFloatDiv(a, b))
	return
}

func PubFloatRoundAuto(a float64) (c float64) {
	// 四舍五入
	r := int32(5)
	n := decimal.NewFromFloat(a).Round(r)
	c, _ = n.Float64()
	return
}
func DecodeToMap(s *structpb.Struct) map[string]interface{} {
	if s == nil {
		return nil
	}
	m := map[string]interface{}{}
	for k, v := range s.Fields {
		m[k] = decodeValue(v)
	}
	return m
}

func decodeValue(v *structpb.Value) interface{} {
	switch k := v.Kind.(type) {
	case *structpb.Value_NullValue:
		return nil
	case *structpb.Value_NumberValue:
		return k.NumberValue
	case *structpb.Value_StringValue:
		return k.StringValue
	case *structpb.Value_BoolValue:
		return k.BoolValue
	case *structpb.Value_StructValue:
		return DecodeToMap(k.StructValue)
	case *structpb.Value_ListValue:
		s := make([]interface{}, len(k.ListValue.Values))
		for i, e := range k.ListValue.Values {
			s[i] = decodeValue(e)
		}
		return s
	default:
		panic("protostruct: unknown kind")
	}
}
