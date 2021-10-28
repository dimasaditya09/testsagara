package utils

type ValidationResponse struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"validator"`
	Data        interface{}  `json:"data"`
}
