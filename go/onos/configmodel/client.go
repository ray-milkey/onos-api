// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configmodel

import (
	"google.golang.org/grpc"
)

// RegistryServiceClientFactory : Default ConfigModelRegistryServiceClient creation - allows it to be mocked
var RegistryServiceClientFactory = func(cc *grpc.ClientConn) ConfigModelRegistryServiceClient {
	return NewConfigModelRegistryServiceClient(cc)
}

// CreateConfigModelRegistryServiceClient creates and returns a new config admin client
func CreateConfigModelRegistryServiceClient(cc *grpc.ClientConn) ConfigModelRegistryServiceClient {
	return RegistryServiceClientFactory(cc)
}
