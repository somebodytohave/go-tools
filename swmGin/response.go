package swmGin

import (
	"github.com/gin-gonic/gin"
	"github.com/sun-wenming/go-tools/e"
	"github.com/sun-wenming/go-tools/swmUtil"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
)

// Gin 实体
type Gin struct {
	C *gin.Context
}

// GetGin 获取Gin
func GetGin(c *gin.Context) Gin {
	return Gin{c}
}

// Response 返回的数据
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}

// ResponseSuc 返回成功
func (g *Gin) ResponseSuc(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
	return
}

// Response400 返回失败
func (g *Gin) Response400(err error) {
	MarkError(err.Error())
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ERROR400,
		"msg":  err.Error(),
		"data": nil,
	})
	return
}

// Response500 返回失败
func (g *Gin) Response500(err error) {
	MarkError(err.Error())
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ERROR500,
		"msg":  err.Error(),
		"data": nil,
	})
	return
}

// ResponseFailError 返回自定义的错误类型
func (g *Gin) ResponseFailError(error swmUtil.Error) {
	msg := error.Error()
	MarkError(msg)
	g.C.JSON(http.StatusOK, gin.H{
		"code": error.Code(),
		"msg":  msg,
		"data": nil,
	})
	return
}

// ResponseFailValidParam 验证参数错误
func (g *Gin) ResponseFailValidParam(err error) {
	errs := err.(validator.ValidationErrors)

	jsonKey := errs[0].Field()
	fieldName, _ := swmUtil.GetTrans().T(jsonKey)
	msg := strings.Replace(errs[0].Translate(swmUtil.GetTrans()), jsonKey, fieldName, -1)
	//jsonKey = jsonKey[2 : len(jsonKey)-2]
	//fmt.Println(jsonKey, ":", msg)

	MarkError(msg)
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ERROR400,
		"msg":  msg,
		"data": nil,
	})
	return
}