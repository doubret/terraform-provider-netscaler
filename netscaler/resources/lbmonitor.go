package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbmonitor() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbmonitor,
		Read:          read_lbmonitor,
		Update:        update_lbmonitor,
		Delete:        delete_lbmonitor,
		Schema: map[string]*schema.Schema{
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"alertretries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"application": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"basedn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"binddn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"customheaders": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"destip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"destport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"deviation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dispatcherip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dispatcherport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"downtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"evalrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"failureretries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"firmwarerevision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"hostipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"httprequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"inbandsecurityid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"iptunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kcdaccount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lasversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"logonpointname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lrtm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxforwards": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mssqlprotocolversion": &schema.Schema{
				Type:     schema.TypeString,
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
			"oraclesid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"originhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"originrealm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"productname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"querytype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radaccountsession": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radaccounttype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radapn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radframedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radmsisdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radnasid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"radnasip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"recv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"resptimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"resptimeoutthresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rtsprequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"scriptargs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"scriptname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"secondarypassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"secure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"send": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sipmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sipreguri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sipuri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sitepath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"snmpcommunity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"snmpoid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"snmpthreshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"snmpversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sqlquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sslprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"storedb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"storefrontacctservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"storefrontcheckbackendservices": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"storename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"successretries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"transparent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"trofscode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"trofsstring": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"units1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"units2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"units3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"units4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"validatecred": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"vendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"vendorspecificvendorid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_lbmonitor(d *schema.ResourceData) nitro.LbmonitorKey {
	key := nitro.LbmonitorKey{
		Monitorname: d.Get("monitorname").(string),
		Type:        d.Get("type").(string),
	}

	return key
}

func get_lbmonitor(d *schema.ResourceData) nitro.Lbmonitor {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbmonitor{
		Monitorname:                    d.Get("monitorname").(string),
		State:                          d.Get("state").(string),
		Action:                         d.Get("action").(string),
		Alertretries:                   d.Get("alertretries").(int),
		Application:                    d.Get("application").(string),
		Attribute:                      d.Get("attribute").(string),
		Basedn:                         d.Get("basedn").(string),
		Binddn:                         d.Get("binddn").(string),
		Customheaders:                  d.Get("customheaders").(string),
		Database:                       d.Get("database").(string),
		Destip:                         d.Get("destip").(string),
		Destport:                       d.Get("destport").(int),
		Deviation:                      d.Get("deviation").(int),
		Dispatcherip:                   d.Get("dispatcherip").(string),
		Dispatcherport:                 d.Get("dispatcherport").(int),
		Domain:                         d.Get("domain").(string),
		Downtime:                       d.Get("downtime").(int),
		Evalrule:                       d.Get("evalrule").(string),
		Failureretries:                 d.Get("failureretries").(int),
		Filename:                       d.Get("filename").(string),
		Filter:                         d.Get("filter").(string),
		Firmwarerevision:               d.Get("firmwarerevision").(int),
		Group:                          d.Get("group").(string),
		Hostipaddress:                  d.Get("hostipaddress").(string),
		Hostname:                       d.Get("hostname").(string),
		Httprequest:                    d.Get("httprequest").(string),
		Inbandsecurityid:               d.Get("inbandsecurityid").(string),
		Interval:                       d.Get("interval").(int),
		Iptunnel:                       d.Get("iptunnel").(string),
		Kcdaccount:                     d.Get("kcdaccount").(string),
		Lasversion:                     d.Get("lasversion").(string),
		Logonpointname:                 d.Get("logonpointname").(string),
		Lrtm:                           d.Get("lrtm").(string),
		Maxforwards:                    d.Get("maxforwards").(int),
		Metrictable:                    d.Get("metrictable").(string),
		Mssqlprotocolversion:           d.Get("mssqlprotocolversion").(string),
		Netprofile:                     d.Get("netprofile").(string),
		Oraclesid:                      d.Get("oraclesid").(string),
		Originhost:                     d.Get("originhost").(string),
		Originrealm:                    d.Get("originrealm").(string),
		Password:                       d.Get("password").(string),
		Productname:                    d.Get("productname").(string),
		Query:                          d.Get("query").(string),
		Querytype:                      d.Get("querytype").(string),
		Radaccountsession:              d.Get("radaccountsession").(string),
		Radaccounttype:                 d.Get("radaccounttype").(int),
		Radapn:                         d.Get("radapn").(string),
		Radframedip:                    d.Get("radframedip").(string),
		Radkey:                         d.Get("radkey").(string),
		Radmsisdn:                      d.Get("radmsisdn").(string),
		Radnasid:                       d.Get("radnasid").(string),
		Radnasip:                       d.Get("radnasip").(string),
		Recv:                           d.Get("recv").(string),
		Resptimeout:                    d.Get("resptimeout").(int),
		Resptimeoutthresh:              d.Get("resptimeoutthresh").(int),
		Retries:                        d.Get("retries").(int),
		Reverse:                        d.Get("reverse").(string),
		Rtsprequest:                    d.Get("rtsprequest").(string),
		Scriptargs:                     d.Get("scriptargs").(string),
		Scriptname:                     d.Get("scriptname").(string),
		Secondarypassword:              d.Get("secondarypassword").(string),
		Secure:                         d.Get("secure").(string),
		Send:                           d.Get("send").(string),
		Sipmethod:                      d.Get("sipmethod").(string),
		Sipreguri:                      d.Get("sipreguri").(string),
		Sipuri:                         d.Get("sipuri").(string),
		Sitepath:                       d.Get("sitepath").(string),
		Snmpcommunity:                  d.Get("snmpcommunity").(string),
		Snmpoid:                        d.Get("snmpoid").(string),
		Snmpthreshold:                  d.Get("snmpthreshold").(string),
		Snmpversion:                    d.Get("snmpversion").(string),
		Sqlquery:                       d.Get("sqlquery").(string),
		Sslprofile:                     d.Get("sslprofile").(string),
		Storedb:                        d.Get("storedb").(string),
		Storefrontacctservice:          d.Get("storefrontacctservice").(string),
		Storefrontcheckbackendservices: d.Get("storefrontcheckbackendservices").(string),
		Storename:                      d.Get("storename").(string),
		Successretries:                 d.Get("successretries").(int),
		Tos:                            d.Get("tos").(string),
		Tosid:                          d.Get("tosid").(int),
		Transparent:                    d.Get("transparent").(string),
		Trofscode:                      d.Get("trofscode").(int),
		Trofsstring:                    d.Get("trofsstring").(string),
		Type:                           d.Get("type").(string),
		Units1:                         d.Get("units1").(string),
		Units2:                         d.Get("units2").(string),
		Units3:                         d.Get("units3").(string),
		Units4:                         d.Get("units4").(string),
		Username:                       d.Get("username").(string),
		Validatecred:                   d.Get("validatecred").(string),
		Vendorid:                       d.Get("vendorid").(int),
		Vendorspecificvendorid:         d.Get("vendorspecificvendorid").(int),
	}

	return resource
}

func set_lbmonitor(d *schema.ResourceData, resource *nitro.Lbmonitor) {
	var _ = strconv.Itoa

	d.Set("monitorname", resource.Monitorname)
	d.Set("state", resource.State)
	d.Set("action", resource.Action)
	d.Set("alertretries", resource.Alertretries)
	d.Set("application", resource.Application)
	d.Set("attribute", resource.Attribute)
	d.Set("basedn", resource.Basedn)
	d.Set("binddn", resource.Binddn)
	d.Set("customheaders", resource.Customheaders)
	d.Set("database", resource.Database)
	d.Set("destip", resource.Destip)
	d.Set("destport", resource.Destport)
	d.Set("deviation", resource.Deviation)
	d.Set("dispatcherip", resource.Dispatcherip)
	d.Set("dispatcherport", resource.Dispatcherport)
	d.Set("domain", resource.Domain)
	d.Set("downtime", resource.Downtime)
	d.Set("evalrule", resource.Evalrule)
	d.Set("failureretries", resource.Failureretries)
	d.Set("filename", resource.Filename)
	d.Set("filter", resource.Filter)
	d.Set("firmwarerevision", resource.Firmwarerevision)
	d.Set("group", resource.Group)
	d.Set("hostipaddress", resource.Hostipaddress)
	d.Set("hostname", resource.Hostname)
	d.Set("httprequest", resource.Httprequest)
	d.Set("inbandsecurityid", resource.Inbandsecurityid)
	d.Set("interval", resource.Interval)
	d.Set("iptunnel", resource.Iptunnel)
	d.Set("kcdaccount", resource.Kcdaccount)
	d.Set("lasversion", resource.Lasversion)
	d.Set("logonpointname", resource.Logonpointname)
	d.Set("lrtm", resource.Lrtm)
	d.Set("maxforwards", resource.Maxforwards)
	d.Set("metrictable", resource.Metrictable)
	d.Set("mssqlprotocolversion", resource.Mssqlprotocolversion)
	d.Set("netprofile", resource.Netprofile)
	d.Set("oraclesid", resource.Oraclesid)
	d.Set("originhost", resource.Originhost)
	d.Set("originrealm", resource.Originrealm)
	d.Set("password", resource.Password)
	d.Set("productname", resource.Productname)
	d.Set("query", resource.Query)
	d.Set("querytype", resource.Querytype)
	d.Set("radaccountsession", resource.Radaccountsession)
	d.Set("radaccounttype", resource.Radaccounttype)
	d.Set("radapn", resource.Radapn)
	d.Set("radframedip", resource.Radframedip)
	d.Set("radkey", resource.Radkey)
	d.Set("radmsisdn", resource.Radmsisdn)
	d.Set("radnasid", resource.Radnasid)
	d.Set("radnasip", resource.Radnasip)
	d.Set("recv", resource.Recv)
	d.Set("resptimeout", resource.Resptimeout)
	d.Set("resptimeoutthresh", resource.Resptimeoutthresh)
	d.Set("retries", resource.Retries)
	d.Set("reverse", resource.Reverse)
	d.Set("rtsprequest", resource.Rtsprequest)
	d.Set("scriptargs", resource.Scriptargs)
	d.Set("scriptname", resource.Scriptname)
	d.Set("secondarypassword", resource.Secondarypassword)
	d.Set("secure", resource.Secure)
	d.Set("send", resource.Send)
	d.Set("sipmethod", resource.Sipmethod)
	d.Set("sipreguri", resource.Sipreguri)
	d.Set("sipuri", resource.Sipuri)
	d.Set("sitepath", resource.Sitepath)
	d.Set("snmpcommunity", resource.Snmpcommunity)
	d.Set("snmpoid", resource.Snmpoid)
	d.Set("snmpthreshold", resource.Snmpthreshold)
	d.Set("snmpversion", resource.Snmpversion)
	d.Set("sqlquery", resource.Sqlquery)
	d.Set("sslprofile", resource.Sslprofile)
	d.Set("storedb", resource.Storedb)
	d.Set("storefrontacctservice", resource.Storefrontacctservice)
	d.Set("storefrontcheckbackendservices", resource.Storefrontcheckbackendservices)
	d.Set("storename", resource.Storename)
	d.Set("successretries", resource.Successretries)
	d.Set("tos", resource.Tos)
	d.Set("tosid", resource.Tosid)
	d.Set("transparent", resource.Transparent)
	d.Set("trofscode", resource.Trofscode)
	d.Set("trofsstring", resource.Trofsstring)
	d.Set("type", resource.Type)
	d.Set("units1", resource.Units1)
	d.Set("units2", resource.Units2)
	d.Set("units3", resource.Units3)
	d.Set("units4", resource.Units4)
	d.Set("username", resource.Username)
	d.Set("validatecred", resource.Validatecred)
	d.Set("vendorid", resource.Vendorid)
	d.Set("vendorspecificvendorid", resource.Vendorspecificvendorid)

	var key []string

	key = append(key, resource.Monitorname)
	key = append(key, resource.Type)
	d.SetId(strings.Join(key, "-"))
}

func create_lbmonitor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmonitor")

	client := meta.(*nitro.NitroClient)

	key := key_lbmonitor(d)

	exists, err := client.ExistsLbmonitor(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmonitor(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbmonitor(d, resource)
	} else {
		err := client.AddLbmonitor(get_lbmonitor(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbmonitor(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbmonitor(d, resource)
	}

	return nil
}

func read_lbmonitor(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbmonitor")

	client := meta.(*nitro.NitroClient)

	key := key_lbmonitor(d)

	exists, err := client.ExistsLbmonitor(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmonitor(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbmonitor(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_lbmonitor(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbmonitor")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateLbmonitor(get_lbmonitor(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_lbmonitor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmonitor")

	client := meta.(*nitro.NitroClient)

	key := key_lbmonitor(d)

	exists, err := client.ExistsLbmonitor(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbmonitor(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
