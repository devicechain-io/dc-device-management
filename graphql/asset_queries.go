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
}) ([]*AssetTypeResolver, error) {
	list := make([]model.AssetType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*AssetTypeResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&AssetTypeResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
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
}) ([]*AssetResolver, error) {
	list := make([]model.Asset, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Debug().Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize))
	if args.Criteria.AssetTypeToken != nil {
		result = result.Joins("AssetType").Where("asset_type_id = (?)",
			rdbmgr.Database.Model(&model.AssetType{}).Select("id").Where("token = ?", args.Criteria.AssetTypeToken)).Find(&list)
	} else {
		result = result.Joins("AssetType").Find(&list)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*AssetResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&AssetResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
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
}) ([]*AssetRelationshipTypeResolver, error) {
	list := make([]model.AssetRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*AssetRelationshipTypeResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&AssetRelationshipTypeResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
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
}) ([]*AssetRelationshipResolver, error) {
	list := make([]model.AssetRelationship, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize))
	result = result.Joins("SourceAsset").Joins("TargetAsset").Joins("RelationshipType")
	result = result.Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*AssetRelationshipResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&AssetRelationshipResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
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
}) ([]*AssetGroupResolver, error) {
	list := make([]model.AssetGroup, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*AssetGroupResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&AssetGroupResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
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
}) ([]*AssetGroupRelationshipTypeResolver, error) {
	list := make([]model.AssetGroupRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*AssetGroupRelationshipTypeResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&AssetGroupRelationshipTypeResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
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
}) ([]*AssetGroupRelationshipResolver, error) {
	list := make([]model.AssetGroupRelationship, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize))
	result = result.Joins("AssetGroup").Joins("Asset").Joins("RelationshipType")
	result = result.Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*AssetGroupRelationshipResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&AssetGroupRelationshipResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
}
