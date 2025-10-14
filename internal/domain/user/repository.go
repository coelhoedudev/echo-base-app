package user

import (
	"infra-base-go/pkg/util"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) (string, error)
	FindAll() ([]User, error)
	Find(id string) (User, error)
	Update(user *User) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(user *User) (string, error) {
	uuid := util.UUID{}.Create()

	err := r.db.Create(user).Error

	return uuid, err
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) Find(id string) (User, error) {
	var user User
	err := r.db.First(user, id).Error

	return user, err

}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
