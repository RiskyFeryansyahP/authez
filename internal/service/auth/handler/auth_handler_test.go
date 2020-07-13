package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/confus1on/authez/internal/model"
	"github.com/confus1on/authez/internal/service/auth/mock"
	fasthttprouter "github.com/fasthttp/router"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestAuthHandler_Signin(t *testing.T) {
	controller := gomock.NewController(t)

	route := fasthttprouter.New()

	authUC := mock.NewMockUsecaseAuth(controller)
	NewAuthHandler(authUC, route)

	t.Run("test signin handler", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "risky",
			Password:  "risky123",
			TableName: "users",
		}

		resultUsecase := map[string]interface{}{
			"username": "risky",
			"password": "risky123",
			"fullname": "risky feryansyah",
		}

		var responseBody map[string]interface{}

		authUC.EXPECT().AuthenticationValidation(input).Return(resultUsecase, nil).Times(1)

		authHandler := AuthHandler{AuthUC: authUC}

		router := fasthttprouter.New()
		router.POST("/v1/signin", authHandler.Signin)

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/v1/signin")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &responseBody)
		require.NoError(t, err)

		require.Equal(t, "risky feryansyah", responseBody["fullname"])
	})

	t.Run("Failed unmarshall body", func(t *testing.T) {
		input := model.InputAuth{
			Username:  "risky",
			Password:  "fail",
			TableName: "users",
		}

		responseError := errors.New("invalid username or password")

		var responseBody model.ResponseError

		authUC.EXPECT().AuthenticationValidation(input).Return(nil, responseError).Times(1)

		authHandler := AuthHandler{AuthUC: authUC}

		router := fasthttprouter.New()
		router.POST("/v1/signin", authHandler.Signin)

		body, err := json.Marshal(input)
		require.NoError(t, err)

		req := fasthttp.AcquireRequest()
		req.Header.SetMethod("POST")
		req.Header.SetContentType("application/json")
		req.Header.SetHost("localhost")
		req.SetRequestURI("/v1/signin")
		req.SetBody(body)

		resp := fasthttp.AcquireResponse()

		err = serve(router.Handler, req, resp)
		require.NoError(t, err)

		err = json.Unmarshal(resp.Body(), &responseBody)
		require.NoError(t, err)

		require.Equal(t, 400, responseBody.Code)
		require.Equal(t, "invalid username or password", responseBody.Message)
	})
}

// serve serves http request using provided fasthttp handler
func serve(handler fasthttp.RequestHandler, req *fasthttp.Request, res *fasthttp.Response) error {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	go func() {
		err := fasthttp.Serve(ln, handler)
		if err != nil {
			panic(fmt.Errorf("failed to serve: %v", err))
		}
	}()

	client := fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	return client.Do(req, res)
}
