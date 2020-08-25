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
	"github.com/disintegration/imaging"
	"github.com/gorilla/mux"
	"image/jpeg"
	"log"
	"net/http"
	"fmt"
	"strconv"
)

func AddImageFormHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	session, _ := sessions.Store.Get(r, "session")
	vars := mux.Vars(r)
	key := vars["albumId"]
	s, err := strconv.Atoi(key)

	if err != nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}

	utils.ExecuteTemplate(w, "add-image.html", struct{
		Alert utils.Alert
		AlbumId int

	}{
		Alert: utils.NewAlert(message, alert),
		AlbumId: s,
	})
}

func AddImageFormPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	//userid := session.Values["USERID"]
	vars := mux.Vars(r)
	key := vars["albumId"]

	s, err := strconv.Atoi(key)

	if err != nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}
	var imageentity models.Image
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("fileupload")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()


	imageentity.ImageName = r.Form.Get("imagename")

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)


	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = imaging.Encode(&buf, img, imaging.JPEG)
	if err != nil {
		log.Println(err)
		return
	}

	image := base64.StdEncoding.EncodeToString(buf.Bytes())
	imageentity.Imagefile = image
	imageentity.AlbumId = s

	db, err:= config.GetConnection()
	defer db.Close()
	if err != nil{
		dto.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	imagerepo := repository.GetImageRespository(db)
	a, err := imagerepo.Save(imageentity)

	if err !=nil{
		fmt.Println("error saving data")
	}else{
		fmt.Println(a)
	}

	images, err:= imagerepo.GetAll(s)

	imagedtos := utils.ToImageDTOs(images)
	utils.ExecuteTemplate(w, "show-images.html", struct{
		Images []dto.ImageDTO
		AlbumId int
	}{
		Images: imagedtos,
		AlbumId: s,
	})
}

func ImagesGETDataHandler (w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	//userid := session.Values["USERID"]

	vars := mux.Vars(r)
	key := vars["albumId"]

	s, err := strconv.Atoi(key)

	if err != nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}


	db, err:= config.GetConnection()
	if err != nil{
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}

	imagerepo := repository.GetImageRespository(db)

	images, err:= imagerepo.GetAll(s)

	imagedtos := utils.ToImageDTOs(images)


	utils.ExecuteTemplate(w, "show-images.html", struct{
		Images []dto.ImageDTO
		AlbumId int
	}{
		Images: imagedtos,
		AlbumId: s,
	})
}

func DeleteImageHandler (w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	//userid := session.Values["USERID"]

	vars := mux.Vars(r)
	key := vars["imageid"]
	s, err := strconv.Atoi(key)

	if err != nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}

	aid := vars["albumId"]
	albumid, err := strconv.Atoi(aid)

	if err != nil {
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}

	db, err:= config.GetConnection()
	if err != nil{
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}

	imagerepo := repository.GetImageRespository(db)

	success, err:= imagerepo.Delete(s)

	if !success{
		session.Values["MESSAGE"] = fmt.Sprintf("%s", err)
		session.Values["ALERT"] = "danger"
		session.Save(r, w)
		http.Redirect(w, r, "/albums", 302)
	}
	images, err:= imagerepo.GetAll(albumid)
	imagedtos := utils.ToImageDTOs(images)

	utils.ExecuteTemplate(w, "show-images.html", struct{
		Images []dto.ImageDTO
		AlbumId int
	}{
		Images: imagedtos,
		AlbumId: albumid,
	})
}
