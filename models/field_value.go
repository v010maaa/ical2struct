package models

type FieldValue struct {
	Field string
	Value string
}

func NewFieldValue() FieldValue {
	return FieldValue{Field: "none", Value: "none"}
}
