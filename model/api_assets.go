/**
 * Copyright ©2022 AssetChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"context"

	"github.com/devicechain-io/dc-microservice/rdb"
	"gorm.io/gorm"
)

// Create a new asset type.
func (api *Api) CreateAssetType(ctx context.Context, request *AssetTypeCreateRequest) (*AssetType, error) {
	created := &AssetType{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		BrandedEntity: rdb.BrandedEntity{
			ImageUrl:        rdb.NullStrOf(request.ImageUrl),
			Icon:            rdb.NullStrOf(request.Icon),
			BackgroundColor: rdb.NullStrOf(request.BackgroundColor),
			ForegroundColor: rdb.NullStrOf(request.ForegroundColor),
			BorderColor:     rdb.NullStrOf(request.BorderColor),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing asset type.
func (api *Api) UpdateAssetType(ctx context.Context, token string,
	request *AssetTypeCreateRequest) (*AssetType, error) {
	found, err := api.AssetTypeByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	found.Token = request.Token
	found.Name = rdb.NullStrOf(request.Name)
	found.Description = rdb.NullStrOf(request.Description)
	found.ImageUrl = rdb.NullStrOf(request.ImageUrl)
	found.Icon = rdb.NullStrOf(request.Icon)
	found.BackgroundColor = rdb.NullStrOf(request.BackgroundColor)
	found.ForegroundColor = rdb.NullStrOf(request.ForegroundColor)
	found.BorderColor = rdb.NullStrOf(request.BorderColor)
	found.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(found)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

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

// Create a new asset.
func (api *Api) CreateAsset(ctx context.Context, request *AssetCreateRequest) (*Asset, error) {
	dtr, err := api.AssetTypeByToken(ctx, request.AssetTypeToken)
	if err != nil {
		return nil, err
	}

	created := &Asset{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
		AssetType: dtr,
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing asset.
func (api *Api) UpdateAsset(ctx context.Context, token string, request *AssetCreateRequest) (*Asset, error) {
	updated, err := api.AssetByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	// Update fields that changed.
	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	// Update asset type if changed.
	if request.AssetTypeToken != updated.AssetType.Token {
		dtr, err := api.AssetTypeByToken(ctx, request.AssetTypeToken)
		if err != nil {
			return nil, err
		}
		updated.AssetType = dtr
	}

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
}

// Get asset by id.
func (api *Api) AssetById(ctx context.Context, id uint) (*Asset, error) {
	found := &Asset{}
	result := api.RDB.Database
	result = result.Preload("AssetType")
	result = result.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get asset by token.
func (api *Api) AssetByToken(ctx context.Context, token string) (*Asset, error) {
	found := &Asset{}
	result := api.RDB.Database
	result = result.Preload("AssetType")
	result = result.First(&found, "token = ?", token)
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

// Create a new asset relationship type.
func (api *Api) CreateAssetRelationshipType(ctx context.Context, request *AssetRelationshipTypeCreateRequest) (*AssetRelationshipType, error) {
	created := &AssetRelationshipType{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing asset relationship type.
func (api *Api) UpdateAssetRelationshipType(ctx context.Context, token string,
	request *AssetRelationshipTypeCreateRequest) (*AssetRelationshipType, error) {
	updated, err := api.AssetRelationshipTypeByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
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

// Create a new asset relationship.
func (api *Api) CreateAssetRelationship(ctx context.Context, request *AssetRelationshipCreateRequest) (*AssetRelationship, error) {
	// Look up token references.
	source, err := api.AssetByToken(ctx, request.SourceAsset)
	if err != nil {
		return nil, err
	}
	target, err := api.AssetByToken(ctx, request.TargetAsset)
	if err != nil {
		return nil, err
	}
	rtype, err := api.AssetRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &AssetRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		SourceAsset:      *source,
		TargetAsset:      *target,
		RelationshipType: *rtype,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
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

// Get asset relationship by token.
func (api *Api) AssetRelationshipByToken(ctx context.Context, token string) (*AssetRelationship, error) {
	found := &AssetRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceAsset").Preload("TargetAsset").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
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

// Create a new asset group.
func (api *Api) CreateAssetGroup(ctx context.Context, request *AssetGroupCreateRequest) (*AssetGroup, error) {
	created := &AssetGroup{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		BrandedEntity: rdb.BrandedEntity{
			ImageUrl:        rdb.NullStrOf(request.ImageUrl),
			Icon:            rdb.NullStrOf(request.Icon),
			BackgroundColor: rdb.NullStrOf(request.BackgroundColor),
			ForegroundColor: rdb.NullStrOf(request.ForegroundColor),
			BorderColor:     rdb.NullStrOf(request.BorderColor),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing asset type.
func (api *Api) UpdateAssetGroup(ctx context.Context, token string,
	request *AssetGroupCreateRequest) (*AssetGroup, error) {
	updated, err := api.AssetGroupByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.ImageUrl = rdb.NullStrOf(request.ImageUrl)
	updated.Icon = rdb.NullStrOf(request.Icon)
	updated.BackgroundColor = rdb.NullStrOf(request.BackgroundColor)
	updated.ForegroundColor = rdb.NullStrOf(request.ForegroundColor)
	updated.BorderColor = rdb.NullStrOf(request.BorderColor)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
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

// Create a new asset group relationship type.
func (api *Api) CreateAssetGroupRelationshipType(ctx context.Context,
	request *AssetGroupRelationshipTypeCreateRequest) (*AssetGroupRelationshipType, error) {
	created := &AssetGroupRelationshipType{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing asset group relationship type.
func (api *Api) UpdateAssetGroupRelationshipType(ctx context.Context, token string,
	request *AssetGroupRelationshipTypeCreateRequest) (*AssetGroupRelationshipType, error) {
	updated, err := api.AssetGroupRelationshipTypeByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
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

// Create a new asset group relationship.
func (api *Api) CreateAssetGroupRelationship(ctx context.Context,
	request *AssetGroupRelationshipCreateRequest) (*AssetGroupRelationship, error) {

	// Look up token references.
	source, err := api.AssetGroupByToken(ctx, request.AssetGroup)
	if err != nil {
		return nil, err
	}
	target, err := api.AssetByToken(ctx, request.Asset)
	if err != nil {
		return nil, err
	}
	rtype, err := api.AssetGroupRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &AssetGroupRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		AssetGroup:       *source,
		Asset:            *target,
		RelationshipType: *rtype,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
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

// Get asset group relationship by token.
func (api *Api) AssetGroupRelationshipByToken(ctx context.Context, token string) (*AssetGroupRelationship, error) {
	found := &AssetGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("AssetGroup").Preload("Asset").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
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
