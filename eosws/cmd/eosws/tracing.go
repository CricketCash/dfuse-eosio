// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/dfuse-io/derr"
	"github.com/dfuse-io/dtracing"
	"go.opencensus.io/trace"
)

func setupTracing(sampler trace.Sampler, chainID string) {
	err := dtracing.SetupTracing("eosws", sampler, stackdriver.Options{
		DefaultTraceAttributes: dtracing.TraceAttributes{
			"chain_id": chainID,
		},
	})
	derr.ErrorCheck("unable to setup tracing correctly", err)
}
