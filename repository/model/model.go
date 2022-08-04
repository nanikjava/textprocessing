package model

//Datarecord model for database
type Datarecord struct {
	DateISO8601  string `json:"eventTime"`
	EmailAddress string `json:"email"`
	SessionID    string `json:"sessionId"`
}

type RequestBody struct {
	Filename string `json:"filename" binding:"required"`
	From     string `json:"from" binding:"required"`
	To       string `json:"to" binding:"required"`
}

//type ResponseBody struct {
//	Date  string `json:"eventTime"`
//	Email string `json:"email"`
//	To    string `json:"sessionId"`
//}

type ResponseError struct {
	Message string `json:"message"`
}
