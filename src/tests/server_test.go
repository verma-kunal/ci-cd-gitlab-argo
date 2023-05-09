package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/devops-projects6943118/go-rest-api/src"
	"gitlab.com/devops-projects6943118/go-rest-api/src/albums"
)

// For getting the list of albums:
func TestGetAlbums(t *testing.T) {

	// create a new Gin router & register the "getAlbums" function:
	router := gin.Default()
	router.GET("/albums", albums.GetAlbums)

	// creates a new HTTP GET request with the path "/albums"
	req, err := http.NewRequest("GET", "/albums", nil)
	assert.NoError(t, err) // checks if the request has no errors (err == nil)

	//  creates a new "httptest.ResponseRecorder", which will capture the response from the server & store in "resp"
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// check that the response code in "resp" matches http.StatusOK (which is 200)
	assert.Equal(t, http.StatusOK, resp.Code)

	// json.NewDecoder - creates a new JSON decoder that will read from the response body in "resp"
	// Decode - converts the JSON response back to the slice (deserialized)
	var albums []src.Album
	err = json.NewDecoder(resp.Body).Decode(&albums)
	assert.NoError(t, err)

	// check that the response contains the expected number of albums
	assert.Len(t, albums, 3)
}

// for adding a new album:
func TestPostAlbums(t *testing.T) {
	router := gin.Default()
	router.POST("/albums", albums.PostAlbums)

	// create a new album to post to the server:
	newAlbum := src.Album{
		ID:     "4",
		Title:  "Test",
		Artist: "Kunal Verma",
		Price:  10.00,
	}

	// json.Marshal - to serialize the newAlbum object into JSON format
	newAlbumJSON, err := json.Marshal(newAlbum)
	assert.NoError(t, err)

	// make a new POST request using the newly added JSON album:
	req, err := http.NewRequest("POST", "/albums", bytes.NewBuffer(newAlbumJSON))
	assert.NoError(t, err)

	// save the response in "resp"
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// check for the equality of status codes:
	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdAlbum src.Album
	err = json.NewDecoder(resp.Body).Decode(&createdAlbum)
	assert.NoError(t, err)

	// check that the created album matches the input:
	assert.Equal(t, newAlbum.Title, createdAlbum.Title)
	assert.Equal(t, newAlbum.Artist, createdAlbum.Artist)
	assert.Equal(t, newAlbum.Price, createdAlbum.Price)
}