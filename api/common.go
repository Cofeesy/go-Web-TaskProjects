package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"memorandumProject/i18n"
	"memorandumProject/serializer"
)

/**
 * @Title ErrorResponse
 * @Description //返回错误信息ErrorResponse
 * @Author Cofeesy 19:32 2022/6/30
 * @Param err error
 * @Return serializer.Response
 **/
func ErrorResponse(err error) serializer.Response {
	//对错误信息进行翻译
	//先对错误信息进行遍历(底层是[]FieldError)
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			filed := i18n.Translation(fmt.Sprintf("Field.%s", e.Field()))
			tag := i18n.Translation(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", filed, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}
	log.Println(err)
	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
