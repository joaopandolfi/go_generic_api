package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/joaopandolfi/go_generic_api/config"
	"github.com/joaopandolfi/go_generic_api/routes"

	"github.com/joaopandolfi/blackwhale/remotes/mongo"
	"github.com/joaopandolfi/blackwhale/remotes/mysql"
	"github.com/joaopandolfi/blackwhale/utils"
)

func configInit() {
	configurations.LoadConfig(config.Load(os.Args[1:]))
	mysql.Init()
	// Precompile html pages
	routes.Precompile()
}

func resilient() {
	utils.Info("[SERVER] - Shutdown")

	if err := recover(); err != nil {
		utils.CriticalError("[SERVER] - Returning from the dark", err)
		main()
	}
}

func gracefullShutdown() {
	mysql.Close()
	mongo.Close()
}

func welcome() {
	// https://patorjk.com/software/taag/#p=display&f=Slant&t=ms%20-%20calendar
	fmt.Println("#####################")
	fmt.Println("#  go_generic_api!  #")
	fmt.Println("#####################")
	fmt.Println("")
}

func main() {
	defer resilient()

	welcome()

	//Init
	configInit()

	// Initialize Mux Router
	r := mux.NewRouter()

	// Security
	routes.Handlers(r)
	r.Use(mux.CORSMethodMiddleware(r))

	// Routes consist of a path and a handler function.
	routes.Register(r)

	//CRON Register
	//cron.Register(os.Args[1:])

	// Bind to a port and pass our router in
	utils.Info("Server listenning on", configurations.Configuration.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         configurations.Configuration.Port,
		WriteTimeout: configurations.Configuration.Timeout.Write,
		ReadTimeout:  configurations.Configuration.Timeout.Read,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		var err error
		if config.Config.Debug {
			err = srv.ListenAndServe()
		} else {
			err = srv.ListenAndServeTLS(config.Config.TLSCert, config.Config.TLSKey)
		}
		if err != nil && err != http.ErrServerClosed {
			utils.CriticalError("Fatal server error", err.Error())
		}
	}()

	<-done
	utils.Info("[SERVER] Gracefully shutdown")
	gracefullShutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		utils.CriticalError("Server Shutdown Failed", err.Error())
	}
	cancel()
}
