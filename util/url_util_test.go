package util

import (
	"testing"
)

func TestUrlValidation(t *testing.T) {
	url := "https://github.com/"
	err := UrlValidation(url)
	if err != nil {
		t.Errorf("Expected valid URL, but got error: %s", err.Error())
	}

	url = "http://github.com/"
	err = UrlValidation(url)
	if err != nil {
		t.Errorf("Expected valid URL, but got error: %s", err.Error())
	}

	url = "http://github.app/"
	err = UrlValidation(url)
	if err != nil {
		t.Errorf("Expected valid URL, but got error: %s", err.Error())
	}

	url = "https://go.dev/"
	err = UrlValidation(url)
	if err != nil {
		t.Errorf("Expected valid URL, but got error: %s", err.Error())
	}

	url = "https://www.google.com/search?q=calculator&oq=&aqs=chrome.1.69i57j69i59j69i65j69i60l3j69i65l2.1071j0j4&sourceid=chrome&ie=UTF-8"
	err = UrlValidation(url)
	if err != nil {
		t.Errorf("Expected valid URL, but got error: %s", err.Error())
	}

	url = "github-com/"
	err = UrlValidation(url)
	if err == nil {
		t.Errorf("Expected invalid URL, but got no error")
	}
}
