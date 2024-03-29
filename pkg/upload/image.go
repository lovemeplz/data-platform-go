package upload

import (
	"fmt"
	util "github.com/lovemeplz/data-platform-go/utils"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/lovemeplz/data-platform-go/pkg/file"
	"github.com/lovemeplz/data-platform-go/pkg/logging"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
)

func GetImagePreviewUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + GetImagePreviewPath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetImageSavePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImagePreviewPath() string {
	return setting.AppSetting.ImagePreviewPath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImageSavePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

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
