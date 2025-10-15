package controller

import (
	"go_gin_mcis/pkg/result"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func init() {
	RegisterPublicRoutes(func(r *gin.RouterGroup) {
		r.POST("/upload", UploadFile)
	})
}

func UploadFile(c *gin.Context) {
	// 限制上传大小：10 MB
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		result.Fail(c, 400, "Failed to parse multipart form: "+err.Error())
		return
	}

	// 获取文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		result.Fail(c, 400, "Failed to retrieve file: "+err.Error())
		return
	}
	defer file.Close()

	// 确保 uploads 目录存在
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		result.Fail(c, 500, "Failed to create uploads directory: "+err.Error())
		return
	}

	// 生成目标路径
	dstPath := filepath.Join(uploadDir, header.Filename)

	// 创建目标文件
	dst, err := os.Create(dstPath)
	if err != nil {
		result.Fail(c, 500, "Failed to create destination file: "+err.Error())
		return
	}
	defer dst.Close()

	// 复制内容
	if _, err := io.Copy(dst, file); err != nil {
		result.Fail(c, 500, "Failed to save file: "+err.Error())
		return
	}
	result.Success(c, gin.H{
		"filename": header.Filename,
		"path":     dstPath,
	})
}
