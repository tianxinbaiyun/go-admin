package cache

import (
	"time"
)

type Adapter interface {
	Connect()
	Get(key string) (string, error)
	Set(key string, val interface{}, expire int) error
	Del(key string) error
	HashGet(hk, key string) (string, error)
	HashDel(hk, key string) error
	Increase(key string) error
	Expire(key string, dur time.Duration) error
	SetQueue(name string, message Message) error
	GetQueue(name string, f func(message Message) error)
}

type Message interface {
	SetID(string)
	SetStream(string)
	SetValue(map[string]interface{})
	GetID() string
	GetStream() string
	GetValue() map[string]interface{}
}