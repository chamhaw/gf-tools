// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gerror

import "fmt"

// LocalCode is an implementer for interface Code for internal usage only.
type LocalCode struct {
	code       string // Error code, usually a string.
	message    string // Brief message for this error code.
	httpStatus int    //  http code  if necessary
}

// Code returns the integer number of current error code.
func (c LocalCode) Code() string {
	return c.code
}

// Message returns the brief message for current error code.
func (c LocalCode) Message() string {
	return c.message
}

// Detail returns the detailed information of current error code,
// which is mainly designed as an extension field for error code.
func (c LocalCode) HttpStatus() int {
	return c.httpStatus
}

// String returns current error code as a string.
func (c LocalCode) String() string {
	if c.httpStatus != 0 {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.HttpStatus())
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
