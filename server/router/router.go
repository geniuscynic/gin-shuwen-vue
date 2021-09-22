package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	InitUserRoute(router.Group(""))

	s := &http.Server{
		Addr:           ":8889",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
