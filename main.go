package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

/*
variable notes is a map of String keys which has
values of another map of string keys to string values
*/
var notes = map[string]map[string]string{

	/*
		Ella is a key, the date is an key, text is the value
	*/
	"Ella": {"2023-02-23": "Is so lonely"},
}

// getNote returns the note of the user by date
func getNote(c *gin.Context) {
	username := c.Param("username")
	date := c.Param("date")
	now := time.Now()
	today := formatDate(now)
	if date == "" {
		//if the frontend did not send a date, default to today
		date = today
	}

	//userArchive is value of the key & found is a boolean
	//when there is a userArchive for the given username & there is a note with date of today
	// Context.IndentedJSON serialize
	//the todayNote into JSON and add it to the response
	if userArchive, found := notes[username]; found == true {
		note, hasNote := userArchive[date]
		if hasNote == true {
			c.IndentedJSON(http.StatusOK, note)
			return
		}
		//when there were no notes for today for the user
		c.IndentedJSON(http.StatusNotFound, "User has no Note for today")
		return
	}
	//when the user do not exist
	c.IndentedJSON(http.StatusNotFound, "Username "+username+" not found")
}

// handler for the endpoint to store slices of notes
func postNotes(c *gin.Context) {

	//A payload struct to map the entire payload
	//in go better to read the json and a way of validation
	type payload struct {
		Username string `json:"username"`
		Text     string `json:"text"`
	}
	var newNote payload

	//c.BindJSON bind the request body to newNotes
	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusCreated, err)
	}

	//Add the newNotes to the notes

	now := time.Now()
	today := formatDate(now)
	notes[newNote.Username][today] = newNote.Text

	c.IndentedJSON(http.StatusCreated, newNote)

}

// date func
func formatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

// associate a handler with an HTTP method-and-path
// route request to a single path based on the methode the client is using
func main() {
	router := gin.Default()
	router.GET("/note", getNote)
	router.GET("/note/:username", getNote)
	router.POST("/note/", postNotes)
	router.GET("/note/:username/:date", getNote)

	database, err := InitDB()

	if err != nil {
		panic(err)
	}

	defer database.Close()

	router.Run("localhost:8080")
}

func InitDB() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", "postgres", "secure_pass_here", "dailynotes", "localhost", "5432")
	var err error
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

//create an endpoint to check the history of the other days (Prio 1)
//create a database(Prio 2)
//https://go.dev/doc/tutorial/database-access
