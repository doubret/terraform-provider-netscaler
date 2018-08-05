package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerVideooptimizationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_videooptimizationpolicylabel,
		Read:          read_videooptimizationpolicylabel,
		Update:        update_videooptimizationpolicylabel,
		Delete:        delete_videooptimizationpolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_videooptimizationpolicylabel(d *schema.ResourceData) string {
	return d.Get("labelname").(string)
}

func get_videooptimizationpolicylabel(d *schema.ResourceData) nitro.Videooptimizationpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Videooptimizationpolicylabel{
		Labelname:       d.Get("labelname").(string),
		Comment:         d.Get("comment").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	return resource
}

func set_videooptimizationpolicylabel(d *schema.ResourceData, resource *nitro.Videooptimizationpolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)
	d.Set("comment", resource.Comment)
	d.Set("policylabeltype", resource.Policylabeltype)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_videooptimizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationpolicylabel(d)

	exists, err := client.ExistsVideooptimizationpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_videooptimizationpolicylabel(d, resource)
	} else {
		err := client.AddVideooptimizationpolicylabel(get_videooptimizationpolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetVideooptimizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_videooptimizationpolicylabel(d, resource)
	}

	return nil
}

func read_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_videooptimizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationpolicylabel(d)

	exists, err := client.ExistsVideooptimizationpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_videooptimizationpolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_videooptimizationpolicylabel")

	return nil
}

func delete_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_videooptimizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_videooptimizationpolicylabel(d)

	exists, err := client.ExistsVideooptimizationpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteVideooptimizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
