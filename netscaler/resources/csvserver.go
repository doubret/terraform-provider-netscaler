package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCsvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_csvserver,
		Read:          read_csvserver,
		Update:        update_csvserver,
		Delete:        delete_csvserver,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"authenticationhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"authn401": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"authnprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"authnvsname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"backupvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"casesensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
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
			"dbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"disableprimaryondown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"httpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"icmpvsrresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"insertvserveripport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ipmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ippattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ipv46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"l2conn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"listenpolicy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"listenpriority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mssqlserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mysqlcharacterset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mysqlprotocolversion": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mysqlservercapabilities": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mysqlserverversion": &schema.Schema{
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
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"oracleserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"precedence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pushlabel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pushmulticlients": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pushvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"redirectportrewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"redirecturl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rhistate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rtspnat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sobackupaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"somethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sopersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sopersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sothreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"stateupdate": &schema.Schema{
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
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_csvserver(d *schema.ResourceData) nitro.Csvserver {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Csvserver{
		Appflowlog:              d.Get("appflowlog").(string),
		Authentication:          d.Get("authentication").(string),
		Authenticationhost:      d.Get("authenticationhost").(string),
		Authn401:                d.Get("authn401").(string),
		Authnprofile:            d.Get("authnprofile").(string),
		Authnvsname:             d.Get("authnvsname").(string),
		Backupvserver:           d.Get("backupvserver").(string),
		Cacheable:               d.Get("cacheable").(string),
		Casesensitive:           d.Get("casesensitive").(string),
		Clttimeout:              d.Get("clttimeout").(int),
		Comment:                 d.Get("comment").(string),
		Dbprofilename:           d.Get("dbprofilename").(string),
		Disableprimaryondown:    d.Get("disableprimaryondown").(string),
		Dnsprofilename:          d.Get("dnsprofilename").(string),
		Downstateflush:          d.Get("downstateflush").(string),
		Httpprofilename:         d.Get("httpprofilename").(string),
		Icmpvsrresponse:         d.Get("icmpvsrresponse").(string),
		Insertvserveripport:     d.Get("insertvserveripport").(string),
		Ipmask:                  d.Get("ipmask").(string),
		Ippattern:               d.Get("ippattern").(string),
		Ipv46:                   d.Get("ipv46").(string),
		L2conn:                  d.Get("l2conn").(string),
		Listenpolicy:            d.Get("listenpolicy").(string),
		Listenpriority:          d.Get("listenpriority").(int),
		Mssqlserverversion:      d.Get("mssqlserverversion").(string),
		Mysqlcharacterset:       d.Get("mysqlcharacterset").(int),
		Mysqlprotocolversion:    d.Get("mysqlprotocolversion").(int),
		Mysqlservercapabilities: d.Get("mysqlservercapabilities").(int),
		Mysqlserverversion:      d.Get("mysqlserverversion").(string),
		Name:                    d.Get("name").(string),
		Netprofile:              d.Get("netprofile").(string),
		Oracleserverversion:     d.Get("oracleserverversion").(string),
		Port:                    d.Get("port").(int),
		Precedence:              d.Get("precedence").(string),
		Push:                    d.Get("push").(string),
		Pushlabel:               d.Get("pushlabel").(string),
		Pushmulticlients:        d.Get("pushmulticlients").(string),
		Pushvserver:             d.Get("pushvserver").(string),
		Range:                   d.Get("range").(int),
		Redirectportrewrite:     d.Get("redirectportrewrite").(string),
		Redirecturl:             d.Get("redirecturl").(string),
		Rhistate:                d.Get("rhistate").(string),
		Rtspnat:                 d.Get("rtspnat").(string),
		Servicetype:             d.Get("servicetype").(string),
		Sobackupaction:          d.Get("sobackupaction").(string),
		Somethod:                d.Get("somethod").(string),
		Sopersistence:           d.Get("sopersistence").(string),
		Sopersistencetimeout:    d.Get("sopersistencetimeout").(int),
		Sothreshold:             d.Get("sothreshold").(int),
		State:                   d.Get("state").(string),
		Stateupdate:             d.Get("stateupdate").(string),
		Tcpprofilename:          d.Get("tcpprofilename").(string),
		Td:                      d.Get("td").(int),
		Vipheader:               d.Get("vipheader").(string),
	}

	return resource
}

func set_csvserver(d *schema.ResourceData, resource *nitro.Csvserver) {
	var _ = strconv.Itoa

	d.Set("appflowlog", resource.Appflowlog)
	d.Set("authentication", resource.Authentication)
	d.Set("authenticationhost", resource.Authenticationhost)
	d.Set("authn401", resource.Authn401)
	d.Set("authnprofile", resource.Authnprofile)
	d.Set("authnvsname", resource.Authnvsname)
	d.Set("backupvserver", resource.Backupvserver)
	d.Set("cacheable", resource.Cacheable)
	d.Set("casesensitive", resource.Casesensitive)
	d.Set("clttimeout", resource.Clttimeout)
	d.Set("comment", resource.Comment)
	d.Set("dbprofilename", resource.Dbprofilename)
	d.Set("disableprimaryondown", resource.Disableprimaryondown)
	d.Set("dnsprofilename", resource.Dnsprofilename)
	d.Set("downstateflush", resource.Downstateflush)
	d.Set("httpprofilename", resource.Httpprofilename)
	d.Set("icmpvsrresponse", resource.Icmpvsrresponse)
	d.Set("insertvserveripport", resource.Insertvserveripport)
	d.Set("ipmask", resource.Ipmask)
	d.Set("ippattern", resource.Ippattern)
	d.Set("ipv46", resource.Ipv46)
	d.Set("l2conn", resource.L2conn)
	d.Set("listenpolicy", resource.Listenpolicy)
	d.Set("listenpriority", resource.Listenpriority)
	d.Set("mssqlserverversion", resource.Mssqlserverversion)
	d.Set("mysqlcharacterset", resource.Mysqlcharacterset)
	d.Set("mysqlprotocolversion", resource.Mysqlprotocolversion)
	d.Set("mysqlservercapabilities", resource.Mysqlservercapabilities)
	d.Set("mysqlserverversion", resource.Mysqlserverversion)
	d.Set("name", resource.Name)
	d.Set("netprofile", resource.Netprofile)
	d.Set("oracleserverversion", resource.Oracleserverversion)
	d.Set("port", resource.Port)
	d.Set("precedence", resource.Precedence)
	d.Set("push", resource.Push)
	d.Set("pushlabel", resource.Pushlabel)
	d.Set("pushmulticlients", resource.Pushmulticlients)
	d.Set("pushvserver", resource.Pushvserver)
	d.Set("range", resource.Range)
	d.Set("redirectportrewrite", resource.Redirectportrewrite)
	d.Set("redirecturl", resource.Redirecturl)
	d.Set("rhistate", resource.Rhistate)
	d.Set("rtspnat", resource.Rtspnat)
	d.Set("servicetype", resource.Servicetype)
	d.Set("sobackupaction", resource.Sobackupaction)
	d.Set("somethod", resource.Somethod)
	d.Set("sopersistence", resource.Sopersistence)
	d.Set("sopersistencetimeout", resource.Sopersistencetimeout)
	d.Set("sothreshold", resource.Sothreshold)
	d.Set("state", resource.State)
	d.Set("stateupdate", resource.Stateupdate)
	d.Set("tcpprofilename", resource.Tcpprofilename)
	d.Set("td", resource.Td)
	d.Set("vipheader", resource.Vipheader)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_csvserver(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_csvserver")

	client := meta.(*nitro.NitroClient)

	resource := get_csvserver(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsvserver(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsvserver(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_csvserver(d, resource)
	} else {
		err := client.AddCsvserver(get_csvserver(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCsvserver(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_csvserver(d, resource)
	}

	return nil
}

func read_csvserver(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_csvserver")

	client := meta.(*nitro.NitroClient)

	resource := get_csvserver(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsvserver(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsvserver(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_csvserver(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_csvserver(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_csvserver")

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateCsvserver(get_csvserver(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_csvserver(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_csvserver")

	client := meta.(*nitro.NitroClient)

	resource := get_csvserver(d)
	key := resource.ToKey()

	exists, err := client.ExistsCsvserver(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCsvserver(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
