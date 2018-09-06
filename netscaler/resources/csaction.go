package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCsaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_csaction,
		Read:          read_csaction,
		Update:        update_csaction,
		Delete:        delete_csaction,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
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
			"targetlbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"targetvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"targetvserverexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_csaction(d *schema.ResourceData) nitro.Csaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Csaction{
		Comment:           d.Get("comment").(string),
		Name:              d.Get("name").(string),
		Targetlbvserver:   d.Get("targetlbvserver").(string),
		Targetvserver:     d.Get("targetvserver").(string),
		Targetvserverexpr: d.Get("targetvserverexpr").(string),
	}

	return resource
}

func set_csaction(d *schema.ResourceData, resource *nitro.Csaction) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("targetlbvserver", resource.Targetlbvserver)
	d.Set("targetvserver", resource.Targetvserver)
	d.Set("targetvserverexpr", resource.Targetvserverexpr)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_csaction_key(d *schema.ResourceData) nitro.CsactionKey {

	key := nitro.CsactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_csaction")

	client := meta.(*nitro.NitroClient)

	resource := get_csaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_csaction(d, resource)
	} else {
		err := client.AddCsaction(get_csaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCsaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_csaction(d, resource)
	}

	return nil
}

func read_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_csaction")

	client := meta.(*nitro.NitroClient)

	resource := get_csaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_csaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_csaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.CsactionUpdate{}
	unset := nitro.CsactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("targetlbvserver") {
		updateFlag = true

		value := d.Get("targetlbvserver").(string)
		update.Targetlbvserver = value

		if value == "" {
			unsetFlag = true

			unset.Targetlbvserver = true
		}

	}
	if d.HasChange("targetvserver") {
		updateFlag = true

		value := d.Get("targetvserver").(string)
		update.Targetvserver = value

		if value == "" {
			unsetFlag = true

			unset.Targetvserver = true
		}

	}
	if d.HasChange("targetvserverexpr") {
		updateFlag = true

		value := d.Get("targetvserverexpr").(string)
		update.Targetvserverexpr = value

		if value == "" {
			unsetFlag = true

			unset.Targetvserverexpr = true
		}

	}
	if d.HasChange("comment") {
		updateFlag = true

		value := d.Get("comment").(string)
		update.Comment = value

		if value == "" {
			unsetFlag = true

			unset.Comment = true
		}

	}
	key := get_csaction_key(d)

	if updateFlag {
		if err := client.UpdateCsaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetCsaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetCsaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_csaction(d, resource)
	}

	return nil
}

func delete_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_csaction")

	client := meta.(*nitro.NitroClient)

	resource := get_csaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCsaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
