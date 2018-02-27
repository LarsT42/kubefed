/*
Copyright 2018 The Kubernetes Authors.

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

package integration

import (
	"fmt"
	"testing"

	"github.com/pborman/uuid"

	"github.com/marun/fnord/pkg/federatedtypes"
	"github.com/marun/fnord/test/common"
	"github.com/marun/fnord/test/integration/framework"
	pkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// TestCrud validates create/read/update/delete operations for federated types.
func TestCrud(t *testing.T) {
	tl := framework.NewIntegrationLogger(t)
	fedFixture := framework.SetUpFederationFixture(tl, 2)
	defer fedFixture.TearDown(tl)

	fedTypeConfigs := federatedtypes.FederatedTypeConfigs()
	for kind, fedTypeConfig := range fedTypeConfigs {
		t.Run(kind, func(t *testing.T) {
			tl := framework.NewIntegrationLogger(t)
			fixture, crudTester, obj, _ := initCrudTest(tl, fedFixture, fedTypeConfig.AdapterFactory, kind)
			defer fixture.TearDown(tl)

			crudTester.CheckLifecycle(obj)
		})
	}
}

// initCrudTest initializes common elements of a crud test
func initCrudTest(tl common.TestLogger, fedFixture *framework.FederationFixture, adapterFactory federatedtypes.AdapterFactory, kind string) (
	*framework.ControllerFixture, *common.FederatedTypeCrudTester, pkgruntime.Object, federatedtypes.FederatedTypeAdapter) {
	// TODO(marun) stop requiring user agent when creating new config or clients
	userAgent := fmt.Sprintf("crud-test-%s", kind)
	fedConfig := fedFixture.FedApi.NewConfig(tl, userAgent)
	kubeConfig := fedFixture.KubeApi.NewConfig(tl, userAgent)
	crConfig := fedFixture.CrApi.NewConfig(tl, userAgent)
	fixture := framework.NewControllerFixture(tl, kind, adapterFactory, fedConfig, kubeConfig, crConfig)

	client := fedFixture.FedApi.NewClient(tl, userAgent)
	adapter := adapterFactory(client)

	clusterClients := fedFixture.ClusterClients(tl, userAgent)
	crudTester := framework.NewFederatedTypeCrudTester(tl, adapter, clusterClients)

	obj := federatedtypes.NewTestObject(kind, uuid.New())

	return fixture, crudTester, obj, adapter
}
