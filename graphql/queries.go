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

// Find device type by unique id.
func (r *SchemaResolver) DeviceType(ctx context.Context, args struct {
	Id string
}) (*DeviceTypeResolver, error) {
	found := model.DeviceType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		m: found,
	}
	return dt, nil
}

// Find device type by unique token.
func (r *SchemaResolver) DeviceTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceTypeResolver, error) {
	found := model.DeviceType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		m: found,
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
		resolvers = append(resolvers, &DeviceTypeResolver{m: current})
	}

	return resolvers, nil
}

// Find device by unique id.
func (r *SchemaResolver) Device(ctx context.Context, args struct {
	Id string
}) (*DeviceResolver, error) {
	found := model.Device{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Joins("DeviceType").First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceResolver{
		m: found,
	}
	return dt, nil
}

// Find device by unique token.
func (r *SchemaResolver) DeviceByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceResolver, error) {
	found := model.Device{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Joins("DeviceType").First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceResolver{
		m: found,
	}
	return dt, nil
}

// List all devices that match the given criteria.
func (r *SchemaResolver) Devices(ctx context.Context, args struct {
	Criteria model.DeviceSearchCriteria
}) ([]*DeviceResolver, error) {
	list := make([]model.Device, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Debug().Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize))
	if args.Criteria.DeviceTypeToken != nil {
		result = result.Joins("DeviceType").Where("device_type_id = (?)",
			rdbmgr.Database.Model(&model.DeviceType{}).Select("id").Where("token = ?", args.Criteria.DeviceTypeToken)).Find(&list)
	} else {
		result = result.Joins("DeviceType").Find(&list)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*DeviceResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers, &DeviceResolver{m: current})
	}

	return resolvers, nil
}
