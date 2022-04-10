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
		M: found,
		S: r,
		C: ctx,
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
		M: found,
		S: r,
		C: ctx,
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
		resolvers = append(resolvers,
			&DeviceTypeResolver{
				M: current,
				S: r,
				C: ctx,
			})
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
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device by unique token.
func (r *SchemaResolver) DeviceByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceResolver, error) {
	found := model.Device{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Joins("DeviceType").First(&found, "\"devices\".token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceResolver{
		M: found,
		S: r,
		C: ctx,
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
		resolvers = append(resolvers,
			&DeviceResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
}

// Find device relationship type by unique id.
func (r *SchemaResolver) DeviceRelationshipType(ctx context.Context, args struct {
	Id string
}) (*DeviceRelationshipTypeResolver, error) {
	found := model.DeviceRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device relationship type by unique token.
func (r *SchemaResolver) DeviceRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceRelationshipTypeResolver, error) {
	found := model.DeviceRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device relationship types that match the given criteria.
func (r *SchemaResolver) DeviceRelationshipTypes(ctx context.Context, args struct {
	Criteria model.DeviceRelationshipTypeSearchCriteria
}) ([]*DeviceRelationshipTypeResolver, error) {
	list := make([]model.DeviceRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*DeviceRelationshipTypeResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&DeviceRelationshipTypeResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
}

// Find device relationship by unique id.
func (r *SchemaResolver) DeviceRelationship(ctx context.Context, args struct {
	Id string
}) (*DeviceRelationshipResolver, error) {
	found := model.DeviceRelationship{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device relationships that match the given criteria.
func (r *SchemaResolver) DeviceRelationships(ctx context.Context, args struct {
	Criteria model.DeviceRelationshipSearchCriteria
}) ([]*DeviceRelationshipResolver, error) {
	list := make([]model.DeviceRelationship, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize))
	result = result.Joins("SourceDevice").Joins("TargetDevice").Joins("RelationshipType")
	result = result.Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*DeviceRelationshipResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&DeviceRelationshipResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
}

// Find device group by unique id.
func (r *SchemaResolver) DeviceGroup(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupResolver, error) {
	found := model.DeviceGroup{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device group by unique token.
func (r *SchemaResolver) DeviceGroupByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceGroupResolver, error) {
	found := model.DeviceGroup{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device groups that match the given criteria.
func (r *SchemaResolver) DeviceGroups(ctx context.Context, args struct {
	Criteria model.DeviceGroupSearchCriteria
}) ([]*DeviceGroupResolver, error) {
	list := make([]model.DeviceGroup, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*DeviceGroupResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&DeviceGroupResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
}

// Find device group relationship type by unique id.
func (r *SchemaResolver) DeviceGroupRelationshipType(ctx context.Context, args struct {
	Id string
}) (*DeviceGroupRelationshipTypeResolver, error) {
	found := model.DeviceGroupRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, args.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Find device group relationship type by unique token.
func (r *SchemaResolver) DeviceGroupRelationshipTypeByToken(ctx context.Context, args struct {
	Token string
}) (*DeviceGroupRelationshipTypeResolver, error) {
	found := model.DeviceGroupRelationshipType{}
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.First(&found, "token = ?", args.Token)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: found,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// List all device group relationship types that match the given criteria.
func (r *SchemaResolver) DeviceGroupRelationshipTypes(ctx context.Context, args struct {
	Criteria model.DeviceGroupRelationshipTypeSearchCriteria
}) ([]*DeviceGroupRelationshipTypeResolver, error) {
	list := make([]model.DeviceGroupRelationshipType, 0)
	rdbmgr := r.GetRdbManager(ctx)
	result := rdbmgr.Database.Limit(int(args.Criteria.PageSize)).Offset(int(args.Criteria.PageNumber) * int(args.Criteria.PageSize)).Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	resolvers := make([]*DeviceGroupRelationshipTypeResolver, 0)
	for _, current := range list {
		resolvers = append(resolvers,
			&DeviceGroupRelationshipTypeResolver{
				M: current,
				S: r,
				C: ctx,
			})
	}

	return resolvers, nil
}
