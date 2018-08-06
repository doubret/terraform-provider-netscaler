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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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

func key_tmtrafficaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_tmtrafficaction(d *schema.ResourceData) nitro.Tmtrafficaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Tmtrafficaction{
		Name:             d.Get("name").(string),
		Apptimeout:       d.Get("apptimeout").(int),
		Forcedtimeout:    d.Get("forcedtimeout").(string),
		Forcedtimeoutval: d.Get("forcedtimeoutval").(int),
		Formssoaction:    d.Get("formssoaction").(string),
		Initiatelogout:   d.Get("initiatelogout").(string),
		Kcdaccount:       d.Get("kcdaccount").(string),
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

	d.Set("name", resource.Name)
	d.Set("apptimeout", resource.Apptimeout)
	d.Set("forcedtimeout", resource.Forcedtimeout)
	d.Set("forcedtimeoutval", resource.Forcedtimeoutval)
	d.Set("formssoaction", resource.Formssoaction)
	d.Set("initiatelogout", resource.Initiatelogout)
	d.Set("kcdaccount", resource.Kcdaccount)
	d.Set("passwdexpression", resource.Passwdexpression)
	d.Set("persistentcookie", resource.Persistentcookie)
	d.Set("samlssoprofile", resource.Samlssoprofile)
	d.Set("sso", resource.Sso)
	d.Set("userexpression", resource.Userexpression)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_tmtrafficaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmtrafficaction")

	client := meta.(*nitro.NitroClient)

	key := key_tmtrafficaction(d)

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

	key := key_tmtrafficaction(d)

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

	err := client.UpdateTmtrafficaction(get_tmtrafficaction(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_tmtrafficaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmtrafficaction")

	client := meta.(*nitro.NitroClient)

	key := key_tmtrafficaction(d)

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
