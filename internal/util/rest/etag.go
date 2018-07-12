package rest

import (
	"github.com/grafodb/grafodb/internal/util/hash"
)

type ETagger interface {
	ETag(data []byte) string
}

func NewMD5HashETagger() ETagger {
	return &md5HashETagger{}
}

type md5HashETagger struct {
}

func (e md5HashETagger) ETag(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	return hash.EncodeMd5String(data)
}

func NewSha1HashETagger() ETagger {
	return &sha1HashETagger{}
}

type sha1HashETagger struct {
}

func (e sha1HashETagger) ETag(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	return hash.EncodeSha1String(data)
}
