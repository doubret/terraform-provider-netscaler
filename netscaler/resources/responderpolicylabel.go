package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
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
		Update:        update_responderpolicylabel,
		Delete:        delete_responderpolicylabel,
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

func key_responderpolicylabel(d *schema.ResourceData) string {
	return d.Get("labelname").(string)
}

func get_responderpolicylabel(d *schema.ResourceData) nitro.Responderpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Responderpolicylabel{
		Labelname:       d.Get("labelname").(string),
		Comment:         d.Get("comment").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	return resource
}

func set_responderpolicylabel(d *schema.ResourceData, resource *nitro.Responderpolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)
	d.Set("comment", resource.Comment)
	d.Set("policylabeltype", resource.Policylabeltype)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_responderpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_responderpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_responderpolicylabel(d)

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

	key := key_responderpolicylabel(d)

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

func update_responderpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_responderpolicylabel")

	return nil
}

func delete_responderpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_responderpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_responderpolicylabel(d)

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
