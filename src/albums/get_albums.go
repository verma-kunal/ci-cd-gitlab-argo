package albums

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gitlab.com/devops-projects6943118/ci-cd/go-rest-api/src"
)

// get list of albums:
func GetAlbums(ctx *gin.Context) {
	statusCode := http.StatusOK // code 200
	ctx.IndentedJSON(statusCode, src.Albums)
}

