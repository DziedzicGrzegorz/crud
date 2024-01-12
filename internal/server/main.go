package server

import (
	"crud/internal/routes"
	"crud/pkg/cpu"
	"fmt"
	"os"
	"strconv"
)

func StartApi() {
	server := New()
	api := GroupPath("/api/v1", server.App)
	//here in loop start all routes available in entire app
	//	bookRepo := book.NewRepo(bookCollection)
	//bookService := book.NewService(bookRepo)
	r := cpu.NewRepository()
	s := cpu.NewService(r)

	routes.CpuRouter(api, s)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
