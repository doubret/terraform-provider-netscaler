package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_dbdbprofile(d *schema.ResourceData) nitro.Dbdbprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dbdbprofile{
		Conmultiplex:           d.Get("conmultiplex").(string),
		Enablecachingconmuxoff: d.Get("enablecachingconmuxoff").(string),
		Interpretquery:         d.Get("interpretquery").(string),
		Kcdaccount:             d.Get("kcdaccount").(string),
		Name:                   d.Get("name").(string),
		Stickiness:             d.Get("stickiness").(string),
	}

	return resource
}

func set_dbdbprofile(d *schema.ResourceData, resource *nitro.Dbdbprofile) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("conmultiplex", resource.Conmultiplex)
	d.Set("enablecachingconmuxoff", resource.Enablecachingconmuxoff)
	d.Set("interpretquery", resource.Interpretquery)
	d.Set("kcdaccount", resource.Kcdaccount)
	d.Set("name", resource.Name)
	d.Set("stickiness", resource.Stickiness)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_dbdbprofile_key(d *schema.ResourceData) nitro.DbdbprofileKey {

	key := nitro.DbdbprofileKey{
		d.Get("name").(string),
	}
	return key
}

func create_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dbdbprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_dbdbprofile(d)
	key := resource.ToKey()

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

	resource := get_dbdbprofile(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.DbdbprofileUpdate{}
	unset := nitro.DbdbprofileUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("interpretquery") {
		updateFlag = true

		value := d.Get("interpretquery").(string)
		update.Interpretquery = value

		if value == "" {
			unsetFlag = true

			unset.Interpretquery = true
		}

	}
	if d.HasChange("stickiness") {
		updateFlag = true

		value := d.Get("stickiness").(string)
		update.Stickiness = value

		if value == "" {
			unsetFlag = true

			unset.Stickiness = true
		}

	}
	if d.HasChange("kcdaccount") {
		updateFlag = true

		value := d.Get("kcdaccount").(string)
		update.Kcdaccount = value

		if value == "" {
			unsetFlag = true

			unset.Kcdaccount = true
		}

	}
	if d.HasChange("conmultiplex") {
		updateFlag = true

		value := d.Get("conmultiplex").(string)
		update.Conmultiplex = value

		if value == "" {
			unsetFlag = true

			unset.Conmultiplex = true
		}

	}
	if d.HasChange("enablecachingconmuxoff") {
		updateFlag = true

		value := d.Get("enablecachingconmuxoff").(string)
		update.Enablecachingconmuxoff = value

		if value == "" {
			unsetFlag = true

			unset.Enablecachingconmuxoff = true
		}

	}
	key := get_dbdbprofile_key(d)

	if updateFlag {
		if err := client.UpdateDbdbprofile(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetDbdbprofile(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetDbdbprofile(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_dbdbprofile(d, resource)
	}

	return nil
}

func delete_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dbdbprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_dbdbprofile(d)
	key := resource.ToKey()

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
