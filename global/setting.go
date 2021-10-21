/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-09-15 16:02:39
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-21 11:46:11
 */
package global

import (
	"go-admin/pkg/setting"

	"github.com/sirupsen/logrus"
)

var (
	ServerSetting   *setting.ServerSetting
	JWTSetting      *setting.JWTSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logrus.Logger
)
