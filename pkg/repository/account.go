package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"testbank/bank"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccount(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}
func (p *AccountPostgres) CreateAccount(idUser string, a bank.Account) (int, error) {
	var id int
	tx, err := p.db.Begin()
	if err != nil {
		logrus.Fatal(err)
		return 0, err
	}
	query := fmt.Sprintf("INSERT INTO %s (currency, total) VALUES ($1, $2) returning id;", accountsTable)
	row := tx.QueryRow(query, a.Currency, a.Total)

	if err := row.Scan(&id); err != nil {
		logrus.Fatal(err)
		tx.Rollback()
		return 0, err
	}
	accountListQuery := fmt.Sprintf("INSERT INTO %s (users_id, account_id) VALUES ($1, $2)", accountsListsTable)
	_, err = tx.Exec(accountListQuery, idUser, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	updateUser := fmt.Sprintf("UPDATE %s SET account=true WHERE id = $1", usersTable)
	_, err = tx.Exec(updateUser, idUser)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (p *AccountPostgres) AccountTransaction(senderId int, a bank.AccountTransaction) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	checkUser, err := getUserByEmailandAccount(tx, a.Email)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	if !checkUser.Account {
		logrus.Fatal("user dont have accounts")
		return errors.New("user dont have accounts")
	}

	logrus.Print("update account")
	//  Sender
	updateSender, err := checkUserAcount(tx, senderId, a.Currency)
	if err != nil {
		logrus.Fatal(err)
		return err
	}

	if updateSender.Total < a.Total {
		logrus.Fatal("dont have money an account")
		return errors.New("dont have money an account")
	}
	updateSender.Total = updateSender.Total - a.Total
	logrus.Print("ok sender")
	checkUserAcc, err := checkUserAcount(tx, checkUser.Id, a.Currency)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	logrus.Print("ok get")
	query := fmt.Sprintf(`update %s set total=$1 where currency=$2 and id=$3`, accountsTable)
	if _, err = tx.Exec(query, updateSender.Total, a.Currency, updateSender.Id); err != nil {
		tx.Rollback()
		return err
	}
	logrus.Print("send")

	//  Getter
	checkUserAcc.Total = checkUserAcc.Total + a.Total
	logrus.Print(checkUserAcc.Total)
	query = fmt.Sprintf(`update %s set total=$1 where currency=$2 and id=$3`, accountsTable)
	if _, err = tx.Exec(query, checkUserAcc.Total, a.Currency, checkUserAcc.Id); err != nil {
		tx.Rollback()
		return err
	}
	logrus.Print("on acc")
	date := time.Now().UTC()
	transactionList := fmt.Sprintf(`INSERT INTO %s (user_id_sender, account_id, user_id_geter ,currency, total, sendsum, date) values($1,$2,$3,$4,$5,$6,$7)`, transactionListTable)
	if _, err = tx.Exec(transactionList, senderId, updateSender.Id, checkUser.Id, a.Currency, updateSender.Total, a.Total, date); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func getUserByEmailandAccount(tx *sql.Tx, email string) (bank.User, error) {
	var user bank.User
	logrus.Print("get user by email")
	query := fmt.Sprintf("SELECT id, email , account FROM %s WHERE email = $1", usersTable)
	row := tx.QueryRow(query, email)
	if err := row.Scan(&user.Id, &user.Email, &user.Account); err != nil {
		tx.Rollback()
		logrus.Errorf("Unable to Select in get user by email: %v\n", err)
		return user, err
	}
	logrus.Print(user)
	return user, nil
}

func checkUserAcount(tx *sql.Tx, id_user int, currency string) (bank.AccountTransaction, error) {
	var at bank.AccountTransaction
	logrus.Print("check user account")
	query := fmt.Sprintf(`SELECT  at.id, at.currency, at.total FROM %s at INNER JOIN %s al on at.id = al.account_id where al.users_id = $1 and at.currency = $2 `, accountsTable, accountsListsTable)
	//   S`ELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`
	row := tx.QueryRow(query, id_user, currency)
	if err := row.Scan(&at.Id, &at.Currency, &at.Total); err != nil {
		logrus.Fatal(err)
		tx.Rollback()
		return at, err
	}
	logrus.Print(at)
	logrus.Print("check user done!")
	return at, nil
}

func (p *AccountPostgres) GetAllUserAccounts(id string) ([]bank.Account, error) {
	var acc []bank.Account

	query := fmt.Sprintf(`select ac.id , ac.currency, ac.total from %s ac inner join %s al on ac.id = al.account_id where al.users_id=$1`, accountsTable, accountsListsTable)
	if err := p.db.Select(&acc, query, id); err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	return acc, nil
}

func (p *AccountPostgres) GetTransaction(id string) ([]bank.TransactionList, error) {
	var tl []bank.TransactionList
	var t bank.TransactionList

	query := fmt.Sprintf(`select id, user_id_sender, account_id, user_id_geter, currency, total, sendsum, date from %s where user_id_sender=$1`, transactionListTable)
	row, err := p.db.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&t.Id, &t.SenderId, &t.AccountId, &t.GeterId, &t.Currency, &t.Total, &t.Sendsum, &t.Date); err != nil {
			logrus.Fatal(err)
			return nil, err
		}
		tl = append(tl, t)
	}

	return tl, nil
}
