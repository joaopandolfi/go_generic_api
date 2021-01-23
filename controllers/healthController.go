package controllers

import (
	"net/http"

	"github.com/joaopandolfi/blackwhale/configurations"
	"github.com/joaopandolfi/blackwhale/handlers"
	"github.com/joaopandolfi/blackwhale/remotes/mongo"
	"github.com/joaopandolfi/blackwhale/utils"
	"github.com/joaopandolfi/go_generic_api/models"
)

// --- Health ---

// BaseController -
type HealthController struct{}

// Health route
func (cc HealthController) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	handlers.Response(w, true, http.StatusOK)
}

// Config database
func (cc HealthController) Config(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	hash := handlers.GetHeader(r, "hash")
	utils.Debug("HEADER", hash, configurations.Configuration.ResetHash)
	if hash != configurations.Configuration.ResetHash {
		handlers.Response(w, "Not cookies for you", http.StatusForbidden)
		return
	}
	config()
	handlers.Response(w, true, http.StatusOK)
}

// ResetDatabase route
func (cc HealthController) ResetDatabase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	hash := handlers.GetHeader(r, "hash")
	if hash != configurations.Configuration.ResetHash {
		handlers.Response(w, "Invalid Hash", http.StatusForbidden)
		return
	}

	var errors []error
	_, err := mongo.GetSession().GetCollection("traffic").RemoveAll(nil)
	if err != nil {
		errors = append(errors, err)
	}

	//Config
	//config()

	if len(errors) > 0 {
		handlers.Response(w, errors, http.StatusInternalServerError)
	} else {
		utils.Info("[ResetService] - Traffic-files database RESETED")
		handlers.Response(w, "Reseted", http.StatusOK)
	}

}

func config() error {
	userService := NewUserService()
	_, err := userService.NewUser(models.User{
		People: models.People{
			Name: "Gandolfo malandrao",
			CPF:  "00000",
		},
		Email:     "",
		Username:  "gandolfo",
		Picture:   "",
		Password:  "e5f895d875f3393654fc9fc5b5eb7a6e", // #G4nd0lf0!hhlm
		Instution: 0,
		Level:     99,
	})
	return err
}
