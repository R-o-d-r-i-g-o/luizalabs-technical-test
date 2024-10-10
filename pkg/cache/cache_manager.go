package cache

import (
	"sync"
	"time"
)

var (
	// globalCacheManager is a singleton instance of the CacheManager
	globalCacheManager CacheManager
	once               sync.Once
)

// InitializeCacheManager initializes the global cache manager with the specified cleanup interval
func InitializeCacheManager(cleanupInterval time.Duration) {
	once.Do(func() {
		globalCacheManager = NewCacheManager(cleanupInterval)
	})
}

// GetGlobalCacheManager returns the initialized global cache manager
func GetGlobalCacheManager() CacheManager {
	return globalCacheManager
}

type CacheManager interface {
	Get(key string) (interface{}, bool)
	Set(key string, data interface{}, expiration time.Duration)
}

type cacheEntry struct {
	Data       interface{}
	Expiration time.Time
}

// cacheManager is responsible for managing the cache
type cacheManager struct {
	cache  map[string]cacheEntry
	mutex  *sync.RWMutex
	ticker *time.Ticker
}

// NewCacheManager creates a new CacheManager with the specified cleanup interval
func NewCacheManager(cleanupInterval time.Duration) CacheManager {
	cacheManager := &cacheManager{
		cache:  make(map[string]cacheEntry),
		mutex:  &sync.RWMutex{},
		ticker: time.NewTicker(cleanupInterval),
	}

	go cacheManager.startCleanupRoutine()

	return cacheManager
}

// Set adds a new entry to the cache with the specified key, data, and expiration time
func (c *cacheManager) Set(key string, data interface{}, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expirationTime := time.Now().Add(expiration)
	c.cache[key] = cacheEntry{Data: data, Expiration: expirationTime}
}

// Get retrieves the cached data associated with the specified key
func (c *cacheManager) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, found := c.cache[key]
	if !found || time.Now().After(entry.Expiration) {
		// Entry not found or has expired
		return nil, false
	}

	return entry.Data, true
}

// startCleanupRoutine periodically cleans up expired entries from the cache
func (c *cacheManager) startCleanupRoutine() {
	for range c.ticker.C {
		c.cleanupExpiredEntries()
	}
}

// cleanupExpiredEntries removes entries from the cache that have expired
func (c *cacheManager) cleanupExpiredEntries() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	currentTime := time.Now()
	for key, entry := range c.cache {
		if currentTime.After(entry.Expiration) {
			delete(c.cache, key)
		}
	}
}