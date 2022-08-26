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
	router := gin.Default()

	router.GET("/", welcomePage)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumsByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func welcomePage(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Welcome to Vintage recordings")
}
