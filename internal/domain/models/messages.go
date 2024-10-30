package models

import "encoding/json"

type EventMessage struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}
