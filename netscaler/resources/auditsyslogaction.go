package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAuditsyslogaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_auditsyslogaction,
		Read:          read_auditsyslogaction,
		Update:        update_auditsyslogaction,
		Delete:        delete_auditsyslogaction,
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
			"dns": &schema.Schema{
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
			"lbvservername": &schema.Schema{
				Type:     schema.TypeString,
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
			"maxlogdatasizetohold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"netprofile": &schema.Schema{
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
			"tcpprofilename": &schema.Schema{
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
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func key_auditsyslogaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_auditsyslogaction(d *schema.ResourceData) nitro.Auditsyslogaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Auditsyslogaction{
		Name:                 d.Get("name").(string),
		Acl:                  d.Get("acl").(string),
		Alg:                  d.Get("alg").(string),
		Appflowexport:        d.Get("appflowexport").(string),
		Dateformat:           d.Get("dateformat").(string),
		Dns:                  d.Get("dns").(string),
		Domainresolveretry:   d.Get("domainresolveretry").(int),
		Lbvservername:        d.Get("lbvservername").(string),
		Logfacility:          d.Get("logfacility").(string),
		Loglevel:             utils.Convert_set_to_string_array(d.Get("loglevel").(*schema.Set)),
		Lsn:                  d.Get("lsn").(string),
		Maxlogdatasizetohold: d.Get("maxlogdatasizetohold").(int),
		Netprofile:           d.Get("netprofile").(string),
		Serverdomainname:     d.Get("serverdomainname").(string),
		Serverip:             d.Get("serverip").(string),
		Serverport:           d.Get("serverport").(int),
		Sslinterception:      d.Get("sslinterception").(string),
		Subscriberlog:        d.Get("subscriberlog").(string),
		Tcp:                  d.Get("tcp").(string),
		Tcpprofilename:       d.Get("tcpprofilename").(string),
		Timezone:             d.Get("timezone").(string),
		Transport:            d.Get("transport").(string),
		Userdefinedauditlog:  d.Get("userdefinedauditlog").(string),
	}

	return resource
}

func set_auditsyslogaction(d *schema.ResourceData, resource *nitro.Auditsyslogaction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("acl", resource.Acl)
	d.Set("alg", resource.Alg)
	d.Set("appflowexport", resource.Appflowexport)
	d.Set("dateformat", resource.Dateformat)
	d.Set("dns", resource.Dns)
	d.Set("domainresolveretry", resource.Domainresolveretry)
	d.Set("lbvservername", resource.Lbvservername)
	d.Set("logfacility", resource.Logfacility)
	d.Set("loglevel", resource.Loglevel)
	d.Set("lsn", resource.Lsn)
	d.Set("maxlogdatasizetohold", resource.Maxlogdatasizetohold)
	d.Set("netprofile", resource.Netprofile)
	d.Set("serverdomainname", resource.Serverdomainname)
	d.Set("serverip", resource.Serverip)
	d.Set("serverport", resource.Serverport)
	d.Set("sslinterception", resource.Sslinterception)
	d.Set("subscriberlog", resource.Subscriberlog)
	d.Set("tcp", resource.Tcp)
	d.Set("tcpprofilename", resource.Tcpprofilename)
	d.Set("timezone", resource.Timezone)
	d.Set("transport", resource.Transport)
	d.Set("userdefinedauditlog", resource.Userdefinedauditlog)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_auditsyslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_auditsyslogaction")

	client := meta.(*nitro.NitroClient)

	key := key_auditsyslogaction(d)

	exists, err := client.ExistsAuditsyslogaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditsyslogaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_auditsyslogaction(d, resource)
	} else {
		err := client.AddAuditsyslogaction(get_auditsyslogaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAuditsyslogaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_auditsyslogaction(d, resource)
	}

	return nil
}

func read_auditsyslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_auditsyslogaction")

	client := meta.(*nitro.NitroClient)

	key := key_auditsyslogaction(d)

	exists, err := client.ExistsAuditsyslogaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditsyslogaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_auditsyslogaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_auditsyslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_auditsyslogaction")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateAuditsyslogaction(get_auditsyslogaction(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_auditsyslogaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_auditsyslogaction")

	client := meta.(*nitro.NitroClient)

	key := key_auditsyslogaction(d)

	exists, err := client.ExistsAuditsyslogaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAuditsyslogaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
