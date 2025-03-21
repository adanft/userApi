package response

type BaseResponse[T any] struct {
	Errors  []string `json:"errors"`
	Data    T        `json:"data"`
	Message string   `json:"message"`
	Success bool     `json:"success"`
}

func NewBaseResponse[T any](data T) *BaseResponse[T] {
	return &BaseResponse[T]{Data: data, Success: true}
}
