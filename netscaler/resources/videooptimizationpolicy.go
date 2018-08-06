package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerVideooptimizationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_videooptimizationpolicy,
		Read:          read_videooptimizationpolicy,
		Update:        update_videooptimizationpolicy,
		Delete:        delete_videooptimizationpolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"logaction": &schema.Schema{
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
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_videooptimizationpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_videooptimizationpolicy(d *schema.ResourceData) nitro.Videooptimizationpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Videooptimizationpolicy{
		Name:        d.Get("name").(string),
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	return resource
}

func set_videooptimizationpolicy(d *schema.ResourceData, resource *nitro.Videooptimizationpolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("rule", resource.Rule)
	d.Set("undefaction", resource.Undefaction)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_videooptimizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_videooptimizationpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationpolicy(d)

	exists, err := client.ExistsVideooptimizationpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_videooptimizationpolicy(d, resource)
	} else {
		err := client.AddVideooptimizationpolicy(get_videooptimizationpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetVideooptimizationpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_videooptimizationpolicy(d, resource)
	}

	return nil
}

func read_videooptimizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_videooptimizationpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationpolicy(d)

	exists, err := client.ExistsVideooptimizationpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_videooptimizationpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_videooptimizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_videooptimizationpolicy")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateVideooptimizationpolicy(get_videooptimizationpolicy(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_videooptimizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_videooptimizationpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationpolicy(d)

	exists, err := client.ExistsVideooptimizationpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteVideooptimizationpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
