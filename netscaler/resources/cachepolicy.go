package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCachepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cachepolicy,
		Read:          read_cachepolicy,
		Update:        update_cachepolicy,
		Delete:        delete_cachepolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"invalgroups": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"invalobjects": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"storeingroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_cachepolicy(d *schema.ResourceData) nitro.Cachepolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cachepolicy{
		Action:       d.Get("action").(string),
		Invalgroups:  utils.Convert_set_to_string_array(d.Get("invalgroups").(*schema.Set)),
		Invalobjects: utils.Convert_set_to_string_array(d.Get("invalobjects").(*schema.Set)),
		Policyname:   d.Get("policyname").(string),
		Rule:         d.Get("rule").(string),
		Storeingroup: d.Get("storeingroup").(string),
		Undefaction:  d.Get("undefaction").(string),
	}

	return resource
}

func set_cachepolicy(d *schema.ResourceData, resource *nitro.Cachepolicy) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("action", resource.Action)
	d.Set("invalgroups", resource.Invalgroups)
	d.Set("invalobjects", resource.Invalobjects)
	d.Set("policyname", resource.Policyname)
	d.Set("rule", resource.Rule)
	d.Set("storeingroup", resource.Storeingroup)
	d.Set("undefaction", resource.Undefaction)

	var key []string

	key = append(key, resource.Policyname)
	d.SetId(strings.Join(key, "-"))
}

func get_cachepolicy_key(d *schema.ResourceData) nitro.CachepolicyKey {

	key := nitro.CachepolicyKey{
		d.Get("policyname").(string),
	}
	return key
}

func create_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cachepolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cachepolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCachepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachepolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cachepolicy(d, resource)
	} else {
		err := client.AddCachepolicy(get_cachepolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCachepolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cachepolicy(d, resource)
	}

	return nil
}

func read_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cachepolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cachepolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCachepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachepolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cachepolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cachepolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.CachepolicyUpdate{}
	unset := nitro.CachepolicyUnset{}

	updateFlag := false
	unsetFlag := false

	update.Policyname = d.Get("policyname").(string)
	unset.Policyname = d.Get("policyname").(string)

	if d.HasChange("rule") {
		updateFlag = true

		value := d.Get("rule").(string)
		update.Rule = value

		if value == "" {
			unsetFlag = true

			unset.Rule = true
		}

	}
	if d.HasChange("action") {
		updateFlag = true

		value := d.Get("action").(string)
		update.Action = value

		if value == "" {
			unsetFlag = true

			unset.Action = true
		}

	}
	if d.HasChange("storeingroup") {
		updateFlag = true

		value := d.Get("storeingroup").(string)
		update.Storeingroup = value

		if value == "" {
			unsetFlag = true

			unset.Storeingroup = true
		}

	}
	if d.HasChange("invalgroups") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("invalgroups").(*schema.Set))
		update.Invalgroups = value

	}
	if d.HasChange("invalobjects") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("invalobjects").(*schema.Set))
		update.Invalobjects = value

	}
	if d.HasChange("undefaction") {
		updateFlag = true

		value := d.Get("undefaction").(string)
		update.Undefaction = value

		if value == "" {
			unsetFlag = true

			unset.Undefaction = true
		}

	}
	key := get_cachepolicy_key(d)

	if updateFlag {
		if err := client.UpdateCachepolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetCachepolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetCachepolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_cachepolicy(d, resource)
	}

	return nil
}

func delete_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cachepolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_cachepolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsCachepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCachepolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
