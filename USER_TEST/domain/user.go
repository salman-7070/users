package domain

type UserPayload struct {
	Username string `json :  "username"`
	Password string `json :  "password"`
	Active   string `json :  "active"`
	Id       string `json :  "id"`
}

type GeneralResponseResource struct {
	ResponseData interface{}

	ResponseStatus      string `json : "reponsestatus"`
	ResponseCode        string `json : "responsecode"`
	ResponseDescription string `json : "responsedescription"`
}
