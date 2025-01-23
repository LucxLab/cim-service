package cdr

type Publisher interface {
	PublishUploadCreated(upload *Upload) error
}
