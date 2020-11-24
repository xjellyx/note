package database

import "gorm.io/gorm"

type Database struct {
	DB *gorm.DB
}

func NewDatabase(dialector gorm.Dialector, c *gorm.Config) (ret *Database, err error) {
	var (
		d = new(Database)
	)
	if d.DB, err = gorm.Open(dialector, c); err != nil {
		return
	}
	ret = d
	return
}

func (d *Database) GetData() (ret interface{}, err error) {
	return
}
