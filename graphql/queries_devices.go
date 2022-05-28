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

// Find device type by unique id.
func (r *SchemaResolver) DeviceType(ctx context.Context, args struct {
	Id string
}) (*DeviceTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}

	found, err := api.DeviceTypeById(ctx, uint(id))
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
	api := r.GetApi(ctx)
	found, err := api.DeviceTypeByToken(ctx, args.Token)
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
	api := r.GetApi(ctx)
	found, err := api.DeviceTypes(ctx, args.Criteria)
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
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.DeviceById(ctx, uint(id))
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
	api := r.GetApi(ctx)
	found, err := api.DeviceByToken(ctx, args.Token)
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
	api := r.GetApi(ctx)
	found, err := api.Devices(ctx, args.Criteria)
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
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.DeviceRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device relationship type by unique token.
func (r *SchemaResolver) DeviceRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device relationship types that match the given criteria.
func (r *SchemaResolver) DeviceRelationshipTypes(ctx context.Context, args struct {
	Criteria model.DeviceRelationshipTypeSearchCriteria
}) (*DeviceRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find device relationship by unique id.
func (r *SchemaResolver) DeviceRelationship(ctx context.Context, args struct {
	Id string
}) (*DeviceRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.DeviceRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device relationship by unique token.
func (r *SchemaResolver) DeviceRelationshipByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceRelationshipResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceRelationshipByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device relationships that match the given criteria.
func (r *SchemaResolver) DeviceRelationships(ctx context.Context, args struct {
	Criteria model.DeviceRelationshipSearchCriteria
}) (*DeviceRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find device group by unique id.
func (r *SchemaResolver) DeviceGroup(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.DeviceGroupById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device group by unique token.
func (r *SchemaResolver) DeviceGroupByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceGroupResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceGroupByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device groups that match the given criteria.
func (r *SchemaResolver) DeviceGroups(ctx context.Context, args struct {
	Criteria model.DeviceGroupSearchCriteria
}) (*DeviceGroupSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceGroups(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceGroupSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find device group relationship type by unique id.
func (r *SchemaResolver) DeviceGroupRelationshipType(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.DeviceGroupRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device group relationship type by unique token.
func (r *SchemaResolver) DeviceGroupRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceGroupRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device group relationship types that match the given criteria.
func (r *SchemaResolver) DeviceGroupRelationshipTypes(ctx context.Context, args struct {
	Criteria model.DeviceGroupRelationshipTypeSearchCriteria
}) (*DeviceGroupRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceGroupRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceGroupRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find device group relationship by unique id.
func (r *SchemaResolver) DeviceGroupRelationship(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.DeviceGroupRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceGroupRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device group relationship by unique token.
func (r *SchemaResolver) DeviceGroupRelationshipByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceGroupRelationshipResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceGroupRelationshipByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceGroupRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device group relationships that match the given criteria.
func (r *SchemaResolver) DeviceGroupRelationships(ctx context.Context, args struct {
	Criteria model.DeviceGroupRelationshipSearchCriteria
}) (*DeviceGroupRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceGroupRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceGroupRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}
