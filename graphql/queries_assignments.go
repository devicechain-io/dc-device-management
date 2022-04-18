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

// Find device assignment status by unique id.
func (r *SchemaResolver) DeviceAssignmentStatus(ctx context.Context, args struct {
	Id string
}) (*DeviceAssignmentStatusResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}

	found, err := api.DeviceAssignmentStatusById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentStatusResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device assignment status by unique token.
func (r *SchemaResolver) DeviceAssignmentStatusByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceAssignmentStatusResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceAssignmentStatusByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentStatusResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device assignment statuses that match the given criteria.
func (r *SchemaResolver) DeviceAssignmentStatuses(ctx context.Context, args struct {
	Criteria model.DeviceAssignmentStatusSearchCriteria
}) (*DeviceAssignmentStatusSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceAssignmentStatuses(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceAssignmentStatusSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}

// Find device assignment by unique id.
func (r *SchemaResolver) DeviceAssignment(ctx context.Context, args struct {
	Id string
}) (*DeviceAssignmentResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}

	found, err := api.DeviceAssignmentById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device assignment by unique token.
func (r *SchemaResolver) DeviceAssignmentByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceAssignmentResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceAssignmentByToken(ctx, args.Token)
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentResolver{
		M: *found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device assignments that match the given criteria.
func (r *SchemaResolver) DeviceAssignments(ctx context.Context, args struct {
	Criteria model.DeviceAssignmentSearchCriteria
}) (*DeviceAssignmentSearchResultsResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceAssignments(ctx, args.Criteria)
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	return &DeviceAssignmentSearchResultsResolver{
		M: *found,
		S: r,
		C: ctx,
	}, nil
}
