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
package kvstore

import "context"

const (
	// Default timeout in seconds when making a kvstore request
	defaultKVGetTimeout = 5
	// Maximum channel buffer between publisher/subscriber goroutines
	maxClientChannelBufferSize = 10
)

// These constants represent the event types returned by the KV client
const (
	PUT = iota
	DELETE
	CONNECTIONDOWN
	UNKNOWN
)

// KVPair is a common wrapper for key-value pairs returned from the KV store
type KVPair struct {
	Key     string
	Value   interface{}
	Version int64
	Session string
	Lease   int64
}

// NewKVPair creates a new KVPair object
func NewKVPair(key string, value interface{}, session string, lease int64, version int64) *KVPair {
	kv := new(KVPair)
	kv.Key = key
	kv.Value = value
	kv.Session = session
	kv.Lease = lease
	kv.Version = version
	return kv
}

// Event is generated by the KV client when a key change is detected
type Event struct {
	EventType int
	Key       interface{}
	Value     interface{}
	Version   int64
}

// NewEvent creates a new Event object
func NewEvent(eventType int, key interface{}, value interface{}, version int64) *Event {
	evnt := new(Event)
	evnt.EventType = eventType
	evnt.Key = key
	evnt.Value = value
	evnt.Version = version

	return evnt
}

// Client represents the set of APIs a KV Client must implement
type Client interface {
	List(ctx context.Context, key string) (map[string]*KVPair, error)
	Get(ctx context.Context, key string) (*KVPair, error)
	Put(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
	Reserve(ctx context.Context, key string, value interface{}, ttl int64) (interface{}, error)
	ReleaseReservation(ctx context.Context, key string) error
	ReleaseAllReservations(ctx context.Context) error
	RenewReservation(ctx context.Context, key string) error
	Watch(ctx context.Context, key string, withPrefix bool) chan *Event
	AcquireLock(ctx context.Context, lockName string, timeout int) error
	ReleaseLock(lockName string) error
	IsConnectionUp(ctx context.Context) bool // timeout in second
	CloseWatch(key string, ch chan *Event)
	Close()
}
