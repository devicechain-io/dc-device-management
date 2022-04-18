/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"

	"github.com/devicechain-io/dc-device-management/model"
)

// Create a new device assignment status.
func (r *SchemaResolver) CreateDeviceAssignmentStatus(ctx context.Context, args struct {
	Request *model.DeviceAssignmentStatusCreateRequest
}) (*DeviceAssignmentStatusResolver, error) {
	api := r.GetApi(ctx)
	created, err := api.CreateDeviceAssignmentStatus(ctx, args.Request)
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentStatusResolver{
		M: *created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing device assignment status.
func (r *SchemaResolver) UpdateDeviceAssignmentStatus(ctx context.Context, args struct {
	Token   string
	Request *model.DeviceAssignmentStatusCreateRequest
}) (*DeviceAssignmentStatusResolver, error) {
	api := r.GetApi(ctx)
	updated, err := api.UpdateDeviceAssignmentStatus(ctx, args.Token, args.Request)
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentStatusResolver{
		M: *updated,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new device assignment.
func (r *SchemaResolver) CreateDeviceAssignment(ctx context.Context, args struct {
	Request *model.DeviceAssignmentCreateRequest
}) (*DeviceAssignmentResolver, error) {
	api := r.GetApi(ctx)
	created, err := api.CreateDeviceAssignment(ctx, args.Request)
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentResolver{
		M: *created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing device assignment.
func (r *SchemaResolver) UpdateDeviceAssignment(ctx context.Context, args struct {
	Token   string
	Request *model.DeviceAssignmentCreateRequest
}) (*DeviceAssignmentResolver, error) {
	api := r.GetApi(ctx)
	updated, err := api.UpdateDeviceAssignment(ctx, args.Token, args.Request)
	if err != nil {
		return nil, err
	}

	dt := &DeviceAssignmentResolver{
		M: *updated,
		S: r,
		C: ctx,
	}
	return dt, nil
}
