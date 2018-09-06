package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbprofile,
		Read:          read_lbprofile,
		Update:        update_lbprofile,
		Delete:        delete_lbprofile,
		Schema: map[string]*schema.Schema{
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
			"lbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_lbprofile(d *schema.ResourceData) nitro.Lbprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbprofile{
		Cookiepassphrase:              d.Get("cookiepassphrase").(string),
		Dbslb:                         d.Get("dbslb").(string),
		Httponlycookieflag:            d.Get("httponlycookieflag").(string),
		Lbprofilename:                 d.Get("lbprofilename").(string),
		Processlocal:                  d.Get("processlocal").(string),
		Useencryptedpersistencecookie: d.Get("useencryptedpersistencecookie").(string),
		Usesecuredpersistencecookie:   d.Get("usesecuredpersistencecookie").(string),
	}

	return resource
}

func set_lbprofile(d *schema.ResourceData, resource *nitro.Lbprofile) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("cookiepassphrase", resource.Cookiepassphrase)
	d.Set("dbslb", resource.Dbslb)
	d.Set("httponlycookieflag", resource.Httponlycookieflag)
	d.Set("lbprofilename", resource.Lbprofilename)
	d.Set("processlocal", resource.Processlocal)
	d.Set("useencryptedpersistencecookie", resource.Useencryptedpersistencecookie)
	d.Set("usesecuredpersistencecookie", resource.Usesecuredpersistencecookie)

	var key []string

	key = append(key, resource.Lbprofilename)
	d.SetId(strings.Join(key, "-"))
}

func get_lbprofile_key(d *schema.ResourceData) nitro.LbprofileKey {

	key := nitro.LbprofileKey{
		d.Get("lbprofilename").(string),
	}
	return key
}

func create_lbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_lbprofile(d)
	key := resource.ToKey()

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

	resource := get_lbprofile(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.LbprofileUpdate{}
	unset := nitro.LbprofileUnset{}

	updateFlag := false
	unsetFlag := false

	update.Lbprofilename = d.Get("lbprofilename").(string)
	unset.Lbprofilename = d.Get("lbprofilename").(string)

	if d.HasChange("cookiepassphrase") {
		updateFlag = true

		value := d.Get("cookiepassphrase").(string)
		update.Cookiepassphrase = value

		if value == "" {
			unsetFlag = true

			unset.Cookiepassphrase = true
		}

	}
	if d.HasChange("dbslb") {
		updateFlag = true

		value := d.Get("dbslb").(string)
		update.Dbslb = value

		if value == "" {
			unsetFlag = true

			unset.Dbslb = true
		}

	}
	if d.HasChange("processlocal") {
		updateFlag = true

		value := d.Get("processlocal").(string)
		update.Processlocal = value

		if value == "" {
			unsetFlag = true

			unset.Processlocal = true
		}

	}
	if d.HasChange("httponlycookieflag") {
		updateFlag = true

		value := d.Get("httponlycookieflag").(string)
		update.Httponlycookieflag = value

		if value == "" {
			unsetFlag = true

			unset.Httponlycookieflag = true
		}

	}
	if d.HasChange("usesecuredpersistencecookie") {
		updateFlag = true

		value := d.Get("usesecuredpersistencecookie").(string)
		update.Usesecuredpersistencecookie = value

		if value == "" {
			unsetFlag = true

			unset.Usesecuredpersistencecookie = true
		}

	}
	if d.HasChange("useencryptedpersistencecookie") {
		updateFlag = true

		value := d.Get("useencryptedpersistencecookie").(string)
		update.Useencryptedpersistencecookie = value

		if value == "" {
			unsetFlag = true

			unset.Useencryptedpersistencecookie = true
		}

	}
	key := get_lbprofile_key(d)

	if updateFlag {
		if err := client.UpdateLbprofile(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetLbprofile(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetLbprofile(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_lbprofile(d, resource)
	}

	return nil
}

func delete_lbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_lbprofile(d)
	key := resource.ToKey()

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
