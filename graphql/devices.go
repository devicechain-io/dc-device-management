/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	_ "embed"
	"fmt"

	"github.com/devicechain-io/dc-devicemanagement/model"
	util "github.com/devicechain-io/dc-microservice/graphql"
	gql "github.com/graph-gophers/graphql-go"
)

type DeviceTypeResolver struct {
	d *model.DeviceType
}

func (r *DeviceTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.d.ID))
}

func (r *DeviceTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.d.CreatedAt)
}

func (r *DeviceTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.d.UpdatedAt)
}

func (r *DeviceTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.d.DeletedAt.Time)
}

func (r *DeviceTypeResolver) Token() string {
	return r.d.Token
}

func (r *DeviceTypeResolver) Name() *string {
	return util.NullStr(r.d.Name)
}

func (r *DeviceTypeResolver) Description() *string {
	return util.NullStr(r.d.Description)
}

func (r *DeviceTypeResolver) ImageUrl() *string {
	return util.NullStr(r.d.ImageUrl)
}

func (r *DeviceTypeResolver) Icon() *string {
	return util.NullStr(r.d.Icon)
}

func (r *DeviceTypeResolver) BackgroundColor() *string {
	return util.NullStr(r.d.BackgroundColor)
}

func (r *DeviceTypeResolver) ForegroundColor() *string {
	return util.NullStr(r.d.ForegroundColor)
}

func (r *DeviceTypeResolver) BorderColor() *string {
	return util.NullStr(r.d.BorderColor)
}
