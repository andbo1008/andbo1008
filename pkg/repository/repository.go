package repository

import (
	"testbank/bank"

	"github.com/jmoiron/sqlx"
)

type (
	Repository struct {
		User
		Account
	}
	User interface {
		CreateUser(u bank.User) (int, error)
		GetUser(id string) (bank.User, error)
	}
	Account interface {
		CreateAccount(userId string, a bank.Account) (int, error)
		AccountTransaction(senderId int, a bank.AccountTransaction) error
		GetAllUserAccounts(id string) ([]bank.Account, error)
		GetTransaction(id string) ([]bank.TransactionList, error)
	}
)

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		User:    NewUser(db),
		Account: NewAccount(db),
	}
}
