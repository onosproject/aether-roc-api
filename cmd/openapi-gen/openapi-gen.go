// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
//

package main

import (
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	openapi "github.com/go-openapi/spec"
	"github.com/onosproject/aether-roc-api/pkg/schemawalker"
	devicetype "github.com/onosproject/onos-config/api/types/device"
	"github.com/onosproject/onos-config/pkg/modelregistry"
	"github.com/onosproject/onos-config/pkg/utils"
	"github.com/openconfig/goyang/pkg/yang"
	"os"
	"strings"
)

// Generate an OpenApi 3.0.0 definition from a config-model
// This is a compile time process that generates a swagger.yaml file from the schema
func main() {
	modelPluginName := flag.String("modelPlugin", "", "names of model plugin (.so) to load")
	flag.Parse()

	if *modelPluginName == "" {
		fmt.Println("-modelPlugin flag must be specified")
		os.Exit(-1)
	}

	if modelPluginName == nil {
		fmt.Println("modelPlugin must be given")
		os.Exit(-1)
	}

	modelRegistry := &modelregistry.ModelRegistry{
		ModelPlugins:        make(map[string]modelregistry.ModelPlugin),
		ModelReadOnlyPaths:  make(map[string]modelregistry.ReadOnlyPathMap),
		ModelReadWritePaths: make(map[string]modelregistry.ReadWritePathMap),
		LocationStore:       make(map[string]string),
	}

	name, version, err := modelRegistry.RegisterModelPlugin(*modelPluginName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	modelName := utils.ToModelName(devicetype.Type(name), devicetype.Version(version))

	fmt.Printf("Loaded model plugin %s:%s as %s\n", name, version, modelName)

	modelPlugin, ok := modelRegistry.ModelPlugins[modelName]
	if !ok {
		fmt.Printf("Unable to get model plugin by name %s\n", modelName)
		os.Exit(-1)
	}
	modelSchema, err := modelPlugin.Schema()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	pathItems := make(map[string]openapi.PathItem)
	defintions := make(openapi.Definitions)

	err = schemawalker.ExtractPaths(modelSchema["Device"], yang.TSUnset, "", "", pathItems, defintions)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	openAPISchema := openapi.SwaggerProps{
		Swagger: "2.0",
		Info: &openapi.Info{
			InfoProps: openapi.InfoProps{
				Description: fmt.Sprintf("Aether ROC API %s", modelName),
				Title:       fmt.Sprintf("aether-roc-api-%s", strings.ToLower(name)),
				Version:     version,
			},
		},
		Paths: &openapi.Paths{
			VendorExtensible: openapi.VendorExtensible{},
			Paths:            pathItems,
		},
		Definitions: defintions,
	}

	data, err := yaml.Marshal(openAPISchema)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(data))
}
