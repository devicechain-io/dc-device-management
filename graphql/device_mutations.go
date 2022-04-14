/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"

	"github.com/devicechain-io/dc-device-management/model"
	"github.com/devicechain-io/dc-microservice/rdb"
)

// Create a new device type.
func (r *SchemaResolver) CreateDeviceType(ctx context.Context, args struct {
	Value *model.DeviceTypeCreateRequest
}) (*DeviceTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.DeviceType{
		TokenReference: rdb.TokenReference{
			Token: args.Value.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(args.Value.Name),
			Description: rdb.NullStrOf(args.Value.Description),
		},
		BrandedEntity: rdb.BrandedEntity{
			ImageUrl:        rdb.NullStrOf(args.Value.ImageUrl),
			Icon:            rdb.NullStrOf(args.Value.Icon),
			BackgroundColor: rdb.NullStrOf(args.Value.BackgroundColor),
			ForegroundColor: rdb.NullStrOf(args.Value.ForegroundColor),
			BorderColor:     rdb.NullStrOf(args.Value.BorderColor),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing device type.
func (r *SchemaResolver) UpdateDeviceType(ctx context.Context, args struct {
	Token string
	Value *model.DeviceTypeCreateRequest
}) (*DeviceTypeResolver, error) {
	found, err := r.DeviceTypeByToken(ctx, struct{ Token string }{Token: args.Token})
	if err != nil {
		return nil, err
	}

	rdbmgr := r.GetRdbManager(ctx)
	upd := found.M
	upd.Token = args.Value.Token
	upd.Name = rdb.NullStrOf(args.Value.Name)
	upd.Description = rdb.NullStrOf(args.Value.Description)
	upd.ImageUrl = rdb.NullStrOf(args.Value.ImageUrl)
	upd.Icon = rdb.NullStrOf(args.Value.Icon)
	upd.BackgroundColor = rdb.NullStrOf(args.Value.BackgroundColor)
	upd.ForegroundColor = rdb.NullStrOf(args.Value.ForegroundColor)
	upd.BorderColor = rdb.NullStrOf(args.Value.BorderColor)
	upd.Metadata = rdb.MetadataStrOf(args.Value.Metadata)

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new device.
func (r *SchemaResolver) CreateDevice(ctx context.Context, args struct {
	Value *model.DeviceCreateRequest
}) (*DeviceResolver, error) {
	dtr, err := r.DeviceTypeByToken(ctx, struct{ Token string }{Token: args.Value.DeviceTypeToken})
	if err != nil {
		return nil, err
	}

	rdbmgr := r.GetRdbManager(ctx)
	created := model.Device{
		TokenReference: rdb.TokenReference{
			Token: args.Value.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(args.Value.Name),
			Description: rdb.NullStrOf(args.Value.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
		DeviceType: &dtr.M,
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing device.
func (r *SchemaResolver) UpdateDevice(ctx context.Context, args struct {
	Token string
	Value *model.DeviceCreateRequest
}) (*DeviceResolver, error) {
	found, err := r.DeviceByToken(ctx, struct{ Token string }{Token: args.Token})
	if err != nil {
		return nil, err
	}

	// Update fields that changed.
	rdbmgr := r.GetRdbManager(ctx)
	upd := found.M
	upd.Token = args.Value.Token
	upd.Name = rdb.NullStrOf(args.Value.Name)
	upd.Description = rdb.NullStrOf(args.Value.Description)
	upd.Metadata = rdb.MetadataStrOf(args.Value.Metadata)

	// Update device type if changed.
	if args.Value.DeviceTypeToken != upd.DeviceType.Token {
		dtr, err := r.DeviceTypeByToken(ctx, struct{ Token string }{Token: args.Value.DeviceTypeToken})
		if err != nil {
			return nil, err
		}
		upd.DeviceType = &dtr.M
	}

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new device relationship type.
func (r *SchemaResolver) CreateDeviceRelationshipType(ctx context.Context, args struct {
	Value *model.DeviceRelationshipTypeCreateRequest
}) (*DeviceRelationshipTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.DeviceRelationshipType{
		TokenReference: rdb.TokenReference{
			Token: args.Value.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(args.Value.Name),
			Description: rdb.NullStrOf(args.Value.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipTypeResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing device relationship type.
func (r *SchemaResolver) UpdateDeviceRelationshipType(ctx context.Context, args struct {
	Token string
	Value *model.DeviceRelationshipTypeCreateRequest
}) (*DeviceRelationshipTypeResolver, error) {
	found, err := r.DeviceRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Token})
	if err != nil {
		return nil, err
	}

	rdbmgr := r.GetRdbManager(ctx)
	upd := found.M
	upd.Token = args.Value.Token
	upd.Name = rdb.NullStrOf(args.Value.Name)
	upd.Description = rdb.NullStrOf(args.Value.Description)
	upd.Metadata = rdb.MetadataStrOf(args.Value.Metadata)

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipTypeResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new device relationship.
func (r *SchemaResolver) CreateDeviceRelationship(ctx context.Context, args struct {
	Value *model.DeviceRelationshipCreateRequest
}) (*DeviceRelationshipResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)

	// Look up token references.
	source, err := r.DeviceByToken(ctx, struct{ Token string }{Token: args.Value.SourceDevice})
	if err != nil {
		return nil, err
	}
	target, err := r.DeviceByToken(ctx, struct{ Token string }{Token: args.Value.TargetDevice})
	if err != nil {
		return nil, err
	}
	rtype, err := r.DeviceRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Value.RelationshipType})
	if err != nil {
		return nil, err
	}

	created := model.DeviceRelationship{
		SourceDevice:     source.M,
		TargetDevice:     target.M,
		RelationshipType: rtype.M,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceRelationshipResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new device group.
func (r *SchemaResolver) CreateDeviceGroup(ctx context.Context, args struct {
	Value *model.DeviceGroupCreateRequest
}) (*DeviceGroupResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.DeviceGroup{
		TokenReference: rdb.TokenReference{
			Token: args.Value.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(args.Value.Name),
			Description: rdb.NullStrOf(args.Value.Description),
		},
		BrandedEntity: rdb.BrandedEntity{
			ImageUrl:        rdb.NullStrOf(args.Value.ImageUrl),
			Icon:            rdb.NullStrOf(args.Value.Icon),
			BackgroundColor: rdb.NullStrOf(args.Value.BackgroundColor),
			ForegroundColor: rdb.NullStrOf(args.Value.ForegroundColor),
			BorderColor:     rdb.NullStrOf(args.Value.BorderColor),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing device type.
func (r *SchemaResolver) UpdateDeviceGroup(ctx context.Context, args struct {
	Token string
	Value *model.DeviceGroupCreateRequest
}) (*DeviceGroupResolver, error) {
	found, err := r.DeviceGroupByToken(ctx, struct{ Token string }{Token: args.Token})
	if err != nil {
		return nil, err
	}

	rdbmgr := r.GetRdbManager(ctx)
	upd := found.M
	upd.Token = args.Value.Token
	upd.Name = rdb.NullStrOf(args.Value.Name)
	upd.Description = rdb.NullStrOf(args.Value.Description)
	upd.ImageUrl = rdb.NullStrOf(args.Value.ImageUrl)
	upd.Icon = rdb.NullStrOf(args.Value.Icon)
	upd.BackgroundColor = rdb.NullStrOf(args.Value.BackgroundColor)
	upd.ForegroundColor = rdb.NullStrOf(args.Value.ForegroundColor)
	upd.BorderColor = rdb.NullStrOf(args.Value.BorderColor)
	upd.Metadata = rdb.MetadataStrOf(args.Value.Metadata)

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new device group relationship type.
func (r *SchemaResolver) CreateDeviceGroupRelationshipType(ctx context.Context, args struct {
	Value *model.DeviceGroupRelationshipTypeCreateRequest
}) (*DeviceGroupRelationshipTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.DeviceGroupRelationshipType{
		TokenReference: rdb.TokenReference{
			Token: args.Value.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(args.Value.Name),
			Description: rdb.NullStrOf(args.Value.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing device group relationship type.
func (r *SchemaResolver) UpdateDeviceGroupRelationshipType(ctx context.Context, args struct {
	Token string
	Value *model.DeviceGroupRelationshipTypeCreateRequest
}) (*DeviceGroupRelationshipTypeResolver, error) {
	found, err := r.DeviceGroupRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Token})
	if err != nil {
		return nil, err
	}

	rdbmgr := r.GetRdbManager(ctx)
	upd := found.M
	upd.Token = args.Value.Token
	upd.Name = rdb.NullStrOf(args.Value.Name)
	upd.Description = rdb.NullStrOf(args.Value.Description)
	upd.Metadata = rdb.MetadataStrOf(args.Value.Metadata)

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipTypeResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new device group relationship.
func (r *SchemaResolver) CreateDeviceGroupRelationship(ctx context.Context, args struct {
	Value *model.DeviceGroupRelationshipCreateRequest
}) (*DeviceGroupRelationshipResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)

	// Look up token references.
	source, err := r.DeviceGroupByToken(ctx, struct{ Token string }{Token: args.Value.DeviceGroup})
	if err != nil {
		return nil, err
	}
	target, err := r.DeviceByToken(ctx, struct{ Token string }{Token: args.Value.Device})
	if err != nil {
		return nil, err
	}
	rtype, err := r.DeviceGroupRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Value.RelationshipType})
	if err != nil {
		return nil, err
	}

	created := model.DeviceGroupRelationship{
		DeviceGroup:      source.M,
		Device:           target.M,
		RelationshipType: rtype.M,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceGroupRelationshipResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}
