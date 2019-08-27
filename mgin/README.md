# go-tools

## mgin

```
func Apply(c *gin.Context) {
	getGin := swmGin.GetGin(c)
	// 1.
	// getGin.Response400Str("错误")
	// 2. 
	// err := errors.New("错误")
	// getGin.Response400(err)
	// 3.
	getGin.ResponseSuc(nil)
}

```