package hello

import (
	"net/http"
)

//Index throws up the SPA
func Index(w http.ResponseWriter, r *http.Request) {
	// http.FileServer(http.Dir("./ng2-app/")).ServeHTTP(w, r)
	http.ServeFile(w, r, "ng2-app/dist/index.html")

}
