/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"

	"github.com/devicechain-io/dc-devicemanagement/model"
	"github.com/devicechain-io/dc-microservice/rdb"
)

func (r *SchemaResolver) CreateDeviceType(ctx context.Context, args struct {
	Value *model.DeviceTypeCreateRequest
}) (*DeviceTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.DeviceType{
		TokenReference: rdb.TokenReference{
			Token: args.Value.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(*args.Value.Name),
			Description: rdb.NullStrOf(*args.Value.Description),
		},
		BrandedEntity: rdb.BrandedEntity{
			ImageUrl:        rdb.NullStrOf(*args.Value.ImageUrl),
			Icon:            rdb.NullStrOf(*args.Value.Icon),
			BackgroundColor: rdb.NullStrOf(*args.Value.BackgroundColor),
			ForegroundColor: rdb.NullStrOf(*args.Value.ForegroundColor),
			BorderColor:     rdb.NullStrOf(*args.Value.BorderColor),
		},
	}
	result := rdbmgr.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &DeviceTypeResolver{
		d: created,
	}
	return dt, nil
}
