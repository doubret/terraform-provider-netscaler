package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerDospolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dospolicy,
		Read:          read_dospolicy,
		Update:        update_dospolicy,
		Delete:        delete_dospolicy,
		Schema: map[string]*schema.Schema{
			"cltdetectrate": &schema.Schema{
				Type:     schema.TypeInt,
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
			"qdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_dospolicy(d *schema.ResourceData) nitro.Dospolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dospolicy{
		Cltdetectrate: d.Get("cltdetectrate").(int),
		Name:          d.Get("name").(string),
		Qdepth:        d.Get("qdepth").(int),
	}

	return resource
}

func set_dospolicy(d *schema.ResourceData, resource *nitro.Dospolicy) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("cltdetectrate", resource.Cltdetectrate)
	d.Set("name", resource.Name)
	d.Set("qdepth", resource.Qdepth)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_dospolicy_key(d *schema.ResourceData) nitro.DospolicyKey {

	key := nitro.DospolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dospolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_dospolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsDospolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDospolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_dospolicy(d, resource)
	} else {
		err := client.AddDospolicy(get_dospolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetDospolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_dospolicy(d, resource)
	}

	return nil
}

func read_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dospolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_dospolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsDospolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDospolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_dospolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dospolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.DospolicyUpdate{}
	unset := nitro.DospolicyUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("qdepth") {
		updateFlag = true

		value := d.Get("qdepth").(int)
		update.Qdepth = value

		if value == 0 {
			unsetFlag = true

			unset.Qdepth = true
		}

	}
	if d.HasChange("cltdetectrate") {
		updateFlag = true

		value := d.Get("cltdetectrate").(int)
		update.Cltdetectrate = value

		if value == 0 {
			unsetFlag = true

			unset.Cltdetectrate = true
		}

	}
	key := get_dospolicy_key(d)

	if updateFlag {
		if err := client.UpdateDospolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetDospolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetDospolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_dospolicy(d, resource)
	}

	return nil
}

func delete_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dospolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_dospolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsDospolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteDospolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
