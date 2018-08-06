package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerRewritepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_rewritepolicylabel,
		Read:          read_rewritepolicylabel,
		Update:        nil,
		Delete:        delete_rewritepolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_rewritepolicylabel(d *schema.ResourceData) string {
	return d.Get("labelname").(string)
}

func get_rewritepolicylabel(d *schema.ResourceData) nitro.Rewritepolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Rewritepolicylabel{
		Labelname: d.Get("labelname").(string),
		Comment:   d.Get("comment").(string),
		Transform: d.Get("transform").(string),
	}

	return resource
}

func set_rewritepolicylabel(d *schema.ResourceData, resource *nitro.Rewritepolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)
	d.Set("comment", resource.Comment)
	d.Set("transform", resource.Transform)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_rewritepolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_rewritepolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_rewritepolicylabel(d)

	exists, err := client.ExistsRewritepolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetRewritepolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_rewritepolicylabel(d, resource)
	} else {
		err := client.AddRewritepolicylabel(get_rewritepolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetRewritepolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_rewritepolicylabel(d, resource)
	}

	return nil
}

func read_rewritepolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_rewritepolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_rewritepolicylabel(d)

	exists, err := client.ExistsRewritepolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetRewritepolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_rewritepolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_rewritepolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_rewritepolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_rewritepolicylabel(d)

	exists, err := client.ExistsRewritepolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteRewritepolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
