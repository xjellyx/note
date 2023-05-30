package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	HTTP HTTP
}

type HTTP struct {
	IP   string
	Port int64 `mapstructure:"port"`
}

func main() {

	// 设置默认值
	viper.SetDefault("http.ip", "0.0.0.0")
	viper.SetDefault("http.port", "9998")
	pflag.String("http.ip", "127.0.0.1", "--http.ip=127.0.0.1")
	pflag.Int64("http.port", 8888, "")
	pflag.Parse()

	// 绑定命令行参数e
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
	// 绑定环境变量
	godotenv.Load()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	// 设置配置文件信息，io.Reader 方式读取配置时，也需要这些配置
	viper.AddConfigPath(".") // // 把当前目录加入到配置文件的搜索路径中，可调用多次添加多个搜索路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//搜索配置文件，获取配置
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 序列化
	var (
		cfg = new(Config)
	)
	viper.Unmarshal(cfg)
	fmt.Println("===== 序列化 =====")
	fmt.Println(cfg)
}
