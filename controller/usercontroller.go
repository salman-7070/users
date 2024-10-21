package controller

import (
	"USER_TEST/common"
	"USER_TEST/domain"
	"USER_TEST/repository"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/goinggo/tracelog"
	"github.com/google/uuid"
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

	getUsername, _ := UserRepo.GetByUsername(ObjReq.Username)

	if getUsername.Username == ObjReq.Username {
		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "Username already exist", nil)

	}

	if Result := UserRepo.Create(&ObjReq); Result {
		ReturnResponse(w, SUCCESSRESULT, TRUERESULT, SUCCESSCODE, "Create Student Successfull ", nil)

	} else {
		tracelog.Error(err, "UserController", "Create Failed")
		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "Update Failed", nil)

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

	// ObjReq := domain.UserPayload{}

	// err := json.NewDecoder(r.Body).Decode(&ObjReq)

	// if err != nil {
	// 	tracelog.Error(err, "UserController", "GetAllUserData: json decoding error")

	// 	ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "json decoding error", nil)
	// }
	myRepo := repository.IUserRepository(&repository.UserRepository{})

	if myResult, status := myRepo.GetAllUserData(); status {

		ReturnResponse(w, myResult, "", "SUCCESSCODE", "GetAllUserData Succesfull", nil)
	} else {

		tracelog.Error(errors.New("error in get all call"), "UserController", "GetAllUserData Failed")

		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "GetAllUserData Failed", nil)
	}

}

func (mController *UserController) GenarateToken(Obj *domain.UserPayload) (string, error) {

	uniqueID := uuid.New().String()
	claim := map[string]interface{}{

		"username": Obj.Username,
		"password": Obj.Password,
		"iat":      time.Now().Unix(), // Update the timestamp
		"jti":      uniqueID,
	}

	jsonData, err := json.Marshal(claim)

	if err != nil {
		return "error in encoding", err
	}

	// header := `{"alg":"none","typ":"JWT"}` // Header indicates no signature
	// headerEncoded := base64.RawURLEncoding.EncodeToString([]byte(header))

	dataencoded := base64.RawURLEncoding.EncodeToString(jsonData)

	return dataencoded, nil
}

func (mController *UserController) Login(w http.ResponseWriter, r *http.Request) {

	common.PanicRecovery("Users", "Login")

	ObjReq := domain.UserPayload{}

	err := json.NewDecoder(r.Body).Decode(&ObjReq)
	if err != nil {
		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "json decoding error", nil)

	}

	userRepo := repository.IUserRepository(&repository.UserRepository{})

	getuserdetails, _ := userRepo.GetByUsername(ObjReq.Username)

	if ObjReq.Username == getuserdetails.Username && ObjReq.Password == getuserdetails.Password {

		ReturnResponse(w, getuserdetails, "", "SUCCESSCODE", "Login Succefull", nil)

	} else {
		ReturnResponse(w, FAILEDRESULT, FALSERESULT, FAILEDCODE, "Login Failed", errors.New("Login Failed"))

	}

}
