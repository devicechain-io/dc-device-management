/**
 * Copyright ©2022 DeviceChain - All Rights Reserved.
 * Unauthorized copying of this file, via any medium is strictly prohibited.
 * Proprietary and confidential.
 */

package graphql

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/devicechain-io/dc-device-management/model"
	util "github.com/devicechain-io/dc-microservice/graphql"
	gql "github.com/graph-gophers/graphql-go"
)

// ---------------------------------
// Device assignment status resolver
// ---------------------------------

type DeviceAssignmentStatusResolver struct {
	M model.DeviceAssignmentStatus
	S *SchemaResolver
	C context.Context
}

func (r *DeviceAssignmentStatusResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceAssignmentStatusResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceAssignmentStatusResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceAssignmentStatusResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceAssignmentStatusResolver) Token() string {
	return r.M.Token
}

func (r *DeviceAssignmentStatusResolver) Name() *string {
	return util.NullStr(r.M.Name)
}

func (r *DeviceAssignmentStatusResolver) Description() *string {
	return util.NullStr(r.M.Description)
}

func (r *DeviceAssignmentStatusResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

// ------------------------------------------------
// Device assignment status search results resolver
// ------------------------------------------------

type DeviceAssignmentStatusSearchResultsResolver struct {
	M model.DeviceAssignmentStatusSearchResults
	S *SchemaResolver
	C context.Context
}

func (r *DeviceAssignmentStatusSearchResultsResolver) Results() []*DeviceAssignmentStatusResolver {
	resolvers := make([]*DeviceAssignmentStatusResolver, 0)
	for _, current := range r.M.Results {
		resolvers = append(resolvers,
			&DeviceAssignmentStatusResolver{
				M: current,
				S: r.S,
				C: r.C,
			})
	}
	return resolvers
}

func (r *DeviceAssignmentStatusSearchResultsResolver) Pagination() *SearchResultsPaginationResolver {
	return &SearchResultsPaginationResolver{
		M: r.M.Pagination,
		S: r.S,
		C: r.C,
	}
}

// --------------------------
// Device assignment resolver
// --------------------------

type DeviceAssignmentResolver struct {
	M model.DeviceAssignment
	S *SchemaResolver
	C context.Context
}

func (r *DeviceAssignmentResolver) Id() gql.ID {
	return gql.ID(fmt.Sprint(r.M.ID))
}

func (r *DeviceAssignmentResolver) CreatedAt() *string {
	return util.FormatTime(r.M.CreatedAt)
}

func (r *DeviceAssignmentResolver) UpdatedAt() *string {
	return util.FormatTime(r.M.UpdatedAt)
}

func (r *DeviceAssignmentResolver) DeletedAt() *string {
	return util.FormatTime(r.M.DeletedAt.Time)
}

func (r *DeviceAssignmentResolver) Token() string {
	return r.M.Token
}

func (r *DeviceAssignmentResolver) Metadata() *string {
	return util.MetadataStr(r.M.Metadata)
}

func (r *DeviceAssignmentResolver) Active() bool {
	return r.M.Active
}

func (r *DeviceAssignmentResolver) Device() *DeviceResolver {
	return &DeviceResolver{
		M: r.M.Device,
		S: r.S,
		C: r.C,
	}
}

func (r *DeviceAssignmentResolver) DeviceGroup() *DeviceGroupResolver {
	if r.M.DeviceGroup != nil {
		return &DeviceGroupResolver{
			M: *r.M.DeviceGroup,
			S: r.S,
			C: r.C,
		}
	} else if r.M.DeviceGroupId != nil {
		rez, err := r.S.DeviceGroup(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.DeviceGroupId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

func (r *DeviceAssignmentResolver) Asset() *AssetResolver {
	if r.M.Asset != nil {
		return &AssetResolver{
			M: *r.M.Asset,
			S: r.S,
			C: r.C,
		}
	} else if r.M.AssetId != nil {
		rez, err := r.S.Asset(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.AssetId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

func (r *DeviceAssignmentResolver) AssetGroup() *AssetGroupResolver {
	if r.M.AssetGroup != nil {
		return &AssetGroupResolver{
			M: *r.M.AssetGroup,
			S: r.S,
			C: r.C,
		}
	} else if r.M.AssetGroupId != nil {
		rez, err := r.S.AssetGroup(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.AssetGroupId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

func (r *DeviceAssignmentResolver) Customer() *CustomerResolver {
	if r.M.Customer != nil {
		return &CustomerResolver{
			M: *r.M.Customer,
			S: r.S,
			C: r.C,
		}
	} else if r.M.CustomerId != nil {
		rez, err := r.S.Customer(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.CustomerId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

func (r *DeviceAssignmentResolver) CustomerGroup() *CustomerGroupResolver {
	if r.M.CustomerGroup != nil {
		return &CustomerGroupResolver{
			M: *r.M.CustomerGroup,
			S: r.S,
			C: r.C,
		}
	} else if r.M.CustomerGroupId != nil {
		rez, err := r.S.CustomerGroup(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.CustomerGroupId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

func (r *DeviceAssignmentResolver) Area() *AreaResolver {
	if r.M.Area != nil {
		return &AreaResolver{
			M: *r.M.Area,
			S: r.S,
			C: r.C,
		}
	} else if r.M.AreaId != nil {
		rez, err := r.S.Area(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.AreaId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

func (r *DeviceAssignmentResolver) AreaGroup() *AreaGroupResolver {
	if r.M.AreaGroup != nil {
		return &AreaGroupResolver{
			M: *r.M.AreaGroup,
			S: r.S,
			C: r.C,
		}
	} else if r.M.AreaGroupId != nil {
		rez, err := r.S.AreaGroup(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.AreaGroupId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

func (r *DeviceAssignmentResolver) DeviceAssignmentStatus() *DeviceAssignmentStatusResolver {
	if r.M.DeviceAssignmentStatus != nil {
		return &DeviceAssignmentStatusResolver{
			M: *r.M.DeviceAssignmentStatus,
			S: r.S,
			C: r.C,
		}
	} else if r.M.DeviceAssignmentStatusId != nil {
		rez, err := r.S.DeviceAssignmentStatus(r.C, struct{ Id string }{Id: fmt.Sprintf("%d", *r.M.DeviceAssignmentStatusId)})
		if err != nil {
			return nil
		}
		return rez
	}
	return nil
}

// -----------------------------------------
// Device assignment search results resolver
// -----------------------------------------

type DeviceAssignmentSearchResultsResolver struct {
	M model.DeviceAssignmentSearchResults
	S *SchemaResolver
	C context.Context
}

func (r *DeviceAssignmentSearchResultsResolver) Results() []*DeviceAssignmentResolver {
	resolvers := make([]*DeviceAssignmentResolver, 0)
	for _, current := range r.M.Results {
		resolvers = append(resolvers,
			&DeviceAssignmentResolver{
				M: current,
				S: r.S,
				C: r.C,
			})
	}
	return resolvers
}

func (r *DeviceAssignmentSearchResultsResolver) Pagination() *SearchResultsPaginationResolver {
	return &SearchResultsPaginationResolver{
		M: r.M.Pagination,
		S: r.S,
		C: r.C,
	}
}
