/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// AuthorizationPolicy defines the control plane contract
// as a duck type for implementors of authorization policies.
type AuthorizationPolicy struct {
	Status AuthPolicyStatus `json:"status"`
}

// AuthPolicyStatus represents the types required to be in
// the status for an authorization policy
type AuthPolicyStatus struct {
	SendBody bool `json:"send_body,omitempty"`

	Address *duckv1.Addressable `json:"address,omitempty"`

	// This may need to be restricted to a single container
	// and config map.
	// TODO: get this shape right
	Template duckv1.PodSpecable `json:"template,omitempty"`
}
