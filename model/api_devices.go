/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"context"

	"gorm.io/gorm"
)

// Get device type by id.
func (api *Api) DeviceTypeById(ctx context.Context, id uint) (*DeviceType, error) {
	found := &DeviceType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device type by token.
func (api *Api) DeviceTypeByToken(ctx context.Context, token string) (*DeviceType, error) {
	found := &DeviceType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device types that meet criteria.
func (api *Api) DeviceTypes(ctx context.Context, criteria DeviceTypeSearchCriteria) (*DeviceTypeSearchResults, error) {
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
func (api *Api) DeviceById(ctx context.Context, id uint) (*Device, error) {
	found := &Device{}
	result := api.RDB.Database.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device by token.
func (api *Api) DeviceByToken(ctx context.Context, token string) (*Device, error) {
	found := &Device{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for devices that meet criteria.
func (api *Api) Devices(ctx context.Context, criteria DeviceSearchCriteria) (*DeviceSearchResults, error) {
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

// Get device relationship type by id.
func (api *Api) DeviceRelationshipTypeById(ctx context.Context, id uint) (*DeviceRelationshipType, error) {
	found := &DeviceRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device relationship type by token.
func (api *Api) DeviceRelationshipTypeByToken(ctx context.Context, token string) (*DeviceRelationshipType, error) {
	found := &DeviceRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device relationship types that meet criteria.
func (api *Api) DeviceRelationshipTypes(ctx context.Context,
	criteria DeviceRelationshipTypeSearchCriteria) (*DeviceRelationshipTypeSearchResults, error) {
	results := make([]DeviceRelationshipType, 0)
	db, pag := api.RDB.ListOf(&DeviceRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get device relationship by id.
func (api *Api) DeviceRelationshipById(ctx context.Context, id uint) (*DeviceRelationship, error) {
	found := &DeviceRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceDevice").Preload("TargetDevice").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device relationships that meet criteria.
func (api *Api) DeviceRelationships(ctx context.Context,
	criteria DeviceRelationshipSearchCriteria) (*DeviceRelationshipSearchResults, error) {
	results := make([]DeviceRelationship, 0)
	db, pag := api.RDB.ListOf(&DeviceRelationship{}, nil, criteria.Pagination)
	db.Preload("SourceDevice").Preload("TargetDevice").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get device group by id.
func (api *Api) DeviceGroupById(ctx context.Context, id uint) (*DeviceGroup, error) {
	found := &DeviceGroup{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device group by token.
func (api *Api) DeviceGroupByToken(ctx context.Context, token string) (*DeviceGroup, error) {
	found := &DeviceGroup{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device groups that meet criteria.
func (api *Api) DeviceGroups(ctx context.Context, criteria DeviceGroupSearchCriteria) (*DeviceGroupSearchResults, error) {
	results := make([]DeviceGroup, 0)
	db, pag := api.RDB.ListOf(&DeviceGroup{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceGroupSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get device group relationship type by id.
func (api *Api) DeviceGroupRelationshipTypeById(ctx context.Context, id uint) (*DeviceGroupRelationshipType, error) {
	found := &DeviceGroupRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device group relationship type by token.
func (api *Api) DeviceGroupRelationshipTypeByToken(ctx context.Context, token string) (*DeviceGroupRelationshipType, error) {
	found := &DeviceGroupRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device group relationship types that meet criteria.
func (api *Api) DeviceGroupRelationshipTypes(ctx context.Context,
	criteria DeviceGroupRelationshipTypeSearchCriteria) (*DeviceGroupRelationshipTypeSearchResults, error) {
	results := make([]DeviceGroupRelationshipType, 0)
	db, pag := api.RDB.ListOf(&DeviceGroupRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceGroupRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get device group relationship by id.
func (api *Api) DeviceGroupRelationshipById(ctx context.Context, id uint) (*DeviceGroupRelationship, error) {
	found := &DeviceGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("DeviceGroup").Preload("Device").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device group relationships that meet criteria.
func (api *Api) DeviceGroupRelationships(ctx context.Context,
	criteria DeviceGroupRelationshipSearchCriteria) (*DeviceGroupRelationshipSearchResults, error) {
	results := make([]DeviceGroupRelationship, 0)
	db, pag := api.RDB.ListOf(&DeviceGroupRelationship{}, nil, criteria.Pagination)
	db.Preload("DeviceGroup").Preload("Device").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceGroupRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}
