package routes

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/prabhatsharma/mango/pkg/auth"
	"github.com/prabhatsharma/mango/pkg/meta"
	"github.com/prabhatsharma/mango/pkg/mindex"
)

// SetSetupRoutes sets up all gi HTTP API endpoints that can be called by front end
func SetRoutes(r *gin.Engine) {

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "authorization", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/healthz", meta.GetHealthz)

	//Standard APIs
	AuthenticatedAPIs := r.Group("/api")
	AuthenticatedAPIs.Use(gzip.Gzip(gzip.DefaultCompression))
	if os.Getenv("JWK_URL") != "" {
		AuthenticatedAPIs.Use(auth.AuthorizeJWT())
	}

	AuthenticatedAPIs.GET("/_list", mindex.ListIndexes)
	AuthenticatedAPIs.PUT("/:indexName", mindex.CreateIndex)
	AuthenticatedAPIs.GET("/:indexName", mindex.GetIndexDetails)
	AuthenticatedAPIs.DELETE("/:indexName", mindex.DeleteIndex)

	AuthenticatedAPIs.PUT("/:indexName/document", mindex.InsertDocument)
	AuthenticatedAPIs.DELETE("/:indexName/:docID", mindex.DeleteDocument)

	AuthenticatedAPIs.POST("/:indexName/_search", mindex.SearchIndex)

}
