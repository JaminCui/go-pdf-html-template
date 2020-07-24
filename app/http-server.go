package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	maxAgeInMinutes  = 60
	timeoutInSeconds = 30
)

// RunHttpServer run server with port
func RunHttpServer(port string) {
	setGinMode()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION"},
		AllowHeaders:     []string{"utoken,x-auth-token,x-request-id,Content-Type,Accept,Origin,Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           maxAgeInMinutes * time.Minute,
	}))
	router.GET("/export/pdf", ExportPDF)
	router.GET("/export/pdfTemplate", ExportPDFTemplate)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
			"data":    "1.0.0",
		})
	})
	StartHTTPServer(router, port)
}

func setGinMode() {
	gin.SetMode(gin.DebugMode)
}

//StartHTTPServer start HTTP server
func StartHTTPServer(router *gin.Engine, port string) {
	// router.Run(":80") 这样写就可以了，下面所有代码（go1.8+）是为了优雅处理重启等动作。可根据实际情况选择。
	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  timeoutInSeconds * time.Second,
		WriteTimeout: timeoutInSeconds * time.Second,
	}

	go func() {
		// service connections
		log.Println("Start Http Server: ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to serve: ", err)
		}
	}()
}
