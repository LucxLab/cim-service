package message

import "time"

const CdrFileProcessCommandType = "cdr_file_process_command"

type CdrFileProcessCommand struct {
	Id               string `json:"id"`
	Type             string `json:"type"`
	NotificationTime string `json:"notification_time"`
}

func NewCdrFileProcessCommand(id string) *CdrFileProcessCommand {
	now := time.Now().Format(time.RFC3339)
	return &CdrFileProcessCommand{
		Id:               id,
		Type:             CdrFileProcessCommandType,
		NotificationTime: now,
	}
}
