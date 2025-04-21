package utils

import "encoding/json"

func Marshal[T any](data T) ([]byte, error) {
	output, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return output, nil
}
