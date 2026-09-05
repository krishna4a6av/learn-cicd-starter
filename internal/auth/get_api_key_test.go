package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	type testStruct struct {
		inp  string
		want string
	}

	tests := []testStruct{
		{
			inp:  "ApiKey my-secret-key",
			want: "my-secret-key",
		},
		{
			inp:  "ApiKey some1234",
			want: "some1234",
		},
	}

	for _, test := range tests {
		headers := http.Header{}
		headers.Set("Authorization", test.inp)

		key, err := auth.GetAPIKey(headers)
		if err != nil {
			t.Fatal(err)
		}

		if key != test.want {
			t.Errorf("Expected %s got %s", test.want, key)
		}
	}
}
