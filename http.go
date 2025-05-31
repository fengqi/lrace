package lrace

import (
	"fmt"
	"net/url"
	"strings"
)

// HttpBuildQuery 模拟PHP的http_build_query
func HttpBuildQuery(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}

	buf := strings.Builder{}
	for k, v := range params {
		switch v := v.(type) {
		case string:
			buf.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(v)))
		case int, int8, int16, int32, int64:
			buf.WriteString(fmt.Sprintf("%s=%d&", k, v))
		case float32, float64:
			buf.WriteString(fmt.Sprintf("%s=%.2f&", k, v))
		default:
			buf.WriteString(fmt.Sprintf("%s=%v&", k, url.QueryEscape(fmt.Sprintf("%v", v))))
		}
	}
	return strings.TrimRight(buf.String(), "&")
}
