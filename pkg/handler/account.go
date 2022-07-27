package handler

import (
	"net/http"
	"strconv"
	"testbank/bank"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateAccount(c echo.Context) error {
	var a bank.Account
	userId := c.Param("id")
	if err := c.Bind(&a); err != nil {
		logrus.Fatal(err)
		return err
	}
	id, err := h.service.CreateAccount(userId, a)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{

		"id": id,
	})
}
func (h *Handler) GetAllUserAccounts(c echo.Context) error {
	userId := c.Param("id")
	getAllUserAccounts, err := h.service.GetAllUserAccounts(userId)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, "dont have an accounts")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Iduser":         userId,
		"users accounts": getAllUserAccounts,
	})
}

func (h *Handler) GetTransaction(c echo.Context) error {
	userId := c.Param("id")
	getTransactions, err := h.service.GetTransaction(userId)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Iduser":            userId,
		"user transactions": getTransactions,
	})

}

func (h *Handler) AccountTransaction(c echo.Context) error {
	id := c.Param("id")
	senderId, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
	}
	var a bank.AccountTransaction
	if err := c.Bind(&a); err != nil {
		logrus.Panic(err)
		return err
	}
	transaction := h.service.AccountTransaction(senderId, a)
	if transaction != nil {
		logrus.Fatal(transaction)
	}
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "succes",
	})
}
