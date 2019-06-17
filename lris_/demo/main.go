package main

import (
	"github.com/LnFen/note/lris-learn/demo/ctrl"
	"github.com/LnFen/note/lris-learn/demo/model"
	"github.com/LnFen/note/lris-learn/demo/serve"
	"github.com/kataras/iris"
	"github.com/kataras/iris/_examples/mvc/overview/web/controllers"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//加载模板文件
	app.RegisterView(iris.HTML("./demo/views", ".html"))
	// 注册控制器
	// mvc.New(app.Party("/movies")).Handle(new(controllers.MovieController))
	//您还可以拆分您编写的代码以配置mvc.Application
	//使用`mvc.Configure`方法，如下所示。
	mvc.Configure(app.Party("/movies"), movies)
	// http://localhost:8080/movies
	// http://localhost:8080/movies/1
	app.Run(
		//开启web服务
		iris.Addr("localhost:8080"),
		// 禁用更新
		iris.WithConfiguration(iris.TOML("./main.tml")),
		// 按下CTRL / CMD + C时跳过错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//实现更快的json序列化和更多优化：
		iris.WithOptimizations,
	)
}

//注意mvc.Application，它不是iris.Application。
func movies(app *mvc.Application) {
	//添加基本身份验证（admin：password）中间件
	//用于基于/电影的请求。
	app.Router.Use(ctrl.BasicAuth)
	// 使用数据源中的一些（内存）数据创建我们的电影资源库。
	repo := model.NewMovieRepository(model.Movies)
	// 创建我们的电影服务，我们将它绑定到电影应用程序的依赖项。
	movieService := serve.NewMovieService(repo)
	app.Register(movieService)
	//为我们的电影控制器服务
	//请注意，您可以为多个控制器提供服务
	//你也可以使用`movies.Party（relativePath）`或`movies.Clone（app.Party（...））创建子mvc应用程序
	app.Handle(new(controllers.MovieController))
}
