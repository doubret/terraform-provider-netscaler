package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppflowpolicylabelAppflowpolicyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowpolicylabel_appflowpolicy_binding,
		Read:          read_appflowpolicylabel_appflowpolicy_binding,
		Update:        nil,
		Delete:        delete_appflowpolicylabel_appflowpolicy_binding,
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
			"invoke_labelname": &schema.Schema{
				Type:     schema.TypeString,
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
		},
	}
}

func get_appflowpolicylabel_appflowpolicy_binding(d *schema.ResourceData) nitro.AppflowpolicylabelAppflowpolicyBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.AppflowpolicylabelAppflowpolicyBinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Invoke_labelname:       d.Get("invoke_labelname").(string),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	return resource
}

func set_appflowpolicylabel_appflowpolicy_binding(d *schema.ResourceData, resource *nitro.AppflowpolicylabelAppflowpolicyBinding) {
	var _ = strconv.Itoa

	d.Set("gotopriorityexpression", resource.Gotopriorityexpression)
	d.Set("invoke", resource.Invoke)
	d.Set("invoke_labelname", resource.Invoke_labelname)
	d.Set("labelname", resource.Labelname)
	d.Set("labeltype", resource.Labeltype)
	d.Set("policyname", resource.Policyname)
	d.Set("priority", resource.Priority)

	var key []string

	key = append(key, resource.Labelname)
	key = append(key, resource.Policyname)
	d.SetId(strings.Join(key, "-"))
}

func create_appflowpolicylabel_appflowpolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowpolicylabel_appflowpolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicylabel_appflowpolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicylabelAppflowpolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicylabelAppflowpolicyBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appflowpolicylabel_appflowpolicy_binding(d, resource)
	} else {
		err := client.AddAppflowpolicylabelAppflowpolicyBinding(get_appflowpolicylabel_appflowpolicy_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppflowpolicylabelAppflowpolicyBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appflowpolicylabel_appflowpolicy_binding(d, resource)
	}

	return nil
}

func read_appflowpolicylabel_appflowpolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowpolicylabel_appflowpolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicylabel_appflowpolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicylabelAppflowpolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicylabelAppflowpolicyBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appflowpolicylabel_appflowpolicy_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_appflowpolicylabel_appflowpolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowpolicylabel_appflowpolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicylabel_appflowpolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicylabelAppflowpolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppflowpolicylabelAppflowpolicyBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
