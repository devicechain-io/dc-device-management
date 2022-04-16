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

// Get asset type by id.
func (api *Api) AssetTypeById(ctx context.Context, id uint) (*AssetType, error) {
	found := &AssetType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get asset type by token.
func (api *Api) AssetTypeByToken(ctx context.Context, token string) (*AssetType, error) {
	found := &AssetType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for asset types that meet criteria.
func (api *Api) AssetTypes(ctx context.Context, criteria AssetTypeSearchCriteria) (*AssetTypeSearchResults, error) {
	results := make([]AssetType, 0)
	db, pag := api.RDB.ListOf(&AssetType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AssetTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get asset by id.
func (api *Api) AssetById(ctx context.Context, id uint) (*Asset, error) {
	found := &Asset{}
	result := api.RDB.Database.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get asset by token.
func (api *Api) AssetByToken(ctx context.Context, token string) (*Asset, error) {
	found := &Asset{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for assets that meet criteria.
func (api *Api) Assets(ctx context.Context, criteria AssetSearchCriteria) (*AssetSearchResults, error) {
	results := make([]Asset, 0)
	db, pag := api.RDB.ListOf(&Asset{}, func(result *gorm.DB) *gorm.DB {
		if criteria.AssetTypeToken != nil {
			result = result.Where("asset_type_id = (?)",
				api.RDB.Database.Model(&AssetType{}).Select("id").Where("token = ?", criteria.AssetTypeToken))
		}
		return result.Preload("AssetType")
	}, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AssetSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get asset relationship type by id.
func (api *Api) AssetRelationshipTypeById(ctx context.Context, id uint) (*AssetRelationshipType, error) {
	found := &AssetRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get asset relationship type by token.
func (api *Api) AssetRelationshipTypeByToken(ctx context.Context, token string) (*AssetRelationshipType, error) {
	found := &AssetRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for asset relationship types that meet criteria.
func (api *Api) AssetRelationshipTypes(ctx context.Context,
	criteria AssetRelationshipTypeSearchCriteria) (*AssetRelationshipTypeSearchResults, error) {
	results := make([]AssetRelationshipType, 0)
	db, pag := api.RDB.ListOf(&AssetRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AssetRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get asset relationship by id.
func (api *Api) AssetRelationshipById(ctx context.Context, id uint) (*AssetRelationship, error) {
	found := &AssetRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceAsset").Preload("TargetAsset").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for asset relationships that meet criteria.
func (api *Api) AssetRelationships(ctx context.Context,
	criteria AssetRelationshipSearchCriteria) (*AssetRelationshipSearchResults, error) {
	results := make([]AssetRelationship, 0)
	db, pag := api.RDB.ListOf(&AssetRelationship{}, nil, criteria.Pagination)
	db.Preload("SourceAsset").Preload("TargetAsset").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AssetRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get asset group by id.
func (api *Api) AssetGroupById(ctx context.Context, id uint) (*AssetGroup, error) {
	found := &AssetGroup{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get asset group by token.
func (api *Api) AssetGroupByToken(ctx context.Context, token string) (*AssetGroup, error) {
	found := &AssetGroup{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for asset groups that meet criteria.
func (api *Api) AssetGroups(ctx context.Context, criteria AssetGroupSearchCriteria) (*AssetGroupSearchResults, error) {
	results := make([]AssetGroup, 0)
	db, pag := api.RDB.ListOf(&AssetGroup{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AssetGroupSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get asset group relationship type by id.
func (api *Api) AssetGroupRelationshipTypeById(ctx context.Context, id uint) (*AssetGroupRelationshipType, error) {
	found := &AssetGroupRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get asset group relationship type by token.
func (api *Api) AssetGroupRelationshipTypeByToken(ctx context.Context, token string) (*AssetGroupRelationshipType, error) {
	found := &AssetGroupRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for asset group relationship types that meet criteria.
func (api *Api) AssetGroupRelationshipTypes(ctx context.Context,
	criteria AssetGroupRelationshipTypeSearchCriteria) (*AssetGroupRelationshipTypeSearchResults, error) {
	results := make([]AssetGroupRelationshipType, 0)
	db, pag := api.RDB.ListOf(&AssetGroupRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AssetGroupRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get asset group relationship by id.
func (api *Api) AssetGroupRelationshipById(ctx context.Context, id uint) (*AssetGroupRelationship, error) {
	found := &AssetGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("AssetGroup").Preload("Asset").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for asset group relationships that meet criteria.
func (api *Api) AssetGroupRelationships(ctx context.Context,
	criteria AssetGroupRelationshipSearchCriteria) (*AssetGroupRelationshipSearchResults, error) {
	results := make([]AssetGroupRelationship, 0)
	db, pag := api.RDB.ListOf(&AssetGroupRelationship{}, nil, criteria.Pagination)
	db.Preload("AssetGroup").Preload("Asset").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AssetGroupRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}
