package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/", HomeGetHandler).Methods("GET")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/register", RegisterGetHandler).Methods("GET")
	r.HandleFunc("/register", RegisterPostHandler).Methods("POST")
	//r.HandleFunc("/albums", middleware.AuthRequired(AlbumsGETDataHandler)).Methods("GET")
	r.HandleFunc("/albums", AlbumsGETDataHandler).Methods("GET")
	r.HandleFunc("/addalbum", AddAlbumFormHandler).Methods("GET")
	r.HandleFunc("/addalbum", AddAlbumFormPostHandler).Methods("POST")
	r.HandleFunc("/images/{albumId}", ImagesGETDataHandler).Methods("GET")
	r.HandleFunc("/images/{albumId}/{imageid}", DeleteImageHandler).Methods("GET")
	r.HandleFunc("/addimage/{albumId}", AddImageFormHandler).Methods("GET")
	r.HandleFunc("/addimage/{albumId}", AddImageFormPostHandler).Methods("POST")
	r.HandleFunc("/logout", logoutGetHandler).Methods("GET")




	r.HandleFunc("/user", UsersGETHandler).Methods("GET")
	r.HandleFunc("/user/{id}", UserGETHandler).Methods("GET")
	r.HandleFunc("/user", CreateNewUser).Methods("POST")
	r.HandleFunc("/album", CreateNewAlbum).Methods("POST")
	r.HandleFunc("/album", AlbumsGETHandler).Methods("GET")



	fileserver := http.FileServer(http.Dir("src" + string(os.PathSeparator) +
		"albumwebapp" + string(os.PathSeparator) + "assets"))
	r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", fileserver))
	return r
}
