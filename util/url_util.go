package util

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net/url"
	"regexp"
	"strings"
)

func UrlValidation(urlStr string) error {
	if strings.Contains(urlStr, " ") {
		return errors.New("invalid URL: Contains spaces")
	}
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return errors.New("invalid URL: " + err.Error())
	}

	// Check Scheme
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("invalid URL: Unsupported scheme")
	}

	// Check Host
	if parsedURL.Host == "" {
		return errors.New("invalid URL: Missing host")
	}

	// Check if host is only periods or starts/ends with a period
	if parsedURL.Host == "." || strings.HasPrefix(parsedURL.Host, ".") || strings.HasSuffix(parsedURL.Host, ".") {
		return errors.New("invalid URL: Hostname starts or ends with a period, or is only a period")
	}

	// Check if host contains illegal characters
	validPath := regexp.MustCompile(`^[a-zA-Z0-9\-\._~:/\?#\[\]@!\$&'\(\)\*\+,;=%]*$`)
	// Check for invalid characters in path
	if !validPath.MatchString(parsedURL.Path) {
		return errors.New("invalid URL: Path contains invalid characters")
	}

	// Check TLD
	r, err := regexp.Compile(`\.[a-zA-Z]{2,}$`)
	if err != nil {
		return errors.New("regex compilation error: " + err.Error())
	}
	if !r.MatchString(parsedURL.Host) {
		return errors.New("invalid URL: Missing or invalid TLD")
	}

	// Check if host contains illegal characters
	if strings.ContainsAny(parsedURL.Host, "<>{}|\\^~[]`") {
		return errors.New("invalid URL: Host contains illegal characters")
	}

	// Check Path
	if parsedURL.Path != "" && strings.Contains(parsedURL.Path, "..") {
		return errors.New("invalid URL: Path contains '..'")
	}

	// Check Query Parameters
	queryParams := parsedURL.Query()
	for key, values := range queryParams {
		if strings.ContainsAny(key, "<>{}|\\^~[]`") {
			return errors.New("invalid URL: Query parameters contain illegal characters")
		}
		for _, value := range values {
			if strings.ContainsAny(value, "<>{}|\\^~[]`") {
				return errors.New("invalid URL: Query parameters contain illegal characters")
			}
		}
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
