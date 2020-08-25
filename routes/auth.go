package routes

import (
	"albumwebapp/config"
	"albumwebapp/repository"
	"albumwebapp/sessions"
	"albumwebapp/utils"
	"fmt"
	"net/http"
	"strings"
)



func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	_, isAuth := sessions.IsLogged(r)
	if isAuth {
		http.Redirect(w, r, "/home", 302)
		return
	}
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "login.html", struct{
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")

	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	err := utils.ValidateFields(strings.ToLower(email), password)
	if err != nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
	repo := repository.GetUserRespository(db)
	user, err := repo.GetEmailById(email)

	if err !=nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
	}
	session.Values["USERID"] = user.UserId
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}


func RegisterGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "register.html", struct{
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}


func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	delete(session.Values, "USERID")
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}


