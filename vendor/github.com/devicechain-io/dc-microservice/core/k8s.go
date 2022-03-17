/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

package core

import (
	"context"

	"github.com/devicechain-io/dc-k8s/api/v1beta1"
)

func (ms *Microservice) getTenantMicroservice() (*v1beta1.TenantMicroservice, error) {
	v1beta1.GetTenantMicroservice(v1beta1.TenantMicroserviceGetRequest{
		InstanceId: instanceId,
		MicroserviceId: ,
	})
}