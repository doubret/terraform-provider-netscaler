package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTransformprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_transformprofile,
		Read:          read_transformprofile,
		Update:        update_transformprofile,
		Delete:        delete_transformprofile,
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
			"onlytransformabsurlinbody": &schema.Schema{
				Type:     schema.TypeString,
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

func key_transformprofile(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_transformprofile(d *schema.ResourceData) nitro.Transformprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Transformprofile{
		Name:                      d.Get("name").(string),
		Comment:                   d.Get("comment").(string),
		Onlytransformabsurlinbody: d.Get("onlytransformabsurlinbody").(string),
		Type: d.Get("type").(string),
	}

	return resource
}

func set_transformprofile(d *schema.ResourceData, resource *nitro.Transformprofile) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("comment", resource.Comment)
	d.Set("onlytransformabsurlinbody", resource.Onlytransformabsurlinbody)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_transformprofile")

	client := meta.(*nitro.NitroClient)

	key := key_transformprofile(d)

	exists, err := client.ExistsTransformprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_transformprofile(d, resource)
	} else {
		err := client.AddTransformprofile(get_transformprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTransformprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_transformprofile(d, resource)
	}

	return nil
}

func read_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_transformprofile")

	client := meta.(*nitro.NitroClient)

	key := key_transformprofile(d)

	exists, err := client.ExistsTransformprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_transformprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_transformprofile")

	return nil
}

func delete_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_transformprofile")

	client := meta.(*nitro.NitroClient)

	key := key_transformprofile(d)

	exists, err := client.ExistsTransformprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTransformprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
