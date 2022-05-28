/**
 * Copyright ©2022 AssetChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"context"

	"github.com/devicechain-io/dc-microservice/rdb"
	"gorm.io/gorm"
)

// Create a new area type.
func (api *Api) CreateAreaType(ctx context.Context, request *AreaTypeCreateRequest) (*AreaType, error) {
	created := &AreaType{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		BrandedEntity: rdb.BrandedEntity{
			ImageUrl:        rdb.NullStrOf(request.ImageUrl),
			Icon:            rdb.NullStrOf(request.Icon),
			BackgroundColor: rdb.NullStrOf(request.BackgroundColor),
			ForegroundColor: rdb.NullStrOf(request.ForegroundColor),
			BorderColor:     rdb.NullStrOf(request.BorderColor),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing area type.
func (api *Api) UpdateAreaType(ctx context.Context, token string,
	request *AreaTypeCreateRequest) (*AreaType, error) {
	found, err := api.AreaTypeByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	found.Token = request.Token
	found.Name = rdb.NullStrOf(request.Name)
	found.Description = rdb.NullStrOf(request.Description)
	found.ImageUrl = rdb.NullStrOf(request.ImageUrl)
	found.Icon = rdb.NullStrOf(request.Icon)
	found.BackgroundColor = rdb.NullStrOf(request.BackgroundColor)
	found.ForegroundColor = rdb.NullStrOf(request.ForegroundColor)
	found.BorderColor = rdb.NullStrOf(request.BorderColor)
	found.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(found)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area type by id.
func (api *Api) AreaTypeById(ctx context.Context, id uint) (*AreaType, error) {
	found := &AreaType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area type by token.
func (api *Api) AreaTypeByToken(ctx context.Context, token string) (*AreaType, error) {
	found := &AreaType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for area types that meet criteria.
func (api *Api) AreaTypes(ctx context.Context, criteria AreaTypeSearchCriteria) (*AreaTypeSearchResults, error) {
	results := make([]AreaType, 0)
	db, pag := api.RDB.ListOf(&AreaType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AreaTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new area.
func (api *Api) CreateArea(ctx context.Context, request *AreaCreateRequest) (*Area, error) {
	dtr, err := api.AreaTypeByToken(ctx, request.AreaTypeToken)
	if err != nil {
		return nil, err
	}

	created := &Area{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
		AreaType: dtr,
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing area.
func (api *Api) UpdateArea(ctx context.Context, token string, request *AreaCreateRequest) (*Area, error) {
	updated, err := api.AreaByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	// Update fields that changed.
	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	// Update area type if changed.
	if request.AreaTypeToken != updated.AreaType.Token {
		dtr, err := api.AreaTypeByToken(ctx, request.AreaTypeToken)
		if err != nil {
			return nil, err
		}
		updated.AreaType = dtr
	}

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
}

// Get area by id.
func (api *Api) AreaById(ctx context.Context, id uint) (*Area, error) {
	found := &Area{}
	result := api.RDB.Database
	result = result.Preload("AreaType")
	result = result.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area by token.
func (api *Api) AreaByToken(ctx context.Context, token string) (*Area, error) {
	found := &Area{}
	result := api.RDB.Database
	result = result.Preload("AreaType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for areas that meet criteria.
func (api *Api) Areas(ctx context.Context, criteria AreaSearchCriteria) (*AreaSearchResults, error) {
	results := make([]Area, 0)
	db, pag := api.RDB.ListOf(&Area{}, func(result *gorm.DB) *gorm.DB {
		if criteria.AreaTypeToken != nil {
			result = result.Where("area_type_id = (?)",
				api.RDB.Database.Model(&AreaType{}).Select("id").Where("token = ?", criteria.AreaTypeToken))
		}
		return result.Preload("AreaType")
	}, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AreaSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new area relationship type.
func (api *Api) CreateAreaRelationshipType(ctx context.Context, request *AreaRelationshipTypeCreateRequest) (*AreaRelationshipType, error) {
	created := &AreaRelationshipType{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing area relationship type.
func (api *Api) UpdateAreaRelationshipType(ctx context.Context, token string,
	request *AreaRelationshipTypeCreateRequest) (*AreaRelationshipType, error) {
	updated, err := api.AreaRelationshipTypeByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
}

// Get area relationship type by id.
func (api *Api) AreaRelationshipTypeById(ctx context.Context, id uint) (*AreaRelationshipType, error) {
	found := &AreaRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area relationship type by token.
func (api *Api) AreaRelationshipTypeByToken(ctx context.Context, token string) (*AreaRelationshipType, error) {
	found := &AreaRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for area relationship types that meet criteria.
func (api *Api) AreaRelationshipTypes(ctx context.Context,
	criteria AreaRelationshipTypeSearchCriteria) (*AreaRelationshipTypeSearchResults, error) {
	results := make([]AreaRelationshipType, 0)
	db, pag := api.RDB.ListOf(&AreaRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AreaRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new area relationship.
func (api *Api) CreateAreaRelationship(ctx context.Context, request *AreaRelationshipCreateRequest) (*AreaRelationship, error) {
	// Look up token references.
	source, err := api.AreaByToken(ctx, request.SourceArea)
	if err != nil {
		return nil, err
	}
	target, err := api.AreaByToken(ctx, request.TargetArea)
	if err != nil {
		return nil, err
	}
	rtype, err := api.AreaRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &AreaRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		SourceArea:       *source,
		TargetArea:       *target,
		RelationshipType: *rtype,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Get area relationship by id.
func (api *Api) AreaRelationshipById(ctx context.Context, id uint) (*AreaRelationship, error) {
	found := &AreaRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceArea").Preload("TargetArea").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area relationship by token.
func (api *Api) AreaRelationshipByToken(ctx context.Context, token string) (*AreaRelationship, error) {
	found := &AreaRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceArea").Preload("TargetArea").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for area relationships that meet criteria.
func (api *Api) AreaRelationships(ctx context.Context,
	criteria AreaRelationshipSearchCriteria) (*AreaRelationshipSearchResults, error) {
	results := make([]AreaRelationship, 0)
	db, pag := api.RDB.ListOf(&AreaRelationship{}, nil, criteria.Pagination)
	db.Preload("SourceArea").Preload("TargetArea").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AreaRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new area group.
func (api *Api) CreateAreaGroup(ctx context.Context, request *AreaGroupCreateRequest) (*AreaGroup, error) {
	created := &AreaGroup{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		BrandedEntity: rdb.BrandedEntity{
			ImageUrl:        rdb.NullStrOf(request.ImageUrl),
			Icon:            rdb.NullStrOf(request.Icon),
			BackgroundColor: rdb.NullStrOf(request.BackgroundColor),
			ForegroundColor: rdb.NullStrOf(request.ForegroundColor),
			BorderColor:     rdb.NullStrOf(request.BorderColor),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing area type.
func (api *Api) UpdateAreaGroup(ctx context.Context, token string,
	request *AreaGroupCreateRequest) (*AreaGroup, error) {
	updated, err := api.AreaGroupByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.ImageUrl = rdb.NullStrOf(request.ImageUrl)
	updated.Icon = rdb.NullStrOf(request.Icon)
	updated.BackgroundColor = rdb.NullStrOf(request.BackgroundColor)
	updated.ForegroundColor = rdb.NullStrOf(request.ForegroundColor)
	updated.BorderColor = rdb.NullStrOf(request.BorderColor)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
}

// Get area group by id.
func (api *Api) AreaGroupById(ctx context.Context, id uint) (*AreaGroup, error) {
	found := &AreaGroup{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area group by token.
func (api *Api) AreaGroupByToken(ctx context.Context, token string) (*AreaGroup, error) {
	found := &AreaGroup{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for area groups that meet criteria.
func (api *Api) AreaGroups(ctx context.Context, criteria AreaGroupSearchCriteria) (*AreaGroupSearchResults, error) {
	results := make([]AreaGroup, 0)
	db, pag := api.RDB.ListOf(&AreaGroup{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AreaGroupSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new area group relationship type.
func (api *Api) CreateAreaGroupRelationshipType(ctx context.Context,
	request *AreaGroupRelationshipTypeCreateRequest) (*AreaGroupRelationshipType, error) {
	created := &AreaGroupRelationshipType{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		NamedEntity: rdb.NamedEntity{
			Name:        rdb.NullStrOf(request.Name),
			Description: rdb.NullStrOf(request.Description),
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing area group relationship type.
func (api *Api) UpdateAreaGroupRelationshipType(ctx context.Context, token string,
	request *AreaGroupRelationshipTypeCreateRequest) (*AreaGroupRelationshipType, error) {
	updated, err := api.AreaGroupRelationshipTypeByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
}

// Get area group relationship type by id.
func (api *Api) AreaGroupRelationshipTypeById(ctx context.Context, id uint) (*AreaGroupRelationshipType, error) {
	found := &AreaGroupRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area group relationship type by token.
func (api *Api) AreaGroupRelationshipTypeByToken(ctx context.Context, token string) (*AreaGroupRelationshipType, error) {
	found := &AreaGroupRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for area group relationship types that meet criteria.
func (api *Api) AreaGroupRelationshipTypes(ctx context.Context,
	criteria AreaGroupRelationshipTypeSearchCriteria) (*AreaGroupRelationshipTypeSearchResults, error) {
	results := make([]AreaGroupRelationshipType, 0)
	db, pag := api.RDB.ListOf(&AreaGroupRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AreaGroupRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new area group relationship.
func (api *Api) CreateAreaGroupRelationship(ctx context.Context,
	request *AreaGroupRelationshipCreateRequest) (*AreaGroupRelationship, error) {

	// Look up token references.
	source, err := api.AreaGroupByToken(ctx, request.AreaGroup)
	if err != nil {
		return nil, err
	}
	target, err := api.AreaByToken(ctx, request.Area)
	if err != nil {
		return nil, err
	}
	rtype, err := api.AreaGroupRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &AreaGroupRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		AreaGroup:        *source,
		Area:             *target,
		RelationshipType: *rtype,
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Get area group relationship by id.
func (api *Api) AreaGroupRelationshipById(ctx context.Context, id uint) (*AreaGroupRelationship, error) {
	found := &AreaGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("AreaGroup").Preload("Area").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get area group relationship by token.
func (api *Api) AreaGroupRelationshipByToken(ctx context.Context, token string) (*AreaGroupRelationship, error) {
	found := &AreaGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("AreaGroup").Preload("Area").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for area group relationships that meet criteria.
func (api *Api) AreaGroupRelationships(ctx context.Context,
	criteria AreaGroupRelationshipSearchCriteria) (*AreaGroupRelationshipSearchResults, error) {
	results := make([]AreaGroupRelationship, 0)
	db, pag := api.RDB.ListOf(&AreaGroupRelationship{}, nil, criteria.Pagination)
	db.Preload("AreaGroup").Preload("Area").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &AreaGroupRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}
