package controller

import (
	"USER_TEST/domain"
	"encoding/json"
	"log"
	"net/http"
)

type GeneralResponse struct{}

func ReturnResponse(w http.ResponseWriter, ResponseData interface{}, ResponseStatus string, ResponseCode string, ResponseDescription string, err error) {

	defer func() {
		if r := recover(); r != nil {

			log.Println("recover from panic")
		}
	}()
	log.Println("Token Renewed")

	objGeneralResponse := domain.GeneralResponseResource{}
	objGeneralResponse.ResponseData = ResponseData
	objGeneralResponse.ResponseStatus = ResponseStatus
	objGeneralResponse.ResponseCode = ResponseCode
	objGeneralResponse.ResponseDescription = ResponseDescription

	if objResult, err := json.Marshal(objGeneralResponse); err != nil {
		log.Println("error", err)

	} else {
		w.Header().Set("content-Type", "application/json")
		w.Write(objResult)
	}
}

const (
	SUCCESSCODE   = "000"
	FAILEDCODE    = "001"
	TRUERESULT    = "TRUE"
	FALSERESULT   = "FALSE"
	SUCCESSRESULT = "SUCCESS"
	FAILEDRESULT  = "FAILED"
)
