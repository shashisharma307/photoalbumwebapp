package routes

import (
	"albumwebapp/config"
	"albumwebapp/dto"
	"albumwebapp/models"
	"albumwebapp/repository"
	"albumwebapp/sessions"
	"albumwebapp/utils"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	//"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateNewAlbum(w http.ResponseWriter, r *http.Request){
	reqBody, err := ioutil.ReadAll(r.Body)

	if err !=nil{
		log.Println("unable to read body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	var albumrequest dto.AlbumRequest
	err = json.Unmarshal(reqBody, &albumrequest)

	if err!=nil{
		log.Println("unable to unmarshal body")
		dto.RespondWithError(w,http.StatusInternalServerError, err.Error())
	}

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	albumrepo := repository.GetAlbumRespository(db)
	user := utils.ToAlbumEntity(albumrequest)

	users, err:= albumrepo.Save(user);

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusCreated, users)
	}
}

func AddAlbumFormHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "add-album.html", struct{
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

//func AddImageFormPostHandler(w http.ResponseWriter, r *http.Request) {
//	session, _ := sessions.Store.Get(r, "session")
//	userid := session.Values["USERID"]
//
//	var Album models.Album
//	r.ParseMultipartForm(10 << 20)
//
//	file, handler, err := r.FormFile("fileupload")
//	if err != nil {
//		fmt.Println("Error Retrieving the File")
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//
//
//	Album.AlbumName = r.Form.Get("albumname")
//	Album.Description = r.PostForm.Get("description")
//	Album.UserId  = userid.(int)
//
//	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
//	fmt.Printf("File Size: %+v\n", handler.Size)
//	fmt.Printf("MIME Header: %+v\n", handler.Header)
//
//
//	img, err := jpeg.Decode(file)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var buf bytes.Buffer
//	err = imaging.Encode(&buf, img, imaging.JPEG)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	image := base64.StdEncoding.EncodeToString(buf.Bytes())
//	Album.AlbumThumbnail = image
//
//	db, err:= config.GetConnection()
//	defer db.Close()
//	if err != nil{
//		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
//	}
//
//	albumrepo := repository.GetAlbumRespository(db)
//	a, err := albumrepo.Save(Album)
//
//	if err !=nil{
//		fmt.Println("error saving data")
//	}else{
//		fmt.Println(a)
//	}
//
//	albums, err:= albumrepo.GetAll(userid)
//
//	albumdtos := utils.ToAlbumDTOs(albums)
//	utils.ExecuteTemplate(w, "show-albums.html", struct{
//		Albums []dto.AlbumDTO
//	}{
//		Albums: albumdtos,
//	})
//}

func AddAlbumFormPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	userid := session.Values["USERID"]

	var Album models.Album
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("fileupload")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	Album.AlbumName = r.Form.Get("albumname")
	Album.Description = r.PostForm.Get("description")
	Album.UserId  = userid.(int)

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fmt.Printf("filename: %+v\n", handler.Header.Get("filename"))

	img, err := png.Decode(file)
	if err != nil {
			log.Fatal(err)
	}

	var buf bytes.Buffer
	err = imaging.Encode(&buf, img, imaging.PNG)
	if err != nil {
			log.Println(err)
			return
	}

	image := base64.StdEncoding.EncodeToString(buf.Bytes())
	Album.AlbumThumbnail = image

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	albumrepo := repository.GetAlbumRespository(db)
	a, err := albumrepo.Save(Album)

	if err !=nil{
		fmt.Println("error saving data")
	}else{
		fmt.Println(a)
	}

	albums, err:= albumrepo.GetAll(userid)

	albumdtos := utils.ToAlbumDTOs(albums)
	utils.ExecuteTemplate(w, "show-albums.html", struct{
		Albums []dto.AlbumDTO
	}{
		Albums: albumdtos,
	})
}


func AlbumsGETDataHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	userid := session.Values["USERID"]
	fmt.Println(userid)
	if userid == 0{
		http.Redirect(w, r, "/login", 302)
		return
	}


	db, err:= config.GetConnection()
	if err != nil{
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	}

	albumrepo := repository.GetAlbumRespository(db)

	albums, err:= albumrepo.GetAll(userid)

	albumdtos := utils.ToAlbumDTOs(albums)

	//
	//if err != nil{
	//	session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
	//	session.Values["ALERT"] = "danger"
	//	session.Save(r, w)
	//	http.Redirect(w, r, "/", 302)
	//}

	utils.ExecuteTemplate(w, "show-albums.html", struct{
		Albums []dto.AlbumDTO
	}{
		Albums: albumdtos,
	})

}

func AlbumsGETHandler(w http.ResponseWriter, r *http.Request){
	db, err:= config.GetConnection()

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	albumrepo := repository.GetAlbumRespository(db)
	albums, err:= albumrepo.GetAll(1);

	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}else {
		dto.RespondWithJSON(w, http.StatusOK, albums)
	}
}
