package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerLbwlm() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbwlm,
		Read:          read_lbwlm,
		Update:        update_lbwlm,
		Delete:        delete_lbwlm,
		Schema: map[string]*schema.Schema{
			"wlmname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"katimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lbuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_lbwlm(d *schema.ResourceData) string {
	return d.Get("wlmname").(string)
}

func get_lbwlm(d *schema.ResourceData) nitro.Lbwlm {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbwlm{
		Wlmname:   d.Get("wlmname").(string),
		Ipaddress: d.Get("ipaddress").(string),
		Katimeout: d.Get("katimeout").(int),
		Lbuid:     d.Get("lbuid").(string),
		Port:      d.Get("port").(int),
	}

	return resource
}

func set_lbwlm(d *schema.ResourceData, resource *nitro.Lbwlm) {
	d.Set("wlmname", resource.Wlmname)
	d.Set("ipaddress", resource.Ipaddress)
	d.Set("katimeout", resource.Katimeout)
	d.Set("lbuid", resource.Lbuid)
	d.Set("port", resource.Port)
	d.SetId(resource.Wlmname)
}

func create_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbwlm")

	client := meta.(*nitro.NitroClient)

	key := key_lbwlm(d)

	exists, err := client.ExistsLbwlm(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbwlm(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbwlm(d, resource)
	} else {
		err := client.AddLbwlm(get_lbwlm(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbwlm(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbwlm(d, resource)
	}

	return nil
}

func read_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbwlm")

	client := meta.(*nitro.NitroClient)

	key := key_lbwlm(d)

	exists, err := client.ExistsLbwlm(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbwlm(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbwlm(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbwlm")

	return nil
}

func delete_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbwlm")

	client := meta.(*nitro.NitroClient)

	key := key_lbwlm(d)

	exists, err := client.ExistsLbwlm(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbwlm(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
