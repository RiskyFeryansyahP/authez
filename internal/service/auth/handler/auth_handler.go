package handler

import (
	"encoding/json"
	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type AuthHandler struct {
	AuthUC auth.UsecaseAuth
}

func NewAuthHandler(authUC auth.UsecaseAuth, router *fasthttprouter.Router) {
	authHandler := &AuthHandler{AuthUC: authUC}

	router.POST("/v1/signin/{type}", authHandler.Signin)
}

func (a *AuthHandler) Signin(ctx *fasthttp.RequestCtx)  {
	var input model.InputAuth

	typeConnection := ctx.UserValue("type").(string)
	
	body := ctx.Request.Body()
	
	ctx.Response.Header.SetContentType("application/json")

	err := json.Unmarshal(body, &input)
	if err != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusBadRequest)
		json.NewEncoder(ctx).Encode(&model.ResponseError{
			Code:    fasthttp.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	result, err := a.AuthUC.AuthenticationValidation(input, typeConnection)
	if err != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusBadRequest)
		json.NewEncoder(ctx).Encode(&model.ResponseError{
			Code:    fasthttp.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}