package request

import "time"


type NewsRequest struct {
	Subject string `form:"subject" binding:"min=3,required"`
	From time.Time `form:"from" binding:"required" time_format:"2006-01-02"`
}