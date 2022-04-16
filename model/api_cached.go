/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
)

type CachedApi struct {
	API *Api
}

// Create a new API instance.
func NewCachedApi(api *Api) *CachedApi {
	capi := &CachedApi{
		API: api,
	}
	return capi
}

// Get device type by id.
func (cached *CachedApi) DeviceTypeById(ctx context.Context, id uint) (*DeviceType, error) {
	// Load from cache if available.
	strid := fmt.Sprintf("%d", id)
	found := &DeviceType{}
	idcache := cached.API.RDB.GetRedisCache(CACHE_NAME_DEVICE_TYPE_BY_ID)
	tokencache := cached.API.RDB.GetRedisCache(CACHE_NAME_DEVICE_TYPE_BY_TOKEN)
	idcache.Get(context.Background(), strid, func(cache *cache.Cache, key string) {
		cache.Get(context.Background(), key, &found)
	})
	if found != nil {
		return found, nil
	}

	// Otherwise search in database.
	result, err := cached.API.DeviceTypeById(ctx, id)
	if err != nil {
		return nil, err
	}

	// Add to cache.
	idcache.Set(context.Background(), fmt.Sprintf("%d", id), found, time.Minute)
	tokencache.Set(context.Background(), found.Token, found, time.Minute)
	return result, nil
}
