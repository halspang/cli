// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package standalone

import "os/exec"

// GetRuntimeVersion returns the version for the local Dapr runtime.
func GetRuntimeVersion() string {
	daprBinDir := defaultDaprBinPath()
	daprCMD := binaryFilePath(daprBinDir, "daprd")

	out, err := exec.Command(daprCMD, "--version").Output()
	if err != nil {
		return "n/a\n"
	}
	return string(out)
}

// GetDashboardVersion returns the version for the local Dapr dashboard.
func GetDashboardVersion() string {
	daprBinDir := defaultDaprBinPath()
	dashboardCMD := binaryFilePath(daprBinDir, "dashboard")

	out, err := exec.Command(dashboardCMD, "--version").Output()
	if err != nil {
		return "n/a\n"
	}
	return string(out)
}
