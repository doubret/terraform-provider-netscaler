package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerScpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_scpolicy,
		Read:          read_scpolicy,
		Update:        update_scpolicy,
		Delete:        delete_scpolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"altcontentpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"altcontentsvcname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxconn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_scpolicy(d *schema.ResourceData) nitro.Scpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Scpolicy{
		Action:            d.Get("action").(string),
		Altcontentpath:    d.Get("altcontentpath").(string),
		Altcontentsvcname: d.Get("altcontentsvcname").(string),
		Delay:             d.Get("delay").(int),
		Maxconn:           d.Get("maxconn").(int),
		Name:              d.Get("name").(string),
		Rule:              d.Get("rule").(string),
		Url:               d.Get("url").(string),
	}

	return resource
}

func set_scpolicy(d *schema.ResourceData, resource *nitro.Scpolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("altcontentpath", resource.Altcontentpath)
	d.Set("altcontentsvcname", resource.Altcontentsvcname)
	d.Set("delay", resource.Delay)
	d.Set("maxconn", resource.Maxconn)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)
	d.Set("url", resource.Url)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_scpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_scpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_scpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsScpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetScpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_scpolicy(d, resource)
	} else {
		err := client.AddScpolicy(get_scpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetScpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_scpolicy(d, resource)
	}

	return nil
}

func read_scpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_scpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_scpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsScpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetScpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_scpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_scpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_scpolicy")

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateScpolicy(get_scpolicy(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_scpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_scpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_scpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsScpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteScpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
