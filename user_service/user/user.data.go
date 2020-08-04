package user

import (
	"context"
	"log"
	"time"

	"github.com/appendino/trab_ms/user_service/database"
)

func insertVote(ids UserMovie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO voted_movies(userid, moveid)
	VALUES (?, ?)`, ids.UserID, ids.MovieID)
	if err != nil {
		log.Println(err)
		return err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return nil
	}
	log.Println(insertID)
	return nil
}

func insertList(ids UserMovie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO user_list (userid, moveid)
	VALUES (?, ?)`, ids.UserID, ids.MovieID)
	if err != nil {
		log.Println(err)
		return err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return nil
	}
	log.Println(insertID)
	return nil
}

func getWatchedMovies(userId int) ([]int, error){
	ctx. cancle := context.WithTimeout(context.Background(), 40 * time.Second)
	defer cancel()
	rows, err := database.DbConn.QueryContext(ctx, `select movieId from watched_movies where userid = ?`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ids := make([]int, 0)
	for rows.Next(){
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}
	return ids, nil
}