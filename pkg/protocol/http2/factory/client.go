/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package factory

import (
	"github.com/cloudwego/hertz/pkg/protocol/client"
	"github.com/cloudwego/hertz/pkg/protocol/http2"
	"github.com/cloudwego/hertz/pkg/protocol/suite"
)

var _ suite.ClientFactory = &clientFactory{}

type clientFactory struct {
	option *http2.ClientOption
}

func (s *clientFactory) NewHostClient() (client client.HostClient, err error) {
	return http2.NewHostClient(s.option), nil
}

func NewClientFactory(option *http2.ClientOption) suite.ClientFactory {
	return &clientFactory{
		option: option,
	}
}
