package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerFilteraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_filteraction,
		Read:          read_filteraction,
		Update:        update_filteraction,
		Delete:        delete_filteraction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"qual": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"respcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_filteraction(d *schema.ResourceData) nitro.Filteraction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Filteraction{
		Name:        d.Get("name").(string),
		Page:        d.Get("page").(string),
		Qual:        d.Get("qual").(string),
		Respcode:    d.Get("respcode").(int),
		Servicename: d.Get("servicename").(string),
		Value:       d.Get("value").(string),
	}

	return resource
}

func set_filteraction(d *schema.ResourceData, resource *nitro.Filteraction) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("name", resource.Name)
	d.Set("page", resource.Page)
	d.Set("qual", resource.Qual)
	d.Set("respcode", resource.Respcode)
	d.Set("servicename", resource.Servicename)
	d.Set("value", resource.Value)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_filteraction_key(d *schema.ResourceData) nitro.FilteractionKey {

	key := nitro.FilteractionKey{
		d.Get("name").(string),
	}
	return key
}

func create_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_filteraction")

	client := meta.(*nitro.NitroClient)

	resource := get_filteraction(d)
	key := resource.ToKey()

	exists, err := client.ExistsFilteraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilteraction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_filteraction(d, resource)
	} else {
		err := client.AddFilteraction(get_filteraction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetFilteraction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_filteraction(d, resource)
	}

	return nil
}

func read_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_filteraction")

	client := meta.(*nitro.NitroClient)

	resource := get_filteraction(d)
	key := resource.ToKey()

	exists, err := client.ExistsFilteraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilteraction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_filteraction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_filteraction")

	client := meta.(*nitro.NitroClient)

	update := nitro.FilteractionUpdate{}
	unset := nitro.FilteractionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("servicename") {
		updateFlag = true

		value := d.Get("servicename").(string)
		update.Servicename = value

		if value == "" {
			unsetFlag = true

			unset.Servicename = true
		}

	}
	if d.HasChange("value") {
		updateFlag = true

		value := d.Get("value").(string)
		update.Value = value

		if value == "" {
			unsetFlag = true

			unset.Value = true
		}

	}
	if d.HasChange("respcode") {
		updateFlag = true

		value := d.Get("respcode").(int)
		update.Respcode = value

		if value == 0 {
			unsetFlag = true

			unset.Respcode = true
		}

	}
	if d.HasChange("page") {
		updateFlag = true

		value := d.Get("page").(string)
		update.Page = value

		if value == "" {
			unsetFlag = true

			unset.Page = true
		}

	}
	key := get_filteraction_key(d)

	if updateFlag {
		if err := client.UpdateFilteraction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetFilteraction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetFilteraction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_filteraction(d, resource)
	}

	return nil
}

func delete_filteraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_filteraction")

	client := meta.(*nitro.NitroClient)

	resource := get_filteraction(d)
	key := resource.ToKey()

	exists, err := client.ExistsFilteraction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteFilteraction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
