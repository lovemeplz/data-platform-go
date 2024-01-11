package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lovemeplz/data-platform-go/pkg/e"
	"github.com/lovemeplz/data-platform-go/pkg/logging"
	"github.com/lovemeplz/data-platform-go/pkg/upload"
)

type File struct {
	FileName string `json:fileName`
	FileUrl  string `json:fileUrl`
}

func Upload(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("file")

	if err != nil {
		logging.Error(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.InvalidParams
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ErrorUploadCheckImageFormat
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ErrorUploadCheckImageFail
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ErrorUploadSaveImageFail
			} else {
				data["fileUrl"] = upload.GetImagePreviewUrl(imageName)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func UploadMultiple(c *gin.Context) {
	var data []File
	code := e.SUCCESS

	form, err := c.MultipartForm()

	if err != nil {
		logging.Error(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	files := form.File["upload[]"]
	for _, file := range files {

		if file == nil {
			code = e.InvalidParams
		} else {
			imageName := upload.GetImageName(file.Filename)
			fullPath := upload.GetImageFullPath()

			src := fullPath + imageName
			if !upload.CheckImageExt(imageName) {
				code = e.ErrorUploadCheckImageFormat
			} else {
				err := upload.CheckImage(fullPath)
				if err != nil {
					logging.Warn(err)
					code = e.ErrorUploadCheckImageFail
				} else if err := c.SaveUploadedFile(file, src); err != nil {
					logging.Warn(err)
					code = e.ErrorUploadSaveImageFail
				} else {
					item := File{
						FileName: "",
						FileUrl:  upload.GetImagePreviewUrl(imageName),
					}
					data = append(data, item)
				}
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
