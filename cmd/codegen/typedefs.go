package main

import "encoding/json"

type Typedefs struct {
	data map[CIdentifier]string
}

func getTypedefs(data []byte) (*Typedefs, error) {
	var ret map[CIdentifier]string
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}

	return &Typedefs{data: ret}, nil
}
