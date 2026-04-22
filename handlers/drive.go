package handlers

import (
	"net/http"

	"gin-drive-manager/services"

	"github.com/gin-gonic/gin"
)

func ListFiles(c *gin.Context) {
	service, err := services.GetDriveService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not authenticated"})
		return
	}

	folderID := c.Query("folderId")
	if folderID == "" {
		folderID = "root"
	}

	res, err := service.Files.List().
		Q("'" + folderID + "' in parents").
		Fields("files(id,name,mimeType)").
		Do()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	c.JSON(http.StatusOK, res.Files)
}
