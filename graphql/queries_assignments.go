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

// Find device assignment statuses by unique id.
func (r *SchemaResolver) DeviceAssignmentStatusesById(ctx context.Context, args struct {
	Ids []string
}) ([]*DeviceAssignmentStatusResolver, error) {
	api := r.GetApi(ctx)
	ids, err := r.asUintIds(args.Ids)
	if err != nil {
		return nil, err
	}

	found, err := api.DeviceAssignmentStatusesById(ctx, ids)
	if err != nil {
		return nil, err
	}

	result := make([]*DeviceAssignmentStatusResolver, 0)
	for _, dt := range found {
		dtr := &DeviceAssignmentStatusResolver{
			M: *dt,
			S: r,
			C: ctx,
		}
		result = append(result, dtr)
	}
	return result, nil
}

// Find device assignment statuses by unique token.
func (r *SchemaResolver) DeviceAssignmentStatusesByToken(ctx context.Context, args struct {
	Tokens []string
}) ([]*DeviceAssignmentStatusResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceAssignmentStatusesByToken(ctx, args.Tokens)
	if err != nil {
		return nil, err
	}

	result := make([]*DeviceAssignmentStatusResolver, 0)
	for _, dt := range found {
		dtr := &DeviceAssignmentStatusResolver{
			M: *dt,
			S: r,
			C: ctx,
		}
		result = append(result, dtr)
	}
	return result, nil
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

// Find device assignments by unique id.
func (r *SchemaResolver) DeviceAssignmentsById(ctx context.Context, args struct {
	Ids []string
}) ([]*DeviceAssignmentResolver, error) {
	api := r.GetApi(ctx)
	ids, err := r.asUintIds(args.Ids)
	if err != nil {
		return nil, err
	}

	found, err := api.DeviceAssignmentsById(ctx, ids)
	if err != nil {
		return nil, err
	}

	result := make([]*DeviceAssignmentResolver, 0)
	for _, dt := range found {
		dtr := &DeviceAssignmentResolver{
			M: *dt,
			S: r,
			C: ctx,
		}
		result = append(result, dtr)
	}
	return result, nil
}

// Find device assignments by unique token.
func (r *SchemaResolver) DeviceAssignmentsByToken(ctx context.Context, args struct {
	Tokens []string
}) ([]*DeviceAssignmentResolver, error) {
	api := r.GetApi(ctx)
	found, err := api.DeviceAssignmentsByToken(ctx, args.Tokens)
	if err != nil {
		return nil, err
	}

	result := make([]*DeviceAssignmentResolver, 0)
	for _, dt := range found {
		dtr := &DeviceAssignmentResolver{
			M: *dt,
			S: r,
			C: ctx,
		}
		result = append(result, dtr)
	}
	return result, nil
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

// List all active device assignments for device with given id.
func (r *SchemaResolver) ActiveDeviceAssignmentsForDevice(ctx context.Context, args struct {
	Id string
}) ([]*DeviceAssignmentResolver, error) {
	api := r.GetApi(ctx)
	id, err := strconv.ParseUint(args.Id, 0, 64)
	if err != nil {
		return nil, err
	}

	found, err := api.ActiveDeviceAssignmentsForDevice(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	// Return as resolver.
	resolvers := make([]*DeviceAssignmentResolver, 0)
	for _, current := range found {
		resolvers = append(resolvers,
			&DeviceAssignmentResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}
	return resolvers, nil
}
