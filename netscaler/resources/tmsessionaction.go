package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTmsessionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_tmsessionaction,
		Read:          read_tmsessionaction,
		Update:        update_tmsessionaction,
		Delete:        delete_tmsessionaction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"defaultauthorizationaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"homepage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"httponlycookie": &schema.Schema{
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
			"persistentcookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistentcookievalidity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sesstimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ssocredential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ssodomain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_tmsessionaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_tmsessionaction(d *schema.ResourceData) nitro.Tmsessionaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Tmsessionaction{
		Name: d.Get("name").(string),
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Homepage:                   d.Get("homepage").(string),
		Httponlycookie:             d.Get("httponlycookie").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Persistentcookie:           d.Get("persistentcookie").(string),
		Persistentcookievalidity:   d.Get("persistentcookievalidity").(int),
		Sesstimeout:                d.Get("sesstimeout").(int),
		Sso:                        d.Get("sso").(string),
		Ssocredential:              d.Get("ssocredential").(string),
		Ssodomain:                  d.Get("ssodomain").(string),
	}

	return resource
}

func set_tmsessionaction(d *schema.ResourceData, resource *nitro.Tmsessionaction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("defaultauthorizationaction", resource.Defaultauthorizationaction)
	d.Set("homepage", resource.Homepage)
	d.Set("httponlycookie", resource.Httponlycookie)
	d.Set("kcdaccount", resource.Kcdaccount)
	d.Set("persistentcookie", resource.Persistentcookie)
	d.Set("persistentcookievalidity", resource.Persistentcookievalidity)
	d.Set("sesstimeout", resource.Sesstimeout)
	d.Set("sso", resource.Sso)
	d.Set("ssocredential", resource.Ssocredential)
	d.Set("ssodomain", resource.Ssodomain)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_tmsessionaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmsessionaction")

	client := meta.(*nitro.NitroClient)

	key := key_tmsessionaction(d)

	exists, err := client.ExistsTmsessionaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmsessionaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_tmsessionaction(d, resource)
	} else {
		err := client.AddTmsessionaction(get_tmsessionaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTmsessionaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_tmsessionaction(d, resource)
	}

	return nil
}

func read_tmsessionaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_tmsessionaction")

	client := meta.(*nitro.NitroClient)

	key := key_tmsessionaction(d)

	exists, err := client.ExistsTmsessionaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmsessionaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_tmsessionaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_tmsessionaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_tmsessionaction")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateTmsessionaction(get_tmsessionaction(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_tmsessionaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmsessionaction")

	client := meta.(*nitro.NitroClient)

	key := key_tmsessionaction(d)

	exists, err := client.ExistsTmsessionaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTmsessionaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
