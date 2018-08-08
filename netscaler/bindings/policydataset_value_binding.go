package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPolicydatasetValueBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policydataset_value_binding,
		Read:          read_policydataset_value_binding,
		Update:        nil,
		Delete:        delete_policydataset_value_binding,
		Schema: map[string]*schema.Schema{
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_policydataset_value_binding(d *schema.ResourceData) nitro.PolicydatasetValueBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.PolicydatasetValueBinding{
		Index: d.Get("index").(int),
		Name:  d.Get("name").(string),
		Value: d.Get("value").(string),
	}

	return resource
}

func set_policydataset_value_binding(d *schema.ResourceData, resource *nitro.PolicydatasetValueBinding) {
	var _ = strconv.Itoa

	d.Set("index", resource.Index)
	d.Set("name", resource.Name)
	d.Set("value", resource.Value)

	var key []string

	key = append(key, resource.Name)
	key = append(key, resource.Value)
	d.SetId(strings.Join(key, "-"))
}

func create_policydataset_value_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policydataset_value_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policydataset_value_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicydatasetValueBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicydatasetValueBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_policydataset_value_binding(d, resource)
	} else {
		err := client.AddPolicydatasetValueBinding(get_policydataset_value_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPolicydatasetValueBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_policydataset_value_binding(d, resource)
	}

	return nil
}

func read_policydataset_value_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policydataset_value_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policydataset_value_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicydatasetValueBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicydatasetValueBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_policydataset_value_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_policydataset_value_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policydataset_value_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policydataset_value_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicydatasetValueBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePolicydatasetValueBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
