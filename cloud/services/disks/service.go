/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package disks

import (
	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/compute/mgmt/compute"
	"github.com/Azure/go-autorest/autorest"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	"sigs.k8s.io/cluster-api-provider-azure/cloud/scope"
)

var _ azure.Service = (*Service)(nil)

// Service provides operations on resource groups
type Service struct {
	Client compute.DisksClient
	Scope  *scope.ClusterScope
}

// getGroupsClient creates a new groups client from subscriptionid.
func getDisksClient(subscriptionID string, authorizer autorest.Authorizer) compute.DisksClient {
	disksClient := compute.NewDisksClient(subscriptionID)
	disksClient.Authorizer = authorizer
	disksClient.AddToUserAgent(azure.UserAgent)
	return disksClient
}

// NewService creates a new groups service.
func NewService(scope *scope.ClusterScope) *Service {
	return &Service{
		Client: getDisksClient(scope.SubscriptionID, scope.Authorizer),
		Scope:  scope,
	}
}
