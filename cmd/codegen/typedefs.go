package main

import "encoding/json"

type Typedefs struct {
	data map[string]string
}

func getTypedefs(data []byte) (*Typedefs, error) {
	var ret map[string]string
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}

	return &Typedefs{data: ret}, nil
}
