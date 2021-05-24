package models

// RawMessage wraps the yaml.Unmarshaller behaviour, exposing an interface to YAML
// that will be familiar to users of json.RawMessage.
type RawMessage struct {
	unmarshal func(interface{}) error
}

func (msg *RawMessage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	msg.unmarshal = unmarshal
	return nil
}

func (msg *RawMessage) Unmarshal(v interface{}) error {
	return msg.unmarshal(v)
}
