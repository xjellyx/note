package model

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/olongfen/note/game"
)

var (
	ModelUser  *gorm.DB
	ModelRedis *redis.Client
)

// InitModel
func InitModel(c struct {
	DBConnArgs string
	RedisHost  string
	RedisPort  string
}) (err error) {
	var (
		db *gorm.DB
	)
	ModelRedis = redis.NewClient(&redis.Options{
		Addr: c.RedisHost + ":" + c.RedisPort,
	})

	if _, err = ModelRedis.Ping().Result(); err != nil {
		return
	}
	if db, err = gorm.Open("postgres", c.DBConnArgs); err != nil {
		return
	}
	ModelUser = db.AutoMigrate(&game.User{}).Model(&game.User{})
	defer db.Close()
	return
}
