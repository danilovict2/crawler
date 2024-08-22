package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute deep nested URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
		<body>
    		<span>
        		<a href="https://other.com/path/one">
            		<span>PATH ONE</span>
        		</a>
        		<div>
            		<div>
                		<div>
                    		<a href="https://other-url.com/path/one">
                        	<span>PATH TWO</span>
                    		</a>
                		</div>
            		</div>
        		</div>
    		</span>
		</body>
		</html>
		`,
			expected: []string{"https://other.com/path/one", "https://other-url.com/path/one"},
		},
		{
			name:     "relative deep nested URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<span>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
            		<div>
                		<div>
                    		<a href="/path/two">
                        	<span>PATH TWO</span>
                    		</a>
                		</div>
            		</div>
				</span>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href=":\\invalidURL">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: nil,
		},
		{
			name:     "handle invalid base URL",
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
		<html>
			<body>
				<a href="/path">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: nil,
			errorContains: "couldn't parse base URL",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)

			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs %v, got URLs %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
