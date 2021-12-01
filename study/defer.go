/*
 * @Description:
 * @Autor: 小明～
 * @Date: 2021-11-08 09:22:59
 * @LastEditors: 小明～
 * @LastEditTime: 2021-11-08 09:24:29
 */

package main

import "fmt"

func f1() int {
	defer f2()
	fmt.Println("f1")
	return 1
}

func f2() {
	fmt.Println("f2")
}

func main() {
	result := f1()
	fmt.Println(result, "result")

}
