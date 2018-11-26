/*
Copyright 2018 The Kubernetes Authors.

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

package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

// SetLogger configures all client-go logging to use the specified logger.  This
// function is not concurrent-safe - callers must be sure to call it from a
// single goroutine.  The default implementation uses Go's standard library via
// the github.com/go-logr/stdr package.
func SetLogger(logger logr.Logger) {
	globalLogger = logger
}

var globalLogger logr.Logger = stdr.New()

// Info logs a non-error message with the given key/value pairs as context.
//
// The msg argument should be used to add some constant description to
// the log line.  The key/value pairs can then be used to add additional
// variable information.  The key/value pairs should alternate string
// keys and arbitrary values.
func Info(msg string, keysAndValues ...interface{}) {
	globalLogger.Info(msg, keysAndValues...)
}

// Enabled tests whether this InfoLogger is enabled.  For example,
// commandline flags might be used to set the logging verbosity and disable
// some info logs.
func Enabled() bool {
	return globalLogger.Enabled()
}

// Error logs an error, with the given message and key/value pairs as context.
// It functions similarly to calling Info with the "error" named value, but may
// have unique behavior, and should be preferred for logging errors (see the
// package documentations for more information).
//
// The msg field should be used to add context to any underlying error,
// while the err field should be used to attach the actual error that
// triggered this log line, if present.
func Error(err error, msg string, keysAndValues ...interface{}) {
	globalLogger.Error(err, msg, keysAndValues...)
}

// V returns an InfoLogger value for a specific verbosity level.  A higher
// verbosity level means a log message is less important.  It's illegal to
// pass a log level less than zero.
func V(level int) logr.InfoLogger {
	return globalLogger.V(level)
}

// WithValues adds some key-value pairs of context to a logger.
// See Info for documentation on how key/value pairs work.
func WithValues(keysAndValues ...interface{}) logr.Logger {
	return globalLogger.WithValues(keysAndValues...)
}

// WithName adds a new element to the logger's name.
// Successive calls with WithName continue to append
// suffixes to the logger's name.  It's strongly reccomended
// that name segments contain only letters, digits, and hyphens
// (see the package documentation for more information).
func WithName(name string) logr.Logger {
	return globalLogger.WithName(name)
}
