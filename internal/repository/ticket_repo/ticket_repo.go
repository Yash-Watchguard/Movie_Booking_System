package ticketrepo

import (
	"database/sql"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type TickketRepo struct {
	db *sql.DB
}

func NewTicketRepo(db *sql.DB)*TickketRepo{
	return &TickketRepo{db: db}
}

func(ticketRepo *TickketRepo)SaveTickets(tickets []model.Ticket)error{
query:=`INSERT INTO tickets (ticket_id, show_id, user_id, booking_time) VALUES (?, ?, ?, ?)`

for _,ticket:=range tickets{
	_,err:=ticketRepo.db.Exec(query,ticket.TicketId,ticket.ShowId,ticket.UserId,ticket.BookingTime)
    
	if err!=nil{
		return err
	}
}
return nil
}
func(ticketRepo *TickketRepo)GetTicketById(ticketId string)(*model.Ticket,error){
	query:=`SELECT ticket_id, show_id, user_id, booking_time FROM tickets WHERE ticket_id = ?`
	var t model.Ticket
	row:=ticketRepo.db.QueryRow(query,ticketId)
	err:=row.Scan(&t.TicketId,&t.ShowId,&t.UserId,&t.BookingTime)

	if err!=nil{
		return nil,err
	}
	return &t,nil
}
func(ticketRepo *TickketRepo)CancleTicket(ticketId string)(error){
	query:=`DELETE FROM tickets WHERE ticket_id = ? `
     
	_,err:=ticketRepo.db.Exec(query,ticketId)

	if err!=nil{
		return err
	}
	return nil
}
