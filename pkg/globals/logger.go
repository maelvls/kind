/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package globals will be deleted in a future commit
// this package contains globals that we've not yet re-worked to not be globals
package globals

import (
	"io"
	"sync"

	"sigs.k8s.io/kind/pkg/log"

	"sigs.k8s.io/kind/pkg/internal/util/cli"
	"sigs.k8s.io/kind/pkg/internal/util/env"
)

// TODO: consider threading through a logger instead of using a global logger
// This version is just a first small step towards making kind easier to import
// in test-harnesses
var globalLoggerMu sync.Mutex
var globalLogger log.Logger = log.NoopLogger{}

// SetLogger sets the standard logger used by this package.
// If not set, log.NoopLogger will be used
func SetLogger(l log.Logger) {
	globalLoggerMu.Lock()
	defer globalLoggerMu.Unlock()
	globalLogger = l
}

// UseCLILogger sets the global logger to the kind CLI's default stderr logger
// if writer is a tty, the CLI spinner will be enabled
//
// Not to be confused with the default if not set of log.NoopLogger
func UseCLILogger(writer io.Writer, verbosity log.Level) {
	if env.IsTerminal(writer) {
		writer = cli.NewSpinner(writer)
	}
	SetLogger(cli.NewLogger(writer, verbosity))
}

// GetLogger returns the standard logger used by this package
func GetLogger() log.Logger {
	globalLoggerMu.Lock()
	defer globalLoggerMu.Unlock()
	return globalLogger
}
