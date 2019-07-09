package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	aai "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai/v20180522"
)

func main() {
	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(
		"AKIDV7DNUB5lZDeB7ru4mvSfzEibJ2JRszky",
		"Ly1jovx4gYyAAlndqap4nYQHA12o3dcO",
	)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	// 如非必要请不要修改默认设置。
	//cpf.HttpProfile.ReqMethod = "GET"
	// SDK有默认的超时时间，如非必要请不要修改默认设置。
	// 如有需要请在代码中查阅以获取最新的默认值。
	//cpf.HttpProfile.ReqTimeout = 10
	// SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务，
	// 则必须手动指定域名，例如云服务器的上海金融区域名： cvm.ap-shanghai-fsi.tencentcloudapi.com
	//cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	// SDK默认用HmacSHA256进行签名，它更安全但是会轻微降低性能。
	// 如非必要请不要修改默认设置。
	//cpf.SignMethod = "HmacSHA1"
	// SDK 默认用 zh-CN 调用返回中文。此外还可以设置 en-US 返回全英文。
	// 但大部分产品或接口并不支持全英文的返回。
	// 如非必要请不要修改默认设置。
	//cpf.Language = "en-US"

	// 实例化要请求产品(以cvm为例)的client对象
	// 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量
	client, _ := aai.NewClient(credential, regions.Guangzhou, cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定DescribeInstancesRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	request := aai.NewTextToVoiceRequest()

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.Chat(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}