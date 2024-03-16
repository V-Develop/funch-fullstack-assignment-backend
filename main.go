package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/V-Develop/funch-fullstack-assignment-backend/app"
	"github.com/V-Develop/funch-fullstack-assignment-backend/handler"
	log "github.com/V-Develop/funch-fullstack-assignment-backend/util"
)

func main() {
	loger := log.NewLog("0", "LOG")
	config := initAppConfig(loger)
	authRoute := handler.ServiceRoute{}.InitAuthController(config)
	authService := &http.Server{
		Addr:    fmt.Sprint(":", 8080),
		Handler: authRoute,
	}

	userRoute := handler.ServiceRoute{}.InitUserController(config)
	userService := &http.Server{
		Addr:    fmt.Sprint(":", 8081),
		Handler: userRoute,
	}

	go func() {
		if err := authService.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loger.LogError("auth service listen: "+err.Error(), "0")
			panic(err.Error())
		} else if err != nil {
			loger.LogError("auth service listen error: "+err.Error(), "0")
			panic(err.Error())
		}
	}()

	go func() {
		if err := userService.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loger.LogError("user service listen: "+err.Error(), "0")
			panic(err.Error())
		} else if err != nil {
			loger.LogError("user service listen error: "+err.Error(), "0")
			panic(err.Error())
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

}

func initAppConfig(loger log.Loger) *app.Config {
	database := app.InitDatabaseConfig(loger)
	errorMessage := app.InitErrorMessage(loger)
	config := &app.Config{
		DataBase:     database,
		ErrorMessage: errorMessage,
	}
	return config
}
