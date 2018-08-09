package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPqpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_pqpolicy,
		Read:          read_pqpolicy,
		Update:        update_pqpolicy,
		Delete:        delete_pqpolicy,
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"polqdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"qdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_pqpolicy(d *schema.ResourceData) nitro.Pqpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Pqpolicy{
		Policyname: d.Get("policyname").(string),
		Polqdepth:  d.Get("polqdepth").(int),
		Priority:   d.Get("priority").(int),
		Qdepth:     d.Get("qdepth").(int),
		Rule:       d.Get("rule").(string),
		Weight:     d.Get("weight").(int),
	}

	return resource
}

func set_pqpolicy(d *schema.ResourceData, resource *nitro.Pqpolicy) {
	var _ = strconv.Itoa

	d.Set("policyname", resource.Policyname)
	d.Set("polqdepth", resource.Polqdepth)
	d.Set("priority", resource.Priority)
	d.Set("qdepth", resource.Qdepth)
	d.Set("rule", resource.Rule)
	d.Set("weight", resource.Weight)

	var key []string

	key = append(key, resource.Policyname)
	d.SetId(strings.Join(key, "-"))
}

func get_pqpolicy_key(d *schema.ResourceData) nitro.PqpolicyKey {

	key := nitro.PqpolicyKey{
		d.Get("policyname").(string),
	}
	return key
}

func create_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_pqpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_pqpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsPqpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPqpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_pqpolicy(d, resource)
	} else {
		err := client.AddPqpolicy(get_pqpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPqpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_pqpolicy(d, resource)
	}

	return nil
}

func read_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_pqpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_pqpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsPqpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPqpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_pqpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_pqpolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.PqpolicyUpdate{}
	unset := nitro.PqpolicyUnset{}

	updateFlag := false
	unsetFlag := false

	update.Policyname = d.Get("policyname").(string)
	unset.Policyname = d.Get("policyname").(string)

	if d.HasChange("weight") {
		updateFlag = true

		value := d.Get("weight").(int)
		update.Weight = value

		if value == 0 {
			unsetFlag = true

			unset.Weight = true
		}

	}
	if d.HasChange("qdepth") {
		updateFlag = true

		value := d.Get("qdepth").(int)
		update.Qdepth = value

		if value == 0 {
			unsetFlag = true

			unset.Qdepth = true
		}

	}
	if d.HasChange("polqdepth") {
		updateFlag = true

		value := d.Get("polqdepth").(int)
		update.Polqdepth = value

		if value == 0 {
			unsetFlag = true

			unset.Polqdepth = true
		}

	}
	key := get_pqpolicy_key(d)

	if updateFlag {
		if err := client.UpdatePqpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetPqpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetPqpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_pqpolicy(d, resource)
	}

	return nil
}

func delete_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_pqpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_pqpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsPqpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePqpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
