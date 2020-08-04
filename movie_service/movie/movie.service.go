package movie

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func SetupRoutes(apiBasePath string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(fmt.Sprintf("%s/%s", apiBasePath, "genre"), genreHandler)
	r.HandleFunc(fmt.Sprintf("%s/%s/{id:[0-9]+}", apiBasePath, "desc"), descriptionHandler)
	r.HandleFunc(fmt.Sprintf("%s/%s", apiBasePath, "search"), keywordHandler)
	r.HandleFunc(fmt.Sprintf("%s/%s", apiBasePath, "top"), topMoviesHandler)
	r.HandleFunc(fmt.Sprintf("%s/{id:[0-9]+}", apiBasePath), movieIdHandler)
	return r
}

func genreHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	genre := keys.Get("genre")
	if genre == "" {
		log.Println("error:: genre deve ser informado")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	genre = strings.Trim(genre, "\t\r ")
	movies, err := getMoviesByGenre(genre)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	moviesJSON, err := json.Marshal(movies)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(moviesJSON)
}

func descriptionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	movie, err := getMovieDetail(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if movie == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	movieJSON, err := json.Marshal(movie)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(movieJSON)
}

func keywordHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	keyword := keys.Get("key")
	if keyword == "" {
		log.Println("error:: key deve ser informado")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keyword = strings.Trim(keyword, "\t\r ")
	movies, err := getMovieByKeyword(keyword)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	moviesJSON, err := json.Marshal(movies)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(moviesJSON)

}

func topMoviesHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	genre := keys.Get("genre")
	if genre == "" {
		log.Println("error:: genre deve ser informado")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	genre = strings.Trim(genre, "\t\r ")
	movies, err := getTopMoviesByGenre(genre)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(movies) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	moviesJSON, err := json.Marshal(movies)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(moviesJSON)
}

func movieIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	movie, err := getMovieById(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if movie == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	movieJSON, err := json.Marshal(movie)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(movieJSON)
}
