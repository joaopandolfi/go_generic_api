package config

import (
	"fmt"
	"time"

	"github.com/flosch/pongo2"
	"github.com/gorilla/sessions"
	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/unrolled/secure"
)

type Configs struct {
	Token string
}

var Config Configs

func Load() configurations.Configurations {

	Config.Token = "$238#!%s@233**#sd*"

	return configurations.Configurations{
		Name: "GENERIC API - GO",

		MysqlUrl: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			"root",      // User
			"usbw",      // password
			"localhost", // host
			"3307",      // port
			"miia"),     // Database

		MongoUrl: "mongodb://127.0.0.1:27017",
		MongoDb:  "gen-api",

		Port:        ":8990",
		CORS:        "*",

		Timeout: configurations.Timeout{
			Write: 60 * time.Second,
			Read:  60 * time.Second,
		},


		StaticDir:      "./views/public/",
		StaticPagesDir: "./views/pages/",
		UploadPath:     "./views/upload/",

		MaxSizeMbUpload: 10 << 55, // min << max

		BCryptSecret: "#1$eY)&4430",


		// Session
		Session: configurations.SessionConfiguration{
			Name:  "A2%!#23dad#32$",
			Store: sessions.NewCookieStore([]byte("_-)(AS(&H:(SD)_:)H@ˆ@@#$##$*{{{$$}}}(U$$#@D)&#Y!)P(@Mkdksdsb321k5*443@@##@$!")),
			Options: &sessions.Options{
				Path:     "/",
				MaxAge:   3600 * 2, //86400 * 7,
				HttpOnly: true,
			},
		},

		Security: configurations.Opsec{
			Options: secure.Options{
				BrowserXssFilter:   true,
				ContentTypeNosniff: false, // Da pau nos js
				SSLHost:            "locahost:443",
				SSLRedirect:        false,
			},
		},

		Templates: make(map[string]*pongo2.Template),

		// Slack
		SlackToken:   "",
		SlackWebHook: []string{"", ""},
		SlackChannel: "",

	}
}
