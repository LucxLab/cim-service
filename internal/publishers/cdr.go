package publishers

import "context"

// Cdr defines methods that a CDR publisher must implement to publish messages to a message broker.
type Cdr interface {
	// PublishFileProcessCommand publishes a message to indicate that a CDR file is ready to be processed.
	PublishFileProcessCommand(ctx context.Context, fileMetadataId string) error
}
