package repository

import (
	"fmt"
	"testbank/bank"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (p *UserPostgres) CreateUser(u bank.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, lastname, email, password) values ($1, $2, $3, $4) RETURNING id;", usersTable)
	row := p.db.QueryRow(query, u.Name, u.LastName, u.Email, u.Password)

	if err := row.Scan(&id); err != nil {
		logrus.Errorf("Unable to INSERT: %v\n", err)
		return 0, nil
	}
	// createAt := time.Now()
	// query = "INSERT INTO tech_info_users (tech_users_id ,create_at) values ($1, $2);"
	// p.db.QueryRow(query, id, createAt)
	// fmt.Println(id)
	return id, nil
}

func (p *UserPostgres) GetUser(id string) (bank.User, error) {
	var user bank.User

	query := fmt.Sprintf("SELECT id, name, lastname, email , account FROM %s WHERE id = $1", usersTable)

	row := p.db.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Account); err != nil {
		logrus.Errorf("Unable to Select in get user: %v\n", err)
		return user, err
	}

	logrus.Print(user)
	return user, nil
}
