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

package config

import "time"

var preDefinedOpts []RequestOption

type RequestOptions struct {
	tags map[string]string
	isSD bool

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// RequestOption is the only struct to set request-level options.
type RequestOption struct {
	F func(o *RequestOptions)
}

// NewRequestOptions create a *RequestOptions according to the given opts.
func NewRequestOptions(opts []RequestOption) *RequestOptions {
	options := &RequestOptions{
		tags: make(map[string]string),
		isSD: false,
	}
	if preDefinedOpts != nil {
		options.Apply(preDefinedOpts)
	}
	options.Apply(opts)
	return options
}

// WithTag set tag in RequestOptions.
func WithTag(k, v string) RequestOption {
	return RequestOption{F: func(o *RequestOptions) {
		o.tags[k] = v
	}}
}

// WithSD set isSD in RequestOptions.
func WithSD(b bool) RequestOption {
	return RequestOption{F: func(o *RequestOptions) {
		o.isSD = b
	}}
}

// WithDialTimeout sets dial timeout.
//
// This is the request level configuration.It has a higher
// priority than the client level configuration
func WithDialTimeout(t time.Duration) RequestOption {
	return RequestOption{F: func(o *RequestOptions) {
		o.DialTimeout = t
	}}
}

// WithReadTimeout sets read timeout.
//
// This is the request level configuration.It has a higher
// priority than the client level configuration
func WithReadTimeout(t time.Duration) RequestOption {
	return RequestOption{F: func(o *RequestOptions) {
		o.ReadTimeout = t
	}}
}

// WithWriteTimeout sets write timeout.
//
// This is the request level configuration.It has a higher
// priority than the client level configuration
func WithWriteTimeout(t time.Duration) RequestOption {
	return RequestOption{F: func(o *RequestOptions) {
		o.WriteTimeout = t
	}}
}

func (o *RequestOptions) Apply(opts []RequestOption) {
	for _, op := range opts {
		op.F(o)
	}
}

func (o *RequestOptions) Tag(k string) string {
	return o.tags[k]
}

func (o *RequestOptions) Tags() map[string]string {
	return o.tags
}

func (o *RequestOptions) IsSD() bool {
	return o.isSD
}

func (o *RequestOptions) CopyTo(dst *RequestOptions) {
	if dst.tags == nil {
		dst.tags = make(map[string]string)
	}

	for k, v := range o.tags {
		dst.tags[k] = v
	}

	dst.isSD = o.isSD
	dst.ReadTimeout = o.ReadTimeout
	dst.WriteTimeout = o.WriteTimeout
	dst.DialTimeout = o.DialTimeout
}

func (o *RequestOptions) Reset() {
	o.tags = make(map[string]string)
	o.isSD = false
	o.DialTimeout = 0
	o.ReadTimeout = 0
	o.WriteTimeout = 0
}

// SetPreDefinedOpts Pre define some RequestOption here
func SetPreDefinedOpts(opts ...RequestOption) {
	preDefinedOpts = nil
	preDefinedOpts = append(preDefinedOpts, opts...)
}
