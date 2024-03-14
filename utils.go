package sumsubgo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
)

func decodeResponseBody(request *http.Response, target interface{}) error {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}
	defer func() {
		_ = request.Body.Close()
	}()

	return json.Unmarshal(body, target)
}

func _sign(ts string, secret string, method string, path string, body *[]byte) string {
	hash := hmac.New(sha256.New, []byte(secret))
	data := []byte(ts + method + path)

	if body != nil {
		data = append(data, *body...)
	}

	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
