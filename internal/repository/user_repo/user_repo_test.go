package userrepo

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/roles"
	"github.com/stretchr/testify/assert"
)

// unit testing for user repo

func TestNewUserRepo(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	if repo == nil {
		t.Error("Expected non-nil UserRepo")
	}
	if repo.db != db {
		t.Error("Expected db to be set correctly")
	}
}

func TestSaveUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	mock.ExpectExec("INSERT INTO users").
		WithArgs("user1", "John Doe", "john@example.com", "1234567890", "password", role.Customer).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.SaveUser("user1", "John Doe", "john@example.com", "1234567890", "password")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveUserError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	mock.ExpectExec("INSERT INTO users").
		WithArgs("user1", "John Doe", "john@example.com", "1234567890", "password", role.Customer).
		WillReturnError(sql.ErrConnDone)

	err = repo.SaveUser("user1", "John Doe", "john@example.com", "1234567890", "password")
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	rows := sqlmock.NewRows([]string{"user_id", "name", "email", "phone_number", "password", "role"}).
		AddRow("user1", "John Doe", "john@example.com", "1234567890", "password", string(role.Customer))

	mock.ExpectQuery("SELECT user_id, name, email, phone_number, password, role FROM users WHERE email = \\?").
		WithArgs("john@example.com").
		WillReturnRows(rows)

	user, err := repo.GetUserByEmail("john@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "user1", user.Id)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "1234567890", user.PhoneNumber)
	assert.Equal(t, string(role.Customer), string(user.Role))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByEmailNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	mock.ExpectQuery("SELECT user_id, name, email, phone_number, password, role FROM users WHERE email = \\?").
		WithArgs("nonexistent@example.com").
		WillReturnError(sql.ErrNoRows)

	user, err := repo.GetUserByEmail("nonexistent@example.com")
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	rows := sqlmock.NewRows([]string{"user_id", "name", "email", "phone_number", "password", "role"}).
		AddRow("user1", "John Doe", "john@example.com", "1234567890", "password", string(role.Customer))

	mock.ExpectQuery("SELECT user_id, name, email, phone_number, password, role FROM users where user_id=\\?").
		WithArgs("user1").
		WillReturnRows(rows)

	user, err := repo.GetUserById("user1")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "user1", user.Id)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "1234567890", user.PhoneNumber)
	assert.Equal(t, string(role.Customer), string(user.Role))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByIdNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	mock.ExpectQuery("SELECT user_id, name, email, phone_number, password, role FROM users where user_id=\\?").
		WithArgs("nonexistent").
		WillReturnError(sql.ErrNoRows)

	user, err := repo.GetUserById("nonexistent")
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}
