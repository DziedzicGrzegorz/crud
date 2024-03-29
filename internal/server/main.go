package server

import (
	"context"
	"crud/internal/routes"
	"crud/pkg/cpu"
	"fmt"
	"os"
	"strconv"
)

func StartApi() {
	dbHandler := cpu.NewDbHandler()
	err := dbHandler.CreateConnection(context.Background())
	if err != nil {
		panic(fmt.Sprintf("cannot connect to database: %s", err))
	}
	dbHandler.CreateDatabase(context.Background())
	defer dbHandler.CloseConnection(context.Background())
	server := New()
	api := GroupPath("/api/v1", server.App)
	//here in loop start all routes available in entire app
	//	bookRepo := book.NewRepo(bookCollection)
	//bookService := book.NewService(bookRepo)
	r := cpu.NewRepository()
	s := cpu.NewService(r)

	routes.CpuRouter(api, s)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err = server.Listen(fmt.Sprintf(":%d", port))

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
