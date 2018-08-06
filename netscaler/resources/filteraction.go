package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerFilteraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_filteraction,
		Read:          read_filteraction,
		Update:        update_filteraction,
		Delete:        delete_filteraction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"qual": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"respcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_filteraction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_filteraction(d *schema.ResourceData) nitro.Filteraction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Filteraction{
		Name:        d.Get("name").(string),
		Page:        d.Get("page").(string),
		Qual:        d.Get("qual").(string),
		Respcode:    d.Get("respcode").(int),
		Servicename: d.Get("servicename").(string),
		Value:       d.Get("value").(string),
	}

	return resource
}

func set_filteraction(d *schema.ResourceData, resource *nitro.Filteraction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("page", resource.Page)
	d.Set("qual", resource.Qual)
	d.Set("respcode", resource.Respcode)
	d.Set("servicename", resource.Servicename)
	d.Set("value", resource.Value)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_filteraction")

	client := meta.(*nitro.NitroClient)

	key := key_filteraction(d)

	exists, err := client.ExistsFilteraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilteraction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_filteraction(d, resource)
	} else {
		err := client.AddFilteraction(get_filteraction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetFilteraction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_filteraction(d, resource)
	}

	return nil
}

func read_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_filteraction")

	client := meta.(*nitro.NitroClient)

	key := key_filteraction(d)

	exists, err := client.ExistsFilteraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilteraction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_filteraction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_filteraction")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateFilteraction(get_filteraction(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_filteraction")

	client := meta.(*nitro.NitroClient)

	key := key_filteraction(d)

	exists, err := client.ExistsFilteraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteFilteraction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
