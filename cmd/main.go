package main

import (
	"fmt"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/config"
)

func main() {
	db, err := config.GetDbInstance()
	if err != nil {
		fmt.Printf(" Unable to connect with database: %v", err)
		return 
	}
    defer func ()  {
      db.Close()
	}()
	
    RunApp(db)
}

