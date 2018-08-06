package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerServer() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_server,
		Read:          read_server,
		Update:        update_server,
		Delete:        delete_server,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"domainresolveretry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ipv6address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"translationip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"translationmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_server(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_server(d *schema.ResourceData) nitro.Server {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Server{
		Name:               d.Get("name").(string),
		State:              d.Get("state").(string),
		Comment:            d.Get("comment").(string),
		Domain:             d.Get("domain").(string),
		Domainresolveretry: d.Get("domainresolveretry").(int),
		Ipaddress:          d.Get("ipaddress").(string),
		Ipv6address:        d.Get("ipv6address").(string),
		Td:                 d.Get("td").(int),
		Translationip:      d.Get("translationip").(string),
		Translationmask:    d.Get("translationmask").(string),
	}

	return resource
}

func set_server(d *schema.ResourceData, resource *nitro.Server) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("state", resource.State)
	d.Set("comment", resource.Comment)
	d.Set("domain", resource.Domain)
	d.Set("domainresolveretry", resource.Domainresolveretry)
	d.Set("ipaddress", resource.Ipaddress)
	d.Set("ipv6address", resource.Ipv6address)
	d.Set("td", resource.Td)
	d.Set("translationip", resource.Translationip)
	d.Set("translationmask", resource.Translationmask)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_server(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_server")

	client := meta.(*nitro.NitroClient)

	key := key_server(d)

	exists, err := client.ExistsServer(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServer(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_server(d, resource)
	} else {
		err := client.AddServer(get_server(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetServer(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_server(d, resource)
	}

	return nil
}

func read_server(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_server")

	client := meta.(*nitro.NitroClient)

	key := key_server(d)

	exists, err := client.ExistsServer(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServer(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_server(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_server(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_server")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateServer(get_server(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_server(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_server")

	client := meta.(*nitro.NitroClient)

	key := key_server(d)

	exists, err := client.ExistsServer(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteServer(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
