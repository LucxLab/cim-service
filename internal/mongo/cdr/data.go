package cdr

import cim "github.com/LucxLab/cim-service/internal"

type uploadData struct {
	Status cim.UploadStatus `bson:"status"`
}
