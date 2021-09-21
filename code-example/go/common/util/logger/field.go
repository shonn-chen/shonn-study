package logger

import "fmt"

type field struct {
	key string
	val interface{}
}

func (f *field) String() string {
	if f == nil {
		return ""
	}
	return fmt.Sprintf("%s=%v", f.key, f.val)
}

func Field(key string, val interface{}) *field {
	return &field{
		key: key,
		val: val,
	}
}
