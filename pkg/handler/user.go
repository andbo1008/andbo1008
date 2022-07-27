package handler

import (
	"net/http"
	"testbank/bank"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateUser(c echo.Context) error {
	var u bank.User
	if err := c.Bind(&u); err != nil {
		logrus.Fatal(err)
		return err
	}
	id, err := h.service.CreateUser(u)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	u, err := h.service.GetUser(id)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":       u.Id,
		"name":     u.Name,
		"lastname": u.LastName,
		"email":    u.Email,
		"account":  u.Account,
	})
}
