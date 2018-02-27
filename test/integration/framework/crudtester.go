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

package framework

import (
	"github.com/marun/fnord/pkg/federatedtypes"
	"github.com/marun/fnord/test/common"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
)

func NewFederatedTypeCrudTester(tl common.TestLogger, adapter federatedtypes.FederatedTypeAdapter, clusterClients []clientset.Interface) *common.FederatedTypeCrudTester {
	return common.NewFederatedTypeCrudTester(tl, adapter, clusterClients, DefaultWaitInterval, wait.ForeverTestTimeout)
}
