package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTmtrafficaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_tmtrafficaction,
		Read:          read_tmtrafficaction,
		Update:        update_tmtrafficaction,
		Delete:        delete_tmtrafficaction,
		Schema: map[string]*schema.Schema{
			"apptimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"forcedtimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"forcedtimeoutval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"formssoaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"initiatelogout": &schema.Schema{
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
			"passwdexpression": &schema.Schema{
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
			"samlssoprofile": &schema.Schema{
				Type:     schema.TypeString,
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
			"userexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_tmtrafficaction(d *schema.ResourceData) nitro.Tmtrafficaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Tmtrafficaction{
		Apptimeout:       d.Get("apptimeout").(int),
		Forcedtimeout:    d.Get("forcedtimeout").(string),
		Forcedtimeoutval: d.Get("forcedtimeoutval").(int),
		Formssoaction:    d.Get("formssoaction").(string),
		Initiatelogout:   d.Get("initiatelogout").(string),
		Kcdaccount:       d.Get("kcdaccount").(string),
		Name:             d.Get("name").(string),
		Passwdexpression: d.Get("passwdexpression").(string),
		Persistentcookie: d.Get("persistentcookie").(string),
		Samlssoprofile:   d.Get("samlssoprofile").(string),
		Sso:              d.Get("sso").(string),
		Userexpression:   d.Get("userexpression").(string),
	}

	return resource
}

func set_tmtrafficaction(d *schema.ResourceData, resource *nitro.Tmtrafficaction) {
	var _ = strconv.Itoa

	d.Set("apptimeout", resource.Apptimeout)
	d.Set("forcedtimeout", resource.Forcedtimeout)
	d.Set("forcedtimeoutval", resource.Forcedtimeoutval)
	d.Set("formssoaction", resource.Formssoaction)
	d.Set("initiatelogout", resource.Initiatelogout)
	d.Set("kcdaccount", resource.Kcdaccount)
	d.Set("name", resource.Name)
	d.Set("passwdexpression", resource.Passwdexpression)
	d.Set("persistentcookie", resource.Persistentcookie)
	d.Set("samlssoprofile", resource.Samlssoprofile)
	d.Set("sso", resource.Sso)
	d.Set("userexpression", resource.Userexpression)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_tmtrafficaction_key(d *schema.ResourceData) nitro.TmtrafficactionKey {

	key := nitro.TmtrafficactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_tmtrafficaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmtrafficaction")

	client := meta.(*nitro.NitroClient)

	resource := get_tmtrafficaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsTmtrafficaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmtrafficaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_tmtrafficaction(d, resource)
	} else {
		err := client.AddTmtrafficaction(get_tmtrafficaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTmtrafficaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_tmtrafficaction(d, resource)
	}

	return nil
}

func read_tmtrafficaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_tmtrafficaction")

	client := meta.(*nitro.NitroClient)

	resource := get_tmtrafficaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsTmtrafficaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmtrafficaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_tmtrafficaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_tmtrafficaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_tmtrafficaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.TmtrafficactionUpdate{}
	unset := nitro.TmtrafficactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("apptimeout") {
		updateFlag = true

		value := d.Get("apptimeout").(int)
		update.Apptimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Apptimeout = true
		}

	}
	if d.HasChange("sso") {
		updateFlag = true

		value := d.Get("sso").(string)
		update.Sso = value

		if value == "" {
			unsetFlag = true

			unset.Sso = true
		}

	}
	if d.HasChange("formssoaction") {
		updateFlag = true

		value := d.Get("formssoaction").(string)
		update.Formssoaction = value

		if value == "" {
			unsetFlag = true

			unset.Formssoaction = true
		}

	}
	if d.HasChange("persistentcookie") {
		updateFlag = true

		value := d.Get("persistentcookie").(string)
		update.Persistentcookie = value

		if value == "" {
			unsetFlag = true

			unset.Persistentcookie = true
		}

	}
	if d.HasChange("initiatelogout") {
		updateFlag = true

		value := d.Get("initiatelogout").(string)
		update.Initiatelogout = value

		if value == "" {
			unsetFlag = true

			unset.Initiatelogout = true
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
	if d.HasChange("samlssoprofile") {
		updateFlag = true

		value := d.Get("samlssoprofile").(string)
		update.Samlssoprofile = value

		if value == "" {
			unsetFlag = true

			unset.Samlssoprofile = true
		}

	}
	if d.HasChange("forcedtimeout") {
		updateFlag = true

		value := d.Get("forcedtimeout").(string)
		update.Forcedtimeout = value

		if value == "" {
			unsetFlag = true

			unset.Forcedtimeout = true
		}

	}
	if d.HasChange("forcedtimeoutval") {
		updateFlag = true

		value := d.Get("forcedtimeoutval").(int)
		update.Forcedtimeoutval = value

		if value == 0 {
			unsetFlag = true

			unset.Forcedtimeoutval = true
		}

	}
	if d.HasChange("userexpression") {
		updateFlag = true

		value := d.Get("userexpression").(string)
		update.Userexpression = value

		if value == "" {
			unsetFlag = true

			unset.Userexpression = true
		}

	}
	if d.HasChange("passwdexpression") {
		updateFlag = true

		value := d.Get("passwdexpression").(string)
		update.Passwdexpression = value

		if value == "" {
			unsetFlag = true

			unset.Passwdexpression = true
		}

	}
	key := get_tmtrafficaction_key(d)

	if updateFlag {
		if err := client.UpdateTmtrafficaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetTmtrafficaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetTmtrafficaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_tmtrafficaction(d, resource)
	}

	return nil
}

func delete_tmtrafficaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmtrafficaction")

	client := meta.(*nitro.NitroClient)

	resource := get_tmtrafficaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsTmtrafficaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTmtrafficaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
