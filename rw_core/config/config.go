/*
 * Copyright 2018-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"flag"
	"os"
	"time"
)

// RW Core service default constants
const (
	EtcdStoreName = "etcd"
)

// RWCoreFlags represents the set of configurations used by the read-write core service
type RWCoreFlags struct {
	// Command line parameters
	RWCoreEndpoint              string
	GrpcAddress                 string
	KafkaAdapterAddress         string
	KafkaClusterAddress         string
	KVStoreType                 string
	KVStoreTimeout              time.Duration
	KVStoreAddress              string
	KVTxnKeyDelTime             int
	CoreTopic                   string
	EventTopic                  string
	LogLevel                    string
	Banner                      bool
	DisplayVersionOnly          bool
	RWCoreKey                   string
	RWCoreCert                  string
	RWCoreCA                    string
	LongRunningRequestTimeout   time.Duration
	DefaultRequestTimeout       time.Duration
	DefaultCoreTimeout          time.Duration
	CoreBindingKey              string
	MaxConnectionRetries        int
	ConnectionRetryInterval     time.Duration
	LiveProbeInterval           time.Duration
	NotLiveProbeInterval        time.Duration
	ProbeAddress                string
	TraceEnabled                bool
	TraceAgentAddress           string
	LogCorrelationEnabled       bool
	VolthaStackID               string
	BackoffRetryInitialInterval time.Duration
	BackoffRetryMaxElapsedTime  time.Duration
	BackoffRetryMaxInterval     time.Duration
}

// ParseCommandArguments parses the arguments when running read-write core service
func (cf *RWCoreFlags) ParseCommandArguments(args []string) {

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	fs.StringVar(&cf.RWCoreEndpoint, "vcore-endpoint",
		"rwcore", "RW core endpoint address")

	fs.StringVar(&cf.GrpcAddress, "grpc_address",
		":50057", "GRPC server - address")

	fs.StringVar(&cf.KafkaAdapterAddress, "kafka_adapter_address",
		"127.0.0.1:9092", "Kafka - Adapter messaging address")

	fs.StringVar(&cf.KafkaClusterAddress, "kafka_cluster_address",
		"127.0.0.1:9094", "Kafka - Cluster messaging address")

	fs.StringVar(&cf.CoreTopic, "rw_core_topic",
		"rwcore", "RW Core topic")

	fs.StringVar(&cf.EventTopic, "event_topic",
		"voltha.events", "RW Core Event topic")

	fs.StringVar(&cf.KVStoreType, "kv_store_type",
		EtcdStoreName, "KV store type")

	fs.DurationVar(&cf.KVStoreTimeout, "kv_store_request_timeout",
		5*time.Second, "The default timeout when making a kv store request")

	fs.StringVar(&cf.KVStoreAddress, "kv_store_address",
		"127.0.0.1:2379", "KV store address")

	fs.IntVar(&cf.KVTxnKeyDelTime, "kv_txn_delete_time",
		60, "The time to wait before deleting a completed transaction key")

	fs.StringVar(&cf.LogLevel, "log_level",
		"warn", "Log level")

	fs.DurationVar(&cf.LongRunningRequestTimeout, "timeout_long_request",
		2000*time.Millisecond, "Timeout for long running request")

	fs.DurationVar(&cf.DefaultRequestTimeout, "timeout_request",
		1000*time.Millisecond, "Default timeout for regular request")

	fs.DurationVar(&cf.DefaultCoreTimeout, "core_timeout",
		1000*time.Millisecond, "Default Core timeout")

	fs.BoolVar(&cf.Banner, "banner",
		false, "Show startup banner log lines")

	fs.BoolVar(&cf.DisplayVersionOnly, "version",
		false, "Show version information and exit")

	fs.StringVar(&cf.CoreBindingKey, "core_binding_key",
		"voltha_backend_name", "The name of the meta-key whose value is the rw-core group to which the ofagent is bound")

	fs.IntVar(&cf.MaxConnectionRetries, "max_connection_retries",
		-1, "The number of retries to connect to a dependent component")

	fs.DurationVar(&cf.ConnectionRetryInterval, "connection_retry_interval",
		2*time.Second, "The number of seconds between each connection retry attempt")

	fs.DurationVar(&cf.LiveProbeInterval, "live_probe_interval",
		60*time.Second, "The number of seconds between liveness probes while in a live state")

	fs.DurationVar(&cf.NotLiveProbeInterval, "not_live_probe_interval",
		5*time.Second, "The number of seconds between liveness probes while in a not live state")

	fs.StringVar(&cf.ProbeAddress, "probe_address",
		":8080", "The address on which to listen to answer liveness and readiness probe queries over HTTP")

	fs.BoolVar(&(cf.TraceEnabled), "trace_enabled",
		false, "Whether to send logs to tracing agent?")

	fs.StringVar(&cf.TraceAgentAddress, "trace_agent_address",
		"127.0.0.1:6831", "The address of tracing agent to which span info should be sent")

	fs.BoolVar(&cf.LogCorrelationEnabled, "log_correlation_enabled",
		true, "Whether to enrich log statements with fields denoting operation being executed for achieving correlation?")

	fs.StringVar(&cf.VolthaStackID, "stack_id",
		"voltha", "ID for the current voltha stack")

	fs.DurationVar(&cf.BackoffRetryInitialInterval,
		"backoff_retry_initial_interval", 500*time.Millisecond,
		"The initial number of milliseconds an exponential backoff will wait before a retry")

	fs.DurationVar(&cf.BackoffRetryMaxElapsedTime,
		"backoff_retry_max_elapsed_time", 0*time.Second,
		"The maximum number of milliseconds an exponential backoff can elasped")

	fs.DurationVar(&cf.BackoffRetryMaxInterval, "backoff_retry_max_interval",
		1*time.Minute, "The maximum number of milliseconds of an exponential backoff interval")

	_ = fs.Parse(args)
}
