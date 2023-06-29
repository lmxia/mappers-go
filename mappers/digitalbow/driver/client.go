/*
Copyright 2021 The KubeEdge Authors.

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

/*
* TODO: This file is for device visit like read, write and get status.
* Please refine those functions.
 */

package driver

import (
	"k8s.io/klog/v2"

	"github.com/kubeedge/mappers-go/mappers/common"
)

// DigitalbowConfig is the structure for client configuration.
type DigitalbowConfig struct {
}

// DigitalbowClient is the structure for Digitalbow client.
type DigitalbowClient struct {
	Client interface{}
	Config DigitalbowConfig
}

var clients map[string]*DigitalbowClient

// NewClient allocate and return a Digitalbow client.
func NewClient(config interface{}) (*DigitalbowClient, error) {
	return nil, nil
}

// GetStatus get device status.
func (c *DigitalbowClient) GetStatus() string {
	return common.DEVSTOK
}

// Get get register.
func (c *DigitalbowClient) Get() (results []byte, err error) {
	klog.V(2).Info("Get result: ", results)
	return results, err
}

// Set set register.
func (c *DigitalbowClient) Set() (results []byte, err error) {
	klog.V(1).Info("Set result:", err, results)
	return results, err
}
