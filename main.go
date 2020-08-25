package main

import (
	"albumwebapp/routes"
	"albumwebapp/sessions"
	"albumwebapp/utils"
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = ":8080"
func main()  {
	sessions.SessionOptions("localhost", "/", 3600, true)

	fmt.Println(os.Getwd())
	fmt.Println("Listening Port %s\n", PORT)

	viewpath := "src" + string(os.PathSeparator) + "albumwebapp" + string(os.PathSeparator) + "views" + string(os.PathSeparator) + "*.html"
	utils.LoadTemplates(viewpath)
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))



}

//func init(){
//	if ok:= config.InitDB(); !ok{
//		fmt.Println("error creating connection")
//	}
//
//	fmt.Println("connection established")
//
//	if ok:= config.TestConnection(); ok{
//		fmt.Println(ok)
//	}else{
//		fmt.Println(ok)
//	}
//}
