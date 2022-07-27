package bank

import "time"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Account  bool   `json:"account"`
}
type Account struct {
	Id       int     `json:"account_id"`
	Currency string  `json:"currency"`
	Total    float64 `json:"total"`
}
type AccountsList struct {
	Id        int `json:"id"`
	UserId    int `json:"users_id"`
	AccountId int `json:"account_id"`
}
type AccountTransaction struct {
	Id       int     `json:"account_transaction_id"`
	Email    string  `json:"email"`
	Currency string  `json:"currency"`
	Total    float64 `json:"total"`
}

type TransactionList struct {
	Id        int       `json:"id"`
	SenderId  int       `json:"user_id_sender"`
	AccountId int       `json:"account_id"`
	GeterId   int       `json:"user_id_geter"`
	Currency  string    `json:"currency"`
	Total     float64   `json:"total"`
	Sendsum   float64   `json:"sendsum"`
	Date      time.Time `json:"date"`
}

// type TechInfoUser struct {
// 	Id       int       `json:"tech_id_user"`
// 	CreateAt time.Time `json:"create_at"`
// 	DeleteAt time.Time `json:"delete_at"`
// }
