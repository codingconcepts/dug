package models

// RawMessage wraps the yaml.Unmarshaller behaviour, exposing an interface to YAML
// that will be familiar to users of json.RawMessage.
type RawMessage struct {
	unmarshal func(interface{}) error
}

// UnmarshalYAML sets the internal function that will be used to provide the two-
// step umarshalling of connector structs.
func (msg *RawMessage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	msg.unmarshal = unmarshal
	return nil
}

// Unmarshal calls the underlying unmarshal function, setting a value for the
// target struct.
func (msg *RawMessage) Unmarshal(v interface{}) error {
	return msg.unmarshal(v)
}
