package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerFilterpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_filterpolicy,
		Read:          read_filterpolicy,
		Update:        update_filterpolicy,
		Delete:        delete_filterpolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reqaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"resaction": &schema.Schema{
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

func key_filterpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_filterpolicy(d *schema.ResourceData) nitro.Filterpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Filterpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Resaction: d.Get("resaction").(string),
		Rule:      d.Get("rule").(string),
	}

	return resource
}

func set_filterpolicy(d *schema.ResourceData, resource *nitro.Filterpolicy) {
	d.Set("name", resource.Name)
	d.Set("reqaction", resource.Reqaction)
	d.Set("resaction", resource.Resaction)
	d.Set("rule", resource.Rule)
	d.SetId(resource.Name)
}

func create_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_filterpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_filterpolicy(d)

	exists, err := client.ExistsFilterpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilterpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_filterpolicy(d, resource)
	} else {
		err := client.AddFilterpolicy(get_filterpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetFilterpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_filterpolicy(d, resource)
	}

	return nil
}

func read_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_filterpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_filterpolicy(d)

	exists, err := client.ExistsFilterpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilterpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_filterpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_filterpolicy")

	return nil
}

func delete_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_filterpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_filterpolicy(d)

	exists, err := client.ExistsFilterpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteFilterpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
