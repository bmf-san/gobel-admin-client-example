package csrf

import (
	"reflect"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	csrfToken, err := GenerateToken()
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(csrfToken) != tokenLength*2 {
		t.Errorf("actual: %v expected: %v\n", len(csrfToken), tokenLength*2)
	}

	newCsrfToken, err := GenerateToken()
	if err != nil {
		t.Errorf("%v", err)
	}

	if reflect.DeepEqual(csrfToken, newCsrfToken) {
		t.Errorf("actual: %v expected: %v\n", csrfToken, newCsrfToken)
	}
}
