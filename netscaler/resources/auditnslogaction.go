package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAuditnslogaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_auditnslogaction,
		Read:          read_auditnslogaction,
		Update:        update_auditnslogaction,
		Delete:        delete_auditnslogaction,
		Schema: map[string]*schema.Schema{
			"acl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"alg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"appflowexport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dateformat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"domainresolveretry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"logfacility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"loglevel": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lsn": &schema.Schema{
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
			"serverdomainname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"serverip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"serverport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sslinterception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"subscriberlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"userdefinedauditlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_auditnslogaction(d *schema.ResourceData) nitro.Auditnslogaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Auditnslogaction{
		Acl:                 d.Get("acl").(string),
		Alg:                 d.Get("alg").(string),
		Appflowexport:       d.Get("appflowexport").(string),
		Dateformat:          d.Get("dateformat").(string),
		Domainresolveretry:  d.Get("domainresolveretry").(int),
		Logfacility:         d.Get("logfacility").(string),
		Loglevel:            utils.Convert_set_to_string_array(d.Get("loglevel").(*schema.Set)),
		Lsn:                 d.Get("lsn").(string),
		Name:                d.Get("name").(string),
		Serverdomainname:    d.Get("serverdomainname").(string),
		Serverip:            d.Get("serverip").(string),
		Serverport:          d.Get("serverport").(int),
		Sslinterception:     d.Get("sslinterception").(string),
		Subscriberlog:       d.Get("subscriberlog").(string),
		Tcp:                 d.Get("tcp").(string),
		Timezone:            d.Get("timezone").(string),
		Userdefinedauditlog: d.Get("userdefinedauditlog").(string),
	}

	return resource
}

func set_auditnslogaction(d *schema.ResourceData, resource *nitro.Auditnslogaction) {
	var _ = strconv.Itoa

	d.Set("acl", resource.Acl)
	d.Set("alg", resource.Alg)
	d.Set("appflowexport", resource.Appflowexport)
	d.Set("dateformat", resource.Dateformat)
	d.Set("domainresolveretry", resource.Domainresolveretry)
	d.Set("logfacility", resource.Logfacility)
	d.Set("loglevel", resource.Loglevel)
	d.Set("lsn", resource.Lsn)
	d.Set("name", resource.Name)
	d.Set("serverdomainname", resource.Serverdomainname)
	d.Set("serverip", resource.Serverip)
	d.Set("serverport", resource.Serverport)
	d.Set("sslinterception", resource.Sslinterception)
	d.Set("subscriberlog", resource.Subscriberlog)
	d.Set("tcp", resource.Tcp)
	d.Set("timezone", resource.Timezone)
	d.Set("userdefinedauditlog", resource.Userdefinedauditlog)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_auditnslogaction_key(d *schema.ResourceData) nitro.AuditnslogactionKey {

	key := nitro.AuditnslogactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_auditnslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_auditnslogaction")

	client := meta.(*nitro.NitroClient)

	resource := get_auditnslogaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuditnslogaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditnslogaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_auditnslogaction(d, resource)
	} else {
		err := client.AddAuditnslogaction(get_auditnslogaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAuditnslogaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_auditnslogaction(d, resource)
	}

	return nil
}

func read_auditnslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_auditnslogaction")

	client := meta.(*nitro.NitroClient)

	resource := get_auditnslogaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuditnslogaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditnslogaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_auditnslogaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_auditnslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_auditnslogaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.AuditnslogactionUpdate{}
	unset := nitro.AuditnslogactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("serverip") {
		updateFlag = true

		value := d.Get("serverip").(string)
		update.Serverip = value

		if value == "" {
			unsetFlag = true

			unset.Serverip = true
		}

	}
	if d.HasChange("serverdomainname") {
		updateFlag = true

		value := d.Get("serverdomainname").(string)
		update.Serverdomainname = value

		if value == "" {
			unsetFlag = true

			unset.Serverdomainname = true
		}

	}
	if d.HasChange("domainresolveretry") {
		updateFlag = true

		value := d.Get("domainresolveretry").(int)
		update.Domainresolveretry = value

		if value == 0 {
			unsetFlag = true

			unset.Domainresolveretry = true
		}

	}
	if d.HasChange("serverport") {
		updateFlag = true

		value := d.Get("serverport").(int)
		update.Serverport = value

		if value == 0 {
			unsetFlag = true

			unset.Serverport = true
		}

	}
	if d.HasChange("loglevel") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("loglevel").(*schema.Set))
		update.Loglevel = value

	}
	if d.HasChange("dateformat") {
		updateFlag = true

		value := d.Get("dateformat").(string)
		update.Dateformat = value

		if value == "" {
			unsetFlag = true

			unset.Dateformat = true
		}

	}
	if d.HasChange("logfacility") {
		updateFlag = true

		value := d.Get("logfacility").(string)
		update.Logfacility = value

		if value == "" {
			unsetFlag = true

			unset.Logfacility = true
		}

	}
	if d.HasChange("tcp") {
		updateFlag = true

		value := d.Get("tcp").(string)
		update.Tcp = value

		if value == "" {
			unsetFlag = true

			unset.Tcp = true
		}

	}
	if d.HasChange("acl") {
		updateFlag = true

		value := d.Get("acl").(string)
		update.Acl = value

		if value == "" {
			unsetFlag = true

			unset.Acl = true
		}

	}
	if d.HasChange("timezone") {
		updateFlag = true

		value := d.Get("timezone").(string)
		update.Timezone = value

		if value == "" {
			unsetFlag = true

			unset.Timezone = true
		}

	}
	if d.HasChange("userdefinedauditlog") {
		updateFlag = true

		value := d.Get("userdefinedauditlog").(string)
		update.Userdefinedauditlog = value

		if value == "" {
			unsetFlag = true

			unset.Userdefinedauditlog = true
		}

	}
	if d.HasChange("appflowexport") {
		updateFlag = true

		value := d.Get("appflowexport").(string)
		update.Appflowexport = value

		if value == "" {
			unsetFlag = true

			unset.Appflowexport = true
		}

	}
	if d.HasChange("lsn") {
		updateFlag = true

		value := d.Get("lsn").(string)
		update.Lsn = value

		if value == "" {
			unsetFlag = true

			unset.Lsn = true
		}

	}
	if d.HasChange("alg") {
		updateFlag = true

		value := d.Get("alg").(string)
		update.Alg = value

		if value == "" {
			unsetFlag = true

			unset.Alg = true
		}

	}
	if d.HasChange("subscriberlog") {
		updateFlag = true

		value := d.Get("subscriberlog").(string)
		update.Subscriberlog = value

		if value == "" {
			unsetFlag = true

			unset.Subscriberlog = true
		}

	}
	if d.HasChange("sslinterception") {
		updateFlag = true

		value := d.Get("sslinterception").(string)
		update.Sslinterception = value

		if value == "" {
			unsetFlag = true

			unset.Sslinterception = true
		}

	}
	key := get_auditnslogaction_key(d)

	if updateFlag {
		if err := client.UpdateAuditnslogaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAuditnslogaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAuditnslogaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_auditnslogaction(d, resource)
	}

	return nil
}

func delete_auditnslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_auditnslogaction")

	client := meta.(*nitro.NitroClient)

	resource := get_auditnslogaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuditnslogaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAuditnslogaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
