package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCspolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cspolicylabel,
		Read:          read_cspolicylabel,
		Update:        nil,
		Delete:        delete_cspolicylabel,
		Schema: map[string]*schema.Schema{
			"cspolicylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_cspolicylabel(d *schema.ResourceData) nitro.Cspolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cspolicylabel{
		Cspolicylabeltype: d.Get("cspolicylabeltype").(string),
		Labelname:         d.Get("labelname").(string),
	}

	return resource
}

func set_cspolicylabel(d *schema.ResourceData, resource *nitro.Cspolicylabel) {
	var _ = strconv.Itoa

	d.Set("cspolicylabeltype", resource.Cspolicylabeltype)
	d.Set("labelname", resource.Labelname)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func get_cspolicylabel_key(d *schema.ResourceData) nitro.CspolicylabelKey {

	key := nitro.CspolicylabelKey{
		d.Get("labelname").(string),
	}
	return key
}

func create_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cspolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_cspolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsCspolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCspolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cspolicylabel(d, resource)
	} else {
		err := client.AddCspolicylabel(get_cspolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCspolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cspolicylabel(d, resource)
	}

	return nil
}

func read_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cspolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_cspolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsCspolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCspolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cspolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cspolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_cspolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsCspolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCspolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
