package order

import "encoding/json"

// Encode dumps an Order to json.
func (o Order) Encode() ([]byte, error) {
	return json.Marshal(o)
}

// Decode loads an Order from json
func (o *Order) Decode(data []byte) error {
	return json.Unmarshal(data, o)
}
