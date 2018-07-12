package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/grafodb/grafodb/internal/util/json"
)

// var defaultETagger = NewMD5HashETagger()
var defaultETagger = NewSha1HashETagger()

// FIXME: Add standard method to send a JSON wrapped error message

// SendStatus sends write a HTTP status code to the response
func SendStatus(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

// TODO: IF the client sends an etag for the previous cached version, we should respond with a 304 and not send data to leverage client caching
// SendJSON sends a JSON marshalled object to the response using jsoniter
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}, eTagger ...ETagger) error {
	if len(eTagger) > 1 {
		return fmt.Errorf("only one etagger is allowed")
	}

	var tagger ETagger

	if len(eTagger) == 0 {
		tagger = defaultETagger
	} else {
		tagger = eTagger[0]
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Add(HeaderContentType, MimeApplicationJSON)
	w.Header().Add(HeaderETag, tagger.ETag(body))

	allowContentLengthHeader := statusCode >= http.StatusOK && statusCode != http.StatusNoContent
	bodyLength := len(body)
	if bodyLength > 0 && allowContentLengthHeader {
		w.Header().Add(HeaderContentLength, strconv.Itoa(bodyLength))
	}

	w.WriteHeader(statusCode)
	w.Write(body)

	return nil
}

// // Bind attempts to decode a JSON request body into a structure
// func Bind(r *http.Request, v interface{}) error {
// 	if reflect.ValueOf(v).IsNil() {
// 		return fmt.Errorf("v cannot be nil")
// 	}
//
// 	if reflect.ValueOf(v).Kind() != reflect.Ptr {
// 		return fmt.Errorf("v must be a pointer type")
// 	}
// 	return json.NewDecoder(r.Body).Decode(v)
// }

func MustQueryParamInt(r *http.Request, key string) (int, error) {
	vs, err := MustQueryParamString(r, key)
	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(vs)
	if err != nil {
		return 0, fmt.Errorf("invalid int value \"%s\" for query param \"%s\"", vs, key)
	}
	return v, nil
}

func QueryParamIntDefault(r *http.Request, key string, defaultValue int) int {
	v, err := MustQueryParamInt(r, key)
	if err != nil {
		return defaultValue
	}
	return v
}

func MustQueryParamString(r *http.Request, key string) (string, error) {
	q := r.URL.Query()
	v := q.Get(key)
	if v == "" {
		return "", fmt.Errorf("missing required query param \"%s\"", key)
	}
	return v, nil
}

func QueryParamStringDefault(r *http.Request, key string, defaultValue string) string {
	v, err := MustQueryParamString(r, key)
	if err != nil {
		return defaultValue
	}
	return v
}

func MustQueryParamTime(r *http.Request, key string) (time.Time, error) {
	vs, err := MustQueryParamString(r, key)
	if err != nil {
		return time.Time{}, err
	}

	v, err := time.Parse(time.RFC3339, vs)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time value \"%s\" for query param \"%s\"", vs, key)
	}
	return v, nil
}

func QueryParamTimeDefault(r *http.Request, key string, defaultValue time.Time) time.Time {
	v, err := MustQueryParamTime(r, key)
	if err != nil {
		return defaultValue
	}
	return v
}
