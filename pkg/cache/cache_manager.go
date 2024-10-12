package cache

import (
	"sync"
	"time"
)

var (
	globalCacheManager Manager
	once               sync.Once
)

// InitializeCacheManager initializes the global cache manager with the specified cleanup interval
func InitializeCacheManager(cleanupInterval time.Duration) {
	once.Do(func() {
		globalCacheManager = NewManager(cleanupInterval)
	})
}

// GetGlobalCacheManager returns the initialized global cache manager
func GetGlobalCacheManager() Manager {
	return globalCacheManager
}

// Manager defines the interface for cache management operations.
type Manager interface {
	Get(key string) (interface{}, bool)
	Set(key string, data interface{}, expiration time.Duration)
}

// entry represents a single entry in the cache.
type entry struct {
	Data       interface{}
	Expiration time.Time
}

// manager is responsible for managing the cache
type manager struct {
	cache  map[string]entry
	mutex  *sync.RWMutex
	ticker *time.Ticker
}

// NewManager creates a new CacheManager with the specified cleanup interval
func NewManager(cleanupInterval time.Duration) Manager {
	cacheManager := &manager{
		cache:  make(map[string]entry),
		mutex:  &sync.RWMutex{},
		ticker: time.NewTicker(cleanupInterval),
	}

	go cacheManager.startCleanupRoutine()

	return cacheManager
}

// Set adds a new entry to the cache with the specified key, data, and expiration time
func (c *manager) Set(key string, data interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expirationTime := time.Now().Add(expiration)
	c.cache[key] = entry{Data: data, Expiration: expirationTime}
}

// Get retrieves the cached data associated with the specified key
func (c *manager) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, found := c.cache[key]
	if !found || time.Now().After(entry.Expiration) {
		return nil, false
	}

	return entry.Data, true
}

// startCleanupRoutine periodically cleans up expired entries from the cache
func (c *manager) startCleanupRoutine() {
	for range c.ticker.C {
		c.cleanupExpiredEntries()
	}
}

// cleanupExpiredEntries removes entries from the cache that have expired
func (c *manager) cleanupExpiredEntries() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	currentTime := time.Now()
	for key, entry := range c.cache {
		if currentTime.After(entry.Expiration) {
			delete(c.cache, key)
		}
	}
}
