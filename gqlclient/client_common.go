/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

// The gqlclient package provides a GraphQL client for commonly used operations
// in the device management API.
package gqlclient

//go:generate go run github.com/Khan/genqlient@v0.4.0

// TEMP: genqlient doesn't handle pointers yet.
func blank(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}
