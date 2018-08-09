package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_tmsessionaction(d *schema.ResourceData) nitro.Tmsessionaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Tmsessionaction{
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Homepage:                   d.Get("homepage").(string),
		Httponlycookie:             d.Get("httponlycookie").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Name:                       d.Get("name").(string),
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

	d.Set("defaultauthorizationaction", resource.Defaultauthorizationaction)
	d.Set("homepage", resource.Homepage)
	d.Set("httponlycookie", resource.Httponlycookie)
	d.Set("kcdaccount", resource.Kcdaccount)
	d.Set("name", resource.Name)
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

func get_tmsessionaction_key(d *schema.ResourceData) nitro.TmsessionactionKey {

	key := nitro.TmsessionactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_tmsessionaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmsessionaction")

	client := meta.(*nitro.NitroClient)

	resource := get_tmsessionaction(d)
	key := resource.ToKey()

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

	resource := get_tmsessionaction(d)
	key := resource.ToKey()

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

	update := nitro.TmsessionactionUpdate{}
	unset := nitro.TmsessionactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("sesstimeout") {
		updateFlag = true

		value := d.Get("sesstimeout").(int)
		update.Sesstimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Sesstimeout = true
		}

	}
	if d.HasChange("defaultauthorizationaction") {
		updateFlag = true

		value := d.Get("defaultauthorizationaction").(string)
		update.Defaultauthorizationaction = value

		if value == "" {
			unsetFlag = true

			unset.Defaultauthorizationaction = true
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
	if d.HasChange("ssocredential") {
		updateFlag = true

		value := d.Get("ssocredential").(string)
		update.Ssocredential = value

		if value == "" {
			unsetFlag = true

			unset.Ssocredential = true
		}

	}
	if d.HasChange("ssodomain") {
		updateFlag = true

		value := d.Get("ssodomain").(string)
		update.Ssodomain = value

		if value == "" {
			unsetFlag = true

			unset.Ssodomain = true
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
	if d.HasChange("httponlycookie") {
		updateFlag = true

		value := d.Get("httponlycookie").(string)
		update.Httponlycookie = value

		if value == "" {
			unsetFlag = true

			unset.Httponlycookie = true
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
	if d.HasChange("persistentcookievalidity") {
		updateFlag = true

		value := d.Get("persistentcookievalidity").(int)
		update.Persistentcookievalidity = value

		if value == 0 {
			unsetFlag = true

			unset.Persistentcookievalidity = true
		}

	}
	if d.HasChange("homepage") {
		updateFlag = true

		value := d.Get("homepage").(string)
		update.Homepage = value

		if value == "" {
			unsetFlag = true

			unset.Homepage = true
		}

	}
	key := get_tmsessionaction_key(d)

	if updateFlag {
		if err := client.UpdateTmsessionaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetTmsessionaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetTmsessionaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_tmsessionaction(d, resource)
	}

	return nil
}

func delete_tmsessionaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmsessionaction")

	client := meta.(*nitro.NitroClient)

	resource := get_tmsessionaction(d)
	key := resource.ToKey()

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
