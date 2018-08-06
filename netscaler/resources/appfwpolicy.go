package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppfwpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appfwpolicy,
		Read:          read_appfwpolicy,
		Update:        update_appfwpolicy,
		Delete:        delete_appfwpolicy,
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
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"profilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_appfwpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_appfwpolicy(d *schema.ResourceData) nitro.Appfwpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appfwpolicy{
		Name:        d.Get("name").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Profilename: d.Get("profilename").(string),
		Rule:        d.Get("rule").(string),
	}

	return resource
}

func set_appfwpolicy(d *schema.ResourceData, resource *nitro.Appfwpolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("profilename", resource.Profilename)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_appfwpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appfwpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_appfwpolicy(d)

	exists, err := client.ExistsAppfwpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppfwpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appfwpolicy(d, resource)
	} else {
		err := client.AddAppfwpolicy(get_appfwpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppfwpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appfwpolicy(d, resource)
	}

	return nil
}

func read_appfwpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appfwpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_appfwpolicy(d)

	exists, err := client.ExistsAppfwpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppfwpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appfwpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appfwpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appfwpolicy")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateAppfwpolicy(get_appfwpolicy(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_appfwpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appfwpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_appfwpolicy(d)

	exists, err := client.ExistsAppfwpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppfwpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
