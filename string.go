package null

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"

	"github.com/volatiletech/null/convert"
	"github.com/volatiletech/sqlboiler/randomize"
)

// StringFrom creates a new String that will never be blank.
func StringFrom(s string) String {
	return NewString(s, true)
}

// StringFromPtr creates a new String that be null if s is nil.
func StringFromPtr(s *string) String {
	if s == nil {
		return NewString("", false)
	}
	return NewString(*s, true)
}

// NewString creates a new String
func NewString(s string, valid bool) String {
	return String{
		Str:   s,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *String) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullBytes) {
		s.Str = ""
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Str); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return NullBytes, nil
	}
	return json.Marshal(s.Str)
}

// MarshalText implements encoding.TextMarshaler.
func (s String) MarshalText() ([]byte, error) {
	if !s.Valid {
		return []byte{}, nil
	}
	return []byte(s.Str), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *String) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		s.Valid = false
		return nil
	}

	s.Str = string(text)
	s.Valid = true
	return nil
}

// SetValid changes this String's value and also sets it to be non-null.
func (s *String) SetValid(v string) {
	s.Str = v
	s.Valid = true
}

// Ptr returns a pointer to this String's value, or a nil pointer if this String is null.
func (s String) Ptr() *string {
	if !s.Valid {
		return nil
	}
	return &s.Str
}

// IsZero returns true for null strings, for potential future omitempty support.
func (s String) IsZero() bool {
	return !s.Valid
}

// Scan implements the Scanner interface.
func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.Str, s.Valid = "", false
		return nil
	}
	s.Valid = true
	return convert.ConvertAssign(&s.Str, value)
}

// Value implements the driver Valuer interface.
func (s String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.Str, nil
}

// Randomize for sqlboiler
func (s *String) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	str, ok := randomize.FormattedString(nextInt, fieldType)
	if ok {
		s.Str = str
		s.Valid = true
		return
	}

	if shouldBeNull {
		s.Str = ""
		s.Valid = false
	} else {
		s.Str = randomize.Str(nextInt, 1)
		s.Valid = true
	}
}
