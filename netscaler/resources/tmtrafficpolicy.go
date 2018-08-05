package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTmtrafficpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_tmtrafficpolicy,
		Read:          read_tmtrafficpolicy,
		Update:        update_tmtrafficpolicy,
		Delete:        delete_tmtrafficpolicy,
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

func key_tmtrafficpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_tmtrafficpolicy(d *schema.ResourceData) nitro.Tmtrafficpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Tmtrafficpolicy{
		Name:   d.Get("name").(string),
		Action: d.Get("action").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_tmtrafficpolicy(d *schema.ResourceData, resource *nitro.Tmtrafficpolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_tmtrafficpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmtrafficpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_tmtrafficpolicy(d)

	exists, err := client.ExistsTmtrafficpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmtrafficpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_tmtrafficpolicy(d, resource)
	} else {
		err := client.AddTmtrafficpolicy(get_tmtrafficpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTmtrafficpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_tmtrafficpolicy(d, resource)
	}

	return nil
}

func read_tmtrafficpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_tmtrafficpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_tmtrafficpolicy(d)

	exists, err := client.ExistsTmtrafficpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmtrafficpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_tmtrafficpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_tmtrafficpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_tmtrafficpolicy")

	return nil
}

func delete_tmtrafficpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmtrafficpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_tmtrafficpolicy(d)

	exists, err := client.ExistsTmtrafficpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTmtrafficpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
