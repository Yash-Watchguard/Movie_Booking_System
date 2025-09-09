package showrepo

import (
	"database/sql"
	"sync"
	"time"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type ShowRepo struct {
	db *sql.DB
	mu *sync.Mutex
}

func NewShowRepo(db *sql.DB)*ShowRepo{
	return &ShowRepo{db: db}
}

func(showRepo *ShowRepo)CreateShow(show *model.Show)(error){
	query:=`INSERT INTO shows (show_id, movie_id, start_time, end_time, total_seats, available_seats) VALUES (?, ?, ?, ?, ?, ?)`
     
	showRepo.mu.Lock()
	_,err:=showRepo.db.Exec(query,show.ShowId,show.MovieId,show.StartTime,show.EndTime,show.TotalSeats,show.AvailableSeat)
    showRepo.mu.Unlock()
	if err!=nil{
		return err
	}

	return nil
}

func(showRepo *ShowRepo)GetAllShow()([]model.Show,error){
	qrery:=`SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows`
    var allShows []model.Show

	showRepo.mu.Lock()
	rows,err:=showRepo.db.Query(qrery)
	showRepo.mu.Unlock()

	if err!=nil{
		return nil,err
	}
    defer rows.Close()
	for rows.Next(){
       var show model.Show

	   rows.Scan(
		&show.ShowId,
		&show.MovieId,
		&show.StartTime,
		&show.EndTime,
		&show.TotalSeats,
		&show.AvailableSeat,
	   )
	   allShows=append(allShows, show)
	}

	if err:=rows.Err();err!=nil{
		return nil,err
	}

	return allShows,nil
}

func(showRepo *ShowRepo)GetShowByMovieId(movieId string)([]model.Show,error){
	qrery:=`SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows WHERE movie_id = ?`
    var allShows []model.Show

	showRepo.mu.Lock()
	rows,err:=showRepo.db.Query(qrery,movieId)
	showRepo.mu.Unlock()

	if err!=nil{
		return nil,err
	}
    defer rows.Close()
	for rows.Next(){
       var show model.Show

	   rows.Scan(
		&show.ShowId,
		&show.MovieId,
		&show.StartTime,
		&show.EndTime,
		&show.TotalSeats,
		&show.AvailableSeat,
	   )
	   allShows=append(allShows, show)
	}

	if err:=rows.Err();err!=nil{
		return nil,err
	}

	return allShows,nil
}

func(showRepo *ShowRepo)UpdateShow(updatedSeat int,showId string)error{
	query:=`UPDATE shows SET available_seats = ? WHERE show_id = ?`

     showRepo.mu.Lock()
	_,err:=showRepo.db.Exec(query,updatedSeat,showId)
	showRepo.mu.Unlock()

	if err!=nil{
		return err
	}
	return nil
}
func(showRepo *ShowRepo)IsConflict(showStartTime,showEndTime time.Time)(bool,error){
	var count int

	showRepo.mu.Lock()
	query:=`SELECT COUNT(*) FROM shows WHERE start_time < ? AND end_time > ?`

	showRepo.mu.Unlock()

	err:=showRepo.db.QueryRow(query,showEndTime,showStartTime).Scan(&count)

	if err!=nil{
		return false,err
	}

	return count>0,nil
}

func(showRepo *ShowRepo)GetShowByShowId(showId string)(*model.Show,error){
	query:=`SELECT show_id, movie_id, start_time, end_time, total_seats, available_seats FROM shows WHERE show_id = ?`
    var oneShow model.Show

	showRepo.mu.Lock()
	row:=showRepo.db.QueryRow(query,showId)
	showRepo.mu.Unlock()

	err:=row.Scan(&oneShow.ShowId,&oneShow.MovieId,&oneShow.StartTime,&oneShow.EndTime,&oneShow.TotalSeats,&oneShow.AvailableSeat)
	if err!=nil{
		return nil,err
	}
	return &oneShow,nil
}
