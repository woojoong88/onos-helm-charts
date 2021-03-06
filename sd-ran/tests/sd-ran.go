// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"github.com/onosproject/helmit/pkg/helm"
	"github.com/onosproject/helmit/pkg/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

// SDRANSuite is the sd-ran chart test suite
type SDRANSuite struct {
	test.Suite
}

// TestInstall tests installing the sd-ran chart
func (s *SDRANSuite) TestInstall(t *testing.T) {
	atomix := helm.Chart("atomix-controller", "https://charts.atomix.io").
		Release("sd-ran-atomix").
		Set("scope", "Namespace")
	assert.NoError(t, atomix.Install(true))

	raft := helm.Chart("raft-storage-controller", "https://charts.atomix.io").
		Release("sd-ran-raft").
		Set("scope", "Namespace")
	assert.NoError(t, raft.Install(true))

	cache := helm.Chart("cache-storage-controller", "https://charts.atomix.io").
		Release("sd-ran-cache").
		Set("scope", "Namespace")
	assert.NoError(t, cache.Install(true))

	sdran := helm.Chart("sd-ran").
		Release("sd-ran").
		Set("global.store.controller", "sd-ran-atomix-atomix-controller:5679").
		Set("import.onos-gui.enabled", false).
		Set("onos-ric.service.external.nodePort", 0).
		Set("onos-ric-ho.service.external.nodePort", 0).
		Set("onos-ric-mlb.service.external.nodePort", 0)
	assert.NoError(t, sdran.Install(true))
}
