/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"github.com/devicechain-io/dc-microservice/rdb"
	"github.com/graphql-go/graphql"
)

// Creates query methods.
func NewQueryObject() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getDevice": &graphql.Field{
					Type: GqlDevice,
					Args: graphql.FieldConfigArgument{
						rdb.FIELD_TOKEN: &graphql.ArgumentConfig{
							Type:        graphql.NewNonNull(graphql.String),
							Description: "Unique token assigned to device",
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// Get device from RDB.
						return "world", nil
					},
				},
			},
		},
	)
}

// Creates mutation methods.
func NewMutationObject() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createDevice": &graphql.Field{
				Type:        GqlDevice,
				Description: "Create a new device",
				Args: graphql.FieldConfigArgument{
					rdb.FIELD_TOKEN: &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "Unique token assigned to device",
					},
					rdb.FIELD_NAME: &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Short name for device",
					},
					rdb.FIELD_DESCRIPTION: &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Longer description for device",
					},
					"device_type_token": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Float),
						Description: "Token indicating device type",
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					// Create device in RDB.
					return nil, nil
				},
			},
		},
	})
}

// Creates a new schema configuration for the microservice.
func NewSchemaConfig() graphql.SchemaConfig {
	return graphql.SchemaConfig{
		Query:    NewQueryObject(),
		Mutation: NewMutationObject(),
	}
}
