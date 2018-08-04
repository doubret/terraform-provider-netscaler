package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerDnspolicy64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dnspolicy64,
		Read:          read_dnspolicy64,
		Update:        update_dnspolicy64,
		Delete:        delete_dnspolicy64,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_dnspolicy64(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_dnspolicy64(d *schema.ResourceData) nitro.Dnspolicy64 {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dnspolicy64{
		Name:   d.Get("name").(string),
		Action: d.Get("action").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_dnspolicy64(d *schema.ResourceData, resource *nitro.Dnspolicy64) {
	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("rule", resource.Rule)
	d.SetId(resource.Name)
}

func create_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dnspolicy64")

	client := meta.(*nitro.NitroClient)

	key := key_dnspolicy64(d)

	exists, err := client.ExistsDnspolicy64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnspolicy64(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_dnspolicy64(d, resource)
	} else {
		err := client.AddDnspolicy64(get_dnspolicy64(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetDnspolicy64(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_dnspolicy64(d, resource)
	}

	return nil
}

func read_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dnspolicy64")

	client := meta.(*nitro.NitroClient)

	key := key_dnspolicy64(d)

	exists, err := client.ExistsDnspolicy64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnspolicy64(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_dnspolicy64(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dnspolicy64")

	return nil
}

func delete_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dnspolicy64")

	client := meta.(*nitro.NitroClient)

	key := key_dnspolicy64(d)

	exists, err := client.ExistsDnspolicy64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteDnspolicy64(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
