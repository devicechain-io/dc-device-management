/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package gqlclient

// Customer type entity.
type ICustomerType interface {
	IModel
	ITokenReference
	INamedEntity
	IBrandedEntity
	IMetadataEntity
}

// Customer entity.
type ICustomer interface {
	IModel
	ITokenReference
	INamedEntity
	IMetadataEntity
	GetCustomerType() DefaultCustomerCustomerType
}

// Customer relationship type entity.
type ICustomerRelationshipType interface {
	IModel
	ITokenReference
	INamedEntity
	IMetadataEntity
}

// Customer relationship entity.
type ICustomerRelationship interface {
	IModel
	ITokenReference
	IMetadataEntity
	GetSourceCustomer() DefaultCustomerRelationshipSourceCustomer
	GetTargetCustomer() DefaultCustomerRelationshipTargetCustomer
	GetRelationshipType() DefaultCustomerRelationshipRelationshipTypeCustomerRelationshipType
}

// Customer group entity.
type ICustomerGroup interface {
	IModel
	ITokenReference
	INamedEntity
	IBrandedEntity
	IMetadataEntity
}

// Customer group relationship type entity.
type ICustomerGroupRelationshipType interface {
	IModel
	ITokenReference
	INamedEntity
	IMetadataEntity
}

// Customer group relationship entity.
type ICustomerGroupRelationship interface {
	IModel
	ITokenReference
	IMetadataEntity
	GetCustomerGroup() DefaultCustomerGroupRelationshipCustomerGroup
	GetCustomer() DefaultCustomerGroupRelationshipCustomer
	GetRelationshipType() DefaultCustomerGroupRelationshipRelationshipTypeCustomerGroupRelationshipType
}
