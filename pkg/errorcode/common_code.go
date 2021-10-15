/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 18:25:32
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-15 10:13:56
 */
package errorcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	NotFoundToken             = NewError(10000008, "请上传Token参数")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	UpdateFail                = NewError(20000000, "更新失败")
	CreateFail                = NewError(20000001, "新增失败")
	DeleteFail                = NewError(20000002, "删除失败")
	SearchFail                = NewError(20000003, "查找失败")
	// TooManyRequests           = NewError(10000007, "请求过多")
	// TooManyRequests           = NewError(10000007, "请求过多")
)
