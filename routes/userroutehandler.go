package routes

import (
	"encoding/json"
	"fmt"
	"albumwebapp/config"
	"albumwebapp/dto"
	"albumwebapp/repository"
	"albumwebapp/utils"
	"github.com/gorilla/mux"
	"go-webapp/sessions"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func UsersGETHandler(w http.ResponseWriter, r *http.Request){
	db, err:= config.GetConnection()

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	userrepository := repository.GetUserRespository(db)
	users, err:= userrepository.GetAll();

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusOK, users)
	}
}

func UserGETHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	s, err := strconv.Atoi(key)

	if err != nil {dto.RespondWithError(w, http.StatusInternalServerError, err.Error())}

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	userrepository := repository.GetUserRespository(db)
	users, err:= userrepository.GetByID(s);

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusOK, users)
	}
}


func CreateNewUser(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)

	if err !=nil{
		log.Println("unable to read body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	var userrequest dto.UserRequest
	err = json.Unmarshal(reqBody, &userrequest)

	if err!=nil{
		log.Println("unable to unmarshal body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	userrepository := repository.GetUserRespository(db)
	user := utils.ToUserEntity(userrequest)

	users, err:= userrepository.Save(user);

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusCreated, users)
	}

}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var userrequest dto.UserRequest
	userrequest.Fname = r.PostForm.Get("firstname")
	userrequest.Lname = r.PostForm.Get("lastname")
	userrequest.Email = r.PostForm.Get("email")
	userrequest.Contact, _ = strconv.ParseInt(r.PostForm.Get("contact"), 10, 64)
	userrequest.Address = r.PostForm.Get("address")
	userrequest.Password = r.PostForm.Get("password")

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		message := fmt.Sprintf("%s", err)
		sessions.Message(message, "danger", r, w)
		http.Redirect(w, r, "/register", 302)
		return	}

	userrepository := repository.GetUserRespository(db)
	user := utils.ToUserEntity(userrequest)

	b, err := utils.Hash(userrequest.Password)

	if err != nil{
		message := fmt.Sprintf("%s", err)
		sessions.Message(message, "danger", r, w)
		http.Redirect(w, r, "/register", 302)
		return
	}

	user.Password = string(b)
	_, err = userrepository.Save(user)

	if err != nil{
		message := fmt.Sprintf("%s", err)
		sessions.Message(message, "danger", r, w)
		http.Redirect(w, r, "/register", 302)
		return
	}
	message := fmt.Sprintf("%s", "User Successfully Created")
	sessions.Message(message, "success", r, w)
	http.Redirect(w, r, "/login", 302)
}



