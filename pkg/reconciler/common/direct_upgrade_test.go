/*
Copyright 2022 The Knative Authors

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

package common

import (
	"os"
	"testing"
)

func Test_canUpgradeDirectly(t *testing.T) {

	t.Log("use default versions")
	envInit()
	if !canUpgradeDirectly("v0.24", "v1.2") {
		t.Error("canUpgradeDirectly should return true")
	}
	if !canUpgradeDirectly("v1.2", "v0.24") {
		t.Error("canUpgradeDirectly should return true")
	}

	t.Log("use specify versions")
	os.Setenv("DIRECT_UPGRADE_VERSIONS", "v0.1-v0.2")
	envInit()
	if canUpgradeDirectly("v0.24", "v1.2") {
		t.Error("canUpgradeDirectly should return false")
	}
	if !canUpgradeDirectly("v0.2", "v0.1") {
		t.Error("canUpgradeDirectly should return true")
	}
	if !canUpgradeDirectly("v0.1", "v0.2") {
		t.Error("canUpgradeDirectly should return true")
	}

	t.Log("use multiple versions")
	os.Setenv("DIRECT_UPGRADE_VERSIONS", "v0.1-v0.2|v1.0-v2.0")
	envInit()
	if canUpgradeDirectly("v0.24", "v1.2") {
		t.Error("canUpgradeDirectly should return false")
	}
	if !canUpgradeDirectly("v0.2", "v0.1") {
		t.Error("canUpgradeDirectly should return true")
	}
	if !canUpgradeDirectly("v1.0", "v2.0") {
		t.Error("canUpgradeDirectly should return true")
	}

	t.Log("use invalid format version")
	os.Setenv("DIRECT_UPGRADE_VERSIONS", "v0.1-v0.2-v0.3")
	envInit()
	if canUpgradeDirectly("v0.24", "v1.2") {
		t.Error("canUpgradeDirectly should return false")
	}
}
