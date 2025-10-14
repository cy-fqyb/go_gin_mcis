package controller

import (
	"go_gin_mcis/pkg/logger"
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
		c.JSON(400, gin.H{"error": "Failed to parse multipart form", "detail": err.Error()})
		return
	}

	// 获取文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to retrieve file", "detail": err.Error()})
		return
	}
	defer file.Close()

	// 确保 uploads 目录存在
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create uploads directory", "detail": err.Error()})
		return
	}

	// 生成目标路径
	dstPath := filepath.Join(uploadDir, header.Filename)

	// 创建目标文件
	dst, err := os.Create(dstPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create destination file", "detail": err.Error(), "path": dstPath})
		return
	}
	defer dst.Close()

	// 复制内容
	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file", "detail": err.Error()})
		return
	}
	logger.Info("File uploaded successfully: " + dstPath)

	c.JSON(200, gin.H{
		"message":  "File uploaded successfully",
		"filename": header.Filename,
		"path":     dstPath,
	})
}
