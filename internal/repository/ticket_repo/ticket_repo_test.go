package ticketrepo

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/stretchr/testify/assert"
)

// unit test for ticket repo
func TestNewTicketRepo(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTicketRepo(db)
	if repo == nil {
		t.Error("Expected non-nil TickketRepo")
	}
	if repo.db != db {
		t.Error("Expected db to be set correctly")
	}
}

func TestTickketRepoSaveTickets(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTicketRepo(db)

	tickets := []model.Ticket{
		{
			TicketId:    "ticket1",
			ShowId:      "show1",
			UserId:      "user1",
			BookingTime: time.Now(),
		},
		{
			TicketId:    "ticket2",
			ShowId:      "show1",
			UserId:      "user2",
			BookingTime: time.Now(),
		},
	}

	mock.ExpectExec("INSERT INTO tickets").
		WithArgs(tickets[0].TicketId, tickets[0].ShowId, tickets[0].UserId, tickets[0].BookingTime).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO tickets").
		WithArgs(tickets[1].TicketId, tickets[1].ShowId, tickets[1].UserId, tickets[1].BookingTime).
		WillReturnResult(sqlmock.NewResult(2, 1))

	err = repo.SaveTickets(tickets)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveTicketsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTicketRepo(db)

	tickets := []model.Ticket{
		{
			TicketId:    "ticket1",
			ShowId:      "show1",
			UserId:      "user1",
			BookingTime: time.Now(),
		},
	}

	mock.ExpectExec("INSERT INTO tickets").
		WithArgs(tickets[0].TicketId, tickets[0].ShowId, tickets[0].UserId, tickets[0].BookingTime).
		WillReturnError(sql.ErrConnDone)

	err = repo.SaveTickets(tickets)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTicketById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTicketRepo(db)

	rows := sqlmock.NewRows([]string{"ticket_id", "show_id", "user_id", "booking_time"}).
		AddRow("ticket1", "show1", "user1", time.Now())

	mock.ExpectQuery("SELECT ticket_id, show_id, user_id, booking_time FROM tickets WHERE ticket_id = \\?").
		WithArgs("ticket1").
		WillReturnRows(rows)

	ticket, err := repo.GetTicketById("ticket1")
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, "ticket1", ticket.TicketId)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTicketByIdNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTicketRepo(db)

	rows := sqlmock.NewRows([]string{"ticket_id", "show_id", "user_id", "booking_time"})

	mock.ExpectQuery("SELECT ticket_id, show_id, user_id, booking_time FROM tickets WHERE ticket_id = \\?").
		WithArgs("nonexistent").
		WillReturnRows(rows)

	ticket, err := repo.GetTicketById("nonexistent")
	assert.Error(t, err)
	assert.Nil(t, ticket)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCancleTicket(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTicketRepo(db)

	mock.ExpectExec("DELETE FROM tickets WHERE ticket_id = \\?").
		WithArgs("ticket1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CancleTicket("ticket1")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCancleTicketError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewTicketRepo(db)

	mock.ExpectExec("DELETE FROM tickets WHERE ticket_id = \\?").
		WithArgs("ticket1").
		WillReturnError(sql.ErrConnDone)

	err = repo.CancleTicket("ticket1")
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
