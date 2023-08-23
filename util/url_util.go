package util

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"log"
	"net/url"
)

func UrlValidation(urlStr string) error {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return errors.New("invalid URL: " + err.Error())
	}
	log.Println("parsedUrl:", parsedURL)
	log.Println("parsedUrl.Scheme:", parsedURL.Scheme)
	log.Println("parsedUrl.Host:", parsedURL.Host)
	if parsedURL.Scheme == "" {
		return errors.New("invalid URL: Missing scheme")
	}

	if parsedURL.Host == "" {
		return errors.New("invalid URL: Missing host")
	}

	_, err = url.ParseRequestURI(urlStr)
	if err != nil {
		return errors.New("invalid URL: " + err.Error())
	}

	return nil
}

func HashUrl(url string) string {
	hasher := sha1.New()
	hasher.Write([]byte(url))
	alias := base64.URLEncoding.EncodeToString(hasher.Sum(nil))[:8]
	return alias
}
