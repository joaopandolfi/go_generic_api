package main

import (
	"net/http"
	"os"

	"./config"
	"./routes"
	"github.com/gorilla/mux"
	"github.com/joaopandolfi/blackwhale/configurations"

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

func main() {
	defer resilient()

	//Init
	configInit()

	// Initialize Mux Router
	r := mux.NewRouter()

	// Security
	routes.Handlers(r)

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

	var err error
	if config.Config.Debug {
		err = srv.ListenAndServe()
	} else {
		err = srv.ListenAndServeTLS(config.Config.TLSCert, config.Config.TLSKey)
	}

	if err != nil {
		utils.CriticalError("Fatal server error", err.Error())
	}
}
