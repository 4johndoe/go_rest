package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
}