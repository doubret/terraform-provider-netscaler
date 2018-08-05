package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerResponderpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_responderpolicy,
		Read:          read_responderpolicy,
		Update:        update_responderpolicy,
		Delete:        delete_responderpolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"appflowaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
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
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_responderpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_responderpolicy(d *schema.ResourceData) nitro.Responderpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Responderpolicy{
		Name:          d.Get("name").(string),
		Action:        d.Get("action").(string),
		Appflowaction: d.Get("appflowaction").(string),
		Comment:       d.Get("comment").(string),
		Logaction:     d.Get("logaction").(string),
		Rule:          d.Get("rule").(string),
		Undefaction:   d.Get("undefaction").(string),
	}

	return resource
}

func set_responderpolicy(d *schema.ResourceData, resource *nitro.Responderpolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("appflowaction", resource.Appflowaction)
	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("rule", resource.Rule)
	d.Set("undefaction", resource.Undefaction)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_responderpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_responderpolicy(d)

	exists, err := client.ExistsResponderpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_responderpolicy(d, resource)
	} else {
		err := client.AddResponderpolicy(get_responderpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetResponderpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_responderpolicy(d, resource)
	}

	return nil
}

func read_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_responderpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_responderpolicy(d)

	exists, err := client.ExistsResponderpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_responderpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_responderpolicy")

	return nil
}

func delete_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_responderpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_responderpolicy(d)

	exists, err := client.ExistsResponderpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteResponderpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
