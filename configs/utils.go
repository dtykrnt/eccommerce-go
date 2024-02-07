package configs

import (
	"encoding/base64"
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	"gorm.io/gorm"
)

type paginate struct {
	limit int
	page  int
}

func Pagination(limit int, page int) *paginate {
	return &paginate{limit: limit, page: page}
}

func (p *paginate) Result(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit
	return db.Offset(offset).Limit(p.limit)
}

func SaveImageBase64ToLocal(b64, filepath string) (*string, error) {
	// Extract the MIME type from the data URI
	mimeRegex := regexp.MustCompile("^data:([a-zA-Z0-9]+/[a-zA-Z0-9]+)")
	matches := mimeRegex.FindStringSubmatch(b64)
	if len(matches) < 2 {
		return nil, errors.New("invalid base64 string")
	}
	mimeType := matches[1]

	// Remove the prefix (e.g., "data:image/png;base64,") from the base64 string
	parts := strings.Split(b64, ",")
	if len(parts) != 2 {
		return nil, errors.New("invalid base64 string")
	}
	img := parts[1]

	// Decode the base64 string
	decoded, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		return nil, err
	}

	// Generate a timestamp for the image file name
	current := time.Now().Format("2006-01-02_15-04-05")

	// Determine the file extension based on the MIME type
	var fileExt string
	switch mimeType {
	case "image/png":
		fileExt = ".png"
	case "image/jpeg":
		fileExt = ".jpg"
	case "image/jpg":
		fileExt = ".jpg"
	// Add cases for other supported MIME types here
	default:
		return nil, errors.New("unsupported MIME type")
	}

	// Construct the file path
	imagePath := filepath + "/" + filepath + current + fileExt

	// Write the decoded data to the file
	err = os.WriteFile("assets/"+imagePath, decoded, 0644)
	if err != nil {
		return nil, err
	}

	return &imagePath, nil
}

func LoadImage(img string) *string {
	localPath := "http://localhost:3000/assets/" + img
	return &localPath
}
