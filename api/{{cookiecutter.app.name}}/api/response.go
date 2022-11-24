package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Ok encodes to JSON and writes the provided response (if any) along with the httpStatus.
func Ok(w http.ResponseWriter, response interface{}) error {
	var headers map[string]string
	return OkWithHeaders(w, headers, response, http.StatusOK)
}

// OkWithStatus encodes to JSON and writes the provided response (if any) along with the httpStatus.
func OkWithStatus(w http.ResponseWriter, response interface{}, httpStatus int) error {
	var headers map[string]string
	return OkWithHeaders(w, headers, response, httpStatus)
}

// OkWithHeaders encodes to JSON and writes the provided response (if any) along with the extra headers and the httpStatus.
func OkWithHeaders(w http.ResponseWriter, headers map[string]string, response interface{}, httpStatus int) error {
	// String values encode as JSON strings coerced to valid UTF-8,
	// replacing invalid bytes with the Unicode replacement rune.
	// So that the JSON will be safe to embed inside HTML <script> tags,
	// the string is encoded using HTMLEscape,
	// which replaces "<", ">", "&", U+2028, and U+2029 are escaped
	// to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029".
	// This replacement can be disabled when using an Encoder,
	// by calling SetEscapeHTML(false).
	// See https://github.com/golang/go/blob/release-branch.go1.14/src/encoding/json/encode.go#L46
	buf := new(bytes.Buffer)
	if response != nil {
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(response)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return fmt.Errorf("encode: %w", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	for k, v := range headers {
		w.Header()[k] = []string{v}
	}
	w.WriteHeader(httpStatus)

	if buf.Len() > 0 {
		// Remove the extra newline json.Encoder.Encode() adds.
		if _, err := w.Write(bytes.TrimRight(buf.Bytes(), "\n")); err != nil {
			return fmt.Errorf("write: %w", err)
		}
	}

	return nil
}
