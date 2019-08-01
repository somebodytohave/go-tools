package mfile

import (
	"mime/multipart"
	"strings"
)

var (
	imageAllowExts = []string{".png", ".jpeg", ".jpg"}
	//MB
	ImageMaxSize = 5
)

func AppendImageAllowExt(ext string) {
	imageAllowExts = append(imageAllowExts, ext)
}

// CheckImageExt 验证图片后缀
func CheckImageExt(fileName string) bool {
	ext := GetExt(fileName)
	for _, allowExt := range imageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImageSize 验证图片大小
func CheckImageSize(f multipart.File) (error, bool) {
	size, err := GetSize(f)
	if err != nil {
		return err, false
	}
	return nil, size <= ImageMaxSize
}
