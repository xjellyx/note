package model

import (
	"github.com/kataras/iris/core/errors"
	"sync"
)

// Query represents the visitor and action queries.
type Query func(Movie) bool

// MovieRepository handles the basic operations of a movie entity/model.
// It's an interface in order to be testable, i.e a memory movie repository or
// a connected to an sql database.
type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (movie Movie, found bool)
	SelectMany(query Query, limit int) (results []Movie)
	InsertOrUpdate(movie Movie) (updatedMovie Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}

// NewMovieRepository returns a new movie memory-based repository,
// the one and only repository type in our example.
func NewMovieRepository(source map[int64]Movie) MovieRepository {
	return &movieMemoryRepository{source: source}
}

// movieMemoryRepository is a "MovieRepository"
// which manages the movies using the memory data source (map).
type movieMemoryRepository struct {
	source map[int64]Movie
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)

func (r *movieMemoryRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loops := 0
	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}
	for _, movie := range r.source {
		ok = query(movie)
		if ok {
			if action(movie) {
				loops++
				if actionLimit >= loops {
					break // break
				}
			}
		}
	}
	return
}

//选择接收查询功能
//为内部的每个电影模型触发
//我们想象中的数据源
//当该函数返回true时，它会停止迭代。

//它返回查询返回的最后一个已知“找到”值
//和最后一个已知的电影模型
//帮助呼叫者减少LOC。

//它实际上是一个简单但非常聪明的原型函数
//自从我第一次想到它以来，我一直在使用它，
//希望你会发现它也很有用。
func (r *movieMemoryRepository) Select(query Query) (movie Movie, found bool) {
	found = r.Exec(query, func(m Movie) bool {
		movie = m
		return true
	}, 1, ReadOnlyMode)

	//如果根本找不到的话,设置一个空的datamodels.Movie，
	if !found {
		movie = Movie{}
	}
	return
}

// SelectMany与Select相同但返回一个或多个datamodels.Movie作为切片。
//如果limit <= 0则返回所有内容
func (r *movieMemoryRepository) SelectMany(query Query, limit int) (results []Movie) {
	r.Exec(query, func(m Movie) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)
	return
}

// InsertOrUpdate将影片添加或更新到（内存）存储。
// 返回新电影，如果有则返回错误。
func (r *movieMemoryRepository) InsertOrUpdate(movie Movie) (Movie, error) {
	id := movie.ID
	if id == 0 { // Create new action
		var lastID int64
		//找到最大的ID，以便不重复
		//在制作应用中，您可以使用第三方
		//库以生成UUID作为字符串。
		r.mu.RLock()
		for _, item := range r.source {
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()
		id = lastID + 1
		movie.ID = id
		// map-specific thing
		r.mu.Lock()
		r.source[id] = movie
		r.mu.Unlock()

		return movie, nil
	}
	//基于movie.ID更新动作，
	//这里我们将允许更新海报和流派，如果不是空的话。
	//或者我们可以做替换：
	// r.source [id] =电影
	//并评论下面的代码;
	current, exists := r.Select(func(m Movie) bool {
		return m.ID == id
	})
	if !exists { //ID不是真实的，返回错误。
		return Movie{}, errors.New("failed to update a nonexistent movie")
	}
	// 或者注释这些和r.source [id] = m进行纯替换
	if movie.Poster != "" {
		current.Poster = movie.Poster
	}
	if movie.Genre != "" {
		current.Genre = movie.Genre
	}
	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()
	return movie, nil
}
func (r *movieMemoryRepository) Delete(query Query, limit int) bool {
	return r.Exec(query, func(m Movie) bool {
		delete(r.source, m.ID)
		return true
	}, limit, ReadWriteMode)
}
