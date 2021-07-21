package main

import (
	"net/http"

	"github.com/einnar82/net-http-api-sample/app/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", controllers.GetPosts)
	http.ListenAndServe(":5000", router)
}
