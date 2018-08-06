package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbmetrictable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbmetrictable,
		Read:          read_lbmetrictable,
		Update:        nil,
		Delete:        delete_lbmetrictable,
		Schema: map[string]*schema.Schema{
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func key_lbmetrictable(d *schema.ResourceData) string {
	return d.Get("metrictable").(string)
}

func get_lbmetrictable(d *schema.ResourceData) nitro.Lbmetrictable {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbmetrictable{
		Metrictable: d.Get("metrictable").(string),
	}

	return resource
}

func set_lbmetrictable(d *schema.ResourceData, resource *nitro.Lbmetrictable) {
	var _ = strconv.Itoa

	d.Set("metrictable", resource.Metrictable)

	var key []string

	key = append(key, resource.Metrictable)
	d.SetId(strings.Join(key, "-"))
}

func create_lbmetrictable(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmetrictable")

	client := meta.(*nitro.NitroClient)

	key := key_lbmetrictable(d)

	exists, err := client.ExistsLbmetrictable(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmetrictable(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbmetrictable(d, resource)
	} else {
		err := client.AddLbmetrictable(get_lbmetrictable(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbmetrictable(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbmetrictable(d, resource)
	}

	return nil
}

func read_lbmetrictable(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbmetrictable")

	client := meta.(*nitro.NitroClient)

	key := key_lbmetrictable(d)

	exists, err := client.ExistsLbmetrictable(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmetrictable(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbmetrictable(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_lbmetrictable(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmetrictable")

	client := meta.(*nitro.NitroClient)

	key := key_lbmetrictable(d)

	exists, err := client.ExistsLbmetrictable(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbmetrictable(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
