package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"path/filepath"
)

func (h HTTPUploadDelivery) saveFile(c *gin.Context, fileHeader *multipart.FileHeader, datatype string) (filename string, err error) {
	ext := filepath.Ext(fileHeader.Filename)
	filename = fmt.Sprintf("%s_%s%s", datatype, uuid.NewString(), ext)
	fileLocation := BasePath + "/" + filename
	err = c.SaveUploadedFile(fileHeader, fileLocation)
	return filename, err
}
