// +build windows

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package firelens

import (
	generator "github.com/awslabs/go-config-generator-for-fluentd-and-fluentbit"
)

const (
	// S3ConfigPathFluentd and S3ConfigPathFluentbit are the paths where we bind mount the config downloaded from S3 to.
	S3ConfigPathFluentd   = `C:\data\fluentd\etc\`
	S3ConfigPathFluentbit = `C:\data\fluent-bit\etc\`

	// inputBindValue is the value for specifying host on Windows this always binds to 0.0.0.0
	inputBridgeBindValue = "0.0.0.0"

	// inputAWSVPCBindValue is the value for specifying host for AWSVPC mode.
	inputAWSVPCBindValue = "0.0.0.0"
)

// Specify log stream input, which is a unix socket that will be used for communication between the Firelens
// container and other containers.
// not supported on Windows
func (firelens *FirelensResource) addSocketInput(config generator.FluentConfig) {

}

// addHealthcheckSections adds a health check input section and a health check output section to the config.
func (firelens *FirelensResource) addHealthcheckSections(config generator.FluentConfig) {

}

func (firelens *FirelensResource) getS3ConfPath() string {
	if firelens.firelensConfigType == FirelensConfigTypeFluentd {
		return S3ConfigPathFluentd + "external.conf"
	} else {
		return S3ConfigPathFluentbit + "external.conf"
	}
}