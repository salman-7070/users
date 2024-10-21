package repository

import (
	"USER_TEST/common"
	"USER_TEST/domain"

	"log"
)

// signature for user methods in repository
// here iam using bool as respose for Create,update and delete methodes
type IUserRepository interface {
	Create(Obj *domain.UserPayload) bool
	Update(Obj *domain.UserPayload) bool
	GetAllUserData() ([]domain.UserPayload, bool)
	Delete(Username string) bool
	GetByUsername(Username string) (domain.UserPayload, bool)
}

type UserRepository struct{}

func (mRepo UserRepository) Create(Obj *domain.UserPayload) bool {

	//this method is for recovering panic error
	defer common.PanicRecovery("UserRepository", "Create")

	//Connect database
	myDb, result := common.GetDBConnection()

	if !result {
		log.Println("UserRepository Create : Failed to connect DataBase")

		return false
	}

	//I used $1 as a placeholder because I am using PostgreSQL as the database driver.
	err := myDb.QueryRow(`INSERT  INTO users ( username,
												password,
												active,
												email)
												VALUES($1,$2,$3,$4) RETURNING id`,

		Obj.Username,
		Obj.Password,
		Obj.Active,
		Obj.Email).Scan(&Obj.Id)

	if err != nil {
		log.Println("UserRepository Create :qStmt Exec ", err)
		return false

	}

	log.Println("UserRepository Create Success")

	return true
}

func (mRepo UserRepository) Update(Obj *domain.UserPayload) bool {

	defer common.PanicRecovery("UserRepository", "Update")

	//Connect database
	myDb, result := common.GetDBConnection()

	if !result {
		log.Println("UserRepository Update : Failed to connect DataBase")
		return false
	}
	//Here iam using id as the unique identifier for updating user details or we can use username also
	QRows, err := myDb.Query(`UPDATE users  SET      username=$1,
												     password=$2,
												     active=$3,
													 email=$4
													 WHERE id=$5 `,
		Obj.Username,
		Obj.Password,
		Obj.Active,
		Obj.Email,
		Obj.Id)

	if err != nil {
		log.Println("UserRepository UpdateBeneficiaryMaster :Query Exec ", err)
		return false
	}
	//defer func will exicute at the end of the method

	defer QRows.Close()

	log.Println("updated successfully")

	return true
}

func (mRepo UserRepository) GetAllUserData() ([]domain.UserPayload, bool) {

	defer common.PanicRecovery("UserRepository", "GetAllUserData")

	//variable for returning array of data
	myResults := []domain.UserPayload{}

	//for datbase connection
	myDb, result := common.GetDBConnection()

	if !result {
		log.Println("UserRepository GetAllUserData :myDb OpenConnection")
		return myResults, false
	}

	QRows, err := myDb.Query(`select    id,
										username,
										password,
										active,
										email
										FROM users`)

	if err != nil {
		log.Println("UserRepository GetAllUserData : myDb.Query Error", err)
		return myResults, false
	}

	for QRows.Next() {
		// variable for scanning objects in the detail table
		myResult := domain.UserPayload{}
		err = QRows.Scan(
			&myResult.Id,
			&myResult.Username,
			&myResult.Password,
			&myResult.Active,
			&myResult.Email)

		if err != nil {
			log.Println("UserRepository GetAllUserData : qResult.Scan", err)
			return myResults, false
		}

		myResults = append(myResults, myResult)

		//defer func will exicute at the end of the method
		defer QRows.Close()
	}

	log.Println("UserRepository GetAllUserData Success")

	return myResults, true
}

func (mRepo UserRepository) GetByUsername(Username string) (domain.UserPayload, bool) {

	defer common.PanicRecovery("UserRepository", "GetAllUserData")

	//variable for returning array of data
	myResult := domain.UserPayload{}

	//for datbase connection
	myDb, result := common.GetDBConnection()

	if !result {
		log.Println("UserRepository GetAllUserData :myDb OpenConnection")
		return myResult, false
	}

	err := myDb.QueryRow(`select    id,
										username,
										password,
										active,
										email
										FROM users
										where username=$1`, Username).Scan(&myResult.Id,
		&myResult.Username,
		&myResult.Password,
		&myResult.Active,
		&myResult.Email)

	if err != nil {
		log.Println("UserRepository GetAllUserData : qResult.Scan", err)
		return myResult, false
	}

	log.Println("UserRepository GetAllUserData Success")

	return myResult, true
}

func (mRepo UserRepository) Delete(Username string) bool {
	defer common.PanicRecovery("UserRepository", "Delete")

	// Connect to the database
	myDb, connected := common.GetDBConnection()
	if !connected {
		log.Println("UserRepository Delete: Failed to connect to the database")
		return false
	}
	defer myDb.Close()

	// Execute the delete query
	result, err := myDb.Exec(`DELETE FROM users WHERE username = $1`, Username)
	if err != nil {
		log.Println("UserRepository Delete: Query Exec error", err)
		return false
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("UserRepository Delete: Error getting rows affected", err)
		return false
	}

	if rowsAffected == 0 {
		log.Println("UserRepository Delete: No user found with the provided username")
		return false
	}

	log.Println("User deleted successfully")
	return true
}
