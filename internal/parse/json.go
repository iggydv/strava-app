package parse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func JSON(response io.Reader, v any) error {
	err := json.NewDecoder(response).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	prettyJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
		return err
	}
	fmt.Println(string(prettyJSON))
	return nil
}
