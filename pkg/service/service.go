package service

import (
	"testbank/bank"
	"testbank/pkg/repository"
)

type (
	Service struct {
		UserService
		AccountService
	}
	UserService interface {
		CreateUser(u bank.User) (int, error)
		GetUser(id string) (bank.User, error)
	}
	AccountService interface {
		CreateAccount(userId string, a bank.Account) (int, error)
		AccountTransaction(senderId int, a bank.AccountTransaction) error
		GetAllUserAccounts(id string) ([]bank.Account, error)
		GetTransaction(id string) ([]bank.TransactionList, error)
	}
)

func NewService(r repository.Repository) Service {
	return Service{
		UserService:    NewUserService(r),
		AccountService: NewAccountService(r),
	}
}
