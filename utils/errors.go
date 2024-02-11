package utils

type GroupNotFound struct{}

func (e *GroupNotFound) Error() string {
	return "group not found"
}

type FlagNotFound struct{}

func (e *FlagNotFound) Error() string {
	return "flag not found"
}

type LimitNotFound struct{}

func (e *LimitNotFound) Error() string {
	return "limit not found"
}

type ValueNotFound struct{}

func (e *ValueNotFound) Error() string {
	return "value not found"
}

type ConfigNotFound struct{}

func (e *ConfigNotFound) Error() string {
	return "config not found"
}

type ErrorResponse struct {
	Error string `json:"error"`
}
