package swmGin

import (
	"github.com/gin-gonic/gin"
	"github.com/sun-wenming/go-tools/swmRegValidUtil"
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
func (g *Gin) Response(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  msg,
		"data": data,
	})

	return
}

// ResponseSuc 返回成功
func (g *Gin) ResponseSuc(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": data,
	})
	return
}

// Response400Str 返回自定义字错误内容
func (g *Gin) Response400Str(errStr string) {
	MarkError(errStr)
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ERROR400,
		"msg":  errStr,
		"data": nil,
	})
	return
}

// Response400 返回失败
func (g *Gin) Response400(err error) {
	MarkError(err.Error())
	g.C.JSON(http.StatusOK, gin.H{
		"code": http.StatusBadRequest,
		"msg":  err.Error(),
		"data": nil,
	})
	return
}

// Response500 返回失败
func (g *Gin) Response500(err error) {
	MarkError(err.Error())
	g.C.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  err.Error(),
		"data": nil,
	})
	return
}

// ResponseFailError 返回自定义的错误类型
func (g *Gin) ResponseFailError(error Error) {
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
	fieldName, _ := swmRegValidUtil.GetTrans().T(jsonKey)
	msg := strings.Replace(errs[0].Translate(swmRegValidUtil.GetTrans()), jsonKey, fieldName, -1)
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
