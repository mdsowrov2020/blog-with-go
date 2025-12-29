package util

import "encoding/base64"

func Base64UrlEncoder(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
