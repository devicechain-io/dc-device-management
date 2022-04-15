/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"
	"strconv"

	"github.com/devicechain-io/dc-device-management/model"
	"gorm.io/gorm"
)

// Find device type by unique id.
func (r *SchemaResolver) DeviceType(ctx context.Context, args struct {
	Id string
}) (*DeviceTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := model.NewApi(rdbmgr).DeviceTypeById(uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device type by unique token.
func (r *SchemaResolver) DeviceTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	found, err := model.NewApi(rdbmgr).DeviceTypeByToken(args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device types that match the given criteria.
func (r *SchemaResolver) DeviceTypes(ctx context.Context, args struct {
	Criteria model.DeviceTypeSearchCriteria
}) (*DeviceTypeSearchResultsResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	found, err := model.NewApi(rdbmgr).DeviceTypes(args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find device by unique id.
func (r *SchemaResolver) Device(ctx context.Context, args struct {
	Id string
}) (*DeviceResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := model.NewApi(rdbmgr).DeviceById(uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device by unique token.
func (r *SchemaResolver) DeviceByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	found, err := model.NewApi(rdbmgr).DeviceByToken(args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all devices that match the given criteria.
func (r *SchemaResolver) Devices(ctx context.Context, args struct {
	Criteria model.DeviceSearchCriteria
}) (*DeviceSearchResultsResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	found, err := model.NewApi(rdbmgr).Devices(args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find device relationship type by unique id.
func (r *SchemaResolver) DeviceRelationshipType(ctx context.Context, args struct {
	Id string
}) (*DeviceRelationshipTypeResolver, error) {
	found := model.DeviceRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device relationship type by unique token.
func (r *SchemaResolver) DeviceRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceRelationshipTypeResolver, error) {
	found := model.DeviceRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device relationship types that match the given criteria.
func (r *SchemaResolver) DeviceRelationshipTypes(ctx context.Context, args struct {
	Criteria model.DeviceRelationshipTypeSearchCriteria
}) (*DeviceRelationshipTypeSearchResultsResolver, error) {
	results := make([]model.DeviceRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.DeviceRelationshipType{}, nil, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.DeviceRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &DeviceRelationshipTypeSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find device relationship by unique id.
func (r *SchemaResolver) DeviceRelationship(ctx context.Context, args struct {
	Id string
}) (*DeviceRelationshipResolver, error) {
	found := model.DeviceRelationship{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Joins("SourceDevice").Joins("TargetDevice").Joins("RelationshipType").First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device relationships that match the given criteria.
func (r *SchemaResolver) DeviceRelationships(ctx context.Context, args struct {
	Criteria model.DeviceRelationshipSearchCriteria
}) (*DeviceRelationshipSearchResultsResolver, error) {
	results := make([]model.DeviceRelationship, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.DeviceRelationship{}, func(db *gorm.DB) *gorm.DB {
		return db.Preload("SourceDevice").Preload("TargetDevice").Preload("RelationshipType")
	}, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.DeviceRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &DeviceRelationshipSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find device group by unique id.
func (r *SchemaResolver) DeviceGroup(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupResolver, error) {
	found := model.DeviceGroup{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device group by unique token.
func (r *SchemaResolver) DeviceGroupByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceGroupResolver, error) {
	found := model.DeviceGroup{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device groups that match the given criteria.
func (r *SchemaResolver) DeviceGroups(ctx context.Context, args struct {
	Criteria model.DeviceGroupSearchCriteria
}) (*DeviceGroupSearchResultsResolver, error) {
	results := make([]model.DeviceGroup, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.DeviceGroup{}, nil, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.DeviceGroupSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &DeviceGroupSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find device group relationship type by unique id.
func (r *SchemaResolver) DeviceGroupRelationshipType(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupRelationshipTypeResolver, error) {
	found := model.DeviceGroupRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device group relationship type by unique token.
func (r *SchemaResolver) DeviceGroupRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceGroupRelationshipTypeResolver, error) {
	found := model.DeviceGroupRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device group relationship types that match the given criteria.
func (r *SchemaResolver) DeviceGroupRelationshipTypes(ctx context.Context, args struct {
	Criteria model.DeviceGroupRelationshipTypeSearchCriteria
}) (*DeviceGroupRelationshipTypeSearchResultsResolver, error) {
	results := make([]model.DeviceGroupRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.DeviceGroupRelationshipType{}, nil, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.DeviceGroupRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &DeviceGroupRelationshipTypeSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}

// Find device group relationship by unique id.
func (r *SchemaResolver) DeviceGroupRelationship(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupRelationshipResolver, error) {
	found := model.DeviceGroupRelationship{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device group relationships that match the given criteria.
func (r *SchemaResolver) DeviceGroupRelationships(ctx context.Context, args struct {
	Criteria model.DeviceGroupRelationshipSearchCriteria
}) (*DeviceGroupRelationshipSearchResultsResolver, error) {
	results := make([]model.DeviceGroupRelationship, 0)
	rdbmgr := r.GetRdbManager(ctx)
	db, pag := rdbmgr.ListOf(&model.DeviceGroupRelationship{}, func(db *gorm.DB) *gorm.DB {
		return db.Preload("DeviceGroup").Preload("Device").Preload("RelationshipType")
	}, args.Criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	found := model.DeviceGroupRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}

	// Return as resolver.
	return &DeviceGroupRelationshipSearchResultsResolver{
		M: found,
		S: r,
		C: ctx,
	}, nil
}
