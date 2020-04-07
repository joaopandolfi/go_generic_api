package main

import (
	"./routes"
	"net/http"
	"./config"
	"github.com/gorilla/mux"
	"github.com/joaopandolfi/blackwhale/configurations"

	"github.com/joaopandolfi/blackwhale/remotes/mysql"
	"github.com/joaopandolfi/blackwhale/utils"
)

func configInit() {
	configurations.LoadConfig(config.Load())
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
	utils.Info("MI server listenning on", configurations.Configuration.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         configurations.Configuration.Port,
		WriteTimeout: configurations.Configuration.Timeout.Write,
		ReadTimeout:  configurations.Configuration.Timeout.Read,
	}

	err := srv.ListenAndServe() //srv.ListenAndServeTLS()
	//"github.com/fvbock/endless"
	///err := endless.ListenAndServeTLS("localhost:4242", "cert.pem", "key.pem", r)

	if err != nil {
		utils.CriticalError("Fatal server error", err.Error())
	}
}
