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

// Create a new customer type.
func (api *Api) CreateCustomerType(ctx context.Context, request *CustomerTypeCreateRequest) (*CustomerType, error) {
	created := &CustomerType{
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

// Update an existing customer type.
func (api *Api) UpdateCustomerType(ctx context.Context, token string,
	request *CustomerTypeCreateRequest) (*CustomerType, error) {
	found, err := api.CustomerTypeByToken(ctx, request.Token)
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

// Get customer type by id.
func (api *Api) CustomerTypeById(ctx context.Context, id uint) (*CustomerType, error) {
	found := &CustomerType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get customer type by token.
func (api *Api) CustomerTypeByToken(ctx context.Context, token string) (*CustomerType, error) {
	found := &CustomerType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for customer types that meet criteria.
func (api *Api) CustomerTypes(ctx context.Context, criteria CustomerTypeSearchCriteria) (*CustomerTypeSearchResults, error) {
	results := make([]CustomerType, 0)
	db, pag := api.RDB.ListOf(&CustomerType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &CustomerTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new customer.
func (api *Api) CreateCustomer(ctx context.Context, request *CustomerCreateRequest) (*Customer, error) {
	dtr, err := api.CustomerTypeByToken(ctx, request.CustomerTypeToken)
	if err != nil {
		return nil, err
	}

	created := &Customer{
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
		CustomerType: dtr,
	}
	result := api.RDB.Database.Create(created)
	if result.Error != nil {
		return nil, result.Error
	}
	return created, nil
}

// Update an existing customer.
func (api *Api) UpdateCustomer(ctx context.Context, token string, request *CustomerCreateRequest) (*Customer, error) {
	updated, err := api.CustomerByToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}

	// Update fields that changed.
	updated.Token = request.Token
	updated.Name = rdb.NullStrOf(request.Name)
	updated.Description = rdb.NullStrOf(request.Description)
	updated.Metadata = rdb.MetadataStrOf(request.Metadata)

	// Update customer type if changed.
	if request.CustomerTypeToken != updated.CustomerType.Token {
		dtr, err := api.CustomerTypeByToken(ctx, request.CustomerTypeToken)
		if err != nil {
			return nil, err
		}
		updated.CustomerType = dtr
	}

	result := api.RDB.Database.Save(updated)
	if result.Error != nil {
		return nil, result.Error
	}
	return updated, nil
}

// Get customer by id.
func (api *Api) CustomerById(ctx context.Context, id uint) (*Customer, error) {
	found := &Customer{}
	result := api.RDB.Database
	result = result.Preload("CustomerType")
	result = result.First(found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get customer by token.
func (api *Api) CustomerByToken(ctx context.Context, token string) (*Customer, error) {
	found := &Customer{}
	result := api.RDB.Database
	result = result.Preload("CustomerType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for customers that meet criteria.
func (api *Api) Customers(ctx context.Context, criteria CustomerSearchCriteria) (*CustomerSearchResults, error) {
	results := make([]Customer, 0)
	db, pag := api.RDB.ListOf(&Customer{}, func(result *gorm.DB) *gorm.DB {
		if criteria.CustomerTypeToken != nil {
			result = result.Where("customer_type_id = (?)",
				api.RDB.Database.Model(&CustomerType{}).Select("id").Where("token = ?", criteria.CustomerTypeToken))
		}
		return result.Preload("CustomerType")
	}, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &CustomerSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new customer relationship type.
func (api *Api) CreateCustomerRelationshipType(ctx context.Context, request *CustomerRelationshipTypeCreateRequest) (*CustomerRelationshipType, error) {
	created := &CustomerRelationshipType{
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

// Update an existing customer relationship type.
func (api *Api) UpdateCustomerRelationshipType(ctx context.Context, token string,
	request *CustomerRelationshipTypeCreateRequest) (*CustomerRelationshipType, error) {
	updated, err := api.CustomerRelationshipTypeByToken(ctx, request.Token)
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

// Get customer relationship type by id.
func (api *Api) CustomerRelationshipTypeById(ctx context.Context, id uint) (*CustomerRelationshipType, error) {
	found := &CustomerRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get customer relationship type by token.
func (api *Api) CustomerRelationshipTypeByToken(ctx context.Context, token string) (*CustomerRelationshipType, error) {
	found := &CustomerRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for customer relationship types that meet criteria.
func (api *Api) CustomerRelationshipTypes(ctx context.Context,
	criteria CustomerRelationshipTypeSearchCriteria) (*CustomerRelationshipTypeSearchResults, error) {
	results := make([]CustomerRelationshipType, 0)
	db, pag := api.RDB.ListOf(&CustomerRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &CustomerRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new customer relationship.
func (api *Api) CreateCustomerRelationship(ctx context.Context, request *CustomerRelationshipCreateRequest) (*CustomerRelationship, error) {
	// Look up token references.
	source, err := api.CustomerByToken(ctx, request.SourceCustomer)
	if err != nil {
		return nil, err
	}
	target, err := api.CustomerByToken(ctx, request.TargetCustomer)
	if err != nil {
		return nil, err
	}
	rtype, err := api.CustomerRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &CustomerRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		SourceCustomer:   *source,
		TargetCustomer:   *target,
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

// Get customer relationship by id.
func (api *Api) CustomerRelationshipById(ctx context.Context, id uint) (*CustomerRelationship, error) {
	found := &CustomerRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceCustomer").Preload("TargetCustomer").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get customer relationship by token.
func (api *Api) CustomerRelationshipByToken(ctx context.Context, token string) (*CustomerRelationship, error) {
	found := &CustomerRelationship{}
	result := api.RDB.Database
	result = result.Preload("SourceCustomer").Preload("TargetCustomer").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for customer relationships that meet criteria.
func (api *Api) CustomerRelationships(ctx context.Context,
	criteria CustomerRelationshipSearchCriteria) (*CustomerRelationshipSearchResults, error) {
	results := make([]CustomerRelationship, 0)
	db, pag := api.RDB.ListOf(&CustomerRelationship{}, nil, criteria.Pagination)
	db.Preload("SourceCustomer").Preload("TargetCustomer").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &CustomerRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new customer group.
func (api *Api) CreateCustomerGroup(ctx context.Context, request *CustomerGroupCreateRequest) (*CustomerGroup, error) {
	created := &CustomerGroup{
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

// Update an existing customer type.
func (api *Api) UpdateCustomerGroup(ctx context.Context, token string,
	request *CustomerGroupCreateRequest) (*CustomerGroup, error) {
	updated, err := api.CustomerGroupByToken(ctx, request.Token)
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

// Get customer group by id.
func (api *Api) CustomerGroupById(ctx context.Context, id uint) (*CustomerGroup, error) {
	found := &CustomerGroup{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get customer group by token.
func (api *Api) CustomerGroupByToken(ctx context.Context, token string) (*CustomerGroup, error) {
	found := &CustomerGroup{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for customer groups that meet criteria.
func (api *Api) CustomerGroups(ctx context.Context, criteria CustomerGroupSearchCriteria) (*CustomerGroupSearchResults, error) {
	results := make([]CustomerGroup, 0)
	db, pag := api.RDB.ListOf(&CustomerGroup{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &CustomerGroupSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new customer group relationship type.
func (api *Api) CreateCustomerGroupRelationshipType(ctx context.Context,
	request *CustomerGroupRelationshipTypeCreateRequest) (*CustomerGroupRelationshipType, error) {
	created := &CustomerGroupRelationshipType{
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

// Update an existing customer group relationship type.
func (api *Api) UpdateCustomerGroupRelationshipType(ctx context.Context, token string,
	request *CustomerGroupRelationshipTypeCreateRequest) (*CustomerGroupRelationshipType, error) {
	updated, err := api.CustomerGroupRelationshipTypeByToken(ctx, request.Token)
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

// Get customer group relationship type by id.
func (api *Api) CustomerGroupRelationshipTypeById(ctx context.Context, id uint) (*CustomerGroupRelationshipType, error) {
	found := &CustomerGroupRelationshipType{}
	result := api.RDB.Database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get customer group relationship type by token.
func (api *Api) CustomerGroupRelationshipTypeByToken(ctx context.Context, token string) (*CustomerGroupRelationshipType, error) {
	found := &CustomerGroupRelationshipType{}
	result := api.RDB.Database.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for customer group relationship types that meet criteria.
func (api *Api) CustomerGroupRelationshipTypes(ctx context.Context,
	criteria CustomerGroupRelationshipTypeSearchCriteria) (*CustomerGroupRelationshipTypeSearchResults, error) {
	results := make([]CustomerGroupRelationshipType, 0)
	db, pag := api.RDB.ListOf(&CustomerGroupRelationshipType{}, nil, criteria.Pagination)
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &CustomerGroupRelationshipTypeSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}

// Create a new customer group relationship.
func (api *Api) CreateCustomerGroupRelationship(ctx context.Context,
	request *CustomerGroupRelationshipCreateRequest) (*CustomerGroupRelationship, error) {

	// Look up token references.
	source, err := api.CustomerGroupByToken(ctx, request.CustomerGroup)
	if err != nil {
		return nil, err
	}
	target, err := api.CustomerByToken(ctx, request.Customer)
	if err != nil {
		return nil, err
	}
	rtype, err := api.CustomerGroupRelationshipTypeByToken(ctx, request.RelationshipType)
	if err != nil {
		return nil, err
	}

	created := &CustomerGroupRelationship{
		TokenReference: rdb.TokenReference{
			Token: request.Token,
		},
		CustomerGroup:    *source,
		Customer:         *target,
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

// Get customer group relationship by id.
func (api *Api) CustomerGroupRelationshipById(ctx context.Context, id uint) (*CustomerGroupRelationship, error) {
	found := &CustomerGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("CustomerGroup").Preload("Customer").Preload("RelationshipType")
	result = result.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Get customer group relationship by token.
func (api *Api) CustomerGroupRelationshipByToken(ctx context.Context, token string) (*CustomerGroupRelationship, error) {
	found := &CustomerGroupRelationship{}
	result := api.RDB.Database
	result = result.Preload("CustomerGroup").Preload("Customer").Preload("RelationshipType")
	result = result.First(&found, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return found, nil
}

// Search for customer group relationships that meet criteria.
func (api *Api) CustomerGroupRelationships(ctx context.Context,
	criteria CustomerGroupRelationshipSearchCriteria) (*CustomerGroupRelationshipSearchResults, error) {
	results := make([]CustomerGroupRelationship, 0)
	db, pag := api.RDB.ListOf(&CustomerGroupRelationship{}, nil, criteria.Pagination)
	db.Preload("CustomerGroup").Preload("Customer").Preload("RelationshipType")
	db.Find(&results)
	if db.Error != nil {
		return nil, db.Error
	}

	// Wrap as search results.
	return &CustomerGroupRelationshipSearchResults{
		Results:    results,
		Pagination: pag,
	}, nil
}
