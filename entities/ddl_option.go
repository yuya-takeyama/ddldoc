package entities

type DDLOption struct {
	withAutoIncrement bool
}

func NewDDLOption(withAutoIncrement bool) *DDLOption {
	return &DDLOption{
		withAutoIncrement,
	}
}

func (option *DDLOption) IsWithAutoIncrement() bool {
	return option.withAutoIncrement
}
