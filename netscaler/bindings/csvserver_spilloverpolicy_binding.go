package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCsvserverSpilloverpolicyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_csvserver_spilloverpolicy_binding,
		Read:          read_csvserver_spilloverpolicy_binding,
		Update:        nil,
		Delete:        delete_csvserver_spilloverpolicy_binding,
		Schema: map[string]*schema.Schema{
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"targetlbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_csvserver_spilloverpolicy_binding(d *schema.ResourceData) nitro.CsvserverSpilloverpolicyBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.CsvserverSpilloverpolicyBinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   d.Get("name").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Targetlbvserver:        d.Get("targetlbvserver").(string),
	}

	return resource
}

func set_csvserver_spilloverpolicy_binding(d *schema.ResourceData, resource *nitro.CsvserverSpilloverpolicyBinding) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("gotopriorityexpression", resource.Gotopriorityexpression)
	d.Set("invoke", resource.Invoke)
	d.Set("labelname", resource.Labelname)
	d.Set("labeltype", resource.Labeltype)
	d.Set("name", resource.Name)
	d.Set("policyname", resource.Policyname)
	d.Set("priority", resource.Priority)
	d.Set("targetlbvserver", resource.Targetlbvserver)

	var key []string

	key = append(key, resource.Name)
	key = append(key, resource.Policyname)
	d.SetId(strings.Join(key, "-"))
}

func get_csvserver_spilloverpolicy_binding_key(d *schema.ResourceData) nitro.CsvserverSpilloverpolicyBindingKey {

	key := nitro.CsvserverSpilloverpolicyBindingKey{
		d.Get("name").(string),
		d.Get("policyname").(string),
	}
	return key
}

func create_csvserver_spilloverpolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_csvserver_spilloverpolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_csvserver_spilloverpolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsvserverSpilloverpolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsvserverSpilloverpolicyBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_csvserver_spilloverpolicy_binding(d, resource)
	} else {
		err := client.AddCsvserverSpilloverpolicyBinding(get_csvserver_spilloverpolicy_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCsvserverSpilloverpolicyBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_csvserver_spilloverpolicy_binding(d, resource)
	}

	return nil
}

func read_csvserver_spilloverpolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_csvserver_spilloverpolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_csvserver_spilloverpolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsvserverSpilloverpolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsvserverSpilloverpolicyBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_csvserver_spilloverpolicy_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_csvserver_spilloverpolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_csvserver_spilloverpolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_csvserver_spilloverpolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsvserverSpilloverpolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCsvserverSpilloverpolicyBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
