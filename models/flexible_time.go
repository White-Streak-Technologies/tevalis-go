package models

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// FlexibleTime handles timestamps with or without timezone and optional nil/empty values.
type FlexibleTime struct {
	time.Time
}

func (t *FlexibleTime) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*t = FlexibleTime{}
		return nil
	}

	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	value := strings.TrimSpace(raw)
	if value == "" {
		*t = FlexibleTime{}
		return nil
	}

	parsed, err := parseFlexibleTime(value)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

func (t FlexibleTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format(time.RFC3339Nano))
}

func (t *FlexibleTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "nil" && (attr.Value == "true" || attr.Value == "1") {
			var discard string
			if err := d.DecodeElement(&discard, &start); err != nil {
				return err
			}
			*t = FlexibleTime{}
			return nil
		}
	}

	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	value := strings.TrimSpace(raw)
	if value == "" {
		*t = FlexibleTime{}
		return nil
	}

	parsed, err := parseFlexibleTime(value)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

func (t FlexibleTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.Time.IsZero() {
		return e.EncodeElement("", start)
	}
	return e.EncodeElement(t.Time.Format(time.RFC3339Nano), start)
}

func parseFlexibleTime(value string) (time.Time, error) {
	layouts := []string{
		time.RFC3339Nano,
		"2006-01-02T15:04:05.9999999",
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05.99999",
		"2006-01-02T15:04:05.9999",
		"2006-01-02T15:04:05.999",
		"2006-01-02T15:04:05.99",
		"2006-01-02T15:04:05.9",
		"2006-01-02T15:04:05",
	}

	for _, layout := range layouts {
		if parsed, err := time.Parse(layout, value); err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid time format: %s", value)
}
