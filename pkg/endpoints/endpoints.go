package endpoints

import (
	"butify/pkg/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type endpoint struct {
	service *services.Service
}

func New(service *services.Service) http.Handler {
	r := mux.NewRouter()

	e := endpoint{service: service}

	r.HandleFunc("/{small:.*}", e.Get)

	return r
}

func (e endpoint) Get(w http.ResponseWriter, r *http.Request) {
	small := mux.Vars(r)["small"]
	url, err := e.service.GetBySmallUrl(r.Context(), small)
	if err != nil {
		log.Println("get", err)
		http.Error(w, "bad", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url.Origin, http.StatusPermanentRedirect)
}
