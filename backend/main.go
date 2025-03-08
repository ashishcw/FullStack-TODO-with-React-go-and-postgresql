package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type ToDo struct {
	Id               int    `json:"id"`
	TaskTitle        string `json:"tasktitle"`
	CompletionStatus string `json:"completionstatus"`
}

var dbTempInstance *sql.DB

// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }

func main() {
	connectToDatabase()

	newEcho := echo.New()

	// wrap the router with CORS and JSON content type middlewares
	enableCORS(jsonContentTypeMiddleware(newEcho))

	newEcho.GET("/api/v1/users/:id", getUsers)
	//newEcho.GET("/users/:id", getUser)
	//newEcho.GET("/api/v1/users", getUser)

	fmt.Print("Hello World")
	newEcho.Logger.Fatal(newEcho.Start("localhost:8000"))

}

func connectToDatabase() {
	dbTempInstance, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}
	defer dbTempInstance.Close()

	// create table if not exists
	_, err = dbTempInstance.Exec("CREATE TABLE IF NOT EXISTS todotasklist (id SERIAL PRIMARY KEY, tasktitle TEXT, completionstatus TEXT)")
	if err != nil {
		log.Fatal(err)
	}

}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// e.GET("/api/v1/users", getUsers)
func getUsers(c echo.Context) error {
	// if dbTempInstance == nil {
	// 	log.Fatal("Db instance can not be null")
	// }

	id := c.Param("id")

	return c.String(http.StatusOK, id)

	// rows, err := dbTempInstance.Query("SELECT * FROM todotasklist")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// users := []ToDo{} // array of users
	// for rows.Next() {
	// 	var u ToDo
	// 	if err := rows.Scan(&u.Id, &u.TaskTitle, &u.CompletionStatus); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	users = append(users, u)
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// var jsonResponse = json.NewEncoder(c.Response().Writer).Encode(users)
	// if jsonResponse == nil {
	// 	return err
	// }
	// return jsonResponse
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
