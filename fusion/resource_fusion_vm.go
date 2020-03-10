package fusion

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVm() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmCreate,
		Read:   resourceVmRead,
		Update: resourceVmUpdate,
		Delete: resourceVmDelete,

		Schema: map[string]*schema.Schema{
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVmCreate(d *schema.ResourceData, m interface{}) error {
	return resourceVmRead(d, m)
}

func resourceVmRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceVmUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceVmRead(d, m)
}

func resourceVmDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}