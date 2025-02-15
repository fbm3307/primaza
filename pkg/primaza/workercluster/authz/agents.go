/*
Copyright 2023 The Primaza Authors.

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

package authz

import (
	"github.com/primaza/primaza/pkg/authz"
)

func GetAgentAppRequiredPermissions() []authz.ResourcePermissions {
	return []authz.ResourcePermissions{
		{
			Verbs:    []string{"create"},
			Version:  "",
			Group:    "apps",
			Resource: "deployments",
		},
		{
			Verbs:    []string{"delete"},
			Version:  "",
			Group:    "apps",
			Resource: "deployments",
			Name:     "primaza-app-agent",
		},
	}
}

func GetAgentSvcRequiredPermissions() []authz.ResourcePermissions {
	return []authz.ResourcePermissions{
		{
			Verbs:    []string{"create"},
			Version:  "",
			Group:    "apps",
			Resource: "deployments",
		},
		{
			Verbs:    []string{"delete"},
			Version:  "",
			Group:    "apps",
			Resource: "deployments",
			Name:     "primaza-svc-agent",
		},
	}
}
