// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

package stats

import (
	"fmt"
	"testing"
)

func TestIsNetworkStatsError(t *testing.T) {
	isNetStatsErr := isNetworkStatsError(fmt.Errorf("no such file or directory"))
	if isNetStatsErr {
		// Expect it to not be a net stats error
		t.Error("Error incorrectly reported as network stats error")
	}

	isNetStatsErr = isNetworkStatsError(fmt.Errorf("open /sys/class/net/veth2f5f3e4/statistics/tx_bytes: no such file or directory"))
	if !isNetStatsErr {
		// Expect this to be a net stats error
		t.Error("Error incorrectly reported as non network stats error")
	}
}
