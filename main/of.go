package main

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jprom "github.com/uber/jaeger-lib/metrics/prometheus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
)

const samplerParam = 1
const reportingHost = "192.168.3.85:6831"

type TraceHandler struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

type CaiPiao struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name" gorm:"type:varchar(36)"`
	Code        string      `json:"code" gorm:"type:varchar(12);uniqueIndex"`
	Date        string      `json:"date" gorm:"type:varchar(36)"`
	Week        string      `json:"week" gorm:"type:varchar(36)"`
	Red         string      `json:"red" gorm:"type:varchar(36)"`
	Blue        string      `json:"blue" gorm:"type:varchar(36)"`
	Blue2       string      `json:"blue2" gorm:"type:varchar(36)"`
	Sales       string      `json:"sales" gorm:"type:varchar(36)"`
	PoolMoney   string      `json:"poolmoney" gorm:"type:varchar(36)"`
	Content     string      `json:"content" `
	AddMoney    string      `json:"addmoney" gorm:"type:varchar(36)"`
	AddMoney2   string      `json:"addmoney2" gorm:"type:varchar(36)"`
	Msg         string      `json:"msg" `
	Z2Add       string      `json:"z2add" gorm:"type:varchar(36)"`
	M2Add       string      `json:"m2add" gorm:"type:varchar(36)"`
	PrizeGrades PrizeGrades `json:"prizegrades" gorm:"type: json"`
	Zj1         string      `json:"zj1,omitempty"`
	Mj1         string      `json:"mj1,omitempty"`
	Zj6         string      `json:"zj6,omitempty"`
	Mj6         string      `json:"mj6,omitempty"`
}

type PrizeGrades []*PrizeGrade

func (p *PrizeGrades) Scan(in interface{}) error {
	return json.Unmarshal(in.([]byte), p)
}

func (p PrizeGrades) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type PrizeGrade struct {
	Type      int    `json:"type"`
	TypeNum   string `json:"type_num"`
	TypeMoney string `json:"type_money"`
}

var globalTracerHandler *TraceHandler
var (
	db *gorm.DB
)

func init() {
	cfg := config.Configuration{ServiceName: "jaeger_test",
		Sampler: &config.SamplerConfig{Type: jaeger.SamplerTypeConst,
			Param: samplerParam},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: reportingHost,
		}}
	me := jprom.New(jprom.WithRegisterer(prometheus.NewPedanticRegistry()))
	tracer, closer, err := cfg.NewTracer(config.Logger(jaegerlog.StdLogger), config.Metrics(me))
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err.Error()}).Errorf("%s", "初始化tracer失败")
	}
	globalTracerHandler = &TraceHandler{Tracer: tracer, Closer: closer}
	opentracing.SetGlobalTracer(globalTracerHandler.Tracer)
	if db, err = gorm.Open(postgres.Open(fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "business",
		"business", "127.0.0.1", "5432", "business"))); err != nil {
		logrus.Panic(err)
	}
	db.Use(&OpentracingPlugin{})
}

func get(c *gin.Context) {
	//span := globalTracerHandler.Tracer.StartSpan("hello")
	span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "获取数据")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(c.Request.Context(), span)
	c.JSON(200, gin.H{"data": srvGet(ctx, c.Query("say"))})
}

func srvGet(ctx context.Context, param string) (res []*CaiPiao) {
	// 创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "返回数据")
	defer func() {
		span.SetTag("recv", param)
		span.Finish()
	}()
	db.WithContext(ctx).Model(&CaiPiao{}).Limit(10).Find(&res)
	return res
}

// 包内静态变量
const gormSpanKey = "__gorm_span"
const (
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
)

type OpentracingPlugin struct{}

// 告诉编译器这个结构体实现了gorm.Plugin接口
var _ gorm.Plugin = &OpentracingPlugin{}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前 - 并不是都用相同的方法，可以自己自定义
	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后 - 并不是都用相同的方法，可以自己自定义
	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

func (op *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}
func before(db *gorm.DB) {
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, "gorm")
	// 利用db实例去传递span
	db.InstanceSet(gormSpanKey, span)
}

func after(db *gorm.DB) {
	_span, exist := db.InstanceGet(gormSpanKey)
	if !exist {
		return
	}
	// 断言类型
	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}

	defer span.Finish()

	if db.Error != nil {
		span.LogFields(tracerLog.Error(db.Error))
	}

	span.LogFields(tracerLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))

}

func main() {

	mux := gin.Default()
	mux.Use(ginhttp.Middleware(globalTracerHandler.Tracer))
	mux.GET("", get)
	mux.Run()
}
