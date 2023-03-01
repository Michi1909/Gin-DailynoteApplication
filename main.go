package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

//this func create JSON from the slice of notes 
func getNotes(c *gin.Context) {
	username := c.Param("username")
	date := c.Param("date")
	now := time.Now()
	today := formatDate(now)

	//userArchive is value of the key & found is a boolean
	//when there is a userArchive for the given username & there is a note with date of today
	// Context.IndentedJSON serialize
	//the todayNote into JSON and add it to the response
	if userArchive, found := notes[username]; found == true {
		todaysNote, hasNote := userArchive[today]
		if hasNote == true {
			c.IndentedJSON(http.StatusOK, todaysNote)
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
	var newNotes []payload
	
	//c.BindJSON bind the request body to newNotes
	if err := c.BindJSON(&newNotes); err != nil {
		c.IndentedJSON(http.StatusCreated, err)
	}

	//Add the newNotes to the notes
	for _, data := range newNotes {
		now := time.Now()
		today := formatDate(now)
		notes[data.Username][today] = data.Text
	}
	c.IndentedJSON(http.StatusCreated, newNotes)

}
//date func 
func formatDate(date time.Time) string {
	return date.Format("2006-01-02")
}
//associate a handler with an HTTP method-and-path
//route request to a single path based on the methode the client is using
func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.GET("/notes/:username", getNotes)
	router.POST("/notes", postNotes)
	router.GET("/notes/:username/:day", getNotes)

	router.Run("localhost:8080")
}

//create a database(Prio 2)
//create an endpoint to check the history of the other days (Prio 1)