package api

import (
	"StadiumSlotBot/internal/dto"
	"StadiumSlotBot/internal/utils"
	"net/http"
	"sync"
	"time"

	"github.com/go-errors/errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// InitServer: starts instance of GIN router
func InitServer(port string, dbConnStr string, isDebug bool, wg *sync.WaitGroup) {
	docs.SwaggerInfo.Schemes = []string{"http"}
	router := initRouter(isDebug)
	api := initAPI(router, dbConnStr)
	serverError := startServer(api, ":" + port, wg)
	if serverError != nil {
		utils.Logger.Error("Web server error: " + serverError.Error())
	}	
}

func errorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	httpResponse := dto.HTTPError{Message: "Internal server error: " + goErr.Error(), Code: http.StatusInternalServerError}
	c.AbortWithStatusJSON(500, httpResponse)
}

func initAPI(router *gin.Engine, dbConnStr string) *gin.Engine {
	api := router.Group("/api")

	//init version
	v1 := api.Group("/v1")

	//init settings controller
	settingsController := &controller{}
	settingsController.initController(v1, "settings", dbConnStr)
	//settingsController.initV1settingsController()

	return router
}

func initRouter(isDebug bool) *gin.Engine {
	if isDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
		
	router := gin.New()

	// middleware
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(gin.CustomRecovery(errorHandler))
	router.Use(static.Serve("/", static.LocalFile("web", false)))
		
	// cors
	//https://github.com/gin-contrib/cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST,PATCH,PUT,GET,DELETE"},
		AllowHeaders:     []string{"Content-Type"}, //"Accept-Encoding", "Authorization", "Cache-Control"},
		MaxAge:           time.Hour,
	}))
	
	//init swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func startServer(router *gin.Engine, listen string, wg *sync.WaitGroup) error {
	defer wg.Done()
	return router.Run(listen)
}