package requests

import (
	"context"
	"errors"
	"github.com/LucxLab/cim-service/internal/upload"
	"mime/multipart"
	"net/http"
	"strings"
)

const OrganizationIdKey = "organization_id"
const UserIdKey = "user_id"
const FileKey = "file"

type Upload struct {
	// Provided through the path parameter.
	OrganizationId string

	// Extracted from the context, which is set by a dedicated middleware.
	UserId string

	// Extracted from the form, which is sent as a multipart request.
	File multipart.File

	// Related metadata of the File.
	FileHeaders *multipart.FileHeader
}

func (u *Upload) Validate() error {
	if u.OrganizationId == "" {
		return errors.New("organization_id is required")
	}
	if u.UserId == "" {
		return errors.New("user_id is required")
	}
	if u.File == nil {
		return errors.New("file is required")
	}
	return nil
}

func (u *Upload) ToFileCreation() *upload.FileCreation {
	fileName := removeFileExtension(u.FileHeaders.Filename)
	return &upload.FileCreation{
		OrganizationId: u.OrganizationId,
		UserId:         u.UserId,
		File:           u.File,
		FileSize:       u.FileHeaders.Size,
		FileName:       fileName,
	}
}

func NewUpload(ctx context.Context, request *http.Request) (*Upload, error) {
	organizationId := request.PathValue(OrganizationIdKey)
	userId := ctx.Value(UserIdKey).(string)

	file, headers, err := request.FormFile(FileKey)
	if err != nil {
		return nil, err
	}
	return &Upload{
		OrganizationId: organizationId,
		UserId:         userId,
		File:           file,
		FileHeaders:    headers,
	}, nil
}

// removeFileExtension returns the file name without the extension.
//
// Example: removeFileExtension("2025-01-01.123456.csv") returns "2025-01-01.123456"
func removeFileExtension(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) == 1 {
		return fileName
	}
	return strings.Join(parts[:len(parts)-1], ".")
}
