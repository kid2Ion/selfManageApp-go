package firebaseauth

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"golang.org/x/exp/slog"
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
		slog.Error("failed to decode:\n %s", err.Error())
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
		slog.Error("failed to credential:\n %s", err.Error())
		return nil, err
	}

	opt := option.WithCredentials(credentials)
	conf := &firebase.Config{
		ProjectID: os.Getenv("FB_PROJECT_ID"),
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		slog.Error("failed to init firebase app:\n %s", err.Error())
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		slog.Error("failed to init firebase client:\n %s", err.Error())
		return nil, err
	}

	return &authClient{ctx: ctx, authClient: client}, nil
}

func (t *authClient) VerifyIDToken(c echo.Context) (string, error) {
	authToken := c.Request().Header.Get("Authorization")
	if authToken == "" {
		slog.Error("empty auth token")
		return "", &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "empty auth token",
		}
	}
	idToken := strings.Replace(authToken, "Bearer ", "", 1)
	token, err := t.authClient.VerifyIDToken(t.ctx, idToken)
	if err != nil {
		slog.Error("invalid auth token:\n %s", err.Error())
		return "", &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid auth token",
		}
	}
	return token.UID, nil
}
