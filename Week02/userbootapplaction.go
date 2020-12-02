package Week02

import (
	"geek_example/week_02/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.POST("/userinfo", controller.GetUserInfoHandler)

	route.Run(":8898")
}
