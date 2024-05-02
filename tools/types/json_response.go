package types

type JsonResponse[DataType interface{}] struct {
	Code    int      `json:"code"`
	Message string   `json:"message,omitempty"`
	Data    DataType `json:"data,omitempty"`
}
