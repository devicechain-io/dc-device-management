/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"

	"github.com/devicechain-io/dc-devicemanagement/model"
)

func (r *SchemaResolver) CreateDeviceType(ctx context.Context, args struct {
	Value *model.DeviceTypeCreateRequest
}) (*DeviceTypeResolver, error) {
	dt := &DeviceTypeResolver{
		d: devicetype,
	}
	return dt, nil
}
