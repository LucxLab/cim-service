package message

import "time"

const CdrFileUploadedType = "cdr_file_uploaded"

type CdrFileUploadedMessage struct {
	Id               string `json:"id"`
	Type             string `json:"type"`
	NotificationTime string `json:"notification_time"`
}

func ToCdrFileUploadedMessage(id string) *CdrFileUploadedMessage {
	now := time.Now().Format(time.RFC3339)
	return &CdrFileUploadedMessage{
		Id:               id,
		Type:             CdrFileUploadedType,
		NotificationTime: now,
	}
}
