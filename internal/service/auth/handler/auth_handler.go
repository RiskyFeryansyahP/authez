package handler

import (
	"encoding/json"
	"github.com/confus1on/authez/config"
	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

// AuthHandler is handler which embedded usecase auth
type AuthHandler struct {
	AuthUC auth.UsecaseAuth
}

// NewAuthHandler initiate router path
func NewAuthHandler(authUC auth.UsecaseAuth, router *fasthttprouter.Router) {
	authHandler := &AuthHandler{AuthUC: authUC}

	router.POST("/v1/signin", authHandler.Signin)
	router.GET("/auth/google/callback", authHandler.HandleGoogleSignin)
	router.GET("/v1/signin/google", signinGoogle)
}

// Signin http method
func (a *AuthHandler) Signin(ctx *fasthttp.RequestCtx) {
	var input model.InputAuth

	body := ctx.Request.Body()

	ctx.Response.Header.SetContentType("application/json")

	_ = json.Unmarshal(body, &input)

	result, err := a.AuthUC.AuthenticationValidation(input)
	if err != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusBadRequest)
		_ = json.NewEncoder(ctx).Encode(&model.ResponseError{
			Code:    fasthttp.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	_ = json.NewEncoder(ctx).Encode(result)
}

// HandleGoogleSignin handle request google signin
func (a *AuthHandler) HandleGoogleSignin(ctx *fasthttp.RequestCtx) {
	byteCode := ctx.FormValue("code")
	byteState := ctx.FormValue("state")

	code := string(byteCode)
	state := string(byteState)

	log.Println("state", state)

	configMap := config.NewConfigMap()

	ctx.Response.Header.SetContentType("application/json")

	response, err := a.AuthUC.GoogleAuthentication(configMap, state, code)
	if err != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusBadRequest)
		_ = json.NewEncoder(ctx).Encode(&model.ResponseError{
			Code:    fasthttp.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	_ = json.NewEncoder(ctx).Encode(response)
}

func signinGoogle(ctx *fasthttp.RequestCtx) {
	conf := config.NewConfigMap()

	url := conf.GoogleOauth.AuthCodeURL("oauthstate")

	ctx.Redirect(url, fasthttp.StatusTemporaryRedirect)
}

/* func generateStateAuthCookie(ctx *fasthttp.RequestCtx) string {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)

	_, _ = rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	cookie := &fasthttp.Cookie{}
	cookie.SetKey("oauthstate")
	cookie.SetValue(state)
	cookie.SetExpire(expiration)
	cookie.SetHTTPOnly(true)

	ctx.Response.Header.SetCookie(cookie)

	return state
} */
