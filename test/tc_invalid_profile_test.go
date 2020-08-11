// +build e2e

/*
Copyright 2020 The Kubernetes Authors.

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

package e2e_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"sigs.k8s.io/seccomp-operator/internal/pkg/controllers/profile"
)

func (e *e2e) testCaseDeployInvalidProfile(nodes []string) {
	const (
		configMapName = "invalid-profile"
		profileName   = "profile-invalid.json"
	)
	invalidProfileContent := fmt.Sprintf(`
apiVersion: v1
kind: ConfigMap
metadata:
  name: %s
  annotations:
    seccomp.security.kubernetes.io/profile: "true"
data:
  %s: |-
    { "defaultAction": true }
`, configMapName, profileName)

	invalidProfile, err := ioutil.TempFile("", configMapName)
	e.Nil(err)

	_, err = invalidProfile.WriteString(invalidProfileContent)
	e.Nil(err)
	e.logf("Deploying invalid configmap")
	e.kubectl("create", "-f", invalidProfile.Name())
	defer func() {
		e.kubectl("delete", "-f", invalidProfile.Name())
		e.Nil(os.RemoveAll(invalidProfile.Name()))
	}()

	// Verify the event
	e.logf("Verifying events")
	eventsOutput := e.kubectl("get", "events")
	for _, s := range []string{
		"Warning",
		"cannot validate profile " + profileName,
		"configmap/" + configMapName,
		"decoding seccomp profile: json: cannot unmarshal bool into " +
			"Go struct field Seccomp.defaultAction of type seccomp.Action",
	} {
		e.Contains(eventsOutput, s)
	}

	// Check that the profile is not reconciled to the node
	e.logf("Verifying node content")
	configMap := e.getConfigMap(configMapName, "default")
	profilePath, err := profile.GetProfilePath(profileName, configMap)
	e.Nil(err)
	for _, node := range nodes {
		e.execNode(node, "bash", "-c", fmt.Sprintf("[ ! -f %s ]", profilePath))
	}
}
