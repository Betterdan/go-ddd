package api

import (
	"github.com/gorilla/mux"
)

func NewRouter(
	userController *UserController,
	orderController *OrderController,
) *mux.Router {
	router := mux.NewRouter()

	// 用户相关的路由
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", userController.RegisterUser).Methods("POST")
	userRouter.HandleFunc("/{id}", userController.GetUser).Methods("GET")

	// 订单相关的路由
	orderRouter := router.PathPrefix("/orders").Subrouter()
	orderRouter.HandleFunc("", orderController.CreateOrder).Methods("POST")
	orderRouter.HandleFunc("/{id}", orderController.GetOrder).Methods("GET")

	return router
}
