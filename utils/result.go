package utils

type Result[T any] struct {
	Code      int         `json:"code"`
	Data      T           `json:"data"`
	IsSuccess bool        `json:"isSuccess"`
	Message   interface{} `json:"message"`
}

var AjaxResult *Result[string]

func init() {
	AjaxResult = &Result[string]{}
}

func (this *Result[T]) Success(data T) *Result[T] {
	return &Result[T]{
		Data:      data,
		Code:      200,
		IsSuccess: true,
		Message:   "success",
	}
}
func (this *Result[T]) SuccessByCode(data T, code int) *Result[T] {
	return &Result[T]{
		Data:      data,
		Code:      code,
		IsSuccess: true,
		Message:   "success",
	}
}
func (this *Result[T]) Error(msg interface{}) *Result[T] {
	return &Result[T]{

		Code:      -200,
		IsSuccess: false,
		Message:   msg,
	}
}
func (this *Result[T]) ErrorByCode(msg interface{}, code int) *Result[T] {
	return &Result[T]{

		Code:      code,
		IsSuccess: false,
		Message:   msg,
	}
}
