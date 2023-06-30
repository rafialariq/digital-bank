package middlewares

import (
	"encoding/json"
	// "io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogEntry struct {
	Timestamp  string `json:"timestamp"`
	IPAddress  string `json:"ip_address"`
	Method     string `json:"method"`
	URL        string `json:"url"`
	Latency    int64  `json:"latency"`
	StatusCode int    `json:"status_code"`
}

func LogMiddleware() gin.HandlerFunc {
	// set JSON format
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// open file
	file, err := os.OpenFile("log.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logrus.Fatal("failed to open log file: ", err)
	}
	defer file.Close()

	// set output destination
	logrus.SetOutput(file)

	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		latency := time.Since(start)

		logEntry := LogEntry{
			Timestamp:  time.Now().Format(time.RFC3339),
			Method:     ctx.Request.Method,
			IPAddress:  ctx.ClientIP(),
			URL:        ctx.Request.URL.Path,
			Latency:    int64(latency),
			StatusCode: ctx.Writer.Status(),
		}

		logData, err := json.MarshalIndent(logEntry, "", "\t")
		if err != nil {
			logrus.Fatal("failed to marshal log entry: ", err)
			return
		}

		// logData = append(logData, []byte("\n")...)

		logrus.Info(string(logData))
	}
}
