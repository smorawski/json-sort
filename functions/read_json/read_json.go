package read_json

import (
	"encoding/json"
	"os"
)

func ReadJson(path string) (map[string]interface{}, error) {
	stream, err := os.OpenFile(path, os.O_RDONLY, 777)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(stream)
	var jsonData map[string]interface{}
	err = decoder.Decode(&jsonData)

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
