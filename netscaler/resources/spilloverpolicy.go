package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerSpilloverpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_spilloverpolicy,
		Read:          read_spilloverpolicy,
		Update:        update_spilloverpolicy,
		Delete:        delete_spilloverpolicy,
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
			"comment": &schema.Schema{
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

func key_spilloverpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_spilloverpolicy(d *schema.ResourceData) nitro.Spilloverpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Spilloverpolicy{
		Name:    d.Get("name").(string),
		Action:  d.Get("action").(string),
		Comment: d.Get("comment").(string),
		Rule:    d.Get("rule").(string),
	}

	return resource
}

func set_spilloverpolicy(d *schema.ResourceData, resource *nitro.Spilloverpolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("comment", resource.Comment)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_spilloverpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_spilloverpolicy(d)

	exists, err := client.ExistsSpilloverpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_spilloverpolicy(d, resource)
	} else {
		err := client.AddSpilloverpolicy(get_spilloverpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_spilloverpolicy(d, resource)
	}

	return nil
}

func read_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_spilloverpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_spilloverpolicy(d)

	exists, err := client.ExistsSpilloverpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_spilloverpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_spilloverpolicy")

	return nil
}

func delete_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_spilloverpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_spilloverpolicy(d)

	exists, err := client.ExistsSpilloverpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
