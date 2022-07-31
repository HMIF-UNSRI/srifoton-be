package http

import (
	"fmt"
	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"path/filepath"
)

func (h HTTPUploadDelivery) saveFile(c *gin.Context, fileHeader *multipart.FileHeader, datatype string) (filename string, err error) {
	if fileHeader.Size > MaxFileSize {
		return filename, errorCommon.NewInvariantError("file size exceeds the maximum limit")
	}

	ext := filepath.Ext(fileHeader.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		return filename, errorCommon.NewInvariantError("only png, jpeg, jpg or pdf extension is supported")
	}

	filename = fmt.Sprintf("%s_%s%s", datatype, uuid.NewString(), ext)
	fileLocation := BasePath + "/" + filename
	err = c.SaveUploadedFile(fileHeader, fileLocation)
	return filename, err
}
