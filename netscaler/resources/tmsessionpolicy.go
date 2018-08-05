package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTmsessionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_tmsessionpolicy,
		Read:          read_tmsessionpolicy,
		Update:        update_tmsessionpolicy,
		Delete:        delete_tmsessionpolicy,
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

func key_tmsessionpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_tmsessionpolicy(d *schema.ResourceData) nitro.Tmsessionpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Tmsessionpolicy{
		Name:   d.Get("name").(string),
		Action: d.Get("action").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_tmsessionpolicy(d *schema.ResourceData, resource *nitro.Tmsessionpolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmsessionpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_tmsessionpolicy(d)

	exists, err := client.ExistsTmsessionpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_tmsessionpolicy(d, resource)
	} else {
		err := client.AddTmsessionpolicy(get_tmsessionpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_tmsessionpolicy(d, resource)
	}

	return nil
}

func read_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_tmsessionpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_tmsessionpolicy(d)

	exists, err := client.ExistsTmsessionpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_tmsessionpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_tmsessionpolicy")

	return nil
}

func delete_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmsessionpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_tmsessionpolicy(d)

	exists, err := client.ExistsTmsessionpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
