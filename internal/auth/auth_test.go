package auth

import (
	"net/http"
	"testing"
)

func TestGetApi(t *testing.T) {
	cases := []struct {
		name string
		h http.Header
		expectedKey string
		wantErr bool
	}{
		{
			name: "Correct Header",
			h: func() http.Header{
				h := http.Header{}
				h.Set("Authorization", "ApiKey CodeHere")
				return h
			}(),
			expectedKey: "CodeHere",
			wantErr: false,
		},
		{
			name: "Right Code, Wrong header",
			h: func() http.Header{
				h := http.Header{}
				h.Set("Code", "ApiKey WrongHeader")
				return h
			}(),
			expectedKey: "CodeHere",
			wantErr: true,
		},
		{
			name: "No Header",
			expectedKey: "",
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T){
			api, err := GetAPIKey(tt.h)
			if err != nil && !tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && api != tt.expectedKey {
				t.Errorf("GetBearerToken() error = %v, token = %v", err, tt.expectedKey)
			}
		})
	}
}