package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string // full Authorization header value; if empty -> header not set
		want    string
		wantErr bool
	}{
		{
			name:    "valid APIkey",
			header:  "ApiKey my-secret-key",
			want:    "my-secret-key",
			wantErr: false,
		},
		{
			name:    "missing Authorization header",
			header:  "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "wrong prefix",
			header:  "Bearer somethingelse",
			want:    "",
			wantErr: true,
		},
		{
			name:    "no key after prefix",
			header:  "ApiKey ",
			want:    "",
			wantErr: true,
		},
		{
			name:    "extra spaces before key",
			header:  "ApiKey   spaced-key",
			want:    "spaced-key",
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			if tc.header != "" {
				req.Header.Set("Authorization", tc.header)
			}

			got, err := GetAPIKey(req)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error but got none (got=%q)", got)
				}
				// OK: error expected
				return
			}
			// no error expected
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("want key %q, got %q", tc.want, got)
			}
		})
	}
}
