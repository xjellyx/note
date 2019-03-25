package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func main() {
	app := iris.New()
	app.Use(logger.New())
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML(`<!DOCTYPE html>
<html>
<head> 
<meta charset="utf-8"> 
<title>哈哈</title> 
</head>
<h1><center>送给生气的小可爱</center></h1>
<body>
<center><img src="https://ss0.bdstatic.com/70cFuHSh_Q1YnxGkpoWK1HF6hhy/it/u=1501065703,1244450463&fm=26&gp=0.jpg"></center>
<p><span>  亲爱的，女生节快乐哦，没什么送给你，送个比较土的礼物，爱你</span></p>
  <style>
            span {
                color: red;
            }
        </style>
	<p>点击播放</p>
<audio controls>
  <source src="http://other.web.np01.sycdn.kuwo.cn/resource/n3/84/40/2548944492.mp3" type="audio/ogg">
  <source src="http://other.web.np01.sycdn.kuwo.cn/resource/n3/84/40/2548944492.mp3" type="audio/mpeg">
</audio>
	<center><img src="https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=1702611453,2473642557&fm=26&gp=0.jpg"></center>
<p>   亲爱的，和你在一起差不多四年了，一直以来都没有跟你告白，因为在你
  心里都是一个你觉得我一直是一个直男,不过估计也一直是吧，至少我也觉得了
  了，哈哈。奈你需浪漫而我视而不见，在一起这么久你一直都没有见到我刻画过
	属于我们的樱花世界。<br/>文学浅，无之想之甜蜜；直癌久，无往日之漫语；时而语之
	凶，恐吓之水心；时而气大，无视之意暖。经日省自身，天长地久，吾之唯心
	疼之之三生；细水长流，吾之长情恋之之三世；吾愿以痞手牵之素手至天涯海角，
	唯之吾姬。今以此告白，吾，梁某，爱之一生一世，山河可鉴。以花送予。
	</p>
	<h2><center>我不要短暂的温存，只要你一世的陪伴。</center></h2>
<center><img src= "http://img0.imgtn.bdimg.com/it/u=1957088614,543784412&fm=26&gp=0.jpg"></center>
</body>
</html>
`)
	})
	app.Get("/", func(context iris.Context) {
		context.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	if err := app.Run(iris.Addr(":8080")); err != nil {
		panic(err)
	}
}
