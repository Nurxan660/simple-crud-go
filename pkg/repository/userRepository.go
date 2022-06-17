package repository

import (
	"fmt"
	"log"
	"testGo/pkg/config"
	"testGo/pkg/models"
)

func CreateTable(){
	db := config.CreateConnection()
	query := "CREATE TABLE  if not exists users (userid SERIAL PRIMARY KEY,firstname varchar,lastname varchar,age int);"

	  res:=db.QueryRow(query)
	  fmt.Print(res)

	  
	defer db.Close()


}

func InsertUser(user models.User) int64 {
	db := config.CreateConnection()

	query := "INSERT INTO users (firstName,lastName,age) values($1,$2,$3) RETURNING userid"

	var id int64

	err := db.QueryRow(query, user.Firstname, user.Lastname, user.Age).Scan(&id)
	if err != nil {
		log.Fatal("cannot execute this one",err)
	}

	return id

}


func GetUser(id int64) (models.User, error) {
	
	db := config.CreateConnection()

	

	var user models.User

	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Age)

	

	defer db.Close()

	return user, err
}


func UpdateUser(id int64, user models.User) int64 {

	db := config.CreateConnection()

	

	sqlStatement := `UPDATE users SET lastname=$3, firstname=$2, age=$4 WHERE userid=$1`

	 res,err := db.Exec(sqlStatement, id, user.Firstname, user.Lastname, user.Age)

	if err != nil {
		log.Fatalf("Cannot execute the query. %v", err)
	}

	fmt.Print(res)

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	defer db.Close()

	return rowsAffected


	

	
}

func DeleteUser(id int64) int64  {

	db := config.CreateConnection()

	

	sqlStatement := `DELETE FROM users WHERE userid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Cannot execute the query. %v", err)
	}

	fmt.Print(res)

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	defer db.Close()

	return rowsAffected


	

	
}
