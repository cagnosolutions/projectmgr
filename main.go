package projectmgr

import (
	"fmt"
	"net/http"

	"github.com/cagnosolutions/web"
)

var mux *web.Mux = web.NewMux()

func init() {
	mux.AddRoutes(idx)
}

func main() {
	http.ListenAndServe(":8080", mux)
}

var idx = web.Route{"GET", "/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "this is the index page")
	return
}}
