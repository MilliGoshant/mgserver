// Copyright 2017 Bo-Yi Wu. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build go_json

package json

import json "github.com/goccy/go-json"

var (
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
