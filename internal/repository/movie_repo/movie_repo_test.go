package movierepo

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/stretchr/testify/assert"
)

// unit testing for movie repo
func TestMovieRepoAddMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewMovieRepo(db)

	movie := model.Movie{
		MovieId:   "movie1",
		MovieName: "Test Movie",
		MovieType: "Action",
		Duration:  120,
	}

	mock.ExpectExec("INSERT INTO movies").
		WithArgs(movie.MovieId, movie.MovieName, movie.MovieType, movie.Duration).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddMovie(movie)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddMovieError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewMovieRepo(db)

	movie := model.Movie{
		MovieId:   "movie1",
		MovieName: "Test Movie",
		MovieType: "Action",
		Duration:  120,
	}

	mock.ExpectExec("INSERT INTO movies").
		WithArgs(movie.MovieId, movie.MovieName, movie.MovieType, movie.Duration).
		WillReturnError(sql.ErrConnDone)

	err = repo.AddMovie(movie)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestViewAllMovies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewMovieRepo(db)

	rows := sqlmock.NewRows([]string{"movie_id", "movie_name", "movie_type", "duration"}).
		AddRow("movie1", "Test Movie", "Action", 120).
		AddRow("movie2", "Another Movie", "Drama", 90)

	mock.ExpectQuery("SELECT movie_id, movie_name, movie_type, duration FROM movies").
		WillReturnRows(rows)

	movies, err := repo.ViewAllMovies()
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
	assert.Equal(t, "movie1", movies[0].MovieId)
	assert.Equal(t, "movie2", movies[1].MovieId)
}

func TestViewAllMoviesError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewMovieRepo(db)

	mock.ExpectQuery("SELECT movie_id, movie_name, movie_type, duration FROM movies").
		WillReturnError(sql.ErrConnDone)

	movies, err := repo.ViewAllMovies()
	assert.Error(t, err)
	assert.Nil(t, movies)
}
