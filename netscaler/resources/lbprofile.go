package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerLbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbprofile,
		Read:          read_lbprofile,
		Update:        update_lbprofile,
		Delete:        delete_lbprofile,
		Schema: map[string]*schema.Schema{
			"lbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cookiepassphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dbslb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"httponlycookieflag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"processlocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"useencryptedpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"usesecuredpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_lbprofile(d *schema.ResourceData) string {
	return d.Get("lbprofilename").(string)
}

func get_lbprofile(d *schema.ResourceData) nitro.Lbprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbprofile{
		Lbprofilename:                 d.Get("lbprofilename").(string),
		Cookiepassphrase:              d.Get("cookiepassphrase").(string),
		Dbslb:                         d.Get("dbslb").(string),
		Httponlycookieflag:            d.Get("httponlycookieflag").(string),
		Processlocal:                  d.Get("processlocal").(string),
		Useencryptedpersistencecookie: d.Get("useencryptedpersistencecookie").(string),
		Usesecuredpersistencecookie:   d.Get("usesecuredpersistencecookie").(string),
	}

	return resource
}

func set_lbprofile(d *schema.ResourceData, resource *nitro.Lbprofile) {
	d.Set("lbprofilename", resource.Lbprofilename)
	d.Set("cookiepassphrase", resource.Cookiepassphrase)
	d.Set("dbslb", resource.Dbslb)
	d.Set("httponlycookieflag", resource.Httponlycookieflag)
	d.Set("processlocal", resource.Processlocal)
	d.Set("useencryptedpersistencecookie", resource.Useencryptedpersistencecookie)
	d.Set("usesecuredpersistencecookie", resource.Usesecuredpersistencecookie)
	d.SetId(resource.Lbprofilename)
}

func create_lbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbprofile")

	client := meta.(*nitro.NitroClient)

	key := key_lbprofile(d)

	exists, err := client.ExistsLbprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbprofile(d, resource)
	} else {
		err := client.AddLbprofile(get_lbprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbprofile(d, resource)
	}

	return nil
}

func read_lbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbprofile")

	client := meta.(*nitro.NitroClient)

	key := key_lbprofile(d)

	exists, err := client.ExistsLbprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_lbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbprofile")

	return nil
}

func delete_lbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbprofile")

	client := meta.(*nitro.NitroClient)

	key := key_lbprofile(d)

	exists, err := client.ExistsLbprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
