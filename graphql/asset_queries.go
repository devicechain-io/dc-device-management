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
)

// Find asset type by unique id.
func (r *SchemaResolver) AssetType(ctx context.Context, args struct {
	Id string
}) (*AssetTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}

	found, err := api.AssetTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AssetTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset type by unique token.
func (r *SchemaResolver) AssetTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AssetTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AssetTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset types that match the given criteria.
func (r *SchemaResolver) AssetTypes(ctx context.Context, args struct {
	Criteria model.AssetTypeSearchCriteria
}) (*AssetTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AssetTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset by unique id.
func (r *SchemaResolver) Asset(ctx context.Context, args struct {
	Id string
}) (*AssetResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AssetById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AssetResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset by unique token.
func (r *SchemaResolver) AssetByToken(ctx context.Context, args struct {
	Token string
}) (*AssetResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AssetResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all assets that match the given criteria.
func (r *SchemaResolver) Assets(ctx context.Context, args struct {
	Criteria model.AssetSearchCriteria
}) (*AssetSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.Assets(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AssetSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset relationship type by unique id.
func (r *SchemaResolver) AssetRelationshipType(ctx context.Context, args struct {
	Id string
}) (*AssetRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AssetRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AssetRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset relationship type by unique token.
func (r *SchemaResolver) AssetRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AssetRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AssetRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset relationship types that match the given criteria.
func (r *SchemaResolver) AssetRelationshipTypes(ctx context.Context, args struct {
	Criteria model.AssetRelationshipTypeSearchCriteria
}) (*AssetRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AssetRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset relationship by unique id.
func (r *SchemaResolver) AssetRelationship(ctx context.Context, args struct {
	Id string
}) (*AssetRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AssetRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AssetRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset relationships that match the given criteria.
func (r *SchemaResolver) AssetRelationships(ctx context.Context, args struct {
	Criteria model.AssetRelationshipSearchCriteria
}) (*AssetRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AssetRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset group by unique id.
func (r *SchemaResolver) AssetGroup(ctx context.Context, args struct {
	Id string
}) (*AssetGroupResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AssetGroupById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AssetGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset group by unique token.
func (r *SchemaResolver) AssetGroupByToken(ctx context.Context, args struct {
	Token string
}) (*AssetGroupResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetGroupByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AssetGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset groups that match the given criteria.
func (r *SchemaResolver) AssetGroups(ctx context.Context, args struct {
	Criteria model.AssetGroupSearchCriteria
}) (*AssetGroupSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetGroups(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AssetGroupSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset group relationship type by unique id.
func (r *SchemaResolver) AssetGroupRelationshipType(ctx context.Context, args struct {
	Id string
}) (*AssetGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AssetGroupRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AssetGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find asset group relationship type by unique token.
func (r *SchemaResolver) AssetGroupRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AssetGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetGroupRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AssetGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset group relationship types that match the given criteria.
func (r *SchemaResolver) AssetGroupRelationshipTypes(ctx context.Context, args struct {
	Criteria model.AssetGroupRelationshipTypeSearchCriteria
}) (*AssetGroupRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetGroupRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AssetGroupRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find asset group relationship by unique id.
func (r *SchemaResolver) AssetGroupRelationship(ctx context.Context, args struct {
	Id string
}) (*AssetGroupRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AssetGroupRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AssetGroupRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all asset group relationships that match the given criteria.
func (r *SchemaResolver) AssetGroupRelationships(ctx context.Context, args struct {
	Criteria model.AssetGroupRelationshipSearchCriteria
}) (*AssetGroupRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AssetGroupRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AssetGroupRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}
