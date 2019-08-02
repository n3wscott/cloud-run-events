/*
Copyright 2019 Google LLC

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

package operations

import (
	"context"
	"fmt"

	"cloud.google.com/go/compute/metadata"

	"github.com/googlecloudplatform/cloud-run-events/pkg/pubsub"
)

type PubSubOps struct {
	// Environment variable containing project id.
	Project string `envconfig:"PROJECT_ID"`

	// Client is the Pub/Sub client used by Ops.
	Client pubsub.Client
}

func (s *PubSubOps) CreateClient(ctx context.Context) error {
	if s.Project == "" {
		project, err := metadata.ProjectID()
		if err != nil {
			return err
		}
		s.Project = project
	}

	client, err := pubsub.NewClient(ctx, s.Project)
	if err != nil {
		return fmt.Errorf("failed to create Pub/Sub client: %s", err)
	}
	s.Client = client
	return nil
}
