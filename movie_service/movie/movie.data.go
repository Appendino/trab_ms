package movie

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/appendino/trab_ms/movie_service/database"
)

func getMoviesByGenre(genre string) ([]Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := database.DbConn.QueryContext(ctx, `SELECT name 
		 FROM movies where upper(genre) = upper(?)`, genre)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	movies := make([]Movie, 0)
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.Name)
		movies = append(movies, movie)
	}
	return movies, nil
}

func getMovieDetail(moveId int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT description from movies where movieId = ?`, moveId)
	movie := &Movie{}
	err := row.Scan(&movie.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return movie, nil
}

func getMovieByKeyword(keyword string) ([]Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	keyword = strings.ToLower(keyword)
	stmt := fmt.Sprintf("SELECT NAME FROM movies where lower(name) like '%%%s%%'", keyword)
	fmt.Println(stmt)
	rows, err := database.DbConn.QueryContext(ctx, stmt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	movies := make([]Movie, 0)
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.Name)
		movies = append(movies, movie)
	}
	return movies, nil
}

func getTopMoviesByGenre(genre string) ([]Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	rows, err := database.DbConn.QueryContext(ctx, `SELECT name, views
		 FROM movies where upper(genre) = upper(?) order by views desc LIMIT 5`, genre)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	movies := make([]Movie, 0)
	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.Name, &movie.TotalViews)
		movies = append(movies, movie)
	}
	return movies, nil
}

func getMovieById(movieId int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT name, description, genre, views 
	from movies where movieId = ?`, movieId)
	movie := &Movie{}
	err := row.Scan(&movie.Name, &movie.Description, &movie.Genre, &movie.TotalViews)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return movie, nil
}
