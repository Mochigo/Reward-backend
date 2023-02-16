package upload

import "mime/multipart"

// TODO考虑怎么支持对象存储，其实差不多，因为本身dst也是配置好的

type Uploader interface {
	Upload(file *multipart.FileHeader) error
}

type LocalUploader struct {
}

func (l *LocalUploader) Upload(file *multipart.FileHeader) error {
	return nil
}

type QiNiuUploader struct {
}
