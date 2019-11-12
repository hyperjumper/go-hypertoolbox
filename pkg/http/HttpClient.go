package http

import (
	"crypto/tls"
	"net/http"
	"time"
)

func GetDefaultHttpClient(insecureSkipVerify bool) *http.Client {
	return GetHttpClient(insecureSkipVerify, 10*time.Second, 30*time.Second, 10*time.Second, 20)
}

func GetHttpClient(insecureSkipVerify bool, tlsHandshakeTimeout, idleConnTimeout, responseHeaderTimeout time.Duration, maxIdleConnections int) *http.Client {
	transport := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		TLSHandshakeTimeout:   tlsHandshakeTimeout,
		MaxIdleConns:          maxIdleConnections,
		IdleConnTimeout:       idleConnTimeout,
		ResponseHeaderTimeout: responseHeaderTimeout,
	}
	return &http.Client{
		Transport: transport,
	}
}
