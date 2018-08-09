package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_lbgroup(d *schema.ResourceData) nitro.Lbgroup {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbgroup{
		Backuppersistencetimeout: d.Get("backuppersistencetimeout").(int),
		Cookiedomain:             d.Get("cookiedomain").(string),
		Cookiename:               d.Get("cookiename").(string),
		Name:                     d.Get("name").(string),
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

	d.Set("backuppersistencetimeout", resource.Backuppersistencetimeout)
	d.Set("cookiedomain", resource.Cookiedomain)
	d.Set("cookiename", resource.Cookiename)
	d.Set("name", resource.Name)
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

func get_lbgroup_key(d *schema.ResourceData) nitro.LbgroupKey {

	key := nitro.LbgroupKey{
		d.Get("name").(string),
	}
	return key
}

func create_lbgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbgroup")

	client := meta.(*nitro.NitroClient)

	resource := get_lbgroup(d)
	key := resource.ToKey()

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

	resource := get_lbgroup(d)
	key := resource.ToKey()

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

	update := nitro.LbgroupUpdate{}
	unset := nitro.LbgroupUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("persistencetype") {
		updateFlag = true

		value := d.Get("persistencetype").(string)
		update.Persistencetype = value

		if value == "" {
			unsetFlag = true

			unset.Persistencetype = true
		}

	}
	if d.HasChange("persistencebackup") {
		updateFlag = true

		value := d.Get("persistencebackup").(string)
		update.Persistencebackup = value

		if value == "" {
			unsetFlag = true

			unset.Persistencebackup = true
		}

	}
	if d.HasChange("backuppersistencetimeout") {
		updateFlag = true

		value := d.Get("backuppersistencetimeout").(int)
		update.Backuppersistencetimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Backuppersistencetimeout = true
		}

	}
	if d.HasChange("persistmask") {
		updateFlag = true

		value := d.Get("persistmask").(string)
		update.Persistmask = value

		if value == "" {
			unsetFlag = true

			unset.Persistmask = true
		}

	}
	if d.HasChange("cookiename") {
		updateFlag = true

		value := d.Get("cookiename").(string)
		update.Cookiename = value

		if value == "" {
			unsetFlag = true

			unset.Cookiename = true
		}

	}
	if d.HasChange("v6persistmasklen") {
		updateFlag = true

		value := d.Get("v6persistmasklen").(int)
		update.V6persistmasklen = value

		if value == 0 {
			unsetFlag = true

			unset.V6persistmasklen = true
		}

	}
	if d.HasChange("cookiedomain") {
		updateFlag = true

		value := d.Get("cookiedomain").(string)
		update.Cookiedomain = value

		if value == "" {
			unsetFlag = true

			unset.Cookiedomain = true
		}

	}
	if d.HasChange("timeout") {
		updateFlag = true

		value := d.Get("timeout").(int)
		update.Timeout = value

		if value == 0 {
			unsetFlag = true

			unset.Timeout = true
		}

	}
	if d.HasChange("rule") {
		updateFlag = true

		value := d.Get("rule").(string)
		update.Rule = value

		if value == "" {
			unsetFlag = true

			unset.Rule = true
		}

	}
	if d.HasChange("usevserverpersistency") {
		updateFlag = true

		value := d.Get("usevserverpersistency").(string)
		update.Usevserverpersistency = value

		if value == "" {
			unsetFlag = true

			unset.Usevserverpersistency = true
		}

	}
	key := get_lbgroup_key(d)

	if updateFlag {
		if err := client.UpdateLbgroup(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetLbgroup(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetLbgroup(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_lbgroup(d, resource)
	}

	return nil
}

func delete_lbgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbgroup")

	client := meta.(*nitro.NitroClient)

	resource := get_lbgroup(d)
	key := resource.ToKey()

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
