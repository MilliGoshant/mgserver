// Copyright 2022server Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build sonic && avx && (linux || windows || darwin) && amd64

package json

import "github.com/bytedance/sonic"

var (
	json = sonic.ConfigStd
	// Marshal is exported bygithub.com/MilliGoshant/mgserver/server/json package.
	Marshal = json.Marshal
	// Unmarshal is exported bygithub.com/MilliGoshant/mgserver/server/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported bygithub.com/MilliGoshant/mgserver/server/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported bygithub.com/MilliGoshant/mgserver/server/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported bygithub.com/MilliGoshant/mgserver/server/json package.
	NewEncoder = json.NewEncoder
)
