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
		ids := []string{fmt.Sprintf("%d", *r.M.DeviceGroupId)}
		matches, err := r.S.DeviceGroupsById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
		ids := []string{fmt.Sprintf("%d", *r.M.AssetId)}
		matches, err := r.S.AssetsById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
		ids := []string{fmt.Sprintf("%d", *r.M.AssetGroupId)}
		matches, err := r.S.AssetGroupsById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
		ids := []string{fmt.Sprintf("%d", *r.M.CustomerId)}
		matches, err := r.S.CustomersById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
		ids := []string{fmt.Sprintf("%d", *r.M.CustomerGroupId)}
		matches, err := r.S.CustomerGroupsById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
		ids := []string{fmt.Sprintf("%d", *r.M.AreaId)}
		matches, err := r.S.AreasById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
		ids := []string{fmt.Sprintf("%d", *r.M.AreaGroupId)}
		matches, err := r.S.AreaGroupsById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
		ids := []string{fmt.Sprintf("%d", *r.M.DeviceAssignmentStatusId)}
		matches, err := r.S.DeviceAssignmentStatusesById(r.C, struct{ Ids []string }{Ids: ids})
		if err != nil {
			return nil
		}
		if len(matches) == 0 {
			return nil
		}
		return matches[0]
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
