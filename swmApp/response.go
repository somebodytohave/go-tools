package swmApp

import (
	"github.com/gin-gonic/gin"
	"github.com/sun-wenming/go-tools/e"
	"github.com/sun-wenming/go-tools/util"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"strings"
)

// Gin 实体
type Gin struct {
	C *gin.Context
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"msg" example:"ok" `
	Data    string `json:"data" example:"null"`
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

// ResponseFail 返回失败
func (g *Gin) ResponseFail() {
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ERROR,
		"msg":  e.GetMsg(e.ERROR),
		"data": nil,
	})
	return
}

// ResponseFailErrCode 返回失败
func (g *Gin) ResponseFailErrCode(errCode int) {
	errMsg := "code : " + strconv.Itoa(errCode) + "msg : " + e.GetMsg(errCode)
	MarkError(errMsg)

	g.C.JSON(http.StatusOK, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": nil,
	})
	return
}

// ResponseFailMsg 返回失败
//func (g *Gin) ResponseFailMsg(msg string) {
//	MarkError(msg)
//	g.C.JSON(http.StatusOK, gin.H{
//		"code": http.StatusBadRequest,
//		"msg":  msg,
//		"data": nil,
//	})
//	return
//}

// ResponseFailError 返回自定义的错误类型
func (g *Gin) ResponseFailError(error cmutil.Error) {
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
	fieldName, _ := cmutil.GetTrans().T(jsonKey)
	msg := strings.Replace(errs[0].Translate(cmutil.GetTrans()), jsonKey, fieldName, -1)
	//jsonKey = jsonKey[2 : len(jsonKey)-2]
	//fmt.Println(jsonKey, ":", msg)

	MarkError(msg)
	g.C.JSON(http.StatusOK, gin.H{
		"code": e.ErrorInvalidParams,
		"msg":  msg,
		"data": nil,
	})
	return
}
