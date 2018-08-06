package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbgroup,
		Read:          read_lbgroup,
		Update:        update_lbgroup,
		Delete:        delete_lbgroup,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"backuppersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cookiedomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cookiename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistencebackup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistencetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistmask": &schema.Schema{
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
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"usevserverpersistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"v6persistmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_lbgroup(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_lbgroup(d *schema.ResourceData) nitro.Lbgroup {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbgroup{
		Name: d.Get("name").(string),
		Backuppersistencetimeout: d.Get("backuppersistencetimeout").(int),
		Cookiedomain:             d.Get("cookiedomain").(string),
		Cookiename:               d.Get("cookiename").(string),
		Persistencebackup:        d.Get("persistencebackup").(string),
		Persistencetype:          d.Get("persistencetype").(string),
		Persistmask:              d.Get("persistmask").(string),
		Rule:                     d.Get("rule").(string),
		Timeout:                  d.Get("timeout").(int),
		Usevserverpersistency:    d.Get("usevserverpersistency").(string),
		V6persistmasklen:         d.Get("v6persistmasklen").(int),
	}

	return resource
}

func set_lbgroup(d *schema.ResourceData, resource *nitro.Lbgroup) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("backuppersistencetimeout", resource.Backuppersistencetimeout)
	d.Set("cookiedomain", resource.Cookiedomain)
	d.Set("cookiename", resource.Cookiename)
	d.Set("persistencebackup", resource.Persistencebackup)
	d.Set("persistencetype", resource.Persistencetype)
	d.Set("persistmask", resource.Persistmask)
	d.Set("rule", resource.Rule)
	d.Set("timeout", resource.Timeout)
	d.Set("usevserverpersistency", resource.Usevserverpersistency)
	d.Set("v6persistmasklen", resource.V6persistmasklen)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_lbgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbgroup")

	client := meta.(*nitro.NitroClient)

	key := key_lbgroup(d)

	exists, err := client.ExistsLbgroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbgroup(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbgroup(d, resource)
	} else {
		err := client.AddLbgroup(get_lbgroup(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbgroup(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbgroup(d, resource)
	}

	return nil
}

func read_lbgroup(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbgroup")

	client := meta.(*nitro.NitroClient)

	key := key_lbgroup(d)

	exists, err := client.ExistsLbgroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbgroup(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbgroup(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_lbgroup(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbgroup")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateLbgroup(get_lbgroup(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_lbgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbgroup")

	client := meta.(*nitro.NitroClient)

	key := key_lbgroup(d)

	exists, err := client.ExistsLbgroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbgroup(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
