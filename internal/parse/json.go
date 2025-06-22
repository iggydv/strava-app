package parse

import (
	"encoding/json"
	"io"
)

func JSON(response io.Reader, v any) error {
	err := json.NewDecoder(response).Decode(&v)
	return err
}

func JSONEncoder(w io.Writer, v any) error {
	err := json.NewEncoder(w).Encode(v)
	return err
}
