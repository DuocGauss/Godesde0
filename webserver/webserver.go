package webserver

import "net/http"

func MiWebServer() {
	http.HandleFunc("/", home)
	http.HandleFunc("/segunda", segundapag)
	http.ListenAndServe(":3000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")

}

func segundapag(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/segunda.html")

}
