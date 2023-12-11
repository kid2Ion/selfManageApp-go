package firebaseauth

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type (
	authClient struct {
		ctx        context.Context
		authClient *auth.Client
	}
	AuthClient interface {
		VerifyIDToken(ctx echo.Context) (string, error)
	}
)

func NewClient() (AuthClient, error) {
	ctx := context.Background()

	// firebaseクレデンシャル環境変数
	encodedKey := os.Getenv("FB_PRIVATE_KEY_ENCODED")
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		fmt.Println("decode失敗")
		return nil, err
	}
	credentials, err := google.CredentialsFromJSON(ctx,
		decodedBytes,
		"https://www.googleapis.com/auth/firebase",
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/identitytoolkit",
		"https://www.googleapis.com/auth/firebase.messaging",
		"https://www.googleapis.com/auth/firebase.database",
	)
	if err != nil {
		fmt.Println("credential 失敗")
		return nil, err
	}

	opt := option.WithCredentials(credentials)
	conf := &firebase.Config{
		ProjectID: os.Getenv("FB_PROJECT_ID"),
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		fmt.Println("firebase 初期化失敗:", err)
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		fmt.Println("auth client 初期化失敗:", err)
		return nil, err
	}

	return &authClient{ctx: ctx, authClient: client}, nil
}

func (t *authClient) VerifyIDToken(c echo.Context) (string, error) {
	authToken := c.Request().Header.Get("Authorization")
	if authToken == "" {
		fmt.Println("auth token が無い")
		return "", &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "empty auth token",
		}
	}
	// 接続確認用　user情報取得
	// a, err := t.authClient.GetUser(t.ctx, "hoge")
	// if err != nil {
	// 	fmt.Println("koko?")
	// 	fmt.Println(err)
	// 	return "", err
	// }
	// fmt.Println(a.Email)
	idToken := strings.Replace(authToken, "Bearer ", "", 1)
	token, err := t.authClient.VerifyIDToken(t.ctx, idToken)
	if err != nil {
		fmt.Println("auth token 無効")
		return "", &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid auth token",
		}
	}
	return token.UID, nil
}
