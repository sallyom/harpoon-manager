/*
Copyright 2022.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make" to regenerate code after modifying this file

// HarpoonSpec defines the desired state of Harpoon
type HarpoonSpec struct {
	// Description of package
	Description string `json:"description"`
	// Image to run or extract artifacts from
	Image *Image `json:"image,omitempty"`
	// GitInfo for this package
	Git *GitInfo `json:"git,omitempty"`
	// Target information for machines to interact with this package
	Targets []TargetDevice `json:"targets,omitempty"`
}

// Image defines this package's image
type Image struct {
	// Image is the image that contains the package artifacts
	Name string `json:"name"`
	// Mapping define artifacts to be extracted
	Mapping []ImageArtifactMapping `json:"mapping,omitempty"`
	// EntryCommand displays an entrypoint command for a package
	EntryCommand string `json:"entryCommand,omitempty"`
	// UnitFile is path within the image to systemd unit file for running package
	UnitFile string `json:"unitFile,omitempty"`
}

// GitInfo describes information required for GitOps
type GitInfo struct {
	URL          string `json:"url"`
	Directory    string `json:"directory"`
	SubDirectory string `json:"subDirectory"`
	Branch       string `json:"branch"`
	Method       string `json:"method"`
	Schedule     string `json:"schedule"`
}

// TargetDevice is information for a machine that will interact with this package
type TargetDevice struct {
	// If managing a kube-native device, path to kubeconfig
	Kubeconfig string `json:"kubeconfig,omitempty"`
	// If managing a kube-native device, kubernetes context from kubeconfig to use
	KubeContext string `json:"kubeContext,omitempty"`
	// IPAddress of target machine to connect to
	IPAddress string `json:"ipAddress,omitempty"`
	// If using SSH, path to private SSH key
	SSHKey string `json:"sshKey,omitempty"`
	// Map of artifact files from image and location they should be extracted to in target
	InstallLocation map[string]string `json:"installLocation,omitempty"`
}

// HarpoonStatus defines the observed state of Harpoon
type HarpoonStatus struct {
	Version string `json:"version,omitempty"`
}

// ImageArtifactMapping holds mapping information from images for extracting artifacts
type ImageArtifactMapping struct {
	// OS is GOOS of an artifact included in an image
	OS string `json:"os,omitempty"`
	// Arch is GOARCH of an artifact included in an image
	Arch string `json:"arch,omitempty"`
	// Location is the file location of an artifact within an image
	Location string `json:"from,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Harpoon is the Schema for the harpoons API
type Harpoon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HarpoonSpec   `json:"spec,omitempty"`
	Status HarpoonStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HarpoonList contains a list of Harpoon
type HarpoonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Harpoon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Harpoon{}, &HarpoonList{})
}
