package utills

import "time"

func ValidateTime(timeString string)(time.Time,bool){
	layout:=time.RFC3339

	parsedTime,err:=time.Parse(layout,timeString)
    if err!=nil{
		return time.Now(),false
	}

	return parsedTime,true
}
