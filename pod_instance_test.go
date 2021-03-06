/*
Copyright 2017 The go-marathon Authors All rights reserved.

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

package marathon

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const fakePodInstanceName = "fake-pod.instance-dc6cfe60-6812-11e7-a18e-70b3d5800003"
const secondFakePodInstanceName = "fake-pod2.instance-63c9735b-5ca0-11e9-98d6-1eb6fdf96d81"

func TestDeletePodInstance(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	podInstance, err := endpoint.Client.DeletePodInstance(fakePodName, fakePodInstanceName)
	require.NoError(t, err)
	assert.Equal(t, podInstance.InstanceID.ID, fakePodInstanceName)

	podInstance2, err := endpoint.Client.DeletePodInstance(secondFakePodName, secondFakePodInstanceName)
	require.NoError(t, err)
	assert.Equal(t, podInstance2.InstanceID.ID, secondFakePodInstanceName)
}

func TestDeletePodInstances(t *testing.T) {
	endpoint := newFakeMarathonEndpoint(t, nil)
	defer endpoint.Close()

	instances := []string{fakePodInstanceName}
	podInstances, err := endpoint.Client.DeletePodInstances(fakePodName, instances)
	require.NoError(t, err)
	assert.Equal(t, podInstances[0].InstanceID.ID, fakePodInstanceName)
}
