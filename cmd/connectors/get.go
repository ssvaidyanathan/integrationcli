// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectors

import (
	"fmt"

	"github.com/srinandan/integrationcli/apiclient"
	"github.com/srinandan/integrationcli/client/connections"

	"github.com/spf13/cobra"
)

// GetCmd to get connection
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get connection details",
	Long:  "Get connection details from a connection created in a region",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		if view != "BASIC" && view != "FULL" {
			return fmt.Errorf("view must be BASIC or FULL")
		}
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = connections.Get(name, view, minimal)
		return
	},
}

var view string
var minimal bool

func init() {
	GetCmd.Flags().StringVarP(&name, "name", "n",
		"", "The name of the connection")
	GetCmd.Flags().StringVarP(&view, "view", "",
		"BASIC", "fields of the Connection to be returned; default is BASIC. FULL is the other option")
	GetCmd.Flags().BoolVarP(&minimal, "minimal", "",
		false, "fields of the Connection to be returned; default is false")

	_ = GetCmd.MarkFlagRequired("name")
}
