package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
/**Struct tags such as json:"artist" specify what a field’s name should be
when the struct’s contents are serialized into JSON. Without them, the JSON
would use the struct’s capitalized field names – a style not as common in JSON.
*/
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Initialize a Gin router using Default.
	router := gin.Default()
	router.GET("/", welcomePage)
	// Use the GET function to associate the GET HTTP method and /albums path with a handler function.
	router.GET("/albums", getAlbums)
	// Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func welcomePage(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Welcome to Vintage recordings")
}
