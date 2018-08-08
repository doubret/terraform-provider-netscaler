package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerResponderpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_responderpolicylabel,
		Read:          read_responderpolicylabel,
		Update:        nil,
		Delete:        delete_responderpolicylabel,
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

func get_responderpolicylabel(d *schema.ResourceData) nitro.Responderpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Responderpolicylabel{
		Comment:         d.Get("comment").(string),
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	return resource
}

func set_responderpolicylabel(d *schema.ResourceData, resource *nitro.Responderpolicylabel) {
	var _ = strconv.Itoa

	d.Set("comment", resource.Comment)
	d.Set("labelname", resource.Labelname)
	d.Set("policylabeltype", resource.Policylabeltype)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_responderpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_responderpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_responderpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderpolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_responderpolicylabel(d, resource)
	} else {
		err := client.AddResponderpolicylabel(get_responderpolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetResponderpolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_responderpolicylabel(d, resource)
	}

	return nil
}

func read_responderpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_responderpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_responderpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderpolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_responderpolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_responderpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_responderpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_responderpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteResponderpolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
