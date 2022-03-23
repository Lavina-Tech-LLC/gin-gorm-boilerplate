package app

import (
	"context"
	"flag"
	"fmt"
	"gin-gorm-boilerplate/internal/auth"
	"gin-gorm-boilerplate/internal/dbCon"
	"gin-gorm-boilerplate/internal/migrate"
	"gin-gorm-boilerplate/internal/network"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var router *gin.Engine

func Main() {
	// setting run mode
	runMode := flag.String("m", "", "[optional] runMode, by default it is dev")
	action := flag.String("act", "app", "[optional] action, [updatePolicies, migrate]")
	configPath := flag.String("conf", "", "[optional] path to configs folder")
	flag.Parse()
	dbCon.RunMode = *runMode
	dbCon.ConfigPath = *configPath

	// Starting viper
	dbCon.StartViper()

	// connecting to db
	dbCon.Connect()

	switch *action {
	case "app":
		startApp()
	case "migrate":
		migrate.Main()
	case "updatePolicies":
		dbCon.PopulateCasbinDefaults()
	}

}

func startApp(){

	// Set the router as the default one provided by Gin
	router = gin.Default()
	router.Use(auth.AuthenticateUser)
	router.Use(auth.AuthorizeUser)
	router.Use(ErrorHandler)

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	//router.LoadHTMLGlob(gin-gorm-boilerplate/internal/templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", Default)
	network.NetworkRoutes(router)

	// Start serving the application
	//router.Run()
 
	srv := &http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: router,
	}
	llog.Info(fmt.Sprintf("Server is listening to %s", srv.Addr))

	go func() {
		var err error
		if viper.GetString("server.cert")==""||viper.GetString("server.key")=="" {
			err = srv.ListenAndServe()
		} else{
			err = srv.ListenAndServeTLS(viper.GetString("server.cert"),viper.GetString("server.key"))
		}
		// service connections
		if err != nil && err != http.ErrServerClosed {
			llog.Error(fmt.Sprintf("listen: %s\n", err))
		} else {
			llog.Info(fmt.Sprintf("Server is listening to %s", srv.Addr))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	llog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		llog.Error(fmt.Sprintf("Server Shutdown: %s", err))
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		llog.Info("timeout of 5 seconds.")
	}
	llog.Info("Server exiting")
}
