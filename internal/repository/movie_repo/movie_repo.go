package movierepo

import (
	"database/sql"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MovieRepo struct {
	db *sql.DB
}

func NewMovieRepo(db * sql.DB)*MovieRepo{
	return &MovieRepo{db: db}
}

func(movieRepo * MovieRepo)AddMovie(newMovie model.Movie)error{
	query:=`INSERT INTO movies (movie_id, movie_name, movie_type, duration) VALUES (?, ?, ?, ?)` 

	_,err:=movieRepo.db.Exec(query,newMovie.MovieId,newMovie.MovieName,newMovie.MovieType,newMovie.Duration)

	if err!=nil{
		return err
	}
	return nil
}

func(movieRepo * MovieRepo)ViewAllMovies()([]model.Movie,error){

	var movies []model.Movie
	query:=`SELECT movie_id, movie_name, movie_type, duration FROM movies`
	rows,err:=movieRepo.db.Query(query)
    
	if err!=nil{
        return nil,err
	}

	defer rows.Close()
    for rows.Next(){
       var movie model.Movie
	   err:=rows.Scan(&movie.MovieId,&movie.MovieName,&movie.MovieType,&movie.Duration)

	   if err!=nil{
		return nil,err

	   }
	   movies=append(movies, movie)
	}

	if err=rows.Err();err!=nil{
		return nil,err
	}
	return movies,nil
}
