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

// Create a new device assignment status.
func (api *Api) CreateDeviceAssignmentStatus(ctx context.Context, request *DeviceAssignmentStatusCreateRequest) (*DeviceAssignmentStatus, error) {
	created := &DeviceAssignmentStatus{
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

// Update an existing device assignment status.
func (api *Api) UpdateDeviceAssignmentStatus(ctx context.Context, token string,
	request *DeviceAssignmentStatusCreateRequest) (*DeviceAssignmentStatus, error) {
	matches, err := api.DeviceAssignmentStatusesByToken(ctx, []string{request.Token})
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	found := matches[0]
	found.Token = request.Token
	found.Name = rdb.NullStrOf(request.Name)
	found.Description = rdb.NullStrOf(request.Description)
	found.Metadata = rdb.MetadataStrOf(request.Metadata)

	result := api.RDB.Database.Save(found)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device assignment statuses by id.
func (api *Api) DeviceAssignmentStatusesById(ctx context.Context, ids []uint) ([]*DeviceAssignmentStatus, error) {
	found := make([]*DeviceAssignmentStatus, 0)
	result := api.RDB.Database.Find(&found, ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device assignment statuses by token.
func (api *Api) DeviceAssignmentStatusesByToken(ctx context.Context, tokens []string) ([]*DeviceAssignmentStatus, error) {
	found := make([]*DeviceAssignmentStatus, 0)
	result := api.RDB.Database.Find(&found, "token in ?", tokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device assignment statuses that meet criteria.
func (api *Api) DeviceAssignmentStatuses(ctx context.Context,
	criteria DeviceAssignmentStatusSearchCriteria) (*DeviceAssignmentStatusSearchResults, error) {
	results := make([]DeviceAssignmentStatus, 0)
	db, pag := api.RDB.ListOf(&DeviceAssignmentStatus{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceAssignmentStatusSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new device assignment.
func (api *Api) CreateDeviceAssignment(ctx context.Context, request *DeviceAssignmentCreateRequest) (*DeviceAssignment, error) {
	created := &DeviceAssignment{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		MetadataEntity: rdb.MetadataEntity{
			Metadata: rdb.MetadataStrOf(request.Metadata),
		},
	}

	// Associate device.
	matches, err := api.DevicesByToken(ctx, []string{request.Device})
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	created.DeviceId = matches[0].ID

	// Associate device group if provided.
	if request.DeviceGroup != nil {
		matches, err := api.DeviceGroupsByToken(ctx, []string{*request.DeviceGroup})
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		created.DeviceGroupId = &matches[0].ID
	}

	// Associate asset if provided.
	if request.Asset != nil {
		matches, err := api.AssetsByToken(ctx, []string{*request.Asset})
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		created.AssetId = &matches[0].ID
	}

	// Associate asset group if provided.
	if request.AssetGroup != nil {
		matches, err := api.AssetGroupsByToken(ctx, []string{*request.AssetGroup})
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		created.AssetGroupId = &matches[0].ID
	}

	// Associate customer if provided.
	if request.Customer != nil {
		matches, err := api.CustomersByToken(ctx, []string{*request.Customer})
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		created.CustomerId = &matches[0].ID
	}

	// Associate customer group if provided.
	if request.CustomerGroup != nil {
		matches, err := api.CustomerGroupsByToken(ctx, []string{*request.CustomerGroup})
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		created.CustomerGroupId = &matches[0].ID
	}

	// Associate area if provided.
	if request.Area != nil {
		matches, err := api.AreasByToken(ctx, []string{*request.Area})
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		created.AreaId = &matches[0].ID
	}

	// Associate area group if provided.
	if request.AreaGroup != nil {
		matches, err := api.AreaGroupsByToken(ctx, []string{*request.AreaGroup})
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			return nil, gorm.ErrRecordNotFound
		}
		created.AreaGroupId = &matches[0].ID
	}

	// Set active flag.
	if request.Active != nil {
		created.Active = *request.Active
	} else {
		created.Active = true
	}

	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing device assignment.
func (api *Api) UpdateDeviceAssignment(ctx context.Context, token string,
	request *DeviceAssignmentCreateRequest) (*DeviceAssignment, error) {
	matches, err := api.DeviceAssignmentsByToken(ctx, []string{request.Token})
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	found := matches[0]
	found.Token = request.Token
	found.Metadata = rdb.MetadataStrOf(request.Metadata)

	// Other references may not be updated after original create.

	result := api.RDB.Database.Save(found)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device assignment by id.
func (api *Api) DeviceAssignmentsById(ctx context.Context, ids []uint) ([]*DeviceAssignment, error) {
	found := make([]*DeviceAssignment, 0)
	result := api.RDB.Database
	result = result.Preload("Device").Preload("DeviceGroup").Preload("Asset").Preload("AssetGroup")
	result = result.Preload("Customer").Preload("CustomerGroup").Preload("Area").Preload("AreaGroup")
	result = result.Find(&found, ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get device assignments by token.
func (api *Api) DeviceAssignmentsByToken(ctx context.Context, tokens []string) ([]*DeviceAssignment, error) {
	found := make([]*DeviceAssignment, 0)
	result := api.RDB.Database
	result = result.Preload("Device").Preload("DeviceGroup").Preload("Asset").Preload("AssetGroup")
	result = result.Preload("Customer").Preload("CustomerGroup").Preload("Area").Preload("AreaGroup")
	result = result.Find(&found, "token in ?", tokens)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for device assignments that meet criteria.
func (api *Api) DeviceAssignments(ctx context.Context,
	criteria DeviceAssignmentSearchCriteria) (*DeviceAssignmentSearchResults, error) {
	results := make([]DeviceAssignment, 0)
	db, pag := api.RDB.ListOf(&DeviceAssignment{}, func(db *gorm.DB) *gorm.DB {
		if criteria.Device != nil {
			db = db.Where("device_id = (?)",
				api.RDB.Database.Model(&Device{}).Select("id").Where("token = ?", criteria.Device))
		}
		if criteria.DeviceGroup != nil {
			db = db.Where("device_group_id = (?)",
				api.RDB.Database.Model(&DeviceGroup{}).Select("id").Where("token = ?", criteria.DeviceGroup))
		}
		db = db.Preload("Device").Preload("DeviceGroup").Preload("Asset").Preload("AssetGroup")
		db = db.Preload("Customer").Preload("CustomerGroup").Preload("Area").Preload("AreaGroup")
		return db
	}, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &DeviceAssignmentSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Get all device assignments for the given device id.
func (api *Api) ActiveDeviceAssignmentsForDevice(ctx context.Context, id uint) ([]DeviceAssignment, error) {
	results := make([]DeviceAssignment, 0)
	db, _ := api.RDB.ListOf(&DeviceAssignment{}, func(db *gorm.DB) *gorm.DB {
		db = db.Where(&DeviceAssignment{DeviceId: id}).Where(&DeviceAssignment{Active: true})
		db = db.Preload("Device").Preload("DeviceGroup").Preload("Asset").Preload("AssetGroup")
		db = db.Preload("Customer").Preload("CustomerGroup").Preload("Area").Preload("AreaGroup")
		return db
	}, rdb.Pagination{PageNumber: 1, PageSize: 0})
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}
	return results, nil
}
