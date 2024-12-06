// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dackerman/terraform-provider-demostore/internal"
	"github.com/dackerman/terraform-provider-demostore/internal/migration"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/tidwall/gjson"
)

var (
	version string = "dev"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	var upgrade bool
	flag.BoolVar(&upgrade, "upgrade", false, "run upgrade check")
	flag.Parse()

	if upgrade {
		checkUpgrade()
		return
	}

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/dackerman/stlstore",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), internal.NewProvider(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func checkUpgrade() {
	fmt.Println("Checking for upgrades...!!")

	oldData, err := os.ReadFile("upgrading/alpha4.json")
	if err != nil {
		panic(err)
	}
	oldJson := string(oldData)
	newData, err := os.ReadFile("upgrading/alpha5.json")
	if err != nil {
		panic(err)
	}
	newJson := string(newData)

	oldSchemas := gjson.Get(oldJson, `provider_schemas.*.resource_schemas`)
	newSchemas := gjson.Get(newJson, `provider_schemas.*.resource_schemas`)

	newSchemas.ForEach(func(key, value gjson.Result) bool {

		oldSchema := oldSchemas.Get(key.String())
		if !oldSchema.Exists() {
			return true
		}

		oldAttributes := oldSchema.Get("block.attributes")
		newAttributes := value.Get("block.attributes")

		notInOld := []string{}
		notInNew := []string{}

		existing := migration.AllMigrations[key.String()]

		newAttributes.ForEach(func(key, value gjson.Result) bool {
			oldAttr := oldAttributes.Get(key.String())
			if !oldAttr.Exists() && existing[key.String()] == nil {
				notInOld = append(notInOld, key.String())
			}
			return true
		})

		oldAttributes.ForEach(func(key, value gjson.Result) bool {
			newAttr := newAttributes.Get(key.String())
			if !newAttr.Exists() && existing[key.String()] == nil {
				notInNew = append(notInNew, key.String())
			}
			return true
		})

		if (len(notInOld) == 0) && (len(notInNew) == 0) {
			fmt.Println("Found no differences in", key)
			return true
		}

		if len(notInOld) > 0 {
			fmt.Println("Found new attributes in", key)
			for _, attr := range notInOld {
				fmt.Println(" ", attr)
			}
		}

		if len(notInNew) > 0 {
			fmt.Println("Found removed attributes in", key)
			for _, attr := range notInNew {
				fmt.Println(" ", attr)
			}
		}

		return true
	})
}
