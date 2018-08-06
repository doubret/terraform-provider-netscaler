package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTransformaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_transformaction,
		Read:          read_transformaction,
		Update:        update_transformaction,
		Delete:        delete_transformaction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cookiedomainfrom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cookiedomaininto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"profilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"requrlfrom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"requrlinto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"resurlfrom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"resurlinto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_transformaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_transformaction(d *schema.ResourceData) nitro.Transformaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Transformaction{
		Name:             d.Get("name").(string),
		Comment:          d.Get("comment").(string),
		Cookiedomainfrom: d.Get("cookiedomainfrom").(string),
		Cookiedomaininto: d.Get("cookiedomaininto").(string),
		Priority:         d.Get("priority").(int),
		Profilename:      d.Get("profilename").(string),
		Requrlfrom:       d.Get("requrlfrom").(string),
		Requrlinto:       d.Get("requrlinto").(string),
		Resurlfrom:       d.Get("resurlfrom").(string),
		Resurlinto:       d.Get("resurlinto").(string),
		State:            d.Get("state").(string),
	}

	return resource
}

func set_transformaction(d *schema.ResourceData, resource *nitro.Transformaction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("comment", resource.Comment)
	d.Set("cookiedomainfrom", resource.Cookiedomainfrom)
	d.Set("cookiedomaininto", resource.Cookiedomaininto)
	d.Set("priority", resource.Priority)
	d.Set("profilename", resource.Profilename)
	d.Set("requrlfrom", resource.Requrlfrom)
	d.Set("requrlinto", resource.Requrlinto)
	d.Set("resurlfrom", resource.Resurlfrom)
	d.Set("resurlinto", resource.Resurlinto)
	d.Set("state", resource.State)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_transformaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_transformaction")

	client := meta.(*nitro.NitroClient)

	key := key_transformaction(d)

	exists, err := client.ExistsTransformaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_transformaction(d, resource)
	} else {
		err := client.AddTransformaction(get_transformaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTransformaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_transformaction(d, resource)
	}

	return nil
}

func read_transformaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_transformaction")

	client := meta.(*nitro.NitroClient)

	key := key_transformaction(d)

	exists, err := client.ExistsTransformaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_transformaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_transformaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_transformaction")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateTransformaction(get_transformaction(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_transformaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_transformaction")

	client := meta.(*nitro.NitroClient)

	key := key_transformaction(d)

	exists, err := client.ExistsTransformaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTransformaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
