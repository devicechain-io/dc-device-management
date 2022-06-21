/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package gqlclient

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/devicechain-io/dc-device-management/model"
)

// Assure that an customer type exists.
func AssureCustomerType(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerTypeCreateRequest,
) (*getCustomerTypesByTokenResponse, *createCustomerTypeResponse, error) {
	gresp, err := GetCustomerTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.CustomerTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateCustomerType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new customer type.
func CreateCustomerType(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerTypeCreateRequest,
) (*createCustomerTypeResponse, error) {
	return createCustomerType(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get customer types by token.
func GetCustomerTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getCustomerTypesByTokenResponse, error) {
	return getCustomerTypesByToken(ctx, client, tokens)
}

// List customer types based on criteria.
func ListCustomerTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listCustomerTypesResponse, error) {
	return listCustomerTypes(ctx, client, pageNumber, pageSize)
}

// Assure that a customer exists.
func AssureCustomer(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerCreateRequest,
) (*getCustomersByTokenResponse, *createCustomerResponse, error) {
	gresp, err := GetCustomersByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.CustomersByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateCustomer(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new customer.
func CreateCustomer(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerCreateRequest,
) (*createCustomerResponse, error) {
	return createCustomer(ctx, client, request.Token, request.CustomerTypeToken,
		blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get customers by token.
func GetCustomersByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getCustomersByTokenResponse, error) {
	return getCustomersByToken(ctx, client, tokens)
}

// List customers based on criteria.
func ListCustomers(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listCustomersResponse, error) {
	return listCustomers(ctx, client, pageNumber, pageSize)
}

// Assure that an customer relationship type exists.
func AssureCustomerRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerRelationshipTypeCreateRequest,
) (*getCustomerRelationshipTypesByTokenResponse, *createCustomerRelationshipTypeResponse, error) {
	gresp, err := GetCustomerRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.CustomerRelationshipTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateCustomerRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new customer relationship type.
func CreateCustomerRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerRelationshipTypeCreateRequest,
) (*createCustomerRelationshipTypeResponse, error) {
	return createCustomerRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get customer relationship types by token.
func GetCustomerRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getCustomerRelationshipTypesByTokenResponse, error) {
	return getCustomerRelationshipTypesByToken(ctx, client, tokens)
}

// List customer relationship types based on criteria.
func ListCustomerRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listCustomerRelationshipTypesResponse, error) {
	return listCustomerRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that a customer relationship exists.
func AssureCustomerRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerRelationshipCreateRequest,
) (*getCustomerRelationshipsByTokenResponse, *createCustomerRelationshipResponse, error) {
	gresp, err := GetCustomerRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.CustomerRelationshipsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateCustomerRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new customer relationship.
func CreateCustomerRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerRelationshipCreateRequest,
) (*createCustomerRelationshipResponse, error) {
	return createCustomerRelationship(ctx, client, request.Token, request.SourceCustomer, request.TargetCustomer, request.RelationshipType)
}

// Get customer relationships by token.
func GetCustomerRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getCustomerRelationshipsByTokenResponse, error) {
	return getCustomerRelationshipsByToken(ctx, client, tokens)
}

// List customer relationships based on criteria.
func ListCustomerRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listCustomerRelationshipsResponse, error) {
	return listCustomerRelationships(ctx, client, pageNumber, pageSize)
}

// Assure that a customer group exists.
func AssureCustomerGroup(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerGroupCreateRequest,
) (*getCustomerGroupsByTokenResponse, *createCustomerGroupResponse, error) {
	gresp, err := GetCustomerGroupsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.CustomerGroupsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateCustomerGroup(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new customer group.
func CreateCustomerGroup(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerGroupCreateRequest,
) (*createCustomerGroupResponse, error) {
	return createCustomerGroup(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get customer groups by token.
func GetCustomerGroupsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getCustomerGroupsByTokenResponse, error) {
	return getCustomerGroupsByToken(ctx, client, tokens)
}

// List customer groups based on criteria.
func ListCustomerGroups(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listCustomerGroupsResponse, error) {
	return listCustomerGroups(ctx, client, pageNumber, pageSize)
}

// Assure that a customer group relationship type exists.
func AssureCustomerGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerGroupRelationshipTypeCreateRequest,
) (*getCustomerGroupRelationshipTypesByTokenResponse, *createCustomerGroupRelationshipTypeResponse, error) {
	gresp, err := GetCustomerGroupRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.CustomerGroupRelationshipTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateCustomerGroupRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new customer group relationship type.
func CreateCustomerGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerGroupRelationshipTypeCreateRequest,
) (*createCustomerGroupRelationshipTypeResponse, error) {
	return createCustomerGroupRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get customer group relationship types by token.
func GetCustomerGroupRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getCustomerGroupRelationshipTypesByTokenResponse, error) {
	return getCustomerGroupRelationshipTypesByToken(ctx, client, tokens)
}

// List customer group relationship types based on criteria.
func ListCustomerGroupRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listCustomerGroupRelationshipTypesResponse, error) {
	return listCustomerGroupRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that a customer group relationship exists.
func AssureCustomerGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerGroupRelationshipCreateRequest,
) (*getCustomerGroupRelationshipsByTokenResponse, *createCustomerGroupRelationshipResponse, error) {
	gresp, err := GetCustomerGroupRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.CustomerGroupRelationshipsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateCustomerGroupRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new customer group relationship.
func CreateCustomerGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.CustomerGroupRelationshipCreateRequest,
) (*createCustomerGroupRelationshipResponse, error) {
	return createCustomerGroupRelationship(ctx, client, request.Token, request.CustomerGroup, request.Customer, request.RelationshipType)
}

// Get customer group relationships by token.
func GetCustomerGroupRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getCustomerGroupRelationshipsByTokenResponse, error) {
	return getCustomerGroupRelationshipsByToken(ctx, client, tokens)
}

// List customer group relationships based on criteria.
func ListCustomerGroupRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listCustomerGroupRelationshipsResponse, error) {
	return listCustomerGroupRelationships(ctx, client, pageNumber, pageSize)
}
