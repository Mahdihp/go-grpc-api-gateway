package data_util

import (
	"encoding/json"
	"fmt"
)

func SerializeMessage(message interface{}) ([]byte, error) {
	// Serialize the message struct to JSON
	serialized, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize message: %w", err)
	}
	return serialized, nil
}
