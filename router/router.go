package router

import (
	"net/http"
	"testbank/pkg/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router struct {
	H handler.Handler
}

func NewRouter(H handler.Handler) *Router {
	return &Router{H: H}
}

func (r *Router) Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": "OK",
			"page":   "MAIN",
		})
	})

	user := e.Group("user")
	user.POST("", r.H.CreateUser)
	user.GET("/:id", r.H.GetUser)

	account := user.Group("/:id/account")
	account.POST("", r.H.CreateAccount)
	account.GET("", r.H.GetAllUserAccounts)
	account.POST("/transaction", r.H.AccountTransaction)
	account.GET("/transaction", r.H.GetTransaction)

	e.Logger.Fatal(e.Start(":1313"))
}
