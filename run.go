package main

import "net/http"
type handler int
func (h handler)ServeHTTP(w http.ResponseWriter,q *http.Request){

}
func main() {
	var x handler
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {})
	http.Handle("/route",x)
}
