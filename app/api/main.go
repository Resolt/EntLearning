package main

import (
	"entlearning/pkg/db/ent/schema"

	"github.com/gin-gonic/gin"
)

func main() {
	user := schema.User{}
	_ = user.Config()
	r := gin.Default()
	r.Run()
}
