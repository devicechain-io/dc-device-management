/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"
	"strconv"

	"github.com/devicechain-io/dc-device-management/model"
	gqlcore "github.com/devicechain-io/dc-microservice/graphql"
	"github.com/devicechain-io/dc-microservice/rdb"
)

//go:embed schema.graphql
var SchemaContent string

type SchemaResolver struct{}

// Get rdb manager from context.
func (s *SchemaResolver) GetRdbManager(ctx context.Context) *rdb.RdbManager {
	return ctx.Value(gqlcore.ContextRdbKey).(*rdb.RdbManager)
}

// Get api from context.
func (s *SchemaResolver) GetApi(ctx context.Context) *model.Api {
	return ctx.Value(gqlcore.ContextApiKey).(*model.Api)
}

// Convert string ids to uint ids.
func (r *SchemaResolver) asUintIds(val []string) ([]uint, error) {
	ids := make([]uint, 0)
	for _, sid := range val {
		id, err := strconv.ParseUint(sid, 0, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, uint(id))
	}
	return ids, nil
}
