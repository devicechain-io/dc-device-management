/**
 * Copyright ©2022 AssetChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

import (
	"context"

	"gorm.io/gorm"
)

// Resolves tokens from entity relationship create request into object references.
func (api *Api) resolveRelationshipTargets(ctx context.Context, req EntityRelationshipCreateRequest, rel *EntityRelationship) error {
	if req.TargetDevice != nil {
		matches, err := api.DevicesByToken(ctx, []string{*req.TargetDevice})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetDevice = matches[0]
	}
	if req.TargetDeviceGroup != nil {
		matches, err := api.DeviceGroupsByToken(ctx, []string{*req.TargetDeviceGroup})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetDeviceGroup = matches[0]
	}
	if req.TargetAsset != nil {
		matches, err := api.AssetsByToken(ctx, []string{*req.TargetAsset})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetAsset = matches[0]
	}
	if req.TargetAssetGroup != nil {
		matches, err := api.AssetGroupsByToken(ctx, []string{*req.TargetAssetGroup})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetAssetGroup = matches[0]
	}
	if req.TargetArea != nil {
		matches, err := api.AreasByToken(ctx, []string{*req.TargetArea})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetArea = matches[0]
	}
	if req.TargetAreaGroup != nil {
		matches, err := api.AreaGroupsByToken(ctx, []string{*req.TargetAreaGroup})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetAreaGroup = matches[0]
	}
	if req.TargetCustomer != nil {
		matches, err := api.CustomersByToken(ctx, []string{*req.TargetCustomer})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetCustomer = matches[0]
	}
	if req.TargetCustomerGroup != nil {
		matches, err := api.CustomerGroupsByToken(ctx, []string{*req.TargetCustomerGroup})
		if err != nil {
			return err
		}
		if len(matches) == 0 {
			return gorm.ErrRecordNotFound
		}
		rel.TargetCustomerGroup = matches[0]
	}
	return nil
}

func preloadRelationshipTargets(db *gorm.DB) *gorm.DB {
	db = db.Preload("TargetDevice").Preload("TargetDeviceGroup").Preload("TargetAsset").Preload("TargetAssetGroup")
	db = db.Preload("TargetArea").Preload("TargetAreaGroup").Preload("TargetCustomer").Preload("TargetCustomerGroup")
	return db
}
