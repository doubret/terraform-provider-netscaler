package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cspolicy,
		Read:          read_cspolicy,
		Update:        update_cspolicy,
		Delete:        delete_cspolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_cspolicy(d *schema.ResourceData) nitro.Cspolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cspolicy{
		Action:     d.Get("action").(string),
		Domain:     d.Get("domain").(string),
		Logaction:  d.Get("logaction").(string),
		Policyname: d.Get("policyname").(string),
		Rule:       d.Get("rule").(string),
		Url:        d.Get("url").(string),
	}

	return resource
}

func set_cspolicy(d *schema.ResourceData, resource *nitro.Cspolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("domain", resource.Domain)
	d.Set("logaction", resource.Logaction)
	d.Set("policyname", resource.Policyname)
	d.Set("rule", resource.Rule)
	d.Set("url", resource.Url)

	var key []string

	key = append(key, resource.Policyname)
	d.SetId(strings.Join(key, "-"))
}

func create_cspolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cspolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cspolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCspolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCspolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cspolicy(d, resource)
	} else {
		err := client.AddCspolicy(get_cspolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCspolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cspolicy(d, resource)
	}

	return nil
}

func read_cspolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cspolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cspolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCspolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCspolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cspolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_cspolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cspolicy")

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateCspolicy(get_cspolicy(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_cspolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cspolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cspolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCspolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCspolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
