package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"template-ulamm-backend-go/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type blacklist struct {
	bodyRequest []string
}

type metricMeta struct {
	Path         string `json:"path"`
	Latency      string `json:"latency"`
	RequestBody  any    `json:"request_body"`
	ResponseBody any    `json:"response_body"`
	ClientIP     string `json:"client_ip"`
	Method       string `json:"method"`
	StatusCode   int    `json:"status_code"`
	BodySize     string `json:"body_size"`
	UserAgent    string `json:"user_agent"`
}

func logMetric(blacklist blacklist) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeStart := time.Now()
		path := ctx.Request.URL.Path
		rawQuery := ctx.Request.URL.RawQuery

		// Read Body
		var buf bytes.Buffer
		temp := io.TeeReader(ctx.Request.Body, &buf)
		requestBody, _ := io.ReadAll(temp)

		// Bypass Writer and Request Body
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		ctx.Request.Body = io.NopCloser(&buf)

		// Process request
		ctx.Next()

		// End Time
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		bodySize := ctx.Writer.Size()
		userAgent := ctx.Request.UserAgent()
		responseBody := blw.body.String()
		if rawQuery != "" {
			path = path + "?" + rawQuery
		}

		data := metricMeta{
			Path:       path,
			ClientIP:   clientIP,
			Method:     method,
			StatusCode: statusCode,
			BodySize:   fmt.Sprintf("%d byte", bodySize),
			UserAgent:  userAgent,
		}

		var parsedRequest map[string]any
		if err := json.Unmarshal(requestBody, &parsedRequest); err == nil {
			for _, val := range blacklist.bodyRequest {
				delete(parsedRequest, val)
			}

			data.RequestBody = parsedRequest
		} else {
			data.RequestBody = string(requestBody)
		}

		var parsedResponse map[string]any
		if err := json.Unmarshal([]byte(responseBody), &parsedResponse); err == nil {
			data.ResponseBody = parsedResponse
		} else {
			data.ResponseBody = responseBody
		}

		data.Latency = fmt.Sprintf("%v", time.Since(timeStart))

		infoBytes, _ := json.MarshalIndent(data, "", "    ")

		utils.GetLogger().Info(string(infoBytes))
	}
}
