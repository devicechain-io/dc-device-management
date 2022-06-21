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

// Assure that an asset type exists.
func AssureAssetType(
	ctx context.Context,
	client graphql.Client,
	request model.AssetTypeCreateRequest,
) (*getAssetTypesByTokenResponse, *createAssetTypeResponse, error) {
	gresp, err := GetAssetTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AssetTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAssetType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new asset type.
func CreateAssetType(
	ctx context.Context,
	client graphql.Client,
	request model.AssetTypeCreateRequest,
) (*createAssetTypeResponse, error) {
	return createAssetType(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get asset types by token.
func GetAssetTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAssetTypesByTokenResponse, error) {
	return getAssetTypesByToken(ctx, client, tokens)
}

// List asset types based on criteria.
func ListAssetTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAssetTypesResponse, error) {
	return listAssetTypes(ctx, client, pageNumber, pageSize)
}

// Assure that an asset exists.
func AssureAsset(
	ctx context.Context,
	client graphql.Client,
	request model.AssetCreateRequest,
) (*getAssetsByTokenResponse, *createAssetResponse, error) {
	gresp, err := GetAssetsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AssetsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAsset(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new asset.
func CreateAsset(
	ctx context.Context,
	client graphql.Client,
	request model.AssetCreateRequest,
) (*createAssetResponse, error) {
	return createAsset(ctx, client, request.Token, request.AssetTypeToken,
		blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get assets by token.
func GetAssetsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAssetsByTokenResponse, error) {
	return getAssetsByToken(ctx, client, tokens)
}

// List assets based on criteria.
func ListAssets(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAssetsResponse, error) {
	return listAssets(ctx, client, pageNumber, pageSize)
}

// Assure that an asset relationship type exists.
func AssureAssetRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AssetRelationshipTypeCreateRequest,
) (*getAssetRelationshipTypesByTokenResponse, *createAssetRelationshipTypeResponse, error) {
	gresp, err := GetAssetRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AssetRelationshipTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAssetRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new asset relationship type.
func CreateAssetRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AssetRelationshipTypeCreateRequest,
) (*createAssetRelationshipTypeResponse, error) {
	return createAssetRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get asset relationship types by token.
func GetAssetRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAssetRelationshipTypesByTokenResponse, error) {
	return getAssetRelationshipTypesByToken(ctx, client, tokens)
}

// List asset relationship types based on criteria.
func ListAssetRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAssetRelationshipTypesResponse, error) {
	return listAssetRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that an asset relationship exists.
func AssureAssetRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AssetRelationshipCreateRequest,
) (*getAssetRelationshipsByTokenResponse, *createAssetRelationshipResponse, error) {
	gresp, err := GetAssetRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AssetRelationshipsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAssetRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new asset relationship.
func CreateAssetRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AssetRelationshipCreateRequest,
) (*createAssetRelationshipResponse, error) {
	return createAssetRelationship(ctx, client, request.Token, request.SourceAsset, request.TargetAsset, request.RelationshipType)
}

// Get asset relationships by token.
func GetAssetRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAssetRelationshipsByTokenResponse, error) {
	return getAssetRelationshipsByToken(ctx, client, tokens)
}

// List asset relationships based on criteria.
func ListAssetRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAssetRelationshipsResponse, error) {
	return listAssetRelationships(ctx, client, pageNumber, pageSize)
}

// Assure that an asset group exists.
func AssureAssetGroup(
	ctx context.Context,
	client graphql.Client,
	request model.AssetGroupCreateRequest,
) (*getAssetGroupsByTokenResponse, *createAssetGroupResponse, error) {
	gresp, err := GetAssetGroupsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AssetGroupsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAssetGroup(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new asset group.
func CreateAssetGroup(
	ctx context.Context,
	client graphql.Client,
	request model.AssetGroupCreateRequest,
) (*createAssetGroupResponse, error) {
	return createAssetGroup(ctx, client, request.Token, blank(request.Name), blank(request.Description),
		blank(request.ImageUrl), blank(request.Icon), blank(request.BackgroundColor), blank(request.ForegroundColor),
		blank(request.BorderColor), blank(request.Metadata))
}

// Get asset groups by token.
func GetAssetGroupsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAssetGroupsByTokenResponse, error) {
	return getAssetGroupsByToken(ctx, client, tokens)
}

// List asset groups based on criteria.
func ListAssetGroups(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAssetGroupsResponse, error) {
	return listAssetGroups(ctx, client, pageNumber, pageSize)
}

// Assure that an asset group relationship type exists.
func AssureAssetGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AssetGroupRelationshipTypeCreateRequest,
) (*getAssetGroupRelationshipTypesByTokenResponse, *createAssetGroupRelationshipTypeResponse, error) {
	gresp, err := GetAssetGroupRelationshipTypesByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AssetGroupRelationshipTypesByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAssetGroupRelationshipType(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new asset group relationship type.
func CreateAssetGroupRelationshipType(
	ctx context.Context,
	client graphql.Client,
	request model.AssetGroupRelationshipTypeCreateRequest,
) (*createAssetGroupRelationshipTypeResponse, error) {
	return createAssetGroupRelationshipType(ctx, client, request.Token, blank(request.Name), blank(request.Description), blank(request.Metadata))
}

// Get asset group relationship types by token.
func GetAssetGroupRelationshipTypesByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAssetGroupRelationshipTypesByTokenResponse, error) {
	return getAssetGroupRelationshipTypesByToken(ctx, client, tokens)
}

// List asset group relationship types based on criteria.
func ListAssetGroupRelationshipTypes(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAssetGroupRelationshipTypesResponse, error) {
	return listAssetGroupRelationshipTypes(ctx, client, pageNumber, pageSize)
}

// Assure that an asset group relationship exists.
func AssureAssetGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AssetGroupRelationshipCreateRequest,
) (*getAssetGroupRelationshipsByTokenResponse, *createAssetGroupRelationshipResponse, error) {
	gresp, err := GetAssetGroupRelationshipsByToken(ctx, client, []string{request.Token})
	if err == nil && len(gresp.AssetGroupRelationshipsByToken) > 0 {
		return gresp, nil, nil
	}
	cresp, err := CreateAssetGroupRelationship(ctx, client, request)
	if err != nil {
		return nil, nil, err
	}
	return nil, cresp, nil
}

// Create a new asset group relationship.
func CreateAssetGroupRelationship(
	ctx context.Context,
	client graphql.Client,
	request model.AssetGroupRelationshipCreateRequest,
) (*createAssetGroupRelationshipResponse, error) {
	return createAssetGroupRelationship(ctx, client, request.Token, request.AssetGroup, request.Asset, request.RelationshipType)
}

// Get asset group relationships by token.
func GetAssetGroupRelationshipsByToken(
	ctx context.Context,
	client graphql.Client,
	tokens []string,
) (*getAssetGroupRelationshipsByTokenResponse, error) {
	return getAssetGroupRelationshipsByToken(ctx, client, tokens)
}

// List asset group relationships based on criteria.
func ListAssetGroupRelationships(
	ctx context.Context,
	client graphql.Client,
	pageNumber int,
	pageSize int,
) (*listAssetGroupRelationshipsResponse, error) {
	return listAssetGroupRelationships(ctx, client, pageNumber, pageSize)
}
