/*
 * Copyright 2019-present Open Networking Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package core

import (
	"github.com/opencord/voltha-lib-go/v2/pkg/log"
)

const (
	logLevel = log.FatalLevel
)

// Unit test initialization. This init() function handles all unit tests in
// the current directory.
func init() {
	// Setup this package so that it's log level can be modified at run time
	_, err := log.AddPackage(log.JSON, logLevel, log.Fields{"instanceId": "mocks"})
	if err != nil {
		panic(err)
	}
}