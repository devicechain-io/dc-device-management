/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"

	"github.com/devicechain-io/dc-device-management/model"
	"gorm.io/gorm"
)

// Find asset type by unique id.
func (r *SchemaResolver) AssetType(ctx context.Context, args struct {
	Id string
}) (*AssetTypeResolver, error) {
	found := model.AssetType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset type by unique token.
func (r *SchemaResolver) AssetTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AssetTypeResolver, error) {
	found := model.AssetType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset types that match the given criteria.
func (r *SchemaResolver) AssetTypes(ctx context.Context, args struct {
	Criteria model.AssetTypeSearchCriteria
}) (*AssetTypeSearchResultsResolver, error) {
	results := make([]model.AssetType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.AssetType{}, nil, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.AssetTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &AssetTypeSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset by unique id.
func (r *SchemaResolver) Asset(ctx context.Context, args struct {
	Id string
}) (*AssetResolver, error) {
	found := model.Asset{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Joins("AssetType").First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset by unique token.
func (r *SchemaResolver) AssetByToken(ctx context.Context, args struct {
	Token string
}) (*AssetResolver, error) {
	found := model.Asset{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Joins("AssetType").First(&found, "\"assets\".token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all assets that match the given criteria.
func (r *SchemaResolver) Assets(ctx context.Context, args struct {
	Criteria model.AssetSearchCriteria
}) (*AssetSearchResultsResolver, error) {
	results := make([]model.Asset, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(
		&model.Device{},
		func(result *gorm.DB) *gorm.DB {
			if args.Criteria.AssetTypeToken != nil {
				return result.Joins("AssetType").Where("asset_type_id = (?)",
					rdbmgr.Database.Model(&model.AssetType{}).Select("id").Where("token = ?", args.Criteria.AssetTypeToken)).Find(&results)
			} else {
				return result.Joins("AssetType").Find(&results)
			}
		}, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.AssetSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &AssetSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset relationship type by unique id.
func (r *SchemaResolver) AssetRelationshipType(ctx context.Context, args struct {
	Id string
}) (*AssetRelationshipTypeResolver, error) {
	found := model.AssetRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset relationship type by unique token.
func (r *SchemaResolver) AssetRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AssetRelationshipTypeResolver, error) {
	found := model.AssetRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset relationship types that match the given criteria.
func (r *SchemaResolver) AssetRelationshipTypes(ctx context.Context, args struct {
	Criteria model.AssetRelationshipTypeSearchCriteria
}) (*AssetRelationshipTypeSearchResultsResolver, error) {
	results := make([]model.AssetRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.AssetRelationshipType{}, nil, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.AssetRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &AssetRelationshipTypeSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset relationship by unique id.
func (r *SchemaResolver) AssetRelationship(ctx context.Context, args struct {
	Id string
}) (*AssetRelationshipResolver, error) {
	found := model.AssetRelationship{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetRelationshipResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset relationships that match the given criteria.
func (r *SchemaResolver) AssetRelationships(ctx context.Context, args struct {
	Criteria model.AssetRelationshipSearchCriteria
}) (*AssetRelationshipSearchResultsResolver, error) {
	results := make([]model.AssetRelationship, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.AssetRelationship{}, func(db *gorm.DB) *gorm.DB {
		return db.Preload("SourceAsset").Preload("TargetAsset").Preload("RelationshipType")
	}, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.AssetRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &AssetRelationshipSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset group by unique id.
func (r *SchemaResolver) AssetGroup(ctx context.Context, args struct {
	Id string
}) (*AssetGroupResolver, error) {
	found := model.AssetGroup{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetGroupResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset group by unique token.
func (r *SchemaResolver) AssetGroupByToken(ctx context.Context, args struct {
	Token string
}) (*AssetGroupResolver, error) {
	found := model.AssetGroup{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetGroupResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset groups that match the given criteria.
func (r *SchemaResolver) AssetGroups(ctx context.Context, args struct {
	Criteria model.AssetGroupSearchCriteria
}) (*AssetGroupSearchResultsResolver, error) {
	results := make([]model.AssetGroup, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.AssetGroup{}, nil, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.AssetGroupSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &AssetGroupSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset group relationship type by unique id.
func (r *SchemaResolver) AssetGroupRelationshipType(ctx context.Context, args struct {
	Id string
}) (*AssetGroupRelationshipTypeResolver, error) {
	found := model.AssetGroupRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetGroupRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset group relationship type by unique token.
func (r *SchemaResolver) AssetGroupRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AssetGroupRelationshipTypeResolver, error) {
	found := model.AssetGroupRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetGroupRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset group relationship types that match the given criteria.
func (r *SchemaResolver) AssetGroupRelationshipTypes(ctx context.Context, args struct {
	Criteria model.AssetGroupRelationshipTypeSearchCriteria
}) (*AssetGroupRelationshipTypeSearchResultsResolver, error) {
	results := make([]model.AssetGroupRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.AssetGroupRelationshipType{}, nil, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.AssetGroupRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &AssetGroupRelationshipTypeSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset group relationship by unique id.
func (r *SchemaResolver) AssetGroupRelationship(ctx context.Context, args struct {
	Id string
}) (*AssetGroupRelationshipResolver, error) {
	found := model.AssetGroupRelationship{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetGroupRelationshipResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset group relationships that match the given criteria.
func (r *SchemaResolver) AssetGroupRelationships(ctx context.Context, args struct {
	Criteria model.AssetGroupRelationshipSearchCriteria
}) (*AssetGroupRelationshipSearchResultsResolver, error) {
	results := make([]model.AssetGroupRelationship, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.AssetGroupRelationship{}, func(db *gorm.DB) *gorm.DB {
		return db.Preload("AssetGroup").Preload("Asset").Preload("RelationshipType")
	}, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.AssetGroupRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &AssetGroupRelationshipSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}
