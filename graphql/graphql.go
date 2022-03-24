/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import "github.com/graphql-go/graphql"

// Creates a new schema configuration for the microservice.
func NewSchemaConfig() graphql.SchemaConfig {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	return graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
}
