package beneficiary_ownership_v1

import (
	"errors"
	models "lexicon/bo-api/beneficiary_ownership/v1/models"
	bo_v1_services "lexicon/bo-api/beneficiary_ownership/v1/services"
	"lexicon/bo-api/common/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {

	r := chi.NewMux()
	r.Get("/search", searchHandler)
	r.Get("/detail/{id}", detailHandler)
	return r
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

	qp := r.URL.Query()
	query := qp.Get("query")
	subjectType := qp.Get("subject_type")
	year := qp.Get("year")
	caseType := qp.Get("type")
	nation := qp.Get("nation")
	page := qp.Get("page")

	years := []string{}

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("page must be a number"))
		return
	}

	if year != "" {
		yearsSplit := strings.Split(year, "-")
		if len(yearsSplit) != 2 {
			utils.WriteError(w, http.StatusBadRequest, errors.New("year must be in the format of year-year"))

			return
		}
		// genereate year between
		yearFrom, err := strconv.Atoi(yearsSplit[0])
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, errors.New("year must be in the format of year-year"))
			return
		}
		yearTo, err := strconv.Atoi(yearsSplit[1])

		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, errors.New("year must be in the format of year-year"))

			return
		}

		for i := yearFrom; i <= yearTo; i++ {
			years = append(years, strconv.Itoa(i))
		}

	}

	req := models.SearchRequest{
		Query:       query,
		SubjectType: subjectType,
		Years:       years,
		Type:        caseType,
		Nation:      nation,
		Page:        int64(pageInt),
	}

	response, err := bo_v1_services.Search(r.Context(), req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if response.Data == nil {
		utils.WriteError(w, http.StatusNotFound, errors.New("data not found"))
		return
	}
	utils.WriteResponse(w, response, http.StatusOK)
}

func detailHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	response, err := bo_v1_services.GetDetail(r.Context(), id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, errors.New("data not found"))
		return
	}

	utils.WriteData(w, response, http.StatusOK)
}
