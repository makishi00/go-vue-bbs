package service

import (
	"github.com/makishi00/go-vue-bbs/model"
	"golang.org/x/crypto/bcrypt"
)


var User = user{}

type user struct {
}

func (u *user) Store(user model.User) model.User {
	user.Password, _ = HashPassword(string(user.Password))
	db.Create(&user)
	return user
}
func (u *user) ExisByEmail(email string) bool {
	var users []model.User
	db.Where("email = ?", email).Find(&users)
	return len(users) != 0
}
func (u *user) ExisByName(name string) bool {
	var users []model.User
	db.Where("name = ?", name).Find(&users)
	return len(users) != 0
}

func (u *user) Login(email, pass string) (*model.User, bool) {
	var users []model.User
	db.Where("email = ?", email).Find(&users)
	if len(users) == 0 {
		return nil, false
	}
	return &users[0], CheckPasswordHash(pass, string(users[0].Password))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}