package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// create a map, the key with type string & values with type string
//type Archive map[string]string
//var notes = map[string]Archive{}

var notes = map[string]map[string]string{
	"Ella": {"2023-02-23": "Is so lonely"},
}

// get information from endpoint
func getNotes(c *gin.Context) {
	username := c.Param("username")
	now := time.Now()
	today := formatDate(now)
	//note is value of the key & found is a boolean
	if userArchive, found := notes[username]; found == true {
		todaysNote, hasNote := userArchive[today]
		if hasNote == true {
			c.IndentedJSON(http.StatusOK, todaysNote)
			return
		}
		c.IndentedJSON(http.StatusNotFound, "User has no Note for today")
		return
	}
	c.IndentedJSON(http.StatusNotFound, "Username "+username+" not found")
}

// save information at endpoint
func postNotes(c *gin.Context) {
	type payload struct {
		Username string `json:"username"`
		Text     string `json:"text"`
	}
	var newNotes []payload

	if err := c.BindJSON(&newNotes); err != nil {
		c.IndentedJSON(http.StatusCreated, err)
	}

	for _, data := range newNotes {
		now := time.Now()
		today := formatDate(now)
		notes[data.Username][today] = data.Text
	}
	c.IndentedJSON(http.StatusCreated, newNotes)

}
func formatDate(date time.Time) string {
	return date.Format("2006-01-02")
}
func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.GET("/notes/:username", getNotes)
	router.POST("/notes", postNotes)
	router.GET("/notes/:username/:day", getNotes)

	router.Run("localhost:8080")
}
