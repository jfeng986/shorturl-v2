package util

import (
	"testing"
)

func TestUrlValidation(t *testing.T) {
	testCases := []struct {
		url            string
		expectingError bool
	}{
		{"https://github.com/", false},
		{"http://github.com/", false},
		{"http://github.app/", false},
		{"https://go.dev/", false},
		{"https://ss-link.netlify.app/", false},
		{"https://www.google.com/search?q=calculator&oq=&aqs=chrome.1.69i57j69i59j69i65j69i60l3j69i65l2.1071j0j4&sourceid=chrome&ie=UTF-8", false},

		// 以下是无效的 URL 用例
		{"", true},                           // empty string
		{"http", true},                       // lack of '//'
		{"http://", true},                    // lack of host
		{"http://github .com", true},         // contains space
		{"http://github.com:abc", true},      // port contains non-numeric characters
		{"https:/github.com", true},          // only one '/'
		{"http//github.com", true},           // lack of ':'
		{"http:/github.com", true},           // only one '/'
		{"http:///github.com", true},         // extra '/'
		{"github.com", true},                 // lack of scheme
		{"http://github/com", true},          // lack of '.'
		{"http://github\\.com", true},        // contains '\\'
		{"http:://github.com", true},         // contains two ':'
		{"http://github#com", true},          // contains '#'
		{"http://.com", true},                // host is only a period
		{"http://com.", true},                // host ends with a period
		{"ftp://github.com", true},           // unsupported scheme
		{"http:///github.com", true},         // host is empty
		{"http://github.com/<>", true},       // contains '<>'
		{"http://github.com//../test", true}, // contains '..' in path

	}

	for _, testCase := range testCases {
		err := UrlValidation(testCase.url)
		if testCase.expectingError && err == nil {
			t.Errorf("Expected invalid URL for '%s', but got no error", testCase.url)
		}
		if !testCase.expectingError && err != nil {
			t.Errorf("Expected valid URL for '%s', but got error: %s", testCase.url, err.Error())
		}
	}
}
