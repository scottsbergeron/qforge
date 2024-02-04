package qforge

type DataType int

const (
	StringDataType DataType = iota
	IntegerDataType
	FloatDataType
	BooleanDataType
	DateDataType
	DateTimeDataType
)

type Column struct {
	Id       string
	Name     string
	DataType DataType
}
