/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/devicechain-io/dc-device-management/model"
	util "github.com/devicechain-io/dc-microservice/graphql"
	gql "github.com/graph-gophers/graphql-go"
)

// --------------------
// Device type resolver
// --------------------

type DeviceTypeResolver struct {
	M model.DeviceType
	S *SchemaResolver
	C context.Context
}

func (r *DeviceTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceTypeResolver) Token() string {
	return r.M.Token
}

func (r *DeviceTypeResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *DeviceTypeResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *DeviceTypeResolver) ImageUrl() *string {
	return util.NullStr(r.M.ImageUrl)
}

func (r *DeviceTypeResolver) Icon() *string {
	return util.NullStr(r.M.Icon)
}

func (r *DeviceTypeResolver) BackgroundColor() *string {
	return util.NullStr(r.M.BackgroundColor)
}

func (r *DeviceTypeResolver) ForegroundColor() *string {
	return util.NullStr(r.M.ForegroundColor)
}

func (r *DeviceTypeResolver) BorderColor() *string {
	return util.NullStr(r.M.BorderColor)
}

// ---------------
// Device resolver
// ---------------

type DeviceResolver struct {
	M model.Device
	S *SchemaResolver
	C context.Context
}

func (r *DeviceResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceResolver) Token() string {
	return r.M.Token
}

func (r *DeviceResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *DeviceResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *DeviceResolver) DeviceType() *DeviceTypeResolver {
	if r.M.DeviceType != nil {
		return &DeviceTypeResolver{
			M: *r.M.DeviceType,
			S: r.S,
			C: r.C,
		}
	} else {
		rez, err := r.S.DeviceType(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", r.M.DeviceTypeId)})
		if err != nil {
			return nil
		}
		return rez
	}
}

// ---------------------------------
// Device relationship type resolver
// ---------------------------------

type DeviceRelationshipTypeResolver struct {
	M model.DeviceRelationshipType
	S *SchemaResolver
	C context.Context
}

func (r *DeviceRelationshipTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceRelationshipTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceRelationshipTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceRelationshipTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceRelationshipTypeResolver) Token() string {
	return r.M.Token
}

func (r *DeviceRelationshipTypeResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *DeviceRelationshipTypeResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

// ----------------------------
// Device relationship resolver
// ----------------------------

type DeviceRelationshipResolver struct {
	M model.DeviceRelationship
	S *SchemaResolver
	C context.Context
}

func (r *DeviceRelationshipResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceRelationshipResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceRelationshipResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceRelationshipResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceRelationshipResolver) SourceDevice() *DeviceResolver {
	return &DeviceResolver{
		M: r.M.SourceDevice,
		S: r.S,
		C: r.C,
	}
}

func (r *DeviceRelationshipResolver) TargetDevice() *DeviceResolver {
	return &DeviceResolver{
		M: r.M.TargetDevice,
		S: r.S,
		C: r.C,
	}
}

func (r *DeviceRelationshipResolver) RelationshipType() *DeviceRelationshipTypeResolver {
	return &DeviceRelationshipTypeResolver{
		M: r.M.RelationshipType,
		S: r.S,
		C: r.C,
	}
}

// ---------------------
// Device group resolver
// ---------------------

type DeviceGroupResolver struct {
	M model.DeviceGroup
	S *SchemaResolver
	C context.Context
}

func (r *DeviceGroupResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceGroupResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceGroupResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceGroupResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceGroupResolver) Token() string {
	return r.M.Token
}

func (r *DeviceGroupResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *DeviceGroupResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *DeviceGroupResolver) ImageUrl() *string {
	return util.NullStr(r.M.ImageUrl)
}

func (r *DeviceGroupResolver) Icon() *string {
	return util.NullStr(r.M.Icon)
}

func (r *DeviceGroupResolver) BackgroundColor() *string {
	return util.NullStr(r.M.BackgroundColor)
}

func (r *DeviceGroupResolver) ForegroundColor() *string {
	return util.NullStr(r.M.ForegroundColor)
}

func (r *DeviceGroupResolver) BorderColor() *string {
	return util.NullStr(r.M.BorderColor)
}

// ---------------------------------------
// Device group relationship type resolver
// ---------------------------------------

type DeviceGroupRelationshipTypeResolver struct {
	M model.DeviceGroupRelationshipType
	S *SchemaResolver
	C context.Context
}

func (r *DeviceGroupRelationshipTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceGroupRelationshipTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceGroupRelationshipTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceGroupRelationshipTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceGroupRelationshipTypeResolver) Token() string {
	return r.M.Token
}

func (r *DeviceGroupRelationshipTypeResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *DeviceGroupRelationshipTypeResolver) Description() *string {
	return util.NullStr(r.M.Description)
}
