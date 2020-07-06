package main

import (
	"github.com/confus1on/authez/internal/service/auth/handler"
	"github.com/confus1on/authez/internal/service/auth/repository"
	"github.com/confus1on/authez/internal/service/auth/usecase"
	"github.com/confus1on/authez/util"
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
	localAddress := util.GetOutboundIP()

	router := fasthttprouter.New()

	authRepo := repository.NewAuthRepository()
	authUsecase := usecase.NewAuthUsecase(authRepo)
	handler.NewAuthHandler(authUsecase, router)

	log.Printf("Server running at http://%v%s or http://127.0.0.1%s", localAddress, port, port)
	log.Fatal(fasthttp.ListenAndServe(port, router.Handler))
}