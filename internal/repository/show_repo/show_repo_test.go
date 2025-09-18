package showrepo

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestNewShowRepo(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)
	if repo == nil {
		t.Error("Expected non-nil ShowRepo")
	}
	if repo.db != db {
		t.Error("Expected db to be set correctly")
	}
}

func TestCreateShow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	show := &model.Show{
		ShowId:        "show1",
		MovieId:       "movie1",
		StartTime:     time.Now(),
		EndTime:       time.Now().Add(time.Hour),
		TotalSeats:    100,
		AvailableSeat: 100,
	}

	mock.ExpectExec("INSERT INTO shows").
		WithArgs(show.ShowId, show.MovieId, show.StartTime, show.EndTime, show.TotalSeats, show.AvailableSeat).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateShow(show)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateShowError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	show := &model.Show{
		ShowId:        "show1",
		MovieId:       "movie1",
		StartTime:     time.Now(),
		EndTime:       time.Now().Add(time.Hour),
		TotalSeats:    100,
		AvailableSeat: 100,
	}

	mock.ExpectExec("INSERT INTO shows").
		WithArgs(show.ShowId, show.MovieId, show.StartTime, show.EndTime, show.TotalSeats, show.AvailableSeat).
		WillReturnError(sql.ErrConnDone)

	err = repo.CreateShow(show)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllShow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	rows := sqlmock.NewRows([]string{"show_id", "movie_id", "start_time", "end_time", "total_seats", "available_seats"}).
		AddRow("show1", "movie1", time.Now(), time.Now().Add(time.Hour), 100, 100).
		AddRow("show2", "movie2", time.Now().Add(2*time.Hour), time.Now().Add(3*time.Hour), 50, 50)

	mock.ExpectQuery("SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows").
		WillReturnRows(rows)

	result, err := repo.GetAllShow()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "show1", result[0].ShowId)
	assert.Equal(t, "show2", result[1].ShowId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllShowError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	mock.ExpectQuery("SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows").
		WillReturnError(sql.ErrConnDone)

	result, err := repo.GetAllShow()
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetShowByMovieId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	rows := sqlmock.NewRows([]string{"show_id", "movie_id", "start_time", "end_time", "total_seats", "available_seats"}).
		AddRow("show1", "movie1", time.Now(), time.Now().Add(time.Hour), 100, 100).
		AddRow("show2", "movie1", time.Now().Add(2*time.Hour), time.Now().Add(3*time.Hour), 50, 50)

	mock.ExpectQuery("SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows WHERE movie_id = \\?").
		WithArgs("movie1").
		WillReturnRows(rows)

	result, err := repo.GetShowByMovieId("movie1")
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "show1", result[0].ShowId)
	assert.Equal(t, "show2", result[1].ShowId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetShowByMovieIdEmpty(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	rows := sqlmock.NewRows([]string{"show_id", "movie_id", "start_time", "end_time", "total_seats", "available_seats"})

	mock.ExpectQuery("SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows WHERE movie_id = \\?").
		WithArgs("movie3").
		WillReturnRows(rows)

	result, err := repo.GetShowByMovieId("movie3")
	assert.NoError(t, err)
	assert.Len(t, result, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateShow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	mock.ExpectExec("UPDATE shows SET available_seats = \\? WHERE show_id = \\?").
		WithArgs(80, "show1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateShow(80, "show1")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateShowError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	mock.ExpectExec("UPDATE shows SET available_seats = \\? WHERE show_id = \\?").
		WithArgs(80, "nonexistent").
		WillReturnError(sql.ErrConnDone)

	err = repo.UpdateShow(80, "nonexistent")
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestIsConflictTrue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	start := time.Now().Add(30 * time.Minute)
	end := time.Now().Add(90 * time.Minute)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM shows WHERE start_time < \\? AND end_time > \\?").
		WithArgs(end, start).
		WillReturnRows(rows)

	conflicts, err := repo.IsConflict(start, end)
	assert.NoError(t, err)
	assert.True(t, conflicts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestIsConflictFalse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	start := time.Now().Add(2 * time.Hour)
	end := time.Now().Add(3 * time.Hour)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM shows WHERE start_time < \\? AND end_time > \\?").
		WithArgs(end, start).
		WillReturnRows(rows)

	conflicts, err := repo.IsConflict(start, end)
	assert.NoError(t, err)
	assert.False(t, conflicts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestIsConflictError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	start := time.Now()
	end := time.Now().Add(time.Hour)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM shows WHERE start_time < \\? AND end_time > \\?").
		WithArgs(end, start).
		WillReturnError(sql.ErrConnDone)

	conflicts, err := repo.IsConflict(start, end)
	assert.Error(t, err)
	assert.False(t, conflicts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetShowByShowId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	rows := sqlmock.NewRows([]string{"show_id", "movie_id", "start_time", "end_time", "total_seats", "available_seats"}).
		AddRow("show1", "movie1", time.Now(), time.Now().Add(time.Hour), 100, 100)

	mock.ExpectQuery("SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows WHERE show_id = \\?").
		WithArgs("show1").
		WillReturnRows(rows)

	result, err := repo.GetShowByShowId("show1")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "show1", result.ShowId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetShowByShowIdNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewShowRepo(db)

	rows := sqlmock.NewRows([]string{"show_id", "movie_id", "start_time", "end_time", "total_seats", "available_seats"})

	mock.ExpectQuery("SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows WHERE show_id = \\?").
		WithArgs("nonexistent").
		WillReturnRows(rows)

	result, err := repo.GetShowByShowId("nonexistent")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}
