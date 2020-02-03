package serve

import "github.com/olefen/note/iris/demo/model"

type MovieService interface {
	GetAll() []model.Movie
	GetByID(id int64) (model.Movie, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndGenreByID(id int64, poster string, genre string) (model.Movie, error)
}

// NewMovieService返回默认 movie service.
func NewMovieService(repo model.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

type movieService struct {
	repo model.MovieRepository
}

// GetAll 获取所有的movie.
func (s *movieService) GetAll() []model.Movie {
	return s.repo.SelectMany(func(_ model.Movie) bool {
		return true
	}, -1)
}

// GetByID 根据其ID返回一行。
func (s *movieService) GetByID(id int64) (model.Movie, bool) {
	return s.repo.Select(func(m model.Movie) bool {
		return m.ID == id
	})
}

// UpdatePosterAndGenreByID更新电影的海报和流派。
func (s *movieService) UpdatePosterAndGenreByID(id int64, poster string, genre string) (model.Movie, error) {
	// update the movie and return it.
	return s.repo.InsertOrUpdate(model.Movie{
		ID:     id,
		Poster: poster,
		Genre:  genre,
	})
}

// DeleteByID按ID删除电影。
//
//如果删除则返回true，否则返回false。
func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m model.Movie) bool {
		return m.ID == id
	}, 1)
}
