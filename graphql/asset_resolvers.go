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

// -------------------
// Asset type resolver
// -------------------

type AssetTypeResolver struct {
	M model.AssetType
	S *SchemaResolver
	C context.Context
}

func (r *AssetTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *AssetTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *AssetTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *AssetTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *AssetTypeResolver) Token() string {
	return r.M.Token
}

func (r *AssetTypeResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *AssetTypeResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *AssetTypeResolver) ImageUrl() *string {
	return util.NullStr(r.M.ImageUrl)
}

func (r *AssetTypeResolver) Icon() *string {
	return util.NullStr(r.M.Icon)
}

func (r *AssetTypeResolver) BackgroundColor() *string {
	return util.NullStr(r.M.BackgroundColor)
}

func (r *AssetTypeResolver) ForegroundColor() *string {
	return util.NullStr(r.M.ForegroundColor)
}

func (r *AssetTypeResolver) BorderColor() *string {
	return util.NullStr(r.M.BorderColor)
}

func (r *AssetTypeResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

// --------------
// Asset resolver
// --------------

type AssetResolver struct {
	M model.Asset
	S *SchemaResolver
	C context.Context
}

func (r *AssetResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *AssetResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *AssetResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *AssetResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *AssetResolver) Token() string {
	return r.M.Token
}

func (r *AssetResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *AssetResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *AssetResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

func (r *AssetResolver) AssetType() *AssetTypeResolver {
	if r.M.AssetType != nil {
		return &AssetTypeResolver{
			M: *r.M.AssetType,
			S: r.S,
			C: r.C,
		}
	} else {
		rez, err := r.S.AssetType(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", r.M.AssetTypeId)})
		if err != nil {
			return nil
		}
		return rez
	}
}

// --------------------------------
// Asset relationship type resolver
// --------------------------------

type AssetRelationshipTypeResolver struct {
	M model.AssetRelationshipType
	S *SchemaResolver
	C context.Context
}

func (r *AssetRelationshipTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *AssetRelationshipTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *AssetRelationshipTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *AssetRelationshipTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *AssetRelationshipTypeResolver) Token() string {
	return r.M.Token
}

func (r *AssetRelationshipTypeResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *AssetRelationshipTypeResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *AssetRelationshipTypeResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

// ---------------------------
// Asset relationship resolver
// ---------------------------

type AssetRelationshipResolver struct {
	M model.AssetRelationship
	S *SchemaResolver
	C context.Context
}

func (r *AssetRelationshipResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *AssetRelationshipResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *AssetRelationshipResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *AssetRelationshipResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *AssetRelationshipResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

func (r *AssetRelationshipResolver) SourceAsset() *AssetResolver {
	return &AssetResolver{
		M: r.M.SourceAsset,
		S: r.S,
		C: r.C,
	}
}

func (r *AssetRelationshipResolver) TargetAsset() *AssetResolver {
	return &AssetResolver{
		M: r.M.TargetAsset,
		S: r.S,
		C: r.C,
	}
}

func (r *AssetRelationshipResolver) RelationshipType() *AssetRelationshipTypeResolver {
	return &AssetRelationshipTypeResolver{
		M: r.M.RelationshipType,
		S: r.S,
		C: r.C,
	}
}

// --------------------
// Asset group resolver
// --------------------

type AssetGroupResolver struct {
	M model.AssetGroup
	S *SchemaResolver
	C context.Context
}

func (r *AssetGroupResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *AssetGroupResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *AssetGroupResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *AssetGroupResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *AssetGroupResolver) Token() string {
	return r.M.Token
}

func (r *AssetGroupResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *AssetGroupResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *AssetGroupResolver) ImageUrl() *string {
	return util.NullStr(r.M.ImageUrl)
}

func (r *AssetGroupResolver) Icon() *string {
	return util.NullStr(r.M.Icon)
}

func (r *AssetGroupResolver) BackgroundColor() *string {
	return util.NullStr(r.M.BackgroundColor)
}

func (r *AssetGroupResolver) ForegroundColor() *string {
	return util.NullStr(r.M.ForegroundColor)
}

func (r *AssetGroupResolver) BorderColor() *string {
	return util.NullStr(r.M.BorderColor)
}

func (r *AssetGroupResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

// --------------------------------------
// Asset group relationship type resolver
// --------------------------------------

type AssetGroupRelationshipTypeResolver struct {
	M model.AssetGroupRelationshipType
	S *SchemaResolver
	C context.Context
}

func (r *AssetGroupRelationshipTypeResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *AssetGroupRelationshipTypeResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *AssetGroupRelationshipTypeResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *AssetGroupRelationshipTypeResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *AssetGroupRelationshipTypeResolver) Token() string {
	return r.M.Token
}

func (r *AssetGroupRelationshipTypeResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *AssetGroupRelationshipTypeResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *AssetGroupRelationshipTypeResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

// ---------------------------------
// Asset group relationship resolver
// ---------------------------------

type AssetGroupRelationshipResolver struct {
	M model.AssetGroupRelationship
	S *SchemaResolver
	C context.Context
}

func (r *AssetGroupRelationshipResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *AssetGroupRelationshipResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *AssetGroupRelationshipResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *AssetGroupRelationshipResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *AssetGroupRelationshipResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

func (r *AssetGroupRelationshipResolver) AssetGroup() *AssetGroupResolver {
	return &AssetGroupResolver{
		M: r.M.AssetGroup,
		S: r.S,
		C: r.C,
	}
}

func (r *AssetGroupRelationshipResolver) Asset() *AssetResolver {
	return &AssetResolver{
		M: r.M.Asset,
		S: r.S,
		C: r.C,
	}
}

func (r *AssetGroupRelationshipResolver) RelationshipType() *AssetGroupRelationshipTypeResolver {
	return &AssetGroupRelationshipTypeResolver{
		M: r.M.RelationshipType,
		S: r.S,
		C: r.C,
	}
}
