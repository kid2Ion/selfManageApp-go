package di

import (
	firebaseauth "github.com/kid2Ion/selfManageApp-go/adapter/firebase"
	"github.com/kid2Ion/selfManageApp-go/domain/repository"
	"github.com/kid2Ion/selfManageApp-go/infra"
	"github.com/kid2Ion/selfManageApp-go/server"
	"github.com/kid2Ion/selfManageApp-go/usecase"
)

// DBのインジェクションは各domainで使い回す想定
func injectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	return *sqlHandler
}

// hello
func injectHelloRepository() repository.HelloRepository {
	sqlHandler := injectDB()
	return infra.NewHelloRepository(sqlHandler)
}

func injectHelloUsecase() usecase.HelloUsecase {
	repo := injectHelloRepository()
	return usecase.NewHelloUsecase(repo)
}

func InjectHelloHandler(authClient firebaseauth.AuthClient) server.HelloHandler {
	usecase := injectHelloUsecase()
	return server.NewHelloHandler(usecase, authClient)
}

// user

func injectUserRepository() repository.UserRepository {
	sqlHandler := injectDB()
	return infra.NewUserRepository(sqlHandler)
}

func injectUserUsecase() usecase.UserUsecase {
	repo := injectUserRepository()
	return usecase.NewUserUsecase(repo)
}

func InjectUserHandler(authClient firebaseauth.AuthClient) server.UserHandler {
	usecase := injectUserUsecase()
	return server.NewUserHandler(usecase, authClient)
}
