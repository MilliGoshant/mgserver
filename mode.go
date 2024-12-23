// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgserver

import (
	"flag"
	"io"
	"os"
	"sync/atomic"

	"github.com/MilliGoshant/mgserver/server/binding"
)

// EnvServerMode indicates environment name forserver mode.
const EnvServerMode = "MGSERVER_MODE"

const (
	// DebugMode indicatesserver mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicatesserver mode is release.
	ReleaseMode = "release"
	// TestMode indicatesserver mode is test.
	TestMode = "test"
)

const (
	debugCode = iota
	releaseCode
	testCode
)

// DefaultWriter is the default io.Writer used byserver for debug output and
// middleware output like Logger() or Recovery().
// Note that both Logger and Recovery provides custom ways to configure their
// output io.Writer.
// To support coloring in Windows use:
//
//	import "github.com/mattn/go-colorable"
//	server.DefaultWriter = colorable.NewColorableStdout()
var DefaultWriter io.Writer = os.Stdout

// DefaultErrorWriter is the default io.Writer used byserver to debug errors
var DefaultErrorWriter io.Writer = os.Stderr

var serverMode int32 = debugCode
var modeName atomic.Value

func init() {
	mode := os.Getenv(EnvServerMode)
	SetMode(mode)
}

// SetMode setsserver mode according to input string.
func SetMode(value string) {
	if value == "" {
		if flag.Lookup("test.v") != nil {
			value = TestMode
		} else {
			value = DebugMode
		}
	}

	switch value {
	case DebugMode, "":
		atomic.StoreInt32(&serverMode, debugCode)
	case ReleaseMode:
		atomic.StoreInt32(&serverMode, releaseCode)
	case TestMode:
		atomic.StoreInt32(&serverMode, testCode)
	default:
		panic("server mode unknown: " + value + " (available mode: debug release test)")
	}
	modeName.Store(value)
}

// DisableBindValidation closes the default validator.
func DisableBindValidation() {
	binding.Validator = nil
}

// EnableJsonDecoderUseNumber sets true for binding.EnableDecoderUseNumber to
// call the UseNumber method on the JSON Decoder instance.
func EnableJsonDecoderUseNumber() {
	binding.EnableDecoderUseNumber = true
}

// EnableJsonDecoderDisallowUnknownFields sets true for binding.EnableDecoderDisallowUnknownFields to
// call the DisallowUnknownFields method on the JSON Decoder instance.
func EnableJsonDecoderDisallowUnknownFields() {
	binding.EnableDecoderDisallowUnknownFields = true
}

// Mode returns currentserver mode.
func Mode() string {
	return modeName.Load().(string)
}
