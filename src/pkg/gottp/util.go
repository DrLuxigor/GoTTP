package gottp

var contentTypes = map[string]string{
	".html":  "text/html",
	".css":   "text/css",
	".js":    "text/javascript",
	".json":  "application/json",
	".png":   "image/png",
	".jpg":   "image/jpeg",
	".jpeg":  "image/jpeg",
	".gif":   "image/gif",
	".webp":  "image/webp",
	".tiff":  "image/tiff",
	".ico":   "image/x-icon",
	".svg":   "image/svg+xml",
	".aac":   "audio/aac",
	".mp3":   "audio/mpeg",
	".wav":   "audio/wav",
	".mp4":   "video/mp4",
	".webm":  "video/webm",
	".csv":   "text/csv",
	".htm":   "text/html",
	".pdf":   "application/pdf",
	".7z":    "application/x-7z-compressed",
	".zip":   "application/zip",
	".xml":   "application/xml",
	".ttf":   "font/ttf",
	".otf":   "font/otf",
	".woff":  "font/woff",
	".woff2": "font/woff2",
}

func GetContentType(fileExtension string) string {
	if val, ok := contentTypes[fileExtension]; ok {
		return val
	}
	return "application/octet-stream"
}
