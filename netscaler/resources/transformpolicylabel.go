package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTransformpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_transformpolicylabel,
		Read:          read_transformpolicylabel,
		Update:        nil,
		Delete:        delete_transformpolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_transformpolicylabel(d *schema.ResourceData) string {
	return d.Get("labelname").(string)
}

func get_transformpolicylabel(d *schema.ResourceData) nitro.Transformpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Transformpolicylabel{
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	return resource
}

func set_transformpolicylabel(d *schema.ResourceData, resource *nitro.Transformpolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)
	d.Set("policylabeltype", resource.Policylabeltype)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_transformpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_transformpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_transformpolicylabel(d)

	exists, err := client.ExistsTransformpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformpolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_transformpolicylabel(d, resource)
	} else {
		err := client.AddTransformpolicylabel(get_transformpolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTransformpolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_transformpolicylabel(d, resource)
	}

	return nil
}

func read_transformpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_transformpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_transformpolicylabel(d)

	exists, err := client.ExistsTransformpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformpolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_transformpolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_transformpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_transformpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_transformpolicylabel(d)

	exists, err := client.ExistsTransformpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTransformpolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
