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
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				ForceNew: true,
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

func get_lbmonitor(d *schema.ResourceData) nitro.Lbmonitor {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbmonitor{
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
		Monitorname:                    d.Get("monitorname").(string),
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
		State:                          d.Get("state").(string),
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
	var _ = strconv.FormatBool

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
	d.Set("monitorname", resource.Monitorname)
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
	d.Set("state", resource.State)
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

func get_lbmonitor_key(d *schema.ResourceData) nitro.LbmonitorKey {

	key := nitro.LbmonitorKey{
		d.Get("monitorname").(string),
		d.Get("type").(string),
	}
	return key
}

func create_lbmonitor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmonitor")

	client := meta.(*nitro.NitroClient)

	resource := get_lbmonitor(d)
	key := resource.ToKey()

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

	resource := get_lbmonitor(d)
	key := resource.ToKey()

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

	update := nitro.LbmonitorUpdate{}
	unset := nitro.LbmonitorUnset{}

	updateFlag := false
	unsetFlag := false

	update.Monitorname = d.Get("monitorname").(string)
	unset.Monitorname = d.Get("monitorname").(string)
	update.Type = d.Get("type").(string)
	unset.Type = d.Get("type").(string)

	if d.HasChange("action") {
		updateFlag = true

		value := d.Get("action").(string)
		update.Action = value

		if value == "" {
			unsetFlag = true

			unset.Action = true
		}

	}
	if d.HasChange("httprequest") {
		updateFlag = true

		value := d.Get("httprequest").(string)
		update.Httprequest = value

		if value == "" {
			unsetFlag = true

			unset.Httprequest = true
		}

	}
	if d.HasChange("rtsprequest") {
		updateFlag = true

		value := d.Get("rtsprequest").(string)
		update.Rtsprequest = value

		if value == "" {
			unsetFlag = true

			unset.Rtsprequest = true
		}

	}
	if d.HasChange("customheaders") {
		updateFlag = true

		value := d.Get("customheaders").(string)
		update.Customheaders = value

		if value == "" {
			unsetFlag = true

			unset.Customheaders = true
		}

	}
	if d.HasChange("maxforwards") {
		updateFlag = true

		value := d.Get("maxforwards").(int)
		update.Maxforwards = value

		if value == 0 {
			unsetFlag = true

			unset.Maxforwards = true
		}

	}
	if d.HasChange("sipmethod") {
		updateFlag = true

		value := d.Get("sipmethod").(string)
		update.Sipmethod = value

		if value == "" {
			unsetFlag = true

			unset.Sipmethod = true
		}

	}
	if d.HasChange("sipreguri") {
		updateFlag = true

		value := d.Get("sipreguri").(string)
		update.Sipreguri = value

		if value == "" {
			unsetFlag = true

			unset.Sipreguri = true
		}

	}
	if d.HasChange("sipuri") {
		updateFlag = true

		value := d.Get("sipuri").(string)
		update.Sipuri = value

		if value == "" {
			unsetFlag = true

			unset.Sipuri = true
		}

	}
	if d.HasChange("send") {
		updateFlag = true

		value := d.Get("send").(string)
		update.Send = value

		if value == "" {
			unsetFlag = true

			unset.Send = true
		}

	}
	if d.HasChange("recv") {
		updateFlag = true

		value := d.Get("recv").(string)
		update.Recv = value

		if value == "" {
			unsetFlag = true

			unset.Recv = true
		}

	}
	if d.HasChange("query") {
		updateFlag = true

		value := d.Get("query").(string)
		update.Query = value

		if value == "" {
			unsetFlag = true

			unset.Query = true
		}

	}
	if d.HasChange("querytype") {
		updateFlag = true

		value := d.Get("querytype").(string)
		update.Querytype = value

		if value == "" {
			unsetFlag = true

			unset.Querytype = true
		}

	}
	if d.HasChange("username") {
		updateFlag = true

		value := d.Get("username").(string)
		update.Username = value

		if value == "" {
			unsetFlag = true

			unset.Username = true
		}

	}
	if d.HasChange("password") {
		updateFlag = true

		value := d.Get("password").(string)
		update.Password = value

		if value == "" {
			unsetFlag = true

			unset.Password = true
		}

	}
	if d.HasChange("secondarypassword") {
		updateFlag = true

		value := d.Get("secondarypassword").(string)
		update.Secondarypassword = value

		if value == "" {
			unsetFlag = true

			unset.Secondarypassword = true
		}

	}
	if d.HasChange("logonpointname") {
		updateFlag = true

		value := d.Get("logonpointname").(string)
		update.Logonpointname = value

		if value == "" {
			unsetFlag = true

			unset.Logonpointname = true
		}

	}
	if d.HasChange("lasversion") {
		updateFlag = true

		value := d.Get("lasversion").(string)
		update.Lasversion = value

		if value == "" {
			unsetFlag = true

			unset.Lasversion = true
		}

	}
	if d.HasChange("radkey") {
		updateFlag = true

		value := d.Get("radkey").(string)
		update.Radkey = value

		if value == "" {
			unsetFlag = true

			unset.Radkey = true
		}

	}
	if d.HasChange("radnasid") {
		updateFlag = true

		value := d.Get("radnasid").(string)
		update.Radnasid = value

		if value == "" {
			unsetFlag = true

			unset.Radnasid = true
		}

	}
	if d.HasChange("radnasip") {
		updateFlag = true

		value := d.Get("radnasip").(string)
		update.Radnasip = value

		if value == "" {
			unsetFlag = true

			unset.Radnasip = true
		}

	}
	if d.HasChange("radaccounttype") {
		updateFlag = true

		value := d.Get("radaccounttype").(int)
		update.Radaccounttype = value

		if value == 0 {
			unsetFlag = true

			unset.Radaccounttype = true
		}

	}
	if d.HasChange("radframedip") {
		updateFlag = true

		value := d.Get("radframedip").(string)
		update.Radframedip = value

		if value == "" {
			unsetFlag = true

			unset.Radframedip = true
		}

	}
	if d.HasChange("radapn") {
		updateFlag = true

		value := d.Get("radapn").(string)
		update.Radapn = value

		if value == "" {
			unsetFlag = true

			unset.Radapn = true
		}

	}
	if d.HasChange("radmsisdn") {
		updateFlag = true

		value := d.Get("radmsisdn").(string)
		update.Radmsisdn = value

		if value == "" {
			unsetFlag = true

			unset.Radmsisdn = true
		}

	}
	if d.HasChange("radaccountsession") {
		updateFlag = true

		value := d.Get("radaccountsession").(string)
		update.Radaccountsession = value

		if value == "" {
			unsetFlag = true

			unset.Radaccountsession = true
		}

	}
	if d.HasChange("lrtm") {
		updateFlag = true

		value := d.Get("lrtm").(string)
		update.Lrtm = value

		if value == "" {
			unsetFlag = true

			unset.Lrtm = true
		}

	}
	if d.HasChange("deviation") {
		updateFlag = true

		value := d.Get("deviation").(int)
		update.Deviation = value

		if value == 0 {
			unsetFlag = true

			unset.Deviation = true
		}

	}
	if d.HasChange("units1") {
		updateFlag = true

		value := d.Get("units1").(string)
		update.Units1 = value

		if value == "" {
			unsetFlag = true

			unset.Units1 = true
		}

	}
	if d.HasChange("scriptname") {
		updateFlag = true

		value := d.Get("scriptname").(string)
		update.Scriptname = value

		if value == "" {
			unsetFlag = true

			unset.Scriptname = true
		}

	}
	if d.HasChange("scriptargs") {
		updateFlag = true

		value := d.Get("scriptargs").(string)
		update.Scriptargs = value

		if value == "" {
			unsetFlag = true

			unset.Scriptargs = true
		}

	}
	if d.HasChange("validatecred") {
		updateFlag = true

		value := d.Get("validatecred").(string)
		update.Validatecred = value

		if value == "" {
			unsetFlag = true

			unset.Validatecred = true
		}

	}
	if d.HasChange("domain") {
		updateFlag = true

		value := d.Get("domain").(string)
		update.Domain = value

		if value == "" {
			unsetFlag = true

			unset.Domain = true
		}

	}
	if d.HasChange("dispatcherip") {
		updateFlag = true

		value := d.Get("dispatcherip").(string)
		update.Dispatcherip = value

		if value == "" {
			unsetFlag = true

			unset.Dispatcherip = true
		}

	}
	if d.HasChange("dispatcherport") {
		updateFlag = true

		value := d.Get("dispatcherport").(int)
		update.Dispatcherport = value

		if value == 0 {
			unsetFlag = true

			unset.Dispatcherport = true
		}

	}
	if d.HasChange("interval") {
		updateFlag = true

		value := d.Get("interval").(int)
		update.Interval = value

		if value == 0 {
			unsetFlag = true

			unset.Interval = true
		}

	}
	if d.HasChange("units3") {
		updateFlag = true

		value := d.Get("units3").(string)
		update.Units3 = value

		if value == "" {
			unsetFlag = true

			unset.Units3 = true
		}

	}
	if d.HasChange("resptimeout") {
		updateFlag = true

		value := d.Get("resptimeout").(int)
		update.Resptimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Resptimeout = true
		}

	}
	if d.HasChange("units4") {
		updateFlag = true

		value := d.Get("units4").(string)
		update.Units4 = value

		if value == "" {
			unsetFlag = true

			unset.Units4 = true
		}

	}
	if d.HasChange("resptimeoutthresh") {
		updateFlag = true

		value := d.Get("resptimeoutthresh").(int)
		update.Resptimeoutthresh = value

		if value == 0 {
			unsetFlag = true

			unset.Resptimeoutthresh = true
		}

	}
	if d.HasChange("retries") {
		updateFlag = true

		value := d.Get("retries").(int)
		update.Retries = value

		if value == 0 {
			unsetFlag = true

			unset.Retries = true
		}

	}
	if d.HasChange("failureretries") {
		updateFlag = true

		value := d.Get("failureretries").(int)
		update.Failureretries = value

		if value == 0 {
			unsetFlag = true

			unset.Failureretries = true
		}

	}
	if d.HasChange("alertretries") {
		updateFlag = true

		value := d.Get("alertretries").(int)
		update.Alertretries = value

		if value == 0 {
			unsetFlag = true

			unset.Alertretries = true
		}

	}
	if d.HasChange("successretries") {
		updateFlag = true

		value := d.Get("successretries").(int)
		update.Successretries = value

		if value == 0 {
			unsetFlag = true

			unset.Successretries = true
		}

	}
	if d.HasChange("downtime") {
		updateFlag = true

		value := d.Get("downtime").(int)
		update.Downtime = value

		if value == 0 {
			unsetFlag = true

			unset.Downtime = true
		}

	}
	if d.HasChange("units2") {
		updateFlag = true

		value := d.Get("units2").(string)
		update.Units2 = value

		if value == "" {
			unsetFlag = true

			unset.Units2 = true
		}

	}
	if d.HasChange("destip") {
		updateFlag = true

		value := d.Get("destip").(string)
		update.Destip = value

		if value == "" {
			unsetFlag = true

			unset.Destip = true
		}

	}
	if d.HasChange("destport") {
		updateFlag = true

		value := d.Get("destport").(int)
		update.Destport = value

		if value == 0 {
			unsetFlag = true

			unset.Destport = true
		}

	}
	if d.HasChange("reverse") {
		updateFlag = true

		value := d.Get("reverse").(string)
		update.Reverse = value

		if value == "" {
			unsetFlag = true

			unset.Reverse = true
		}

	}
	if d.HasChange("transparent") {
		updateFlag = true

		value := d.Get("transparent").(string)
		update.Transparent = value

		if value == "" {
			unsetFlag = true

			unset.Transparent = true
		}

	}
	if d.HasChange("iptunnel") {
		updateFlag = true

		value := d.Get("iptunnel").(string)
		update.Iptunnel = value

		if value == "" {
			unsetFlag = true

			unset.Iptunnel = true
		}

	}
	if d.HasChange("tos") {
		updateFlag = true

		value := d.Get("tos").(string)
		update.Tos = value

		if value == "" {
			unsetFlag = true

			unset.Tos = true
		}

	}
	if d.HasChange("tosid") {
		updateFlag = true

		value := d.Get("tosid").(int)
		update.Tosid = value

		if value == 0 {
			unsetFlag = true

			unset.Tosid = true
		}

	}
	if d.HasChange("secure") {
		updateFlag = true

		value := d.Get("secure").(string)
		update.Secure = value

		if value == "" {
			unsetFlag = true

			unset.Secure = true
		}

	}
	if d.HasChange("group") {
		updateFlag = true

		value := d.Get("group").(string)
		update.Group = value

		if value == "" {
			unsetFlag = true

			unset.Group = true
		}

	}
	if d.HasChange("filename") {
		updateFlag = true

		value := d.Get("filename").(string)
		update.Filename = value

		if value == "" {
			unsetFlag = true

			unset.Filename = true
		}

	}
	if d.HasChange("basedn") {
		updateFlag = true

		value := d.Get("basedn").(string)
		update.Basedn = value

		if value == "" {
			unsetFlag = true

			unset.Basedn = true
		}

	}
	if d.HasChange("binddn") {
		updateFlag = true

		value := d.Get("binddn").(string)
		update.Binddn = value

		if value == "" {
			unsetFlag = true

			unset.Binddn = true
		}

	}
	if d.HasChange("filter") {
		updateFlag = true

		value := d.Get("filter").(string)
		update.Filter = value

		if value == "" {
			unsetFlag = true

			unset.Filter = true
		}

	}
	if d.HasChange("attribute") {
		updateFlag = true

		value := d.Get("attribute").(string)
		update.Attribute = value

		if value == "" {
			unsetFlag = true

			unset.Attribute = true
		}

	}
	if d.HasChange("database") {
		updateFlag = true

		value := d.Get("database").(string)
		update.Database = value

		if value == "" {
			unsetFlag = true

			unset.Database = true
		}

	}
	if d.HasChange("oraclesid") {
		updateFlag = true

		value := d.Get("oraclesid").(string)
		update.Oraclesid = value

		if value == "" {
			unsetFlag = true

			unset.Oraclesid = true
		}

	}
	if d.HasChange("sqlquery") {
		updateFlag = true

		value := d.Get("sqlquery").(string)
		update.Sqlquery = value

		if value == "" {
			unsetFlag = true

			unset.Sqlquery = true
		}

	}
	if d.HasChange("evalrule") {
		updateFlag = true

		value := d.Get("evalrule").(string)
		update.Evalrule = value

		if value == "" {
			unsetFlag = true

			unset.Evalrule = true
		}

	}
	if d.HasChange("snmpoid") {
		updateFlag = true

		value := d.Get("snmpoid").(string)
		update.Snmpoid = value

		if value == "" {
			unsetFlag = true

			unset.Snmpoid = true
		}

	}
	if d.HasChange("snmpcommunity") {
		updateFlag = true

		value := d.Get("snmpcommunity").(string)
		update.Snmpcommunity = value

		if value == "" {
			unsetFlag = true

			unset.Snmpcommunity = true
		}

	}
	if d.HasChange("snmpthreshold") {
		updateFlag = true

		value := d.Get("snmpthreshold").(string)
		update.Snmpthreshold = value

		if value == "" {
			unsetFlag = true

			unset.Snmpthreshold = true
		}

	}
	if d.HasChange("snmpversion") {
		updateFlag = true

		value := d.Get("snmpversion").(string)
		update.Snmpversion = value

		if value == "" {
			unsetFlag = true

			unset.Snmpversion = true
		}

	}
	if d.HasChange("metrictable") {
		updateFlag = true

		value := d.Get("metrictable").(string)
		update.Metrictable = value

		if value == "" {
			unsetFlag = true

			unset.Metrictable = true
		}

	}
	if d.HasChange("application") {
		updateFlag = true

		value := d.Get("application").(string)
		update.Application = value

		if value == "" {
			unsetFlag = true

			unset.Application = true
		}

	}
	if d.HasChange("sitepath") {
		updateFlag = true

		value := d.Get("sitepath").(string)
		update.Sitepath = value

		if value == "" {
			unsetFlag = true

			unset.Sitepath = true
		}

	}
	if d.HasChange("storename") {
		updateFlag = true

		value := d.Get("storename").(string)
		update.Storename = value

		if value == "" {
			unsetFlag = true

			unset.Storename = true
		}

	}
	if d.HasChange("storefrontacctservice") {
		updateFlag = true

		value := d.Get("storefrontacctservice").(string)
		update.Storefrontacctservice = value

		if value == "" {
			unsetFlag = true

			unset.Storefrontacctservice = true
		}

	}
	if d.HasChange("storefrontcheckbackendservices") {
		updateFlag = true

		value := d.Get("storefrontcheckbackendservices").(string)
		update.Storefrontcheckbackendservices = value

		if value == "" {
			unsetFlag = true

			unset.Storefrontcheckbackendservices = true
		}

	}
	if d.HasChange("hostname") {
		updateFlag = true

		value := d.Get("hostname").(string)
		update.Hostname = value

		if value == "" {
			unsetFlag = true

			unset.Hostname = true
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
	if d.HasChange("mssqlprotocolversion") {
		updateFlag = true

		value := d.Get("mssqlprotocolversion").(string)
		update.Mssqlprotocolversion = value

		if value == "" {
			unsetFlag = true

			unset.Mssqlprotocolversion = true
		}

	}
	if d.HasChange("originhost") {
		updateFlag = true

		value := d.Get("originhost").(string)
		update.Originhost = value

		if value == "" {
			unsetFlag = true

			unset.Originhost = true
		}

	}
	if d.HasChange("originrealm") {
		updateFlag = true

		value := d.Get("originrealm").(string)
		update.Originrealm = value

		if value == "" {
			unsetFlag = true

			unset.Originrealm = true
		}

	}
	if d.HasChange("hostipaddress") {
		updateFlag = true

		value := d.Get("hostipaddress").(string)
		update.Hostipaddress = value

		if value == "" {
			unsetFlag = true

			unset.Hostipaddress = true
		}

	}
	if d.HasChange("vendorid") {
		updateFlag = true

		value := d.Get("vendorid").(int)
		update.Vendorid = value

		if value == 0 {
			unsetFlag = true

			unset.Vendorid = true
		}

	}
	if d.HasChange("productname") {
		updateFlag = true

		value := d.Get("productname").(string)
		update.Productname = value

		if value == "" {
			unsetFlag = true

			unset.Productname = true
		}

	}
	if d.HasChange("firmwarerevision") {
		updateFlag = true

		value := d.Get("firmwarerevision").(int)
		update.Firmwarerevision = value

		if value == 0 {
			unsetFlag = true

			unset.Firmwarerevision = true
		}

	}
	if d.HasChange("vendorspecificvendorid") {
		updateFlag = true

		value := d.Get("vendorspecificvendorid").(int)
		update.Vendorspecificvendorid = value

		if value == 0 {
			unsetFlag = true

			unset.Vendorspecificvendorid = true
		}

	}
	if d.HasChange("kcdaccount") {
		updateFlag = true

		value := d.Get("kcdaccount").(string)
		update.Kcdaccount = value

		if value == "" {
			unsetFlag = true

			unset.Kcdaccount = true
		}

	}
	if d.HasChange("storedb") {
		updateFlag = true

		value := d.Get("storedb").(string)
		update.Storedb = value

		if value == "" {
			unsetFlag = true

			unset.Storedb = true
		}

	}
	if d.HasChange("trofscode") {
		updateFlag = true

		value := d.Get("trofscode").(int)
		update.Trofscode = value

		if value == 0 {
			unsetFlag = true

			unset.Trofscode = true
		}

	}
	if d.HasChange("trofsstring") {
		updateFlag = true

		value := d.Get("trofsstring").(string)
		update.Trofsstring = value

		if value == "" {
			unsetFlag = true

			unset.Trofsstring = true
		}

	}
	if d.HasChange("sslprofile") {
		updateFlag = true

		value := d.Get("sslprofile").(string)
		update.Sslprofile = value

		if value == "" {
			unsetFlag = true

			unset.Sslprofile = true
		}

	}
	key := get_lbmonitor_key(d)

	if updateFlag {
		if err := client.UpdateLbmonitor(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetLbmonitor(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetLbmonitor(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_lbmonitor(d, resource)
	}

	return nil
}

func delete_lbmonitor(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmonitor")

	client := meta.(*nitro.NitroClient)

	resource := get_lbmonitor(d)
	key := resource.ToKey()

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
