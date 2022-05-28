/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"context"

	"github.com/devicechain-io/dc-microservice/rdb"
	"gorm.io/gorm"
)

// Create a new device type.
func (api *Api) CreateDeviceType(ctx context.Context, request *DeviceTypeCreateRequest) (*DeviceType, error) {
	created := &DeviceType{
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

// Update an existing device type.
func (api *Api) UpdateDeviceType(ctx context.Context, token string,
	request *DeviceTypeCreateRequest) (*DeviceType, error) {
	found, err := api.DeviceTypeByToken(ctx, request.Token)
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

// Get device type by id.
func (api *Api) DeviceTypeById(ctx context.Context, id uint) (*DeviceType, error) {
	found := &DeviceType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device type by token.
func (api *Api) DeviceTypeByToken(ctx context.Context, token string) (*DeviceType, error) {
	found := &DeviceType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device types that meet criteria.
func (api *Api) DeviceTypes(ctx context.Context, criteria DeviceTypeSearchCriteria) (*DeviceTypeSearchResults, error) {
	results := make([]DeviceType, 0)
	db, pag := api.RDB.ListOf(&DeviceType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new device.
func (api *Api) CreateDevice(ctx context.Context, request *DeviceCreateRequest) (*Device, error) {
	dtr, err := api.DeviceTypeByToken(ctx, request.DeviceTypeToken)
	if err != nil {
		return nil, err
	}

	created := &Device{
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
		DeviceType: dtr,
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing device.
func (api *Api) UpdateDevice(ctx context.Context, token string, request *DeviceCreateRequest) (*Device, error) {
	updated, err := api.DeviceByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	// Update fields that changed.
	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	// Update device type if changed.
	if request.DeviceTypeToken != updated.DeviceType.Token {
		dtr, err := api.DeviceTypeByToken(ctx, request.DeviceTypeToken)
		if err != nil {
			return nil, err
		}
		updated.DeviceType = dtr
	}

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
}

// Get device by id.
func (api *Api) DeviceById(ctx context.Context, id uint) (*Device, error) {
	found := &Device{}
	result := api.RDB.Database
	result = result.Preload("DeviceType")
	result = result.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device by token.
func (api *Api) DeviceByToken(ctx context.Context, token string) (*Device, error) {
	found := &Device{}
	result := api.RDB.Database
	result = result.Preload("DeviceType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for devices that meet criteria.
func (api *Api) Devices(ctx context.Context, criteria DeviceSearchCriteria) (*DeviceSearchResults, error) {
	results := make([]Device, 0)
	db, pag := api.RDB.ListOf(&Device{}, func(result *gorm.DB) *gorm.DB {
		if criteria.DeviceType != nil {
			result = result.Where("device_type_id = (?)",
				api.RDB.Database.Model(&DeviceType{}).Select("id").Where("token = ?", criteria.DeviceType))
		}
		return result.Preload("DeviceType")
	}, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new device relationship type.
func (api *Api) CreateDeviceRelationshipType(ctx context.Context, request *DeviceRelationshipTypeCreateRequest) (*DeviceRelationshipType, error) {
	created := &DeviceRelationshipType{
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

// Update an existing device relationship type.
func (api *Api) UpdateDeviceRelationshipType(ctx context.Context, token string,
	request *DeviceRelationshipTypeCreateRequest) (*DeviceRelationshipType, error) {
	updated, err := api.DeviceRelationshipTypeByToken(ctx, request.Token)
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

// Get device relationship type by id.
func (api *Api) DeviceRelationshipTypeById(ctx context.Context, id uint) (*DeviceRelationshipType, error) {
	found := &DeviceRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device relationship type by token.
func (api *Api) DeviceRelationshipTypeByToken(ctx context.Context, token string) (*DeviceRelationshipType, error) {
	found := &DeviceRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device relationship types that meet criteria.
func (api *Api) DeviceRelationshipTypes(ctx context.Context,
	criteria DeviceRelationshipTypeSearchCriteria) (*DeviceRelationshipTypeSearchResults, error) {
	results := make([]DeviceRelationshipType, 0)
	db, pag := api.RDB.ListOf(&DeviceRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new device relationship.
func (api *Api) CreateDeviceRelationship(ctx context.Context, request *DeviceRelationshipCreateRequest) (*DeviceRelationship, error) {
	// Look up token references.
	source, err := api.DeviceByToken(ctx, request.SourceDevice)
	if err != nil {
		return nil, err
	}
	target, err := api.DeviceByToken(ctx, request.TargetDevice)
	if err != nil {
		return nil, err
	}
	rtype, err := api.DeviceRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &DeviceRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		SourceDevice:     *source,
		TargetDevice:     *target,
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

// Get device relationship by id.
func (api *Api) DeviceRelationshipById(ctx context.Context, id uint) (*DeviceRelationship, error) {
	found := &DeviceRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceDevice").Preload("TargetDevice").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device relationship by token.
func (api *Api) DeviceRelationshipByToken(ctx context.Context, token string) (*DeviceRelationship, error) {
	found := &DeviceRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceDevice").Preload("TargetDevice").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device relationships that meet criteria.
func (api *Api) DeviceRelationships(ctx context.Context,
	criteria DeviceRelationshipSearchCriteria) (*DeviceRelationshipSearchResults, error) {
	results := make([]DeviceRelationship, 0)
	db, pag := api.RDB.ListOf(&DeviceRelationship{}, nil, criteria.Pagination)
	db.Preload("SourceDevice").Preload("TargetDevice").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new device group.
func (api *Api) CreateDeviceGroup(ctx context.Context, request *DeviceGroupCreateRequest) (*DeviceGroup, error) {
	created := &DeviceGroup{
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

// Update an existing device type.
func (api *Api) UpdateDeviceGroup(ctx context.Context, token string,
	request *DeviceGroupCreateRequest) (*DeviceGroup, error) {
	updated, err := api.DeviceGroupByToken(ctx, request.Token)
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

// Get device group by id.
func (api *Api) DeviceGroupById(ctx context.Context, id uint) (*DeviceGroup, error) {
	found := &DeviceGroup{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device group by token.
func (api *Api) DeviceGroupByToken(ctx context.Context, token string) (*DeviceGroup, error) {
	found := &DeviceGroup{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device groups that meet criteria.
func (api *Api) DeviceGroups(ctx context.Context, criteria DeviceGroupSearchCriteria) (*DeviceGroupSearchResults, error) {
	results := make([]DeviceGroup, 0)
	db, pag := api.RDB.ListOf(&DeviceGroup{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceGroupSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new device group relationship type.
func (api *Api) CreateDeviceGroupRelationshipType(ctx context.Context,
	request *DeviceGroupRelationshipTypeCreateRequest) (*DeviceGroupRelationshipType, error) {
	created := &DeviceGroupRelationshipType{
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

// Update an existing device group relationship type.
func (api *Api) UpdateDeviceGroupRelationshipType(ctx context.Context, token string,
	request *DeviceGroupRelationshipTypeCreateRequest) (*DeviceGroupRelationshipType, error) {
	updated, err := api.DeviceGroupRelationshipTypeByToken(ctx, request.Token)
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

// Get device group relationship type by id.
func (api *Api) DeviceGroupRelationshipTypeById(ctx context.Context, id uint) (*DeviceGroupRelationshipType, error) {
	found := &DeviceGroupRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device group relationship type by token.
func (api *Api) DeviceGroupRelationshipTypeByToken(ctx context.Context, token string) (*DeviceGroupRelationshipType, error) {
	found := &DeviceGroupRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device group relationship types that meet criteria.
func (api *Api) DeviceGroupRelationshipTypes(ctx context.Context,
	criteria DeviceGroupRelationshipTypeSearchCriteria) (*DeviceGroupRelationshipTypeSearchResults, error) {
	results := make([]DeviceGroupRelationshipType, 0)
	db, pag := api.RDB.ListOf(&DeviceGroupRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceGroupRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new device group relationship.
func (api *Api) CreateDeviceGroupRelationship(ctx context.Context,
	request *DeviceGroupRelationshipCreateRequest) (*DeviceGroupRelationship, error) {

	// Look up token references.
	source, err := api.DeviceGroupByToken(ctx, request.DeviceGroup)
	if err != nil {
		return nil, err
	}
	target, err := api.DeviceByToken(ctx, request.Device)
	if err != nil {
		return nil, err
	}
	rtype, err := api.DeviceGroupRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &DeviceGroupRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		DeviceGroup:      *source,
		Device:           *target,
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

// Get device group relationship by id.
func (api *Api) DeviceGroupRelationshipById(ctx context.Context, id uint) (*DeviceGroupRelationship, error) {
	found := &DeviceGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("DeviceGroup").Preload("Device").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device group relationship by token.
func (api *Api) DeviceGroupRelationshipByToken(ctx context.Context, token string) (*DeviceGroupRelationship, error) {
	found := &DeviceGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("DeviceGroup").Preload("Device").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device group relationships that meet criteria.
func (api *Api) DeviceGroupRelationships(ctx context.Context,
	criteria DeviceGroupRelationshipSearchCriteria) (*DeviceGroupRelationshipSearchResults, error) {
	results := make([]DeviceGroupRelationship, 0)
	db, pag := api.RDB.ListOf(&DeviceGroupRelationship{}, nil, criteria.Pagination)
	db.Preload("DeviceGroup").Preload("Device").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceGroupRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}
