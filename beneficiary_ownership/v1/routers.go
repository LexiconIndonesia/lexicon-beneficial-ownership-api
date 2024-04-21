package beneficiary_ownership_v1

import (
	"lexicon/bo-api/common/utils"
	"net/http"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {

	r := chi.NewMux()
	r.Get("/search", searchHandler)
	r.Get("/detail/{id}", detailHandler)
	return r
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteMessage(w, http.StatusOK, "search")
}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteMessage(w, http.StatusOK, "detail")
}
