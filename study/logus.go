/*
 * @Description:logus 学习
 * @Autor: 小明～
 * @Date: 2021-10-21 11:08:23
 * @LastEditors: 小明～
 * @LastEditTime: 2021-10-21 11:38:27
 */
package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {

	// h := lineHook{}
	// log.AddHook(h)
	log.SetReportCaller(true) //输出日志中添加文件名和方法信息
	log.Formatter = &logrus.JSONFormatter{}
	log.SetOutput(io.MultiWriter(os.Stdout))
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	log.Warn(".....warn")
	log.Fatal(".....Fatal")
}
