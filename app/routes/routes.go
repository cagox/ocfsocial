package routes

import (
	"fmt"
	"github.com/cagox/ocfsocial/app/util/config"
	"net/http"
)

//Routes calles the Routes() functions in all off the packages that require routing.
func Routes() {
	specialRoutes()
	tempRoutes()

}

func specialRoutes() {
	staticDir := config.Config.StaticPath
	//This will route to /static/, and should keep things going during dev.
	config.Config.Router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir(config.Config.StaticPath))))

}

func tempRoutes() {
	config.Config.Router.HandleFunc("/", indexHandler)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
