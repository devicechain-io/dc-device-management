/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"
	"time"

	"github.com/devicechain-io/dc-devicemanagement/model"
	"github.com/devicechain-io/dc-microservice/rdb"
	gql "github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

var devicetype = &model.DeviceType{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	TokenReference: rdb.TokenReference{
		Token: "rpi",
	},
	NamedEntity: rdb.NamedEntity{
		Name:        rdb.NullStrOf("Raspberry Pi"),
		Description: rdb.NullStrOf("This is a small computer that does stuff"),
	},
	BrandedEntity: rdb.BrandedEntity{
		ImageUrl:        rdb.NullStrOf("http://www.google.com"),
		Icon:            rdb.NullStrOf("http://www.google.com"),
		BackgroundColor: rdb.NullStrOf("0x000000"),
		ForegroundColor: rdb.NullStrOf("0x333333"),
		BorderColor:     rdb.NullStrOf("0x999999"),
	},
}

func (r *SchemaResolver) DeviceType(ctx context.Context, args struct {
	Id gql.ID
}) (*DeviceTypeResolver, error) {
	//rdbmgr := r.GetRdbManager(ctx)

	dt := &DeviceTypeResolver{
		d: devicetype,
	}
	return dt, nil
}
