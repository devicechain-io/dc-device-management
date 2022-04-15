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

	"github.com/devicechain-io/dc-microservice/rdb"
	"github.com/go-redis/cache/v8"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Api struct {
	RDB *rdb.RdbManager
}

// Create a new API instance.
func NewApi(rdb *rdb.RdbManager) *Api {
	api := &Api{}
	api.RDB = rdb
	return api
}

// Get device type by id.
func (api *Api) DeviceTypeById(id uint) (*DeviceType, error) {
	// Load from cache if available.
	strid := fmt.Sprintf("%d", id)
	found := &DeviceType{}
	idcache := api.RDB.GetRedisCache(CACHE_NAME_DEVICE_TYPE_BY_ID)
	idcache.Get(context.Background(), strid, func(cache *cache.Cache, key string) {
		cache.Get(context.Background(), key, &found)
	})
	if found != nil {
		return found, nil
	}

	// Otherwise search in database.
	log.Info().Msg("Loading device type by id from database.")
	result := api.RDB.Database.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Add to cache.
	tokencache := api.RDB.GetRedisCache(CACHE_NAME_DEVICE_TYPE_BY_TOKEN)
	idcache.Set(context.Background(), fmt.Sprintf("%d", id), found, time.Minute)
	tokencache.Set(context.Background(), found.Token, found, time.Minute)
	return found, nil
}

// Get device type by token.
func (api *Api) DeviceTypeByToken(token string) (*DeviceType, error) {
	found := &DeviceType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device types that meet criteria.
func (api *Api) DeviceTypes(criteria DeviceTypeSearchCriteria) (*DeviceTypeSearchResults, error) {
	results := make([]DeviceType, 0)
	db, pag := api.RDB.ListOf(&DeviceType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get device by id.
func (api *Api) DeviceById(id uint) (*Device, error) {
	found := &Device{}
	result := api.RDB.Database.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device by token.
func (api *Api) DeviceByToken(token string) (*Device, error) {
	found := &Device{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for devices that meet criteria.
func (api *Api) Devices(criteria DeviceSearchCriteria) (*DeviceSearchResults, error) {
	results := make([]Device, 0)
	db, pag := api.RDB.ListOf(&Device{}, func(result *gorm.DB) *gorm.DB {
		if criteria.DeviceTypeToken != nil {
			result = result.Where("device_type_id = (?)",
				api.RDB.Database.Model(&DeviceType{}).Select("id").Where("token = ?", criteria.DeviceTypeToken))
		}
		return result.Preload("DeviceType")
	}, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}
