package context

import (
	"MyBalance/internal/http/requesto/rand"
	"MyBalance/internal/projkeys"
	"context"
	"fmt"
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

func NewFromPrev(oldCtx Context, name string) Context {
	mtx := sync.RWMutex{}
	ctx := Context{Context: context.Background(), dataMtx: &mtx, dataMap: map[string]interface{}{}}
	ctx.Set(projkeys.Name, name)

	for _, key := range oldCtx.GetKeys() {
		val, _ := oldCtx.Get(key)
		ctx.Set(key, val)
	}
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

func (c *Context) GetString(key string) (string, error) {
	value, exists := c.Get(key)
	if exists {
		str, ok := value.(string)
		if !ok {
			return "", fmt.Errorf("key %s is not string", key)
		}
		return str, nil
	}
	return "", fmt.Errorf("key %s not found", key)
}

func (c *Context) GetStringOptional(key, defaultValue string) string {
	value, exists := c.Get(key)
	if exists {
		str, ok := value.(string)
		if !ok {
			return defaultValue
		}
		return str
	}
	return defaultValue
}

func (c *Context) GetIntOptional(key string, defaultValue int) int {
	value, exists := c.Get(key)
	if exists {
		valConverted, ok := value.(int)
		if !ok {
			return defaultValue
		}
		return valConverted
	}
	return defaultValue
}

func (c *Context) GetKeys() []string {
	if c.dataMtx != nil {
		c.dataMtx.RLock()
		defer c.dataMtx.RUnlock()
		keys := make([]string, 0, len(c.dataMap))
		for key := range c.dataMap {
			keys = append(keys, key)
		}
		return keys
	}
	return nil
}
