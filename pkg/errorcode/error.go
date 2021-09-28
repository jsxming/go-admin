/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:22:14
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-16 10:42:17
 */
package errorcode

import "fmt"

type Error struct {
	// 错误码
	code int `json:"code"`
	// 错误消息
	msg string `json:"msg"`

	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) AppendDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}
