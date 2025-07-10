package main

import (
	"UrlShort/models"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	conn := models.ConnectDB()
	
	defer conn.Close(context.Background())

	e := echo.New()

	e.Use(middleware.CORS()) // allow all origins for dev
	e.POST("/api/shorten", func(c echo.Context) error{
		//this is what will be found at users where the front end sends request at 
		type Link struct{
			Url string `json:"url"`
		}
		//
		// //if the json is not a string or url
		var link Link
		err := c.Bind(&link)
		if err != nil{
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid json",
			})
		}
		short := link.Url
		models.InsertInto(conn, short, "jdioejidej")
		log.Println(short)
		//this is what the front end will get and "test" will be the shortend url 
		return c.JSON(http.StatusOK, map[string]interface{}{
			"short_url": short,
		})
	})
	e.Logger.Fatal(e.Start(":8081"))
}





