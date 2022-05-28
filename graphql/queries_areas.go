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

// Find area type by unique id.
func (r *SchemaResolver) AreaType(ctx context.Context, args struct {
	Id string
}) (*AreaTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}

	found, err := api.AreaTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AreaTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find area type by unique token.
func (r *SchemaResolver) AreaTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AreaTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AreaTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all area types that match the given criteria.
func (r *SchemaResolver) AreaTypes(ctx context.Context, args struct {
	Criteria model.AreaTypeSearchCriteria
}) (*AreaTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AreaTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find area by unique id.
func (r *SchemaResolver) Area(ctx context.Context, args struct {
	Id string
}) (*AreaResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AreaById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AreaResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find area by unique token.
func (r *SchemaResolver) AreaByToken(ctx context.Context, args struct {
	Token string
}) (*AreaResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AreaResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all areas that match the given criteria.
func (r *SchemaResolver) Areas(ctx context.Context, args struct {
	Criteria model.AreaSearchCriteria
}) (*AreaSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.Areas(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AreaSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find area relationship type by unique id.
func (r *SchemaResolver) AreaRelationshipType(ctx context.Context, args struct {
	Id string
}) (*AreaRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AreaRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AreaRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find area relationship type by unique token.
func (r *SchemaResolver) AreaRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AreaRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AreaRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all area relationship types that match the given criteria.
func (r *SchemaResolver) AreaRelationshipTypes(ctx context.Context, args struct {
	Criteria model.AreaRelationshipTypeSearchCriteria
}) (*AreaRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AreaRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find area relationship by unique id.
func (r *SchemaResolver) AreaRelationship(ctx context.Context, args struct {
	Id string
}) (*AreaRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AreaRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AreaRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find area relationship by unique token.
func (r *SchemaResolver) AreaRelationshipByToken(ctx context.Context, args struct {
	Token string
}) (*AreaRelationshipResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaRelationshipByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AreaRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all area relationships that match the given criteria.
func (r *SchemaResolver) AreaRelationships(ctx context.Context, args struct {
	Criteria model.AreaRelationshipSearchCriteria
}) (*AreaRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AreaRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find area group by unique id.
func (r *SchemaResolver) AreaGroup(ctx context.Context, args struct {
	Id string
}) (*AreaGroupResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AreaGroupById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AreaGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find area group by unique token.
func (r *SchemaResolver) AreaGroupByToken(ctx context.Context, args struct {
	Token string
}) (*AreaGroupResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaGroupByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AreaGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all area groups that match the given criteria.
func (r *SchemaResolver) AreaGroups(ctx context.Context, args struct {
	Criteria model.AreaGroupSearchCriteria
}) (*AreaGroupSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaGroups(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AreaGroupSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find area group relationship type by unique id.
func (r *SchemaResolver) AreaGroupRelationshipType(ctx context.Context, args struct {
	Id string
}) (*AreaGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AreaGroupRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AreaGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find area group relationship type by unique token.
func (r *SchemaResolver) AreaGroupRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*AreaGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaGroupRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AreaGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all area group relationship types that match the given criteria.
func (r *SchemaResolver) AreaGroupRelationshipTypes(ctx context.Context, args struct {
	Criteria model.AreaGroupRelationshipTypeSearchCriteria
}) (*AreaGroupRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaGroupRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AreaGroupRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find area group relationship by unique id.
func (r *SchemaResolver) AreaGroupRelationship(ctx context.Context, args struct {
	Id string
}) (*AreaGroupRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.AreaGroupRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &AreaGroupRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find area group relationship by unique token.
func (r *SchemaResolver) AreaGroupRelationshipByToken(ctx context.Context, args struct {
	Token string
}) (*AreaGroupRelationshipResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaGroupRelationshipByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &AreaGroupRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all area group relationships that match the given criteria.
func (r *SchemaResolver) AreaGroupRelationships(ctx context.Context, args struct {
	Criteria model.AreaGroupRelationshipSearchCriteria
}) (*AreaGroupRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.AreaGroupRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &AreaGroupRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}
