package mUploadUtil

import (
	"fmt"
	"github.com/sun-wenming/go-tools/mConstant"
	"github.com/sun-wenming/go-tools/mFile"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

var (
	imageAllowExts = []string{".png", ".jpeg", ".jpg"}
	//MB
	ImageMaxSize = 5
)

func AppendImageAllowExt(ext string) {
	imageAllowExts = append(imageAllowExts, ext)
}

// GetImageName 获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	//fileName := strings.TrimSuffix(name, ext)
	fileName := string(time.Now().UnixNano())
	return fileName + ext
}

// GetImagePath 获取图片路径
func GetImagePath() string {
	return mConstant.ImageSavePath
}

// GetImageFullPath 获取图片全路径
func GetImageFullPath() string {
	return mConstant.RuntimeRootPath + mConstant.ImageSavePath
}

// CheckImageExt 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := mFile.GetExt(fileName)
	for _, allowExt := range imageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize 检查图片大小
func CheckImageSize(f multipart.File) (error, bool) {
	size, err := mFile.GetSize(f)
	if err != nil {
		return err, false
	}
	return nil, size <= ImageMaxSize
}

// CheckImage 检查上传图片所需（权限、文件夹）
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = mFile.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("mFile.IsNotExistMkDir err: %v", err)
	}

	perm := mFile.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("mFile.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
