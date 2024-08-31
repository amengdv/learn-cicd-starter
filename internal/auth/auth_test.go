package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
    tests := []struct{
        name string
        input http.Header
        expected string
    }{
        {
            name: "Test against random",
            input: http.Header{
                "Authorization": []string{"ApiKey test123"},
            },
            expected: "test123",
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            actual, err := GetAPIKey(tc.input)
            if err != nil {
                t.Errorf("Error: %v\n", err)
            }

            if actual != tc.expected {
                t.Errorf("FAIL! Expect %v Got %v\n", tc.expected, actual)
            }
        })
    }
}
