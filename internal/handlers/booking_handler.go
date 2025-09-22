package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/contextkey"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/response"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/bookingservice"
)

type BookingHandler struct {
	bookingService bookingservice.BookingServiceInterface
}

func NewBookingHandler(bookingService bookingservice.BookingServiceInterface)*BookingHandler{
	return &BookingHandler{bookingService: bookingService}
}

func(bookingHandler * BookingHandler)BookTicket(w http.ResponseWriter,r *http.Request){
	userId:=r.Context().Value(contextkey.UserId).(string)
	
    showId:=r.PathValue("show_id")
	if showId==""{
        response.ErrorResponse(w,"Invalid request",http.StatusBadRequest)
		return
	}
    
	type RequestBody struct{
		NumberOfSeat int `json:"numberofseat"`
	}
	var seats RequestBody

	err:=json.NewDecoder(r.Body).Decode(&seats)

	if err!=nil{
		response.ErrorResponse(w,"Invalid request body",http.StatusBadRequest)
		return
	}

	tickets,err:=bookingHandler.bookingService.BookTicket(showId,userId,seats.NumberOfSeat)

	if err!=nil{
		response.ErrorResponse(w,"Ticket booking failed",http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(w,tickets,"Tickets Booked Successfully",http.StatusOK)
	
}

func(bookingHandler *BookingHandler)CancelTicket(w http.ResponseWriter,r *http.Request){
	ticketId:=r.PathValue("ticket_id")
    if ticketId==""{
		response.ErrorResponse(w,"Invalid request",http.StatusBadRequest)
		return
	}
    
	err:=bookingHandler.bookingService.CancelTicket(r.Context(),ticketId)

	if err!=nil{
		response.ErrorResponse(w,"Ticket Cancellation failed",http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(w,nil,"ticket cancelled",http.StatusOK)
}