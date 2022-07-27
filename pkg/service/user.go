package service

import (
	"crypto/md5"
	"encoding/hex"
	"testbank/bank"
	"testbank/pkg/repository"
)

type User struct {
	repository repository.User
}

func NewUserService(r repository.User) *User {
	return &User{repository: r}
}

func (s *User) CreateUser(u bank.User) (int, error) {
	u.Password = generetePasswordHash(u.Password)
	return s.repository.CreateUser(u)
}
func (s *User) GetUser(id string) (bank.User, error) {
	return s.repository.GetUser(id)
}

func generetePasswordHash(pwd string) string {
	hash := md5.Sum([]byte(pwd))
	return hex.EncodeToString(hash[:])
}
