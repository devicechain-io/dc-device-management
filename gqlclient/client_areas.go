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

// Assure that a area type exists.
func AssureAreaType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaTypeCreateRequest,
) (IAreaType, bool, error) {
	gresp, err := GetAreaTypesByToken(ctx, client, []string{request.Token})
	if err != nil {
		return nil, false, err
	}
	if gresp[request.Token] != nil {
		return gresp[request.Token], false, nil
	}
	cresp, err := CreateAreaType(ctx, client, request)
	if err != nil {
		return nil, false, err
	}
	return cresp, true, nil
}

// Create a new area type.
func CreateAreaType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaTypeCreateRequest,
) (IAreaType, error) {
	cresp, err := createAreaType(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
	if err != nil {
		return nil, err
	}
	return &cresp.CreateAreaType, nil
}

// Get area types by token.
func GetAreaTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (map[string]IAreaType, error) {
	gresp, err := getAreaTypesByToken(ctx, client, tokens)
	if err != nil {
		return nil, err
	}
	itypes := make(map[string]IAreaType)
	if gresp != nil {
		for _, res := range gresp.AreaTypesByToken {
			itypes[res.Token] = IAreaType(&res)
		}
	}
	return itypes, nil
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

// Assure that a area exists.
func AssureArea(
	ctx context.Context,
	client graphql.Client,
	request model.AreaCreateRequest,
) (IArea, bool, error) {
	gresp, err := GetAreasByToken(ctx, client, []string{request.Token})
	if err != nil {
		return nil, false, err
	}
	if gresp[request.Token] != nil {
		return gresp[request.Token], false, nil
	}
	cresp, err := CreateArea(ctx, client, request)
	if err != nil {
		return nil, false, err
	}
	return cresp, true, nil
}

// Create a new area.
func CreateArea(
	ctx context.Context,
	client graphql.Client,
	request model.AreaCreateRequest,
) (IArea, error) {
	cresp, err := createArea(ctx, client, request.Token, request.AreaTypeToken,
		blank(request.Name), blank(request.Description), blank(request.Metadata))
	if err != nil {
		return nil, err
	}
	return &cresp.CreateArea, nil
}

// Get areas by token.
func GetAreasByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (map[string]IArea, error) {
	gresp, err := getAreasByToken(ctx, client, tokens)
	if err != nil {
		return nil, err
	}
	itypes := make(map[string]IArea)
	if gresp != nil {
		for _, res := range gresp.AreasByToken {
			itypes[res.Token] = IArea(&res)
		}
	}
	return itypes, nil
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

// Assure that a area relationship type exists.
func AssureAreaRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipTypeCreateRequest,
) (IAreaRelationshipType, bool, error) {
	gresp, err := GetAreaRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err != nil {
		return nil, false, err
	}
	if gresp[request.Token] != nil {
		return gresp[request.Token], false, nil
	}
	cresp, err := CreateAreaRelationshipType(ctx, client, request)
	if err != nil {
		return nil, false, err
	}
	return cresp, true, nil
}

// Create a new area relationship type.
func CreateAreaRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipTypeCreateRequest,
) (IAreaRelationshipType, error) {
	cresp, err := createAreaRelationshipType(ctx, client, request.Token, blank(request.Name),
		blank(request.Description), blank(request.Metadata))
	if err != nil {
		return nil, err
	}
	return &cresp.CreateAreaRelationshipType, nil
}

// Get area relationship types by token.
func GetAreaRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (map[string]IAreaRelationshipType, error) {
	gresp, err := getAreaRelationshipTypesByToken(ctx, client, tokens)
	if err != nil {
		return nil, err
	}
	itypes := make(map[string]IAreaRelationshipType)
	if gresp != nil {
		for _, res := range gresp.AreaRelationshipTypesByToken {
			itypes[res.Token] = IAreaRelationshipType(&res)
		}
	}
	return itypes, nil
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

// Assure that a area relationship exists.
func AssureAreaRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipCreateRequest,
) (IAreaRelationship, bool, error) {
	gresp, err := GetAreaRelationshipsByToken(ctx, client, []string{request.Token})
	if err != nil {
		return nil, false, err
	}
	if gresp[request.Token] != nil {
		return gresp[request.Token], false, nil
	}
	cresp, err := CreateAreaRelationship(ctx, client, request)
	if err != nil {
		return nil, false, err
	}
	return cresp, true, nil
}

// Create a new area relationship.
func CreateAreaRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaRelationshipCreateRequest,
) (IAreaRelationship, error) {
	cresp, err := createAreaRelationship(ctx, client, request.Token, request.SourceArea,
		request.TargetArea, request.RelationshipType)
	if err != nil {
		return nil, err
	}
	return &cresp.CreateAreaRelationship, nil
}

// Get area relationships by token.
func GetAreaRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (map[string]IAreaRelationship, error) {
	gresp, err := getAreaRelationshipsByToken(ctx, client, tokens)
	if err != nil {
		return nil, err
	}
	itypes := make(map[string]IAreaRelationship)
	if gresp != nil {
		for _, res := range gresp.AreaRelationshipsByToken {
			itypes[res.Token] = IAreaRelationship(&res)
		}
	}
	return itypes, nil
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

// Assure that a area group exists.
func AssureAreaGroup(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupCreateRequest,
) (IAreaGroup, bool, error) {
	gresp, err := GetAreaGroupsByToken(ctx, client, []string{request.Token})
	if err != nil {
		return nil, false, err
	}
	if gresp[request.Token] != nil {
		return gresp[request.Token], false, nil
	}
	cresp, err := CreateAreaGroup(ctx, client, request)
	if err != nil {
		return nil, false, err
	}
	return cresp, true, nil
}

// Create a new area group.
func CreateAreaGroup(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupCreateRequest,
) (IAreaGroup, error) {
	cresp, err := createAreaGroup(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
	if err != nil {
		return nil, err
	}
	return &cresp.CreateAreaGroup, nil
}

// Get area groups by token.
func GetAreaGroupsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (map[string]IAreaGroup, error) {
	gresp, err := getAreaGroupsByToken(ctx, client, tokens)
	if err != nil {
		return nil, err
	}
	itypes := make(map[string]IAreaGroup)
	if gresp != nil {
		for _, res := range gresp.AreaGroupsByToken {
			itypes[res.Token] = IAreaGroup(&res)
		}
	}
	return itypes, nil
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

// Assure that a area group relationship type exists.
func AssureAreaGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipTypeCreateRequest,
) (IAreaGroupRelationshipType, bool, error) {
	gresp, err := GetAreaGroupRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err != nil {
		return nil, false, err
	}
	if gresp[request.Token] != nil {
		return gresp[request.Token], false, nil
	}
	cresp, err := CreateAreaGroupRelationshipType(ctx, client, request)
	if err != nil {
		return nil, false, err
	}
	return cresp, true, nil
}

// Create a new area group relationship type.
func CreateAreaGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipTypeCreateRequest,
) (IAreaGroupRelationshipType, error) {
	cresp, err := createAreaGroupRelationshipType(ctx, client, request.Token, blank(request.Name),
		blank(request.Description), blank(request.Metadata))
	if err != nil {
		return nil, err
	}
	return &cresp.CreateAreaGroupRelationshipType, nil
}

// Get a area group relationship types by token.
func GetAreaGroupRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (map[string]IAreaGroupRelationshipType, error) {
	gresp, err := getAreaGroupRelationshipTypesByToken(ctx, client, tokens)
	if err != nil {
		return nil, err
	}
	itypes := make(map[string]IAreaGroupRelationshipType)
	if gresp != nil {
		for _, res := range gresp.AreaGroupRelationshipTypesByToken {
			itypes[res.Token] = IAreaGroupRelationshipType(&res)
		}
	}
	return itypes, nil
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

// Assure that a area group relationship exists.
func AssureAreaGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipCreateRequest,
) (IAreaGroupRelationship, bool, error) {
	gresp, err := GetAreaGroupRelationshipsByToken(ctx, client, []string{request.Token})
	if err != nil {
		return nil, false, err
	}
	if gresp[request.Token] != nil {
		return gresp[request.Token], false, nil
	}
	cresp, err := CreateAreaGroupRelationship(ctx, client, request)
	if err != nil {
		return nil, false, err
	}
	return cresp, true, nil
}

// Create a new area group relationship.
func CreateAreaGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AreaGroupRelationshipCreateRequest,
) (IAreaGroupRelationship, error) {
	cresp, err := createAreaGroupRelationship(ctx, client, request.Token, request.AreaGroup, request.Area,
		request.RelationshipType)
	if err != nil {
		return nil, err
	}
	return &cresp.CreateAreaGroupRelationship, nil
}

// Get a area group relationships by token.
func GetAreaGroupRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (map[string]IAreaGroupRelationship, error) {
	gresp, err := getAreaGroupRelationshipsByToken(ctx, client, tokens)
	if err != nil {
		return nil, err
	}
	itypes := make(map[string]IAreaGroupRelationship)
	if gresp != nil {
		for _, res := range gresp.AreaGroupRelationshipsByToken {
			itypes[res.Token] = IAreaGroupRelationship(&res)
		}
	}
	return itypes, nil
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
