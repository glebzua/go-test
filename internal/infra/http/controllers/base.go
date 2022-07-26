package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/test_server/internal/domain"
	"log"
	"net/http"
	"strconv"
)

func success(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Print(err)
	}

}

func created(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if body != nil {
		err := json.NewEncoder(w).Encode(map[string]interface{}{"created": body})

		if err != nil {
			log.Print(err)
		}
	}
}

func badRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if err != nil {
		log.Print(err)
	}
}

func internalServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if err != nil {
		log.Print(err)
	}
}

func ok(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func notFound(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	err = json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if err != nil {
		log.Print(err)
	}
}
func parseUrlQuery(r *http.Request) (*domain.UrlQueryParams, error) {
	params := domain.UrlQueryParams{}
	q := r.URL.Query()
	if q.Has("page") {
		if page, err := strconv.ParseUint(q.Get("page"), 10, 32); err == nil {
			params.Page = uint(page)
		} else {
			return nil, fmt.Errorf("parse Url Query page %w ", err)
		}
	}

	if q.Has("pageSize") {
		if size, err := strconv.ParseUint(q.Get("pageSize"), 10, 32); err == nil {
			params.PageSize = uint(size)
		} else {
			return nil, fmt.Errorf("parse Url Query page size %w ", err)
		}
	}

	if q.Has("showDeleted") {
		if show, err := strconv.ParseUint(q.Get("showDeleted"), 10, 32); err == nil {
			params.ShowDeleted = show > 0
		} else {
			return nil, fmt.Errorf("expected 'showDeleted' to be an unsigned integer, '%s' was given: %w", q.Get("showDeleted"), err)
		}
	}

	return &params, nil
}
