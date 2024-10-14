package controller

import (
	"USER_TEST/common"
	"USER_TEST/domain"
	"USER_TEST/repository"
	"encoding/json"
	"net/http"

	"github.com/goinggo/tracelog"
)

type UserController struct{}

// create method by using http
func (mController *UserController) Create(w http.ResponseWriter, r *http.Request) {

	common.PanicRecovery("UserController", "Create")

	ObjReq := domain.UserPayload{}

	err := json.NewDecoder(r.Body).Decode(&ObjReq)

	if err != nil {
		tracelog.Error(err, "UserController", "Create: json decoding error")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "json decoding error", nil)
	}

	UserRepo := repository.IUserRepository(&repository.UserRepository{})

	if Result := UserRepo.Create(&ObjReq); Result {
		ReturnResponse(w, SUCCESSRESULT, TRUERESULT, SUCCESSCODE, "Create Student Successfull ", nil)

	} else {
		tracelog.Error(err, "UserController", "Create Failed")
		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "Upadate Failed", nil)

	}

}

func (mController *UserController) Update(w http.ResponseWriter, r *http.Request) {

	common.PanicRecovery("UserController", "Update")

	ObjReq := domain.UserPayload{}

	err := json.NewDecoder(r.Body).Decode(&ObjReq)

	if err != nil {
		tracelog.Error(err, "UserController", "Update: json decoding error")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "json decoding error", nil)
	}

	myRepo := repository.IUserRepository(&repository.UserRepository{})

	if myResult := myRepo.Update(&ObjReq); myResult {

		ReturnResponse(w, SUCCESSRESULT, TRUERESULT, SUCCESSCODE, "Update Succesfull", nil)
	} else {

		tracelog.Error(err, "UserController", "Update Failed")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "Upadate Failed", nil)
	}

}

func (mController *UserController) Delete(w http.ResponseWriter, r *http.Request) {

	common.PanicRecovery("UserController", "Delete")

	ObjReq := domain.UserPayload{}

	err := json.NewDecoder(r.Body).Decode(&ObjReq)

	if err != nil {
		tracelog.Error(err, "UserController", "Delete: json decoding error")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "json decoding error", nil)
	}

	myRepo := repository.IUserRepository(&repository.UserRepository{})

	if myResult := myRepo.Delete(ObjReq.Username); myResult {

		ReturnResponse(w, SUCCESSRESULT, TRUERESULT, SUCCESSCODE, "Delete Succesfull", nil)
	} else {

		tracelog.Error(err, "UserController", "Delete Failed")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "Delete Failed", nil)
	}

}

func (mController *UserController) GetAllUserData(w http.ResponseWriter, r *http.Request) {

	common.PanicRecovery("UserController", "GetAllUserData")

	ObjReq := domain.UserPayload{}

	err := json.NewDecoder(r.Body).Decode(&ObjReq)

	if err != nil {
		tracelog.Error(err, "UserController", "GetAllUserData: json decoding error")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "json decoding error", nil)
	}
	myRepo := repository.IUserRepository(&repository.UserRepository{})

	if myResult, status := myRepo.GetAllUserData(); status {

		ReturnResponse(w, myResult, TRUERESULT, SUCCESSCODE, "GetAllUserData Succesfull", nil)
	} else {

		tracelog.Error(err, "UserController", "GetAllUserData Failed")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "GetAllUserData Failed", nil)
	}

}
