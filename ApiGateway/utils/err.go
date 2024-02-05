package utils

import (
	"fmt"
	"my-package/models"
	"net/http"
)

type MockResponse interface {
	GetError() bool
	GetStatus() int64
	GetMessage() string
}

func HandlerErrGrpcCleint[T MockResponse](result T, err error) models.Response {
	fmt.Println("----", err)
	if err != nil {
		res := models.Response{
			Error:   true,
			Status:  http.StatusInternalServerError,
			Massage: err.Error(),
		}
		return res
	}
	res := models.Response{
		Error:   result.GetError(),
		Status:  result.GetStatus(),
		Massage: result.GetMessage(),
	}
	return res
}

//

// func HandlerErrGrpcCleint(result interface{}, err error) models.Response {
// 	if err != nil {
// 		res := models.Response{
// 			Error:   true,
// 			Status:  http.StatusInternalServerError,
// 			Massage: err.Error(),
// 		}
// 		return res
// 	}
// 	if result == nil {
// 		res := models.Response{
// 			Error:   true,
// 			Status:  http.StatusInternalServerError,
// 			Massage: "error result is nil",
// 		}
// 		return res
// 	}
// 	val := reflect.ValueOf(result).Elem()
// 	// check
// 	// fmt.Println(reflect.TypeOf(result))
// 	// fmt.Println(val.Field(5).Interface().(int))
// 	// fmt.Println(reflect.TypeOf(val.Field(5)))
// 	// for i := 0; i < val.NumField(); i++ {
// 	// 	fmt.Println(val.Type().Field(i).Name)
// 	// }
// 	// check interface null
// 	if !(val.IsValid()) {
// 		res := models.Response{
// 			Error:   true,
// 			Status:  http.StatusInternalServerError,
// 			Massage: "error result is nil",
// 		}
// 		return res
// 	}

// 	res := models.Response{
// 		Error:   val.Field(3).Interface().(bool),
// 		Status:  val.Field(4).Interface().(int64),
// 		Massage: val.Field(5).Interface().(string),
// 	}

// 	return res
// }
