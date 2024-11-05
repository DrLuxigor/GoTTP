package gottp

import (
	"strings"
)

func GetContentType(fileExtension string) string {
	switch strings.ToLower(fileExtension) {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	case ".json":
		return "application/json"
	case ".png":
		return "image/png"
	case ".jpg":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".tiff":
		return "image/tiff"
	case ".ico":
		return "image/x-icon"
	case ".svg":
		return "image/svg+xml"
	case ".aac":
		return "audio/aac"
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".mp4":
		return "video/mp4"
	case ".webm":
		return "video/webm"
	case ".csv":
		return "text/csv"
	case ".htm":
		return "text/html"
	case ".pdf":
		return "application/pdf"
	case ".7z":
		return "application/x-7z-compressed"
	case ".zip":
		return "application/zip"
	case ".xml":
		return "application/xml"
	case ".ttf":
		return "font/ttf"
	case ".otf":
		return "font/otf"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	}
	return "text/plain"
}
