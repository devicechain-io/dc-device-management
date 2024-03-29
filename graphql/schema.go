/**
 * Copyright © 2022 DeviceChain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
