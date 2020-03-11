package fusion

import (
	"github.com/hashicorp/terraform/helper/schema"
	fusionRest "github.com/jovalle/fusion-rest-go"
)

func resourceVm() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmCreate,
		Read:   resourceVmRead,
		Update: resourceVmUpdate,
		Delete: resourceVmDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVmCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*fusionRest.Client)
	name := d.Get("name").(string)
	parentId := d.Get("parent_id").(string)

	newVm := &fusionRest.NewVm{
		Name: name,
		ParentId: parentId,
	}

	vm, err := client.CreateVm(newVm)
	if err != nil {
		return err
	}

	d.SetId(vm.Id)
	return resourceVmRead(d, m)
}

func resourceVmRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*fusionRest.Client)
	id := d.Id()
	var v fusionRest.Vm

	vm, err := client.GetVm(id)
	if err == nil {
		v.Id = vm.Id
	}

	d.Set("id", v.Id)

	return nil
}

func resourceVmUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*fusionRest.Client)
	id := d.Id()
	cpu := d.Get("cpu").(string)
	memory := d.Get("memory").(string)

	if d.HasChange("cpu") || d.HasChange("memory") {
		v := &fusionRest.Vm{
			Id:     id,
			Cpu:    cpu,
			Memory: memory,
		}
		err := client.UpdateVm(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceVmDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*fusionRest.Client)
	id := d.Id()

	err := client.DeleteVm(id)
	if err != nil {
		return err
	}

	return nil
}