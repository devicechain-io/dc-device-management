/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	"fmt"

	"github.com/devicechain-io/dc-device-management/model"
	"github.com/devicechain-io/dc-microservice/rdb"
	"github.com/rs/zerolog/log"
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
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		m: created,
	}
	return dt, nil
}

// Update an existing device type.
func (r *SchemaResolver) UpdateDeviceType(ctx context.Context, args struct {
	Id    string
	Value *model.DeviceTypeCreateRequest
}) (*DeviceTypeResolver, error) {
	dtr, err := r.DeviceType(ctx, struct{ Id string }{Id: args.Id})
	if err != nil {
		return nil, err
	}
	log.Info().Msg(fmt.Sprintf("Found device type %+v", dtr.m))

	rdbmgr := r.GetRdbManager(ctx)
	upd := dtr.m
	upd.Token = args.Value.Token
	upd.Name = rdb.NullStrOf(args.Value.Name)
	upd.Description = rdb.NullStrOf(args.Value.Description)
	upd.ImageUrl = rdb.NullStrOf(args.Value.ImageUrl)
	upd.Icon = rdb.NullStrOf(args.Value.Icon)
	upd.BackgroundColor = rdb.NullStrOf(args.Value.BackgroundColor)
	upd.ForegroundColor = rdb.NullStrOf(args.Value.ForegroundColor)
	upd.BorderColor = rdb.NullStrOf(args.Value.BorderColor)
	log.Info().Msg(fmt.Sprintf("Updated all fields %+v", upd))

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		m: upd,
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
		DeviceType: dtr.m,
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceResolver{
		m: created,
	}
	return dt, nil
}

// Update an existing device.
func (r *SchemaResolver) UpdateDevice(ctx context.Context, args struct {
	Id    string
	Value *model.DeviceCreateRequest
}) (*DeviceResolver, error) {
	dr, err := r.Device(ctx, struct{ Id string }{Id: args.Id})
	if err != nil {
		return nil, err
	}

	// Update fields that changed.
	rdbmgr := r.GetRdbManager(ctx)
	upd := dr.m
	upd.Token = args.Value.Token
	upd.Name = rdb.NullStrOf(args.Value.Name)
	upd.Description = rdb.NullStrOf(args.Value.Description)

	// Update device type if changed.
	if args.Value.DeviceTypeToken != upd.DeviceType.Token {
		dtr, err := r.DeviceTypeByToken(ctx, struct{ Token string }{Token: args.Value.DeviceTypeToken})
		if err != nil {
			return nil, err
		}
		upd.DeviceType = dtr.m
	}

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceResolver{
		m: upd,
	}
	return dt, nil
}
