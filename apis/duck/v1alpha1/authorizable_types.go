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

import "knative.dev/pkg/tracker"

// +genduck

// Authorizable provides way for an obect to advertise that
// accepts authoirization policy objects (that conform to
// AuthorizationPolicy).
//
// Object that implement Authorizable agree to confirm to the
// data plane contract.
type Authorizable struct {
	Spec AuthorizableSpec `json:"spec"`
}

// AuthorizableSpec describes what must be present in the spec
// of objects sthat implement Authoirzable
type AuthorizableSpec struct {
	AuthPolicy tracker.Reference `json:"auth_policy"`
}
