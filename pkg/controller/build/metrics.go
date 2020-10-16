/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package build

import (
	"time"

	"sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

const buildResultLabel = "result"

var (
	buildAttempt = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "camel_k_build_attempt",
			Help: "Camel K build attempt",
		},
		[]string{
			buildResultLabel,
		})

	buildDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "camel_k_build_duration_seconds",
			Help: "Camel K build duration",
			Buckets: []float64{
				30 * time.Second.Seconds(),
				1 * time.Minute.Seconds(),
				1.5 * time.Minute.Seconds(),
				2 * time.Minute.Seconds(),
				5 * time.Minute.Seconds(),
				10 * time.Minute.Seconds(),
			},
		},
		[]string{
			buildResultLabel,
		},
	)

	queueDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "camel_k_build_queue_duration_seconds",
			Help: "Camel K build queue duration",
			Buckets: []float64{
				5 * time.Second.Seconds(),
				15 * time.Second.Seconds(),
				30 * time.Second.Seconds(),
				1 * time.Minute.Seconds(),
				5 * time.Minute.Seconds(),
			},
		},
	)
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(buildAttempt, buildDuration, queueDuration)
}