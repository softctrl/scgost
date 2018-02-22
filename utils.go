//
// Author:
//  Carlos Timoshenko
//  carlostimoshenkorodrigueslopes@gmail.com
//
//  https://github.com/softctrl
//
// This project is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
package scgost

import (
	"fmt"
	"net/http"
)

const (
	SUCCESS  = 0
	ERROR    = 1
	RESPONSE = `{"_0":%d,%s`
)

func ReturnError(__resp http.ResponseWriter, __err error) {
	ReturnErrorString(__resp, __err.Error())
}

func ReturnErrorString(__resp http.ResponseWriter, __err string) {
	__resp.Write([]byte(fmt.Sprintf(RESPONSE, ERROR,
		fmt.Sprintf(`"_1":"%s"}`, __err))))
}

func ReturnMessage(__resp http.ResponseWriter, __code int, __msg []byte) {

	__resp.Write([]byte(fmt.Sprintf(RESPONSE, __code, "")))
	__resp.Write([]byte(`"_1":"`))
	__resp.Write(__msg)
	__resp.Write([]byte(`"}`))

}

func ReturnMessageString(__resp http.ResponseWriter, __code int, __msg string) {
	ReturnMessage(__resp, __code, []byte(__msg))
}

func ReturnJsonMessage(__resp http.ResponseWriter, __code int, __msg []byte) {

	__resp.Write([]byte(fmt.Sprintf(RESPONSE, __code, "")))
	__resp.Write([]byte(`"_1":`))
	__resp.Write(__msg)
	__resp.Write([]byte(`}`))

}
