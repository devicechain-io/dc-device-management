/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package model

type CachedApi struct {
	API *Api
}

// Create a new API instance.
func NewCachedApi(api *Api) *CachedApi {
	capi := &CachedApi{
		API: api,
	}
	return capi
}
