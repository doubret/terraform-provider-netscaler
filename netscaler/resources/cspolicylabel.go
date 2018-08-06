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
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cspolicylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_cspolicylabel(d *schema.ResourceData) string {
	return d.Get("labelname").(string)
}

func get_cspolicylabel(d *schema.ResourceData) nitro.Cspolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cspolicylabel{
		Labelname:         d.Get("labelname").(string),
		Cspolicylabeltype: d.Get("cspolicylabeltype").(string),
	}

	return resource
}

func set_cspolicylabel(d *schema.ResourceData, resource *nitro.Cspolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)
	d.Set("cspolicylabeltype", resource.Cspolicylabeltype)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cspolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_cspolicylabel(d)

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

	key := key_cspolicylabel(d)

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

	key := key_cspolicylabel(d)

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
