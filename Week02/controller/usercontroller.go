package controller

import (
	"errors"
	"geek_example/week_02/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type UserRequest struct {
	KeyWord int `json:"keyWord"  binding:"required"`
}

const (
	demotion  = "服务器开小差了"
	emptyTips = "数据为空"
)

//获取用户信息
func GetUserInfoHandler(ctx *gin.Context) {
	var request UserRequest
	if err := ctx.BindJSON(&request); nil != err {
		log.Printf("searh faild error is %+v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": demotion})
	}

	person, err := service.GetUserInfoWithoutSomeThing(request.KeyWord)

	if err != nil {
		//找不到记录不能当做异常用日志记录,给用户端暴露数据为空的提示即可
		if errors.Is(errors.Unwrap(err), gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, emptyTips)
			return
		}

		//如果是其他的异常错误,日志记录
		log.Printf("%+v", err)
		ctx.JSON(http.StatusInternalServerError, demotion)
		return
	}

	ctx.JSON(http.StatusOK, person)
}
