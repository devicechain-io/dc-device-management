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

// Assure that an area type exists.
func AssureAreaType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaTypeCreateRequest,
) (*getAreaTypesByTokenResponse, *createAreaTypeResponse, error) {
	gresp, err := GetAreaTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AreaTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAreaType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new area type.
func CreateAreaType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaTypeCreateRequest,
) (*createAreaTypeResponse, error) {
	return createAreaType(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get area types by token.
func GetAreaTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAreaTypesByTokenResponse, error) {
	return getAreaTypesByToken(ctx, client, tokens)
}

// List area types based on criteria.
func ListAreaTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAreaTypesResponse, error) {
	return listAreaTypes(ctx, client, pageNumber, pageSize)
}

// Assure that an area exists.
func AssureArea(
	ctx context.Context,
	client graphql.Client,
	request model.AreaCreateRequest,
) (*getAreasByTokenResponse, *createAreaResponse, error) {
	gresp, err := GetAreasByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AreasByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateArea(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new area.
func CreateArea(
	ctx context.Context,
	client graphql.Client,
	request model.AreaCreateRequest,
) (*createAreaResponse, error) {
	return createArea(ctx, client, request.Token, request.AreaTypeToken,
		blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get areas by token.
func GetAreasByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAreasByTokenResponse, error) {
	return getAreasByToken(ctx, client, tokens)
}

// List areas based on criteria.
func ListAreas(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAreasResponse, error) {
	return listAreas(ctx, client, pageNumber, pageSize)
}

// Assure that an area relationship type exists.
func AssureAreaRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipTypeCreateRequest,
) (*getAreaRelationshipTypesByTokenResponse, *createAreaRelationshipTypeResponse, error) {
	gresp, err := GetAreaRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AreaRelationshipTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAreaRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new area relationship type.
func CreateAreaRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipTypeCreateRequest,
) (*createAreaRelationshipTypeResponse, error) {
	return createAreaRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get area relationship types by token.
func GetAreaRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAreaRelationshipTypesByTokenResponse, error) {
	return getAreaRelationshipTypesByToken(ctx, client, tokens)
}

// List area relationship types based on criteria.
func ListAreaRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAreaRelationshipTypesResponse, error) {
	return listAreaRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that an area relationship exists.
func AssureAreaRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipCreateRequest,
) (*getAreaRelationshipsByTokenResponse, *createAreaRelationshipResponse, error) {
	gresp, err := GetAreaRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AreaRelationshipsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAreaRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new area relationship.
func CreateAreaRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipCreateRequest,
) (*createAreaRelationshipResponse, error) {
	return createAreaRelationship(ctx, client, request.Token, request.SourceArea, request.TargetArea, request.RelationshipType)
}

// Get area relationships by token.
func GetAreaRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAreaRelationshipsByTokenResponse, error) {
	return getAreaRelationshipsByToken(ctx, client, tokens)
}

// List area relationships based on criteria.
func ListAreaRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAreaRelationshipsResponse, error) {
	return listAreaRelationships(ctx, client, pageNumber, pageSize)
}

// Assure that an area group exists.
func AssureAreaGroup(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupCreateRequest,
) (*getAreaGroupsByTokenResponse, *createAreaGroupResponse, error) {
	gresp, err := GetAreaGroupsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AreaGroupsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAreaGroup(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new area group.
func CreateAreaGroup(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupCreateRequest,
) (*createAreaGroupResponse, error) {
	return createAreaGroup(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get area groups by token.
func GetAreaGroupsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAreaGroupsByTokenResponse, error) {
	return getAreaGroupsByToken(ctx, client, tokens)
}

// List area groups based on criteria.
func ListAreaGroups(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAreaGroupsResponse, error) {
	return listAreaGroups(ctx, client, pageNumber, pageSize)
}

// Assure that an area group relationship type exists.
func AssureAreaGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipTypeCreateRequest,
) (*getAreaGroupRelationshipTypesByTokenResponse, *createAreaGroupRelationshipTypeResponse, error) {
	gresp, err := GetAreaGroupRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AreaGroupRelationshipTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAreaGroupRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new area group relationship type.
func CreateAreaGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipTypeCreateRequest,
) (*createAreaGroupRelationshipTypeResponse, error) {
	return createAreaGroupRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get area group relationship types by token.
func GetAreaGroupRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAreaGroupRelationshipTypesByTokenResponse, error) {
	return getAreaGroupRelationshipTypesByToken(ctx, client, tokens)
}

// List area group relationship types based on criteria.
func ListAreaGroupRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAreaGroupRelationshipTypesResponse, error) {
	return listAreaGroupRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that an area group relationship exists.
func AssureAreaGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipCreateRequest,
) (*getAreaGroupRelationshipsByTokenResponse, *createAreaGroupRelationshipResponse, error) {
	gresp, err := GetAreaGroupRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AreaGroupRelationshipsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAreaGroupRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new area group relationship.
func CreateAreaGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipCreateRequest,
) (*createAreaGroupRelationshipResponse, error) {
	return createAreaGroupRelationship(ctx, client, request.Token, request.AreaGroup, request.Area, request.RelationshipType)
}

// Get area group relationships by token.
func GetAreaGroupRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAreaGroupRelationshipsByTokenResponse, error) {
	return getAreaGroupRelationshipsByToken(ctx, client, tokens)
}

// List area group relationships based on criteria.
func ListAreaGroupRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAreaGroupRelationshipsResponse, error) {
	return listAreaGroupRelationships(ctx, client, pageNumber, pageSize)
}
