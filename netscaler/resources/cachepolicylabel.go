package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCachepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cachepolicylabel,
		Read:          read_cachepolicylabel,
		Update:        nil,
		Delete:        delete_cachepolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"evaluates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_cachepolicylabel(d *schema.ResourceData) string {
	return d.Get("labelname").(string)
}

func get_cachepolicylabel(d *schema.ResourceData) nitro.Cachepolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cachepolicylabel{
		Labelname: d.Get("labelname").(string),
		Evaluates: d.Get("evaluates").(string),
	}

	return resource
}

func set_cachepolicylabel(d *schema.ResourceData, resource *nitro.Cachepolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)
	d.Set("evaluates", resource.Evaluates)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_cachepolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cachepolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_cachepolicylabel(d)

	exists, err := client.ExistsCachepolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachepolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cachepolicylabel(d, resource)
	} else {
		err := client.AddCachepolicylabel(get_cachepolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCachepolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cachepolicylabel(d, resource)
	}

	return nil
}

func read_cachepolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cachepolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_cachepolicylabel(d)

	exists, err := client.ExistsCachepolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachepolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cachepolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_cachepolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cachepolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_cachepolicylabel(d)

	exists, err := client.ExistsCachepolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCachepolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
