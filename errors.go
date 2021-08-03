//  Copyright (c) 2021. The EFF Team Authors.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  See the License for the specific language governing permissions and
//  limitations under the License.

package errors

import (
	"fmt"
)

type ErgoError struct {
	Message string
}

func (ee *ErgoError) Error() string {
	return ee.Message
}

func (ee *ErgoError) String() string {
	return ee.Message
}

func Bomb(format string, args ...interface{}) {
	panic(ErgoError{Message: fmt.Sprintf(format, args...)})
}

func Dangerous(v interface{}) {
	if v == nil {
		return
	}

	switch t := v.(type) {
	case string:
		if t != "" {
			panic(ErgoError{Message: t})
		}
	case error:
		panic(ErgoError{Message: t.Error()})
	}
}

func Boka(value string, v interface{}) {
	if v == nil {
		return
	}
	Bomb(value)
}
