package http

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func GetDefaultHttpClient(insecureSkipVerify bool) *http.Client {
	return GetHttpClient(insecureSkipVerify, 10*time.Second, 30*time.Second, 10*time.Second, 20)
}

func GetHttpClient(insecureSkipVerify bool``, tlsHandshakeTimeout, idleConnTimeout, responseHeaderTimeout time.Duration, maxIdleConnections int) *http.Client {
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


func GetHttpPathParam(r *http.Request, template, name string) (string, error) {
	pmap, err := ParsePathParams(template, r.URL.Path)
	if err != nil {
		return "", err
	}
	if val, ok := pmap[name]; ok {
		return val, nil
	}
	return "", errors.New(fmt.Sprintf("There are no param %s in path %s", name, r.URL.Path))
}

func GetHttpRequestParam(r *http.Request, name string) ([]string, error) {
	vals, ok := r.URL.Query()[name]
	if !ok || len(vals[0]) < 1 {
		return nil, errors.New("request key not found")
	}
	ret := make([]string, 0)
	for _, v := range vals {
		ret = append(ret, string(v))
	}
	return ret, nil
}

func GetHttpRequestBody(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	} else {
		return string(body), nil
	}
}

func IsHttpContentJson(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")
	if len(contentType) == 0 {
		return false
	}
	if contentType != "application/json" {
		return false
	}
	return true
}

func ParsePathParams(template, path string) (map[string]string, error) {
	var pth string
	if strings.Contains(path, "?") {
		pth = path[:strings.Index(path, "?")]
	} else {
		pth = path
	}
	templatePaths := strings.Split(template, "/")
	pathPaths := strings.Split(pth, "/")
	if len(templatePaths) != len(pathPaths) {
		return nil, errors.New(fmt.Sprintf("pathElement length not equals to templateElement length."))
	}
	ret := make(map[string]string)
	for idx, templateElement := range templatePaths {
		pathElement := pathPaths[idx]
		if len(templateElement) > 0 && len(pathElement) > 0 {
			if templateElement[:1] == "{" && templateElement[len(templateElement)-1:] == "}" {
				tKey := templateElement[1 : len(templateElement)-1]
				ret[tKey] = pathElement
			} else if templateElement != pathElement {
				return nil, errors.New(fmt.Sprintf("Template %s not compatible with path"))
			}
		}
	}
	return ret, nil
}
