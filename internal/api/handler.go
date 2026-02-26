package api

import (
	"fastmail/internal/service"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	emailService *service.EmailService
}

func NewHandler(svc *service.EmailService) *Handler {
	return &Handler{
		emailService: svc,
	}
}

// SendEmail handles the email sending request
// POST /api/v1/send
func (h *Handler) SendEmail(c *gin.Context) {
	// Parse multipart form
	// Max 32MB
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	to := c.PostForm("to")
	subject := c.PostForm("subject")
	body := c.PostForm("body")

	if to == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recipient 'to' is required"})
		return
	}

	// Parse 'to' field (comma separated)
	recipientsRaw := strings.Split(to, ",")
	var recipients []string
	for _, r := range recipientsRaw {
		trimmed := strings.TrimSpace(r)
		if trimmed != "" {
			recipients = append(recipients, trimmed)
		}
	}

	if len(recipients) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "valid recipient 'to' is required"})
		return
	}

	// Handle attachments
	form, _ := c.MultipartForm()
	files := form.File["attachments"]
	var attachmentPaths []string

	if len(files) > 0 {
		// Create a temporary directory for this request
		tempDir, err := os.MkdirTemp("", "fastmail-attachments-*")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temp directory"})
			return
		}
		// Clean up temp dir after sending (or on error)
		defer os.RemoveAll(tempDir)

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			dst := filepath.Join(tempDir, filename)

			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save attachment %s", filename)})
				return
			}
			attachmentPaths = append(attachmentPaths, dst)
		}
	}

	// Call service
	if err := h.emailService.SendEmail(recipients, subject, body, attachmentPaths); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send email: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

// Health checks the health of the service
// GET /health
func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
