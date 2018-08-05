package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerVideooptimizationaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_videooptimizationaction,
		Read:          read_videooptimizationaction,
		Update:        update_videooptimizationaction,
		Delete:        delete_videooptimizationaction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_videooptimizationaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_videooptimizationaction(d *schema.ResourceData) nitro.Videooptimizationaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Videooptimizationaction{
		Name:    d.Get("name").(string),
		Comment: d.Get("comment").(string),
		Rate:    d.Get("rate").(int),
		Type:    d.Get("type").(string),
	}

	return resource
}

func set_videooptimizationaction(d *schema.ResourceData, resource *nitro.Videooptimizationaction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("comment", resource.Comment)
	d.Set("rate", resource.Rate)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationaction(d)

	exists, err := client.ExistsVideooptimizationaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_videooptimizationaction(d, resource)
	} else {
		err := client.AddVideooptimizationaction(get_videooptimizationaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_videooptimizationaction(d, resource)
	}

	return nil
}

func read_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationaction(d)

	exists, err := client.ExistsVideooptimizationaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_videooptimizationaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_videooptimizationaction")

	return nil
}

func delete_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationaction(d)

	exists, err := client.ExistsVideooptimizationaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
