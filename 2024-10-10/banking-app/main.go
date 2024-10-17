package main

import (
	"bankingApp/helper"
	"bankingApp/routers"
	userController "bankingApp/user/controller"
	user "bankingApp/user/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Shape interface {
	abc()
}
type Circle struct {
}

func (c *Circle) abc() {

}

var logger *log.Logger

func main() {

	logger = log.New(log.Default().Writer(), "=============================================================\nLOG:", 0)
	user.NewAdminUserWithIdPass("Test", "Test", "Test", "TestTest")
	mainRouter := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	mainRouter.HandleFunc("/login", userController.LoginUser).Methods(http.MethodPost)
	routers.RegisterCusomterRouter(mainRouter)
	routers.RegisterBankRouter(mainRouter)

	http.ListenAndServe(":4000", ErrorHandlerMiddleware(LoggerMiddleware(mainRouter)))

}

func LoggerMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger.Println(r.RequestURI)

			next.ServeHTTP(w, r)
		})
}

func ErrorHandlerMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func(w http.ResponseWriter) {
				err := recover()
				if err != nil {
					fmt.Println("Internal Server Error:", err)
					helper.SendErrorWithCustomMessage(w, "Internal Server Error:")
				}
			}(w)

			next.ServeHTTP(w, r)
		})
}
