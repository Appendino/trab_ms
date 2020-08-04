package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/appendino/trab_ms/user_service/config"
	"github.com/appendino/trab_ms/user_service/discovery"
)

func SetupRoutes(apiBasePath string) {
	handleVote := http.HandlerFunc(voteHandler)
	handleList := http.HandlerFunc(addListHandler)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "vote"), handleVote)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "addlist"), handleList)
}

func getMovie(movieId int) (*Movie, error) {
	movieService, err := discovery.GetService(config.CONSUL_SERVER, config.CONSUL_SERVICE_NAME,
		config.HOSTNAME, config.PORT, config.MOVIE_CONSUL)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//movieService = config.MOVIE_PROTOCOL + "://" + movieService + config.MOVIE_ENDPOINT + "/" + strconv.Itoa(movieId)
	movieService = config.MOVIE_PROTOCOL + "://" + "localhost:8088" + config.MOVIE_ENDPOINT + "/" + strconv.Itoa(movieId)
	resp, err := http.Get(movieService)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var movie Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var ids UserMovie
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &ids)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(ids)
		movie, err := getMovie(ids.MovieID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if movie == nil {
			log.Println("invalid movie")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = insertVote(ids)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		movieJson, err := json.Marshal(movie)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(movieJson)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func addListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var ids UserMovie
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &ids)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(ids)
		movie, err := getMovie(ids.MovieID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if movie == nil {
			log.Println("invalid movie")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = insertVote(ids)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		movieJson, err := json.Marshal(movie)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(movieJson)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
