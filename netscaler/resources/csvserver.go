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

func get_csvserver_key(d *schema.ResourceData) nitro.CsvserverKey {

	key := nitro.CsvserverKey{
		d.Get("name").(string),
	}
	return key
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

	client := meta.(*nitro.NitroClient)

	update := nitro.CsvserverUpdate{}
	unset := nitro.CsvserverUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("ipv46") {
		updateFlag = true

		value := d.Get("ipv46").(string)
		update.Ipv46 = value

		if value == "" {
			unsetFlag = true

			unset.Ipv46 = true
		}

	}
	if d.HasChange("ippattern") {
		updateFlag = true

		value := d.Get("ippattern").(string)
		update.Ippattern = value

		if value == "" {
			unsetFlag = true

			unset.Ippattern = true
		}

	}
	if d.HasChange("ipmask") {
		updateFlag = true

		value := d.Get("ipmask").(string)
		update.Ipmask = value

		if value == "" {
			unsetFlag = true

			unset.Ipmask = true
		}

	}
	if d.HasChange("stateupdate") {
		updateFlag = true

		value := d.Get("stateupdate").(string)
		update.Stateupdate = value

		if value == "" {
			unsetFlag = true

			unset.Stateupdate = true
		}

	}
	if d.HasChange("precedence") {
		updateFlag = true

		value := d.Get("precedence").(string)
		update.Precedence = value

		if value == "" {
			unsetFlag = true

			unset.Precedence = true
		}

	}
	if d.HasChange("casesensitive") {
		updateFlag = true

		value := d.Get("casesensitive").(string)
		update.Casesensitive = value

		if value == "" {
			unsetFlag = true

			unset.Casesensitive = true
		}

	}
	if d.HasChange("backupvserver") {
		updateFlag = true

		value := d.Get("backupvserver").(string)
		update.Backupvserver = value

		if value == "" {
			unsetFlag = true

			unset.Backupvserver = true
		}

	}
	if d.HasChange("redirecturl") {
		updateFlag = true

		value := d.Get("redirecturl").(string)
		update.Redirecturl = value

		if value == "" {
			unsetFlag = true

			unset.Redirecturl = true
		}

	}
	if d.HasChange("cacheable") {
		updateFlag = true

		value := d.Get("cacheable").(string)
		update.Cacheable = value

		if value == "" {
			unsetFlag = true

			unset.Cacheable = true
		}

	}
	if d.HasChange("clttimeout") {
		updateFlag = true

		value := d.Get("clttimeout").(int)
		update.Clttimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Clttimeout = true
		}

	}
	if d.HasChange("somethod") {
		updateFlag = true

		value := d.Get("somethod").(string)
		update.Somethod = value

		if value == "" {
			unsetFlag = true

			unset.Somethod = true
		}

	}
	if d.HasChange("sopersistence") {
		updateFlag = true

		value := d.Get("sopersistence").(string)
		update.Sopersistence = value

		if value == "" {
			unsetFlag = true

			unset.Sopersistence = true
		}

	}
	if d.HasChange("sopersistencetimeout") {
		updateFlag = true

		value := d.Get("sopersistencetimeout").(int)
		update.Sopersistencetimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Sopersistencetimeout = true
		}

	}
	if d.HasChange("sothreshold") {
		updateFlag = true

		value := d.Get("sothreshold").(int)
		update.Sothreshold = value

		if value == 0 {
			unsetFlag = true

			unset.Sothreshold = true
		}

	}
	if d.HasChange("sobackupaction") {
		updateFlag = true

		value := d.Get("sobackupaction").(string)
		update.Sobackupaction = value

		if value == "" {
			unsetFlag = true

			unset.Sobackupaction = true
		}

	}
	if d.HasChange("redirectportrewrite") {
		updateFlag = true

		value := d.Get("redirectportrewrite").(string)
		update.Redirectportrewrite = value

		if value == "" {
			unsetFlag = true

			unset.Redirectportrewrite = true
		}

	}
	if d.HasChange("downstateflush") {
		updateFlag = true

		value := d.Get("downstateflush").(string)
		update.Downstateflush = value

		if value == "" {
			unsetFlag = true

			unset.Downstateflush = true
		}

	}
	if d.HasChange("disableprimaryondown") {
		updateFlag = true

		value := d.Get("disableprimaryondown").(string)
		update.Disableprimaryondown = value

		if value == "" {
			unsetFlag = true

			unset.Disableprimaryondown = true
		}

	}
	if d.HasChange("insertvserveripport") {
		updateFlag = true

		value := d.Get("insertvserveripport").(string)
		update.Insertvserveripport = value

		if value == "" {
			unsetFlag = true

			unset.Insertvserveripport = true
		}

	}
	if d.HasChange("vipheader") {
		updateFlag = true

		value := d.Get("vipheader").(string)
		update.Vipheader = value

		if value == "" {
			unsetFlag = true

			unset.Vipheader = true
		}

	}
	if d.HasChange("rtspnat") {
		updateFlag = true

		value := d.Get("rtspnat").(string)
		update.Rtspnat = value

		if value == "" {
			unsetFlag = true

			unset.Rtspnat = true
		}

	}
	if d.HasChange("authenticationhost") {
		updateFlag = true

		value := d.Get("authenticationhost").(string)
		update.Authenticationhost = value

		if value == "" {
			unsetFlag = true

			unset.Authenticationhost = true
		}

	}
	if d.HasChange("authentication") {
		updateFlag = true

		value := d.Get("authentication").(string)
		update.Authentication = value

		if value == "" {
			unsetFlag = true

			unset.Authentication = true
		}

	}
	if d.HasChange("listenpolicy") {
		updateFlag = true

		value := d.Get("listenpolicy").(string)
		update.Listenpolicy = value

		if value == "" {
			unsetFlag = true

			unset.Listenpolicy = true
		}

	}
	if d.HasChange("listenpriority") {
		updateFlag = true

		value := d.Get("listenpriority").(int)
		update.Listenpriority = value

		if value == 0 {
			unsetFlag = true

			unset.Listenpriority = true
		}

	}
	if d.HasChange("authn401") {
		updateFlag = true

		value := d.Get("authn401").(string)
		update.Authn401 = value

		if value == "" {
			unsetFlag = true

			unset.Authn401 = true
		}

	}
	if d.HasChange("authnvsname") {
		updateFlag = true

		value := d.Get("authnvsname").(string)
		update.Authnvsname = value

		if value == "" {
			unsetFlag = true

			unset.Authnvsname = true
		}

	}
	if d.HasChange("push") {
		updateFlag = true

		value := d.Get("push").(string)
		update.Push = value

		if value == "" {
			unsetFlag = true

			unset.Push = true
		}

	}
	if d.HasChange("pushvserver") {
		updateFlag = true

		value := d.Get("pushvserver").(string)
		update.Pushvserver = value

		if value == "" {
			unsetFlag = true

			unset.Pushvserver = true
		}

	}
	if d.HasChange("pushlabel") {
		updateFlag = true

		value := d.Get("pushlabel").(string)
		update.Pushlabel = value

		if value == "" {
			unsetFlag = true

			unset.Pushlabel = true
		}

	}
	if d.HasChange("pushmulticlients") {
		updateFlag = true

		value := d.Get("pushmulticlients").(string)
		update.Pushmulticlients = value

		if value == "" {
			unsetFlag = true

			unset.Pushmulticlients = true
		}

	}
	if d.HasChange("tcpprofilename") {
		updateFlag = true

		value := d.Get("tcpprofilename").(string)
		update.Tcpprofilename = value

		if value == "" {
			unsetFlag = true

			unset.Tcpprofilename = true
		}

	}
	if d.HasChange("httpprofilename") {
		updateFlag = true

		value := d.Get("httpprofilename").(string)
		update.Httpprofilename = value

		if value == "" {
			unsetFlag = true

			unset.Httpprofilename = true
		}

	}
	if d.HasChange("dbprofilename") {
		updateFlag = true

		value := d.Get("dbprofilename").(string)
		update.Dbprofilename = value

		if value == "" {
			unsetFlag = true

			unset.Dbprofilename = true
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
	if d.HasChange("l2conn") {
		updateFlag = true

		value := d.Get("l2conn").(string)
		update.L2conn = value

		if value == "" {
			unsetFlag = true

			unset.L2conn = true
		}

	}
	if d.HasChange("mssqlserverversion") {
		updateFlag = true

		value := d.Get("mssqlserverversion").(string)
		update.Mssqlserverversion = value

		if value == "" {
			unsetFlag = true

			unset.Mssqlserverversion = true
		}

	}
	if d.HasChange("mysqlprotocolversion") {
		updateFlag = true

		value := d.Get("mysqlprotocolversion").(int)
		update.Mysqlprotocolversion = value

		if value == 0 {
			unsetFlag = true

			unset.Mysqlprotocolversion = true
		}

	}
	if d.HasChange("oracleserverversion") {
		updateFlag = true

		value := d.Get("oracleserverversion").(string)
		update.Oracleserverversion = value

		if value == "" {
			unsetFlag = true

			unset.Oracleserverversion = true
		}

	}
	if d.HasChange("mysqlserverversion") {
		updateFlag = true

		value := d.Get("mysqlserverversion").(string)
		update.Mysqlserverversion = value

		if value == "" {
			unsetFlag = true

			unset.Mysqlserverversion = true
		}

	}
	if d.HasChange("mysqlcharacterset") {
		updateFlag = true

		value := d.Get("mysqlcharacterset").(int)
		update.Mysqlcharacterset = value

		if value == 0 {
			unsetFlag = true

			unset.Mysqlcharacterset = true
		}

	}
	if d.HasChange("mysqlservercapabilities") {
		updateFlag = true

		value := d.Get("mysqlservercapabilities").(int)
		update.Mysqlservercapabilities = value

		if value == 0 {
			unsetFlag = true

			unset.Mysqlservercapabilities = true
		}

	}
	if d.HasChange("appflowlog") {
		updateFlag = true

		value := d.Get("appflowlog").(string)
		update.Appflowlog = value

		if value == "" {
			unsetFlag = true

			unset.Appflowlog = true
		}

	}
	if d.HasChange("netprofile") {
		updateFlag = true

		value := d.Get("netprofile").(string)
		update.Netprofile = value

		if value == "" {
			unsetFlag = true

			unset.Netprofile = true
		}

	}
	if d.HasChange("authnprofile") {
		updateFlag = true

		value := d.Get("authnprofile").(string)
		update.Authnprofile = value

		if value == "" {
			unsetFlag = true

			unset.Authnprofile = true
		}

	}
	if d.HasChange("icmpvsrresponse") {
		updateFlag = true

		value := d.Get("icmpvsrresponse").(string)
		update.Icmpvsrresponse = value

		if value == "" {
			unsetFlag = true

			unset.Icmpvsrresponse = true
		}

	}
	if d.HasChange("rhistate") {
		updateFlag = true

		value := d.Get("rhistate").(string)
		update.Rhistate = value

		if value == "" {
			unsetFlag = true

			unset.Rhistate = true
		}

	}
	if d.HasChange("dnsprofilename") {
		updateFlag = true

		value := d.Get("dnsprofilename").(string)
		update.Dnsprofilename = value

		if value == "" {
			unsetFlag = true

			unset.Dnsprofilename = true
		}

	}
	key := get_csvserver_key(d)

	if updateFlag {
		if err := client.UpdateCsvserver(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetCsvserver(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetCsvserver(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_csvserver(d, resource)
	}

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
