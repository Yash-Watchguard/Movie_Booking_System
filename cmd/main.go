package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/config"
)

func main() {
	db, err := config.GetDbInstance()
	if err != nil {
		log.Fatalf(" Unable to connect with database: %v", err)
	}
    defer func ()  {
      db.Close()
	}()
	

	c:=make(chan os.Signal,1)

	signal.Notify(c,os.Interrupt,syscall.SIGTERM)

	go func(){
       <-c
	   fmt.Println("Disconnectiong from mysql")
	   db.Close()
	   os.Exit(1)
	}()

    RunApp(db)
}
