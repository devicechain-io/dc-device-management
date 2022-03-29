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

// -------------------
// Base model resolver
// -------------------

type DeviceTypeResolver struct {
	m model.DeviceType
}

func (r *DeviceTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.m.ID))
}

func (r *DeviceTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.m.CreatedAt)
}

func (r *DeviceTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.m.UpdatedAt)
}

func (r *DeviceTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.m.DeletedAt.Time)
}

func (r *DeviceTypeResolver) Token() string {
	return r.m.Token
}

func (r *DeviceTypeResolver) Name() *string {
	return util.NullStr(r.m.Name)
}

func (r *DeviceTypeResolver) Description() *string {
	return util.NullStr(r.m.Description)
}

func (r *DeviceTypeResolver) ImageUrl() *string {
	return util.NullStr(r.m.ImageUrl)
}

func (r *DeviceTypeResolver) Icon() *string {
	return util.NullStr(r.m.Icon)
}

func (r *DeviceTypeResolver) BackgroundColor() *string {
	return util.NullStr(r.m.BackgroundColor)
}

func (r *DeviceTypeResolver) ForegroundColor() *string {
	return util.NullStr(r.m.ForegroundColor)
}

func (r *DeviceTypeResolver) BorderColor() *string {
	return util.NullStr(r.m.BorderColor)
}

// --------------------
// Device resolver
// --------------------

type DeviceResolver struct {
	m model.Device
}

func (r *DeviceResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.m.ID))
}

func (r *DeviceResolver) CreatedAt() *string {
	return util.FormatTime(r.m.CreatedAt)
}

func (r *DeviceResolver) UpdatedAt() *string {
	return util.FormatTime(r.m.UpdatedAt)
}

func (r *DeviceResolver) DeletedAt() *string {
	return util.FormatTime(r.m.DeletedAt.Time)
}

func (r *DeviceResolver) Token() string {
	return r.m.Token
}

func (r *DeviceResolver) Name() *string {
	return util.NullStr(r.m.Name)
}

func (r *DeviceResolver) Description() *string {
	return util.NullStr(r.m.Description)
}

func (r *DeviceResolver) DeviceType() *DeviceTypeResolver {
	return &DeviceTypeResolver{
		m: r.m.DeviceType,
	}
}
