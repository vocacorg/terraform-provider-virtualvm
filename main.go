package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vocacorg/terraform-provider-template/template"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: template.Provider})
}
