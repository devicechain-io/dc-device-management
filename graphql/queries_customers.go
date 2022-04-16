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

// Find customer type by unique id.
func (r *SchemaResolver) CustomerType(ctx context.Context, args struct {
	Id string
}) (*CustomerTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}

	found, err := api.CustomerTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &CustomerTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find customer type by unique token.
func (r *SchemaResolver) CustomerTypeByToken(ctx context.Context, args struct {
	Token string
}) (*CustomerTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &CustomerTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all customer types that match the given criteria.
func (r *SchemaResolver) CustomerTypes(ctx context.Context, args struct {
	Criteria model.CustomerTypeSearchCriteria
}) (*CustomerTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &CustomerTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find customer by unique id.
func (r *SchemaResolver) Customer(ctx context.Context, args struct {
	Id string
}) (*CustomerResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.CustomerById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &CustomerResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find customer by unique token.
func (r *SchemaResolver) CustomerByToken(ctx context.Context, args struct {
	Token string
}) (*CustomerResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &CustomerResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all customers that match the given criteria.
func (r *SchemaResolver) Customers(ctx context.Context, args struct {
	Criteria model.CustomerSearchCriteria
}) (*CustomerSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.Customers(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &CustomerSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find customer relationship type by unique id.
func (r *SchemaResolver) CustomerRelationshipType(ctx context.Context, args struct {
	Id string
}) (*CustomerRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.CustomerRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &CustomerRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find customer relationship type by unique token.
func (r *SchemaResolver) CustomerRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*CustomerRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &CustomerRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all customer relationship types that match the given criteria.
func (r *SchemaResolver) CustomerRelationshipTypes(ctx context.Context, args struct {
	Criteria model.CustomerRelationshipTypeSearchCriteria
}) (*CustomerRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &CustomerRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find customer relationship by unique id.
func (r *SchemaResolver) CustomerRelationship(ctx context.Context, args struct {
	Id string
}) (*CustomerRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.CustomerRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &CustomerRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all customer relationships that match the given criteria.
func (r *SchemaResolver) CustomerRelationships(ctx context.Context, args struct {
	Criteria model.CustomerRelationshipSearchCriteria
}) (*CustomerRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &CustomerRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find customer group by unique id.
func (r *SchemaResolver) CustomerGroup(ctx context.Context, args struct {
	Id string
}) (*CustomerGroupResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.CustomerGroupById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &CustomerGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find customer group by unique token.
func (r *SchemaResolver) CustomerGroupByToken(ctx context.Context, args struct {
	Token string
}) (*CustomerGroupResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerGroupByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &CustomerGroupResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all customer groups that match the given criteria.
func (r *SchemaResolver) CustomerGroups(ctx context.Context, args struct {
	Criteria model.CustomerGroupSearchCriteria
}) (*CustomerGroupSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerGroups(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &CustomerGroupSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find customer group relationship type by unique id.
func (r *SchemaResolver) CustomerGroupRelationshipType(ctx context.Context, args struct {
	Id string
}) (*CustomerGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.CustomerGroupRelationshipTypeById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &CustomerGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find customer group relationship type by unique token.
func (r *SchemaResolver) CustomerGroupRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*CustomerGroupRelationshipTypeResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerGroupRelationshipTypeByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &CustomerGroupRelationshipTypeResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all customer group relationship types that match the given criteria.
func (r *SchemaResolver) CustomerGroupRelationshipTypes(ctx context.Context, args struct {
	Criteria model.CustomerGroupRelationshipTypeSearchCriteria
}) (*CustomerGroupRelationshipTypeSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerGroupRelationshipTypes(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &CustomerGroupRelationshipTypeSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find customer group relationship by unique id.
func (r *SchemaResolver) CustomerGroupRelationship(ctx context.Context, args struct {
	Id string
}) (*CustomerGroupRelationshipResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}
	found, err := api.CustomerGroupRelationshipById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &CustomerGroupRelationshipResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all customer group relationships that match the given criteria.
func (r *SchemaResolver) CustomerGroupRelationships(ctx context.Context, args struct {
	Criteria model.CustomerGroupRelationshipSearchCriteria
}) (*CustomerGroupRelationshipSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.CustomerGroupRelationships(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &CustomerGroupRelationshipSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}
