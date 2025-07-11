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

	// e.GET("/:code", func(c echo.Context) error {
	// 	code := c.Param("code")
	// 	log.Println(code)
	// 	Long_Url := "https://www.youtube.com/"
	// 	if Long_Url == ""{
	// 		return c.String(http.StatusNotFound, "Url not found")
	// 	}
	//
	// 	return c.Redirect(http.StatusFound, Long_Url)
	// })


	gets := GetReq(e)
	log.Println(gets)
	e.Logger.Fatal(e.Start(":8081"))
}

func GetReq(e *echo.Echo) string{
	//come back here for notes will need to fill the table full of links and add a case where the link doesnt 
	//exist and makes a random code and adds it to the db and now the code will exist for real in the db
	db := map[string]string{
		"yt": "https://www.youtube.com/",
		"yt2": "https://www.youtube.com/watch?v=5Welk51oDWs&t=175s",
	}
	e.GET("/:code", func(c echo.Context) error {
		code := c.Param("code")
		log.Println(code)
		Long_Url := db[code]
		if Long_Url == ""{
			return c.String(http.StatusNotFound, "Url not found")
		}

		return c.Redirect(http.StatusFound, Long_Url)
	})
	return "Good get Request"
}
