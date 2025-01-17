/*
Copyright 2024 Drewbernetes.

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

package flags

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

type BuildOptions struct {
	BaseOptions
	KubernetesClusterFlags
	S3Flags
	OpenStackFlags
	KubeVirtFlags

	Verbose                 bool
	BuildOS                 string
	ImagePrefix             string
	ImageRepo               string
	ImageRepoBranch         string
	ContainerdSHA256        string
	ContainerdVersion       string
	CrictlVersion           string
	CniVersion              string
	CniDebVersion           string
	KubeVersion             string
	KubeRpmVersion          string
	KubeDebVersion          string
	ExtraDebs               string
	AdditionalImages        []string
	AdditionalMetadata      map[string]string
	AddFalco                bool
	AddTrivy                bool
	AddGpuSupport           bool
	GpuVendor               string
	GpuModelSupport         string
	GpuInstanceSupport      string
	AMDVersion              string
	AMDDebVersion           string
	AMDUseCase              string
	NvidiaVersion           string
	NvidiaBucket            string
	NvidiaInstallerLocation string
	NvidiaTOKLocation       string
	NvidiaGriddFeatureType  int
}

func (o *BuildOptions) SetOptionsFromViper() {
	// General Flags
	o.Verbose = viper.GetBool(fmt.Sprintf("%s.verbose", viperBuildPrefix))
	o.BuildOS = viper.GetString(fmt.Sprintf("%s.build-os", viperBuildPrefix))
	o.ImagePrefix = viper.GetString(fmt.Sprintf("%s.image-prefix", viperBuildPrefix))
	o.ImageRepo = viper.GetString(fmt.Sprintf("%s.image-repo", viperBuildPrefix))
	o.ImageRepoBranch = viper.GetString(fmt.Sprintf("%s.image-repo-branch", viperBuildPrefix))
	o.ContainerdSHA256 = viper.GetString(fmt.Sprintf("%s.containerd-sha256", viperBuildPrefix))
	o.ContainerdVersion = viper.GetString(fmt.Sprintf("%s.containerd-version", viperBuildPrefix))
	o.CrictlVersion = viper.GetString(fmt.Sprintf("%s.crictl-version", viperBuildPrefix))
	o.CniVersion = viper.GetString(fmt.Sprintf("%s.cni-version", viperBuildPrefix))
	o.CniDebVersion = viper.GetString(fmt.Sprintf("%s.cni-deb-version", viperBuildPrefix))
	o.KubeVersion = viper.GetString(fmt.Sprintf("%s.kubernetes-version", viperBuildPrefix))
	o.KubeDebVersion = viper.GetString(fmt.Sprintf("%s.kubernetes-deb-version", viperBuildPrefix))
	o.KubeRpmVersion = viper.GetString(fmt.Sprintf("%s.kubernetes-rpm-version", viperBuildPrefix))
	o.ExtraDebs = viper.GetString(fmt.Sprintf("%s.extra-debs", viperBuildPrefix))
	o.AdditionalImages = viper.GetStringSlice(fmt.Sprintf("%s.additional-images", viperBuildPrefix))
	o.AdditionalMetadata = viper.GetStringMapString(fmt.Sprintf("%s.additional-metadata", viperBuildPrefix))
	o.AddFalco = viper.GetBool(fmt.Sprintf("%s.add-falco", viperBuildPrefix))
	o.AddTrivy = viper.GetBool(fmt.Sprintf("%s.add-trivy", viperBuildPrefix))

	// GPU
	o.AddGpuSupport = viper.GetBool(fmt.Sprintf("%s.enable-gpu-support", viperGpuPrefix))
	o.GpuVendor = viper.GetString(fmt.Sprintf("%s.gpu-vendor", viperGpuPrefix))
	o.GpuModelSupport = viper.GetString(fmt.Sprintf("%s.gpu-model-support", viperGpuPrefix))
	o.GpuInstanceSupport = viper.GetString(fmt.Sprintf("%s.gpu-instance-support", viperGpuPrefix))
	// AMD
	o.AMDVersion = viper.GetString(fmt.Sprintf("%s.amd-driver-version", viperGpuPrefix))
	o.AMDDebVersion = viper.GetString(fmt.Sprintf("%s.amd-deb-version", viperGpuPrefix))
	o.AMDUseCase = viper.GetString(fmt.Sprintf("%s.amd-usecase", viperGpuPrefix))
	// NVIDIA
	o.NvidiaVersion = viper.GetString(fmt.Sprintf("%s.nvidia-driver-version", viperGpuPrefix))
	o.NvidiaBucket = viper.GetString(fmt.Sprintf("%s.nvidia-bucket", viperGpuPrefix))
	o.NvidiaInstallerLocation = viper.GetString(fmt.Sprintf("%s.nvidia-installer-location", viperGpuPrefix))
	o.NvidiaTOKLocation = viper.GetString(fmt.Sprintf("%s.nvidia-tok-location", viperGpuPrefix))
	o.NvidiaGriddFeatureType = viper.GetInt(fmt.Sprintf("%s.nvidia-gridd-feature-type", viperGpuPrefix))

	o.BaseOptions.SetOptionsFromViper()
	o.KubernetesClusterFlags.SetOptionsFromViper()
	o.S3Flags.SetOptionsFromViper()
	o.OpenStackFlags.SetOptionsFromViper()
	o.KubeVirtFlags.SetOptionsFromViper()
}

func (o *BuildOptions) AddFlags(cmd *cobra.Command, imageBuilderRepo string) {
	// Build flags
	BoolVarWithViper(cmd, &o.Verbose, viperBuildPrefix, "verbose", false, "--DEPRECATED-- Enable verbose output to see the information from packer. Not turning this on will mean the process appears to hang while the image build happens")
	StringVarWithViper(cmd, &o.BuildOS, viperBuildPrefix, "build-os", "ubuntu-2204", "--DEPRECATED-- USE THE CONFIG FILE. This is the target os to build. Valid values are currently: ubuntu-2004 and ubuntu-2204")
	StringVarWithViper(cmd, &o.ImagePrefix, viperBuildPrefix, "image-prefix", "kube", "--DEPRECATED-- USE THE CONFIG FILE. This will prefix the image with the value provided. Defaults to 'kube' producing an image name of kube-yymmdd-xxxxxxxx")
	StringVarWithViper(cmd, &o.ImageRepo, viperBuildPrefix, "image-repo", strings.Join([]string{imageBuilderRepo, "git"}, "."), "--DEPRECATED-- USE THE CONFIG FILE. The imageRepo from which the image builder should be deployed")
	StringVarWithViper(cmd, &o.ImageRepoBranch, viperBuildPrefix, "image-repo-branch", "main", "--DEPRECATED-- USE THE CONFIG FILE. The branch to checkout from the cloned imageRepo")
	StringVarWithViper(cmd, &o.ContainerdSHA256, viperBuildPrefix, "containerd-sha256", "9be621c0206b5c20a1dea05fae12fc698e5083cc81f65c9d918c644090696d19", "--DEPRECATED-- USE THE CONFIG FILE. The sha256 of containerd - required when setting contained")
	StringVarWithViper(cmd, &o.ContainerdVersion, viperBuildPrefix, "containerd-version", "1.7.13", "--DEPRECATED-- USE THE CONFIG FILE. The containerd version to include in the image")
	StringVarWithViper(cmd, &o.CniVersion, viperBuildPrefix, "cni-version", "1.2.0", "--DEPRECATED-- USE THE CONFIG FILE. The CNI plugins version to include to the built image")
	StringVarWithViper(cmd, &o.CrictlVersion, viperBuildPrefix, "crictl-version", "1.25.0", "--DEPRECATED-- USE THE CONFIG FILE. The crictl-tools version to add to the built image")
	StringVarWithViper(cmd, &o.KubeVersion, viperBuildPrefix, "kubernetes-version", "1.25.3", "--DEPRECATED-- USE THE CONFIG FILE. The Kubernetes version to add to the built image")
	StringVarWithViper(cmd, &o.ExtraDebs, viperBuildPrefix, "extra-debs", "", "--DEPRECATED-- USE THE CONFIG FILE. A space-seperated list of any extra (Debian / Ubuntu) packages that should be installed")
	StringSliceVarWithViper(cmd, &o.AdditionalImages, viperBuildPrefix, "additional-images", nil, "--DEPRECATED-- USE THE CONFIG FILE. Add any additional container images which should be baked into the image")
	StringMapVarWithViper(cmd, &o.AdditionalMetadata, viperBuildPrefix, "additional-metadata", nil, "--DEPRECATED-- USE THE CONFIG FILE. Add any additional metadata to tag the image with.")
	BoolVarWithViper(cmd, &o.AddFalco, viperBuildPrefix, "add-falco", false, "--DEPRECATED-- USE THE CONFIG FILE. If enabled, will install Falco onto the image")
	BoolVarWithViper(cmd, &o.AddTrivy, viperBuildPrefix, "add-trivy", false, "--DEPRECATED-- USE THE CONFIG FILE. If enabled, will install Trivy onto the image")

	// GPU
	BoolVarWithViper(cmd, &o.AddGpuSupport, viperGpuPrefix, "enable-gpu-support", false, "--DEPRECATED-- USE THE CONFIG FILE. This will configure GPU support in the image")
	StringVarWithViper(cmd, &o.GpuVendor, viperGpuPrefix, "gpu-vendor", "", "--DEPRECATED-- USE THE CONFIG FILE. The architecture of the GPU (currently supported: nvidia, amd)")
	// AMD
	StringVarWithViper(cmd, &o.AMDVersion, viperGpuPrefix, "amd-driver-version", "6.0.2", "--DEPRECATED-- USE THE CONFIG FILE. The AMD driver version")
	StringVarWithViper(cmd, &o.AMDDebVersion, viperGpuPrefix, "amd-deb-version", "6.0.60002-1", "--DEPRECATED-- USE THE CONFIG FILE. The AMD deb version")
	StringVarWithViper(cmd, &o.AMDUseCase, viperGpuPrefix, "amd-usecase", "dkms", "--DEPRECATED-- USE THE CONFIG FILE. A comma-dliminated string of usecases for the AMDGPU installer")
	// NVIDIA flags
	StringVarWithViper(cmd, &o.NvidiaVersion, viperGpuPrefix, "nvidia-driver-version", "525.129.03", "--DEPRECATED-- USE THE CONFIG FILE. The NVIDIA driver version")
	StringVarWithViper(cmd, &o.NvidiaBucket, viperGpuPrefix, "nvidia-bucket", "", "--DEPRECATED-- USE THE CONFIG FILE. The bucket name in which the NVIDIA components are stored")
	StringVarWithViper(cmd, &o.NvidiaInstallerLocation, viperGpuPrefix, "nvidia-installer-location", "", "--DEPRECATED-- USE THE CONFIG FILE. The NVIDIA installer location in the bucket - this must be acquired from NVIDIA and uploaded to your bucket")
	StringVarWithViper(cmd, &o.NvidiaTOKLocation, viperGpuPrefix, "nvidia-tok-location", "", "--DEPRECATED-- USE THE CONFIG FILE. The NVIDIA .tok file location in the bucket - this must be acquired from NVIDIA and uploaded to your bucket")
	IntVarWithViper(cmd, &o.NvidiaGriddFeatureType, viperGpuPrefix, "nvidia-gridd-feature-type", -1, "--DEPRECATED-- USE THE CONFIG FILE. The gridd feature type - See https://docs.nvidia.com/license-system/latest/nvidia-license-system-quick-start-guide/index.html#configuring-nls-licensed-client-on-linux for more information")

	o.BaseOptions.AddFlags(cmd)
	o.KubernetesClusterFlags.AddFlags(cmd)
	o.S3Flags.AddFlags(cmd)
	o.OpenStackFlags.AddFlags(cmd, viperOpenStackPrefix)
	o.KubeVirtFlags.AddFlags(cmd, viperKubeVirtPrefix)

	cmd.MarkFlagsRequiredTogether("nvidia-driver-version", "nvidia-bucket", "nvidia-installer-location", "nvidia-tok-location", "nvidia-gridd-feature-type")
	cmd.MarkFlagsRequiredTogether("cni-version", "crictl-version", "kubernetes-version")
	cmd.MarkFlagsRequiredTogether("containerd-version", "containerd-sha256")
}
