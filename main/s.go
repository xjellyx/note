package main

import "gorm.io/gorm"

// Add add one record
func (t *User) Add(db *gorm.DB) (err error) {
	if err = db.Create(t).Error; err != nil {
		return
	}
	return
}

// Delete delete record
func (t *User) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		return
	}
	return
}

// Updates update record
func (t *User) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Where("id = ?", t.ID).Updates(m).Error; err != nil {
		return
	}
	return
}

// GetUserAll get all record
func GetUserAll(db *gorm.DB) (ret []*User, err error) {
	if err = db.Find(&ret).Error; err != nil {
		return
	}
	return
}

// GetUserCount get count
func GetUserCount(db *gorm.DB) (ret int64) {
	db.Model(&User{}).Count(&ret)
	return
}

// GetByID get one record by ID
func (t *User) GetByID(db *gorm.DB) (err error) {
	if err = db.First(t, "id = ?", t.ID).Error; err != nil {
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *User) DeleteByID(db *gorm.DB) (err error) {
	if err = db.Delete(t, "id = ?", t.ID).Error; err != inl {
		return
	}
	return
}

// GetByName get one record by Name
func (t *User) GetByName(db *gorm.DB) (err error) {
	if err = db.First(t, "name = ?", t.Name).Error; err != nil {
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *User) DeleteByName(db *gorm.DB) (err error) {
	if err = db.Delete(t, "name = ?", t.Name).Error; err != inl {
		return
	}
	return
}
