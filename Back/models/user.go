package main //for testing can be main once its all working change to models

import (
	"UrlShort/env"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)


func ConnectDB()  *pgx.Conn{
	user := env.DBUSER()
	pass := env.DBPASS()
	dbname := env.DBNAME()
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", user, pass, dbname) //windows url user and pass for now 
	ctx := context.Background()

  conn, err := pgx.Connect(ctx, url)
  if err != nil {
		log.Println("Faild to connect")
    panic(err)
  }
	// defer conn.Close(ctx)
	fmt.Println("Connected DB")
	return conn
}

	// defer conn.Close(ctx)
	//use this in the main.go file when using these function refer to tracker backend for more 
func main()  {
	ConnectDB()
}

//just use this function as an example moving forward 
// func InsertIntoSearchedFood(conn *pgx.Conn, nameOfFood string, calsOfFood int) string{
// 	var food Food
// 	err := conn.QueryRow(
// 		//changed this name and from food_data to searched_food
// 		context.Background(),
// 		`INSERT INTO searched_food(food_name, food_cals, day_of_food)
// 		 VALUES ($1, $2, CURRENT_DATE)
// 		 RETURNING id, food_name, food_cals, day_of_food`,
// 		nameOfFood, calsOfFood,
// 	).Scan(&food.Id, &food.Food_name, &food.Food_Cals, &food.Date)
//
// 	foodDate := food.Date.Format("02-01-2006")
// 	foodData := fmt.Sprintf(" Id: %d | %s | %d | %s", food.Id, food.Food_name, food.Food_Cals, foodDate)
// 	if err != nil {
// 		panic(err)
// 	}
// 	returned := fmt.Sprintf("\nInserted: %s", foodData)
// 	return returned
// }
//
