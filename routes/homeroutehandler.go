package routes

import (
	"albumwebapp/utils"
	"net/http"
)

func HomeGetHandler(w http.ResponseWriter, r *http.Request){
	utils.ExecuteTemplate(w, "home.html", nil)
}