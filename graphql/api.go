/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"

	"github.com/devicechain-io/dc-device-management/model"
	gqlcore "github.com/devicechain-io/dc-microservice/graphql"
	"github.com/devicechain-io/dc-microservice/rdb"
)

type Api struct {
	Context  context.Context
	Resolver SchemaResolver
}

// Create a new API instance.
func NewApi(rdb *rdb.RdbManager) Api {
	api := Api{}
	api.Context = context.WithValue(context.Background(), gqlcore.ContextRdbKey, rdb)
	api.Resolver = SchemaResolver{}
	return api
}

// Get device type by token.
func (api *Api) DeviceTypeByToken(token string) (*model.DeviceType, error) {
	rez, err := api.Resolver.DeviceTypeByToken(api.Context, struct{ Token string }{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return &rez.M, nil
}
