package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
		Update:        nil,
		Delete:        delete_videooptimizationpolicylabel,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
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

func get_videooptimizationpolicylabel(d *schema.ResourceData) nitro.Videooptimizationpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Videooptimizationpolicylabel{
		Comment:         d.Get("comment").(string),
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	return resource
}

func set_videooptimizationpolicylabel(d *schema.ResourceData, resource *nitro.Videooptimizationpolicylabel) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("comment", resource.Comment)
	d.Set("labelname", resource.Labelname)
	d.Set("policylabeltype", resource.Policylabeltype)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func get_videooptimizationpolicylabel_key(d *schema.ResourceData) nitro.VideooptimizationpolicylabelKey {

	key := nitro.VideooptimizationpolicylabelKey{
		d.Get("labelname").(string),
	}
	return key
}

func create_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_videooptimizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_videooptimizationpolicylabel(d)
	key := resource.ToKey()

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

	resource := get_videooptimizationpolicylabel(d)
	key := resource.ToKey()

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

func delete_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_videooptimizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_videooptimizationpolicylabel(d)
	key := resource.ToKey()

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
