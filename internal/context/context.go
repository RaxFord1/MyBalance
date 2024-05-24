package context

import (
	"MyBalance/internal/http/requesto/rand"
	"MyBalance/internal/projkeys"
	"context"
	"log"
	"sync"
)

type Context struct {
	context.Context
	dataMtx *sync.RWMutex
	dataMap map[string]interface{}
}

type Config struct {
	Version string
	Tag     string
}

var (
	cfg Config
)

func Init(config Config) {
	cfg = config
}

func New() Context {
	mtx := sync.RWMutex{}
	ctx := Context{Context: context.Background(), dataMtx: &mtx, dataMap: map[string]interface{}{}}
	druid, err := rand.GenerateTraceId()
	if err != nil {
		log.Println("error create druid for context. err:", err)
	}
	ctx.SetString(projkeys.Druid, druid)
	ctx.SetString(projkeys.Version, cfg.Version)
	ctx.SetString(projkeys.Tag, cfg.Tag)
	return ctx
}

func Named(name string) Context {
	ctx := New()
	ctx.Set(projkeys.Name, name)
	return ctx
}

func (c *Context) SetString(key, value string) {
	c.Set(key, value)
}

func (c *Context) Set(key string, value interface{}) {
	if c.dataMtx != nil {
		c.dataMtx.Lock()
		defer c.dataMtx.Unlock()
		c.dataMap[key] = value
	}
}

func (c *Context) Get(key string) (interface{}, bool) {
	if c.dataMtx != nil {
		c.dataMtx.RLock()
		defer c.dataMtx.RUnlock()
		value, exists := c.dataMap[key]
		return value, exists
	}
	return nil, false
}

func (c *Context) GetString(key string) (string, bool) {
	if c.dataMtx != nil {
		c.dataMtx.RLock()
		defer c.dataMtx.RUnlock()
		value, exists := c.dataMap[key]
		if exists {
			str, ok := value.(string)
			if !ok {
				return "", false
			}
			return str, exists
		}
	}
	return "", false
}
