/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"

	"github.com/devicechain-io/dc-devicemanagement/model"
	gql "github.com/graph-gophers/graphql-go"
)

// Find device type by unique id.
func (r *SchemaResolver) DeviceType(ctx context.Context, args struct {
	Id gql.ID
}) (*DeviceTypeResolver, error) {
	found := model.DeviceType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		d: found,
	}
	return dt, nil
}

// Find device type by unique token.
func (r *SchemaResolver) DeviceTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceTypeResolver, error) {
	found := model.DeviceType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		d: found,
	}
	return dt, nil
}

// List all device types that match the given criteria.
func (r *SchemaResolver) DeviceTypes(ctx context.Context, args struct {
	Criteria model.DeviceTypeSearchCriteria
}) ([]*DeviceTypeResolver, error) {
	list := make([]model.DeviceType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*DeviceTypeResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers, &DeviceTypeResolver{d: current})
	}

	return resolvers, nil
}
