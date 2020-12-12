package helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
)

// EncodeToBase64 encode to base64
func EncodeToBase64(v interface{}) (string, error) {
    var buf bytes.Buffer
    encoder := base64.NewEncoder(base64.StdEncoding, &buf)
    err := json.NewEncoder(encoder).Encode(v)
    if err != nil {
        return "", err
    }
    encoder.Close()
    return buf.String(), nil
}

// DecodeFromBase64 decode from base64
func DecodeFromBase64(v interface{}, enc string) error {
    return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}