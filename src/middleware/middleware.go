package middleware

import (
	"bytes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)

		bw := &bodyWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bw

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()
		log.Printf("Response status: %d, duration: %v, body: %s", statusCode, duration, bw.body.String())
	}
}
