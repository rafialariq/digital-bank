package middlewares

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Timestamp   time.Time `json:"timestamp"`
	IPAddress   string    `json:"ip_address"`
	Method      string    `json:"method"`
	URL         string    `json:"url"`
	RequestBody string    `json:"request_body"`
	Latency     int64     `json:"latency"`
	StatusCode  int       `json:"status_code"`
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mencatat waktu mulai request
		start := time.Now()

		// Mencatat informasi dari request
		entry := LogEntry{
			Timestamp:   time.Now(),
			IPAddress:   c.ClientIP(),
			Method:      c.Request.Method,
			URL:         c.Request.URL.Path,
			RequestBody: "", // Isi dengan body request jika diperlukan
		}

		// Membaca body request jika ada
		body, err := io.ReadAll(c.Request.Body)
		if err == nil {
			entry.RequestBody = string(body)
		}

		// Melanjutkan eksekusi handler berikutnya
		c.Next()

		// Menghitung latency
		latency := time.Since(start)

		// Mencatat informasi response
		entry.Latency = latency.Milliseconds()
		entry.StatusCode = c.Writer.Status()

		// Melakukan logging ke file
		file, err := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// log.Error("Failed to open log file:", err)
			panic(err)
		} else {
			defer file.Close()

			// Mengubah entry log menjadi format JSON
			entryJSON, err := json.Marshal(entry)
			if err != nil {
				// log.Error("Failed to marshal log entry to JSON:", err)
				panic(err)
			} else {
				// Menulis log ke file
				if _, err := file.Write(entryJSON); err != nil {
					// log.Error("Failed to write log entry:", err)
					panic(err)
				}
			}
		}
	}
}
