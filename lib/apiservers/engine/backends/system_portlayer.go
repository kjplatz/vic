// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vicbackends

//****
// system_portlayer.go
//
// Contains all code that touches the portlayer for system operations and all
// code that converts swagger based returns to docker personality backend structs.
// The goal is to make the backend code that implements the docker engine-api
// interfaces be as simple as possible and contain no swagger or portlayer code.
//
// Rule for code to be in here:
// 1. touches VIC portlayer
// 2. converts swagger to docker engine-api structs
// 3. errors MUST be docker engine-api compatible errors.  DO NOT return arbitrary errors!
//		- Do NOT return portlayer errors
//		- Do NOT return fmt.Errorf()
//		- Do NOT return errors.New()
//		- DO USE the aliased docker error package 'derr'

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	derr "github.com/docker/docker/errors"

	"github.com/vmware/vic/lib/apiservers/portlayer/client/containers"
	"github.com/vmware/vic/lib/apiservers/portlayer/client/misc"
	"github.com/vmware/vic/lib/apiservers/portlayer/models"
	"github.com/vmware/vic/pkg/trace"
)

type VicSystemProxy interface {
	PingPortlayer() bool
	ContainerCount() (int, int, int, error)
	VCHInfo() (*models.VCHInfo, error)
}

type SystemProxy struct{}

func (s *SystemProxy) PingPortlayer() bool {
	defer trace.End(trace.Begin("PingPortlayer"))

	plClient := PortLayerClient()
	if plClient == nil {
		return false
	}

	if plClient != nil {
		pingParams := misc.NewPingParamsWithContext(ctx)
		_, err := plClient.Misc.Ping(pingParams)
		if err != nil {
			log.Info("Ping to portlayer failed")
			return false
		}
		return true
	}

	log.Errorf("Portlayer client is invalid")
	return false
}

// Use the Portlayer's support for docker ps to get the container count
//   return order: running, paused, stopped counts
func (s *SystemProxy) ContainerCount() (int, int, int, error) {
	defer trace.End(trace.Begin("ContainerCount"))

	var running, paused, stopped int

	plClient := PortLayerClient()
	if plClient == nil {
		return 0, 0, 0, derr.NewErrorWithStatusCode(fmt.Errorf("ContainerCount failed to create a portlayer client"),
			http.StatusInternalServerError)
	}

	all := true
	containList, err := plClient.Containers.GetContainerList(containers.NewGetContainerListParamsWithContext(ctx).WithAll(&all))
	if err != nil {
		return 0, 0, 0, derr.NewErrorWithStatusCode(fmt.Errorf("Failed to get container list: %s", err), http.StatusInternalServerError)
	}

	for _, t := range containList.Payload {
		if *t.Status == "Running" {
			running++
		} else if *t.Status == "Stopped" || *t.Status == "Created" {
			stopped++
		}
	}

	return running, paused, stopped, nil
}

func (s *SystemProxy) VCHInfo() (*models.VCHInfo, error) {
	defer trace.End(trace.Begin("VCHInfo"))

	plClient := PortLayerClient()
	if plClient == nil {
		return nil, derr.NewErrorWithStatusCode(fmt.Errorf("VCHInfo failed to create a portlayer client"),
			http.StatusInternalServerError)
	}

	params := misc.NewGetVCHInfoParamsWithContext(ctx)
	resp, err := plClient.Misc.GetVCHInfo(params)
	if err != nil {
		//There are no custom error for this operation.  If we get back an error, it's
		//unknown.
		return nil, derr.NewErrorWithStatusCode(fmt.Errorf("Unknown error from port layer: %s", err),
			http.StatusInternalServerError)
	}

	return resp.Payload, nil
}
