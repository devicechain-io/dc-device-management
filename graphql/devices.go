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

var GqlDeviceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "DeviceType",
		Fields: graphql.Fields{
			rdb.FIELD_ID: &graphql.Field{
				Type:        graphql.Int,
				Description: "Unique database id",
			},
			rdb.FIELD_CREATED_AT: &graphql.Field{
				Type:        graphql.DateTime,
				Description: "Creation timestamp",
			},
			rdb.FIELD_UPDATED_AT: &graphql.Field{
				Type:        graphql.DateTime,
				Description: "Update timestamp",
			},
			rdb.FIELD_DELETED_AT: &graphql.Field{
				Type:        graphql.DateTime,
				Description: "Deletion timestamp",
			},
			rdb.FIELD_TOKEN: &graphql.Field{
				Type:        graphql.String,
				Description: "Unique token",
			},
			rdb.FIELD_NAME: &graphql.Field{
				Type:        graphql.String,
				Description: "Short name",
			},
			rdb.FIELD_DESCRIPTION: &graphql.Field{
				Type:        graphql.String,
				Description: "Longer description",
			},
			rdb.FIELD_IMAGE_URL: &graphql.Field{
				Type:        graphql.String,
				Description: "URL for branding image",
			},
			rdb.FIELD_ICON: &graphql.Field{
				Type:        graphql.String,
				Description: "Icon indicator for branding",
			},
			rdb.FIELD_BACKGROUND_COLOR: &graphql.Field{
				Type:        graphql.String,
				Description: "Background color for branding",
			},
			rdb.FIELD_FOREGROUND_COLOR: &graphql.Field{
				Type:        graphql.String,
				Description: "Foreground color for branding",
			},
			rdb.FIELD_BORDER_COLOR: &graphql.Field{
				Type:        graphql.String,
				Description: "Border color for branding",
			},
		},
	},
)

var GqlDevice = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.Fields{
			rdb.FIELD_ID: &graphql.Field{
				Type: graphql.Int,
			},
			rdb.FIELD_CREATED_AT: &graphql.Field{
				Type: graphql.DateTime,
			},
			rdb.FIELD_UPDATED_AT: &graphql.Field{
				Type: graphql.DateTime,
			},
			rdb.FIELD_DELETED_AT: &graphql.Field{
				Type: graphql.DateTime,
			},
			rdb.FIELD_TOKEN: &graphql.Field{
				Type: graphql.String,
			},
			rdb.FIELD_NAME: &graphql.Field{
				Type: graphql.String,
			},
			rdb.FIELD_DESCRIPTION: &graphql.Field{
				Type: graphql.String,
			},
			"device_type": &graphql.Field{
				Type: GqlDeviceType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return func() (interface{}, error) {
						return nil, nil
					}, nil
				},
			},
		},
	},
)
