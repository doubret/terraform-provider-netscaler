package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerSpilloveraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_spilloveraction,
		Read:          read_spilloveraction,
		Update:        nil,
		Delete:        delete_spilloveraction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_spilloveraction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_spilloveraction(d *schema.ResourceData) nitro.Spilloveraction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Spilloveraction{
		Name:   d.Get("name").(string),
		Action: d.Get("action").(string),
	}

	return resource
}

func set_spilloveraction(d *schema.ResourceData, resource *nitro.Spilloveraction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_spilloveraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_spilloveraction")

	client := meta.(*nitro.NitroClient)

	key := key_spilloveraction(d)

	exists, err := client.ExistsSpilloveraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSpilloveraction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_spilloveraction(d, resource)
	} else {
		err := client.AddSpilloveraction(get_spilloveraction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetSpilloveraction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_spilloveraction(d, resource)
	}

	return nil
}

func read_spilloveraction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_spilloveraction")

	client := meta.(*nitro.NitroClient)

	key := key_spilloveraction(d)

	exists, err := client.ExistsSpilloveraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSpilloveraction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_spilloveraction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_spilloveraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_spilloveraction")

	client := meta.(*nitro.NitroClient)

	key := key_spilloveraction(d)

	exists, err := client.ExistsSpilloveraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteSpilloveraction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
