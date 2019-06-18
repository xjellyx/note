package ctrl

import (
	"github.com/srlemon/note/iris_/demo/model"
	"github.com/srlemon/note/iris_/demo/serve"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
)

type MovieCTRL struct {
	Serve serve.MovieService
}

func (c *MovieCTRL) Get() (results []model.Movie) {
	return c.Serve.GetAll()
}

func (c *MovieCTRL) GetBy(id int64) (movie model.Movie, found bool) {
	return c.Serve.GetByID(id) // it will throw 404 if not found.
}

// 用put请求更新一部电影
// Demo:
// curl -i -X PUT -F "genre=Thriller" -F "poster=@/Users/kataras/Downloads/out.gif" http://localhost:8080/movies/1
func (c *MovieCTRL) PutBy(ctx iris.Context, id int64) (model.Movie, error) {
	// get the request data for poster and genre
	file, info, err := ctx.FormFile("poster")
	if err != nil {
		return model.Movie{}, errors.New("failed due form file 'poster' missing")
	}
	// 不需要文件所以关闭他
	file.Close()
	//想象一下，这是上传文件的网址......
	poster := info.Filename
	genre := ctx.FormValue("genre")
	return c.Serve.UpdatePosterAndGenreByID(id, poster, genre)
}

// Delete请求删除一部电影
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *MovieCTRL) DeleteBy(id int64) interface{} {
	wasDel := c.Serve.DeleteByID(id)
	if wasDel {
		// 返回删除的id
		return iris.Map{"deleted": id}
	}
	//在这里我们可以看到方法函数可以返回这两种类型中的任何一种（map或int），
	//我们不必将返回类型指定为特定类型。
	return iris.StatusBadRequest
}
