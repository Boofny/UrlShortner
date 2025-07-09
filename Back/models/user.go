package modles

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

func InsertInto(conn *pgx.Conn, short_url string, long_url string) string{
	_, err := conn.Exec(context.Background(), "insert into urls (short_url, long_url) values ($1, $2)", short_url, long_url)

	if err != nil {
		panic(err)
	}
	return "Inserted long and short"
}

func GetLongId(conn *pgx.Conn, short_url string)string{


	long_url_from_DB := "Test"
	return  long_url_from_DB
}


//just use this function as an example moving forward 
// func CreateTable(conn *pgx.Conn) string{
// 	_, err := conn.Exec(context.Background(), "CREATE TABLE urls (id SERIAL PRIMARY KEY, short_url TEXT NOT NULL, long_url TEXT NOT NULL)")
//
// 	if err != nil {
// 		panic(err)
// 	}
// 	return "Created table"
// }

// use this code to get the long url from the short one 
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
