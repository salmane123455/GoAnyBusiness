package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const RequestIDKey = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader(RequestIDKey)
		if id == "" {
			id = generateID()
		}
		c.Set(RequestIDKey, id)
		c.Header(RequestIDKey, id)
		c.Next()
	}
}

func generateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
