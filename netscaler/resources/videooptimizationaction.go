package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_videooptimizationaction(d *schema.ResourceData) nitro.Videooptimizationaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Videooptimizationaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Rate:    d.Get("rate").(int),
		Type:    d.Get("type").(string),
	}

	return resource
}

func set_videooptimizationaction(d *schema.ResourceData, resource *nitro.Videooptimizationaction) {
	var _ = strconv.Itoa

	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("rate", resource.Rate)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	resource := get_videooptimizationaction(d)
	key := resource.ToKey()

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

	resource := get_videooptimizationaction(d)
	key := resource.ToKey()

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

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateVideooptimizationaction(get_videooptimizationaction(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	resource := get_videooptimizationaction(d)
	key := resource.ToKey()

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
