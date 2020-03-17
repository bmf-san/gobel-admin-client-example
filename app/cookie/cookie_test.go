package cookie

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewCookie(t *testing.T) {
	expected := &Cookie{
		Flash: &Flash{
			Messages: make(map[string]string),
		},
	}

	actual := NewCookie()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual: %v expected: %v\n", actual, expected)
	}
}

func TestSetFlash(t *testing.T) {
	cookie := NewCookie()

	recorder := httptest.NewRecorder()

	key := "test"
	value := "test message"
	cookie.SetFlash(recorder, key, []byte(value))

	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}

	reqC, err := request.Cookie(key)
	if err != nil {
		t.Errorf("%v", err)
	}

	actual := reqC.Value
	expected := base64.RawURLEncoding.EncodeToString([]byte(value))
	if actual != expected {
		t.Errorf("actual: %v expected: %v\n", actual, expected)
	}

	actualMap := cookie.Flash.Messages
	expectedMap := map[string]string{key: value}
	if !reflect.DeepEqual(actualMap, expectedMap) {
		t.Errorf("actual: %v expected: %v\n", actualMap, expectedMap)
	}
}

func TestGetFlash(t *testing.T) {
	flash := NewCookie()

	recorder := httptest.NewRecorder()

	key := "test"
	value := "test message"
	flash.SetFlash(recorder, key, []byte(value))

	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}

	expected := &Flash{
		Messages: map[string]string{key: value},
	}

	actual, err := flash.GetFlash(recorder, request)
	if err != nil {
		t.Errorf("%v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual: %v expected: %v\n", actual, expected)
	}
}
