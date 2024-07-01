package ip

import (
	"context"
	"log"
	"net/http"
	"strings"
)

// GetRemoteIP
//
//	@Description: 获取IP
//	@param ctx
//	@param request
//	@return string
func GetRemoteIP(ctx context.Context, request *http.Request) string {
	customIpAddresses := request.Header.Get("Render-X-Forwarded-For")
	if customIpAddresses != "" {
		log.Printf("customIpAddresses from Render-X-Forwarded-For: %s", customIpAddresses)
		// 可能存在多个地址，取第一个ip地址
		addresses := strings.Split(customIpAddresses, ",")
		customIpAddresses = strings.TrimSpace(addresses[0])
		log.Printf("get remote ip is: %s", customIpAddresses)
		return customIpAddresses
	}

	customIpAddresses = request.Header.Get("X-Forwarded-For")
	if customIpAddresses != "" {
		log.Printf("customIpAddresses from X-Forwarded-For: %s", customIpAddresses)
		return customIpAddresses
	}

	customIpAddresses = request.RemoteAddr
	log.Printf("get remote ip from RemoteAddr, is: %s", customIpAddresses)
	return customIpAddresses
}
