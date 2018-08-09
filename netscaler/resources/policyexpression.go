package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPolicyexpression() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policyexpression,
		Read:          read_policyexpression,
		Update:        update_policyexpression,
		Delete:        delete_policyexpression,
		Schema: map[string]*schema.Schema{
			"clientsecuritymessage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"description": &schema.Schema{
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
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_policyexpression(d *schema.ResourceData) nitro.Policyexpression {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Policyexpression{
		Clientsecuritymessage: d.Get("clientsecuritymessage").(string),
		Comment:               d.Get("comment").(string),
		Description:           d.Get("description").(string),
		Name:                  d.Get("name").(string),
		Value:                 d.Get("value").(string),
	}

	return resource
}

func set_policyexpression(d *schema.ResourceData, resource *nitro.Policyexpression) {
	var _ = strconv.Itoa

	d.Set("clientsecuritymessage", resource.Clientsecuritymessage)
	d.Set("comment", resource.Comment)
	d.Set("description", resource.Description)
	d.Set("name", resource.Name)
	d.Set("value", resource.Value)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_policyexpression_key(d *schema.ResourceData) nitro.PolicyexpressionKey {

	key := nitro.PolicyexpressionKey{
		d.Get("name").(string),
	}
	return key
}

func create_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policyexpression")

	client := meta.(*nitro.NitroClient)

	resource := get_policyexpression(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicyexpression(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicyexpression(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_policyexpression(d, resource)
	} else {
		err := client.AddPolicyexpression(get_policyexpression(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPolicyexpression(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_policyexpression(d, resource)
	}

	return nil
}

func read_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policyexpression")

	client := meta.(*nitro.NitroClient)

	resource := get_policyexpression(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicyexpression(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicyexpression(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_policyexpression(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_policyexpression")

	client := meta.(*nitro.NitroClient)

	update := nitro.PolicyexpressionUpdate{}
	unset := nitro.PolicyexpressionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("value") {
		updateFlag = true

		value := d.Get("value").(string)
		update.Value = value

		if value == "" {
			unsetFlag = true

			unset.Value = true
		}

	}
	if d.HasChange("description") {
		updateFlag = true

		value := d.Get("description").(string)
		update.Description = value

		if value == "" {
			unsetFlag = true

			unset.Description = true
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
	if d.HasChange("clientsecuritymessage") {
		updateFlag = true

		value := d.Get("clientsecuritymessage").(string)
		update.Clientsecuritymessage = value

		if value == "" {
			unsetFlag = true

			unset.Clientsecuritymessage = true
		}

	}
	key := get_policyexpression_key(d)

	if updateFlag {
		if err := client.UpdatePolicyexpression(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetPolicyexpression(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetPolicyexpression(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_policyexpression(d, resource)
	}

	return nil
}

func delete_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policyexpression")

	client := meta.(*nitro.NitroClient)

	resource := get_policyexpression(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicyexpression(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePolicyexpression(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
