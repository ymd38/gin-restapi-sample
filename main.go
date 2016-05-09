package main

import (
	. "gin-restapi-sample/db"
	"gin-restapi-sample/issue"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	InitDB()
	defer Dbm.Db.Close()

	router.GET("/api/issues", issue.List)
	router.GET("/api/issues/:id", issue.Show)

	router.Run(":8000")
}
