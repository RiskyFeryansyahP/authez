package main

import (
	"github.com/confus1on/authez/internal/service/auth/handler"
	"github.com/confus1on/authez/internal/service/auth/repository"
	"github.com/confus1on/authez/internal/service/auth/usecase"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func main()  {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	port = ":" + port

	router := fasthttprouter.New()

	authRepo := repository.NewAuthRepository()
	authUsecase := usecase.NewAuthUsecase(authRepo)
	handler.NewAuthHandler(authUsecase, router)

	log.Printf("Server running at 127.0.0.1%s", port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}