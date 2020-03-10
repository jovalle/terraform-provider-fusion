package fusion

import (
	"github.com/hashicorp/terraform/helper/schema"
	fusionRest "github.com/jovalle/fusion-rest-go"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Your VMware Fusion REST API key",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VMWARE_FUSION_REST_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fusion_vm": resourceVm(),
		},
	}
}