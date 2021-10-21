/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:25:32
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-20 10:53:10
 */
package errorcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000, "服务内部错误")
	InvalidParams             = NewError(1001, "入参错误")
	NotFound                  = NewError(1002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(1004, "鉴权失败，Token错误")
	NotFoundToken             = NewError(1008, "请上传Token参数")
	UnauthorizedTokenTimeout  = NewError(1005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(1006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(1007, "请求过多")
	UpdateFail                = NewError(2000, "更新失败")
	CreateFail                = NewError(2001, "新增失败")
	DeleteFail                = NewError(2002, "删除失败")
	SearchFail                = NewError(2003, "查找失败")
	// TooManyRequests           = NewError(10000007, "请求过多")
	// TooManyRequests           = NewError(10000007, "请求过多")
)
