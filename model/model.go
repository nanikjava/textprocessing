package model

//Datarecord model for database and user's response
type Datarecord struct {
	DateISO8601  string `json:"eventTime"`
	EmailAddress string `json:"email"`
	SessionID    string `json:"sessionId"`

	// Use omitempty to make sure that this will not be sent as part of
	// user's response
	FileName string `json:",omitempty"`
}

//RequestBody populated with data from request
type RequestBody struct {
	Filename string `json:"filename" binding:"required"`
	From     string `json:"from" binding:"required"`
	To       string `json:"to" binding:"required"`
}

//ResponseError for error reporting to user
type ResponseError struct {
	Message string `json:"message"`
}
