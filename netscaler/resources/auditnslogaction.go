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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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

func key_auditnslogaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_auditnslogaction(d *schema.ResourceData) nitro.Auditnslogaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Auditnslogaction{
		Name:                d.Get("name").(string),
		Acl:                 d.Get("acl").(string),
		Alg:                 d.Get("alg").(string),
		Appflowexport:       d.Get("appflowexport").(string),
		Dateformat:          d.Get("dateformat").(string),
		Domainresolveretry:  d.Get("domainresolveretry").(int),
		Logfacility:         d.Get("logfacility").(string),
		Loglevel:            utils.Convert_set_to_string_array(d.Get("loglevel").(*schema.Set)),
		Lsn:                 d.Get("lsn").(string),
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

	d.Set("name", resource.Name)
	d.Set("acl", resource.Acl)
	d.Set("alg", resource.Alg)
	d.Set("appflowexport", resource.Appflowexport)
	d.Set("dateformat", resource.Dateformat)
	d.Set("domainresolveretry", resource.Domainresolveretry)
	d.Set("logfacility", resource.Logfacility)
	d.Set("loglevel", resource.Loglevel)
	d.Set("lsn", resource.Lsn)
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

func create_auditnslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_auditnslogaction")

	client := meta.(*nitro.NitroClient)

	key := key_auditnslogaction(d)

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

	key := key_auditnslogaction(d)

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

	err := client.UpdateAuditnslogaction(get_auditnslogaction(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_auditnslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_auditnslogaction")

	client := meta.(*nitro.NitroClient)

	key := key_auditnslogaction(d)

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
