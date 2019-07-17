package swmUploadUtil

import (
	"fmt"
	"github.com/sun-wenming/go-tools/constant"
	"github.com/sun-wenming/go-tools/file"
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
	return constant.ImageSavePath
}

// CheckImageExt 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range imageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize 检查图片大小
func CheckImageSize(f multipart.File) (error, bool) {
	size, err := file.GetSize(f)
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

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
