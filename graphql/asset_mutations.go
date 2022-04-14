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

// Create a new asset type.
func (r *SchemaResolver) CreateAssetType(ctx context.Context, args struct {
	Value *model.AssetTypeCreateRequest
}) (*AssetTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.AssetType{
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

	dt := &AssetTypeResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing asset type.
func (r *SchemaResolver) UpdateAssetType(ctx context.Context, args struct {
	Token string
	Value *model.AssetTypeCreateRequest
}) (*AssetTypeResolver, error) {
	found, err := r.AssetTypeByToken(ctx, struct{ Token string }{Token: args.Token})
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

	dt := &AssetTypeResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new asset.
func (r *SchemaResolver) CreateAsset(ctx context.Context, args struct {
	Value *model.AssetCreateRequest
}) (*AssetResolver, error) {
	found, err := r.AssetTypeByToken(ctx, struct{ Token string }{Token: args.Value.AssetTypeToken})
	if err != nil {
		return nil, err
	}

	rdbmgr := r.GetRdbManager(ctx)
	created := model.Asset{
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
		AssetType: &found.M,
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing asset.
func (r *SchemaResolver) UpdateAsset(ctx context.Context, args struct {
	Token string
	Value *model.AssetCreateRequest
}) (*AssetResolver, error) {
	found, err := r.AssetByToken(ctx, struct{ Token string }{Token: args.Token})
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

	// Update asset type if changed.
	if args.Value.AssetTypeToken != upd.AssetType.Token {
		found, err := r.AssetTypeByToken(ctx, struct{ Token string }{Token: args.Value.AssetTypeToken})
		if err != nil {
			return nil, err
		}
		upd.AssetType = &found.M
	}

	result := rdbmgr.Database.Save(&upd)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new asset relationship type.
func (r *SchemaResolver) CreateAssetRelationshipType(ctx context.Context, args struct {
	Value *model.AssetRelationshipTypeCreateRequest
}) (*AssetRelationshipTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.AssetRelationshipType{
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

	dt := &AssetRelationshipTypeResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing asset relationship type.
func (r *SchemaResolver) UpdateAssetRelationshipType(ctx context.Context, args struct {
	Token string
	Value *model.AssetRelationshipTypeCreateRequest
}) (*AssetRelationshipTypeResolver, error) {
	found, err := r.AssetRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Token})
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

	dt := &AssetRelationshipTypeResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new asset relationship.
func (r *SchemaResolver) CreateAssetRelationship(ctx context.Context, args struct {
	Value *model.AssetRelationshipCreateRequest
}) (*AssetRelationshipResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)

	// Look up token references.
	source, err := r.AssetByToken(ctx, struct{ Token string }{Token: args.Value.SourceAsset})
	if err != nil {
		return nil, err
	}
	target, err := r.AssetByToken(ctx, struct{ Token string }{Token: args.Value.TargetAsset})
	if err != nil {
		return nil, err
	}
	rtype, err := r.AssetRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Value.RelationshipType})
	if err != nil {
		return nil, err
	}

	created := model.AssetRelationship{
		SourceAsset:      source.M,
		TargetAsset:      target.M,
		RelationshipType: rtype.M,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetRelationshipResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new asset group.
func (r *SchemaResolver) CreateAssetGroup(ctx context.Context, args struct {
	Value *model.AssetGroupCreateRequest
}) (*AssetGroupResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.AssetGroup{
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

	dt := &AssetGroupResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing asset group.
func (r *SchemaResolver) UpdateAssetGroup(ctx context.Context, args struct {
	Token string
	Value *model.AssetGroupCreateRequest
}) (*AssetGroupResolver, error) {
	found, err := r.AssetGroupByToken(ctx, struct{ Token string }{Token: args.Token})
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

	dt := &AssetGroupResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new asset group relationship type.
func (r *SchemaResolver) CreateAssetGroupRelationshipType(ctx context.Context, args struct {
	Value *model.AssetGroupRelationshipTypeCreateRequest
}) (*AssetGroupRelationshipTypeResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)
	created := model.AssetGroupRelationshipType{
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

	dt := &AssetGroupRelationshipTypeResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Update an existing asset group relationship type.
func (r *SchemaResolver) UpdateAssetGroupRelationshipType(ctx context.Context, args struct {
	Token string
	Value *model.AssetGroupRelationshipTypeCreateRequest
}) (*AssetGroupRelationshipTypeResolver, error) {
	found, err := r.AssetGroupRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Token})
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

	dt := &AssetGroupRelationshipTypeResolver{
		M: upd,
		S: r,
		C: ctx,
	}
	return dt, nil
}

// Create a new asset group relationship.
func (r *SchemaResolver) CreateAssetGroupRelationship(ctx context.Context, args struct {
	Value *model.AssetGroupRelationshipCreateRequest
}) (*AssetGroupRelationshipResolver, error) {
	rdbmgr := r.GetRdbManager(ctx)

	// Look up token references.
	source, err := r.AssetGroupByToken(ctx, struct{ Token string }{Token: args.Value.AssetGroup})
	if err != nil {
		return nil, err
	}
	target, err := r.AssetByToken(ctx, struct{ Token string }{Token: args.Value.Asset})
	if err != nil {
		return nil, err
	}
	rtype, err := r.AssetGroupRelationshipTypeByToken(ctx, struct{ Token string }{Token: args.Value.RelationshipType})
	if err != nil {
		return nil, err
	}

	created := model.AssetGroupRelationship{
		AssetGroup:       source.M,
		Asset:            target.M,
		RelationshipType: rtype.M,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(args.Value.Metadata),
		},
	}
	result := rdbmgr.Database.Create(&created)
	if result.Error != nil {
		return nil, result.Error
	}

	dt := &AssetGroupRelationshipResolver{
		M: created,
		S: r,
		C: ctx,
	}
	return dt, nil
}
