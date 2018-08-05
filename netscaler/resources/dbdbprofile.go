package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerDbdbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dbdbprofile,
		Read:          read_dbdbprofile,
		Update:        update_dbdbprofile,
		Delete:        delete_dbdbprofile,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"conmultiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"enablecachingconmuxoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"interpretquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kcdaccount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"stickiness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_dbdbprofile(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_dbdbprofile(d *schema.ResourceData) nitro.Dbdbprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dbdbprofile{
		Name:                   d.Get("name").(string),
		Conmultiplex:           d.Get("conmultiplex").(string),
		Enablecachingconmuxoff: d.Get("enablecachingconmuxoff").(string),
		Interpretquery:         d.Get("interpretquery").(string),
		Kcdaccount:             d.Get("kcdaccount").(string),
		Stickiness:             d.Get("stickiness").(string),
	}

	return resource
}

func set_dbdbprofile(d *schema.ResourceData, resource *nitro.Dbdbprofile) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("conmultiplex", resource.Conmultiplex)
	d.Set("enablecachingconmuxoff", resource.Enablecachingconmuxoff)
	d.Set("interpretquery", resource.Interpretquery)
	d.Set("kcdaccount", resource.Kcdaccount)
	d.Set("stickiness", resource.Stickiness)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dbdbprofile")

	client := meta.(*nitro.NitroClient)

	key := key_dbdbprofile(d)

	exists, err := client.ExistsDbdbprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDbdbprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_dbdbprofile(d, resource)
	} else {
		err := client.AddDbdbprofile(get_dbdbprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetDbdbprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_dbdbprofile(d, resource)
	}

	return nil
}

func read_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dbdbprofile")

	client := meta.(*nitro.NitroClient)

	key := key_dbdbprofile(d)

	exists, err := client.ExistsDbdbprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDbdbprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_dbdbprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dbdbprofile")

	return nil
}

func delete_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dbdbprofile")

	client := meta.(*nitro.NitroClient)

	key := key_dbdbprofile(d)

	exists, err := client.ExistsDbdbprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteDbdbprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
