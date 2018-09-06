package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbvserver,
		Read:          read_lbvserver,
		Update:        update_lbvserver,
		Delete:        delete_lbvserver,
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
			"backuplbmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"backuppersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
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
			"bypassaaaa": &schema.Schema{
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
			"connfailover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cookiename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"datalength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dataoffset": &schema.Schema{
				Type:     schema.TypeInt,
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
			"dbslb": &schema.Schema{
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
			"dns64": &schema.Schema{
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
			"hashlength": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"healththreshold": &schema.Schema{
				Type:     schema.TypeInt,
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
			"httpsredirecturl": &schema.Schema{
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
			"lbmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lbprofilename": &schema.Schema{
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
			"m": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"macmoderetainvlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxautoscalemembers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"minautoscalemembers": &schema.Schema{
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
			"netmask": &schema.Schema{
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
			"newservicerequest": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"newservicerequestincrementinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"newservicerequestunit": &schema.Schema{
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
			"persistavpno": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistencebackup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistencetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistmask": &schema.Schema{
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
			"pq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"processlocal": &schema.Schema{
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
			"recursionavailable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"redirectfromport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"redirectportrewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"redirurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"resrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"retainconnectionsoncluster": &schema.Schema{
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
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sc": &schema.Schema{
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
			"sessionless": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"skippersistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
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
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
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
			"trofspersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"v6netmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"v6persistmasklen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
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

func get_lbvserver(d *schema.ResourceData) nitro.Lbvserver {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbvserver{
		Appflowlog:                         d.Get("appflowlog").(string),
		Authentication:                     d.Get("authentication").(string),
		Authenticationhost:                 d.Get("authenticationhost").(string),
		Authn401:                           d.Get("authn401").(string),
		Authnprofile:                       d.Get("authnprofile").(string),
		Authnvsname:                        d.Get("authnvsname").(string),
		Backuplbmethod:                     d.Get("backuplbmethod").(string),
		Backuppersistencetimeout:           d.Get("backuppersistencetimeout").(int),
		Backupvserver:                      d.Get("backupvserver").(string),
		Bypassaaaa:                         d.Get("bypassaaaa").(string),
		Cacheable:                          d.Get("cacheable").(string),
		Clttimeout:                         d.Get("clttimeout").(int),
		Comment:                            d.Get("comment").(string),
		Connfailover:                       d.Get("connfailover").(string),
		Cookiename:                         d.Get("cookiename").(string),
		Datalength:                         d.Get("datalength").(int),
		Dataoffset:                         d.Get("dataoffset").(int),
		Dbprofilename:                      d.Get("dbprofilename").(string),
		Dbslb:                              d.Get("dbslb").(string),
		Disableprimaryondown:               d.Get("disableprimaryondown").(string),
		Dns64:                              d.Get("dns64").(string),
		Dnsprofilename:                     d.Get("dnsprofilename").(string),
		Downstateflush:                     d.Get("downstateflush").(string),
		Hashlength:                         d.Get("hashlength").(int),
		Healththreshold:                    d.Get("healththreshold").(int),
		Httpprofilename:                    d.Get("httpprofilename").(string),
		Httpsredirecturl:                   d.Get("httpsredirecturl").(string),
		Icmpvsrresponse:                    d.Get("icmpvsrresponse").(string),
		Insertvserveripport:                d.Get("insertvserveripport").(string),
		Ipmask:                             d.Get("ipmask").(string),
		Ippattern:                          d.Get("ippattern").(string),
		Ipv46:                              d.Get("ipv46").(string),
		L2conn:                             d.Get("l2conn").(string),
		Lbmethod:                           d.Get("lbmethod").(string),
		Lbprofilename:                      d.Get("lbprofilename").(string),
		Listenpolicy:                       d.Get("listenpolicy").(string),
		Listenpriority:                     d.Get("listenpriority").(int),
		M:                                  d.Get("m").(string),
		Macmoderetainvlan:                  d.Get("macmoderetainvlan").(string),
		Maxautoscalemembers:                d.Get("maxautoscalemembers").(int),
		Minautoscalemembers:                d.Get("minautoscalemembers").(int),
		Mssqlserverversion:                 d.Get("mssqlserverversion").(string),
		Mysqlcharacterset:                  d.Get("mysqlcharacterset").(int),
		Mysqlprotocolversion:               d.Get("mysqlprotocolversion").(int),
		Mysqlservercapabilities:            d.Get("mysqlservercapabilities").(int),
		Mysqlserverversion:                 d.Get("mysqlserverversion").(string),
		Name:                               d.Get("name").(string),
		Netmask:                            d.Get("netmask").(string),
		Netprofile:                         d.Get("netprofile").(string),
		Newservicerequest:                  d.Get("newservicerequest").(int),
		Newservicerequestincrementinterval: d.Get("newservicerequestincrementinterval").(int),
		Newservicerequestunit:              d.Get("newservicerequestunit").(string),
		Oracleserverversion:                d.Get("oracleserverversion").(string),
		Persistavpno:                       utils.Convert_set_to_int_array(d.Get("persistavpno").(*schema.Set)),
		Persistencebackup:                  d.Get("persistencebackup").(string),
		Persistencetype:                    d.Get("persistencetype").(string),
		Persistmask:                        d.Get("persistmask").(string),
		Port:                               d.Get("port").(int),
		Pq:                                 d.Get("pq").(string),
		Processlocal:                       d.Get("processlocal").(string),
		Push:                               d.Get("push").(string),
		Pushlabel:                          d.Get("pushlabel").(string),
		Pushmulticlients:                   d.Get("pushmulticlients").(string),
		Pushvserver:                        d.Get("pushvserver").(string),
		Range:                              d.Get("range").(int),
		Recursionavailable:                 d.Get("recursionavailable").(string),
		Redirectfromport:                   d.Get("redirectfromport").(int),
		Redirectportrewrite:                d.Get("redirectportrewrite").(string),
		Redirurl:                           d.Get("redirurl").(string),
		Resrule:                            d.Get("resrule").(string),
		Retainconnectionsoncluster:         d.Get("retainconnectionsoncluster").(string),
		Rhistate:                           d.Get("rhistate").(string),
		Rtspnat:                            d.Get("rtspnat").(string),
		Rule:                               d.Get("rule").(string),
		Sc:                                 d.Get("sc").(string),
		Servicetype:                        d.Get("servicetype").(string),
		Sessionless:                        d.Get("sessionless").(string),
		Skippersistency:                    d.Get("skippersistency").(string),
		Sobackupaction:                     d.Get("sobackupaction").(string),
		Somethod:                           d.Get("somethod").(string),
		Sopersistence:                      d.Get("sopersistence").(string),
		Sopersistencetimeout:               d.Get("sopersistencetimeout").(int),
		Sothreshold:                        d.Get("sothreshold").(int),
		Tcpprofilename:                     d.Get("tcpprofilename").(string),
		Td:                                 d.Get("td").(int),
		Timeout:                            d.Get("timeout").(int),
		Tosid:                              d.Get("tosid").(int),
		Trofspersistence:                   d.Get("trofspersistence").(string),
		V6netmasklen:                       d.Get("v6netmasklen").(int),
		V6persistmasklen:                   d.Get("v6persistmasklen").(int),
		Vipheader:                          d.Get("vipheader").(string),
	}

	return resource
}

func set_lbvserver(d *schema.ResourceData, resource *nitro.Lbvserver) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("appflowlog", resource.Appflowlog)
	d.Set("authentication", resource.Authentication)
	d.Set("authenticationhost", resource.Authenticationhost)
	d.Set("authn401", resource.Authn401)
	d.Set("authnprofile", resource.Authnprofile)
	d.Set("authnvsname", resource.Authnvsname)
	d.Set("backuplbmethod", resource.Backuplbmethod)
	d.Set("backuppersistencetimeout", resource.Backuppersistencetimeout)
	d.Set("backupvserver", resource.Backupvserver)
	d.Set("bypassaaaa", resource.Bypassaaaa)
	d.Set("cacheable", resource.Cacheable)
	d.Set("clttimeout", resource.Clttimeout)
	d.Set("comment", resource.Comment)
	d.Set("connfailover", resource.Connfailover)
	d.Set("cookiename", resource.Cookiename)
	d.Set("datalength", resource.Datalength)
	d.Set("dataoffset", resource.Dataoffset)
	d.Set("dbprofilename", resource.Dbprofilename)
	d.Set("dbslb", resource.Dbslb)
	d.Set("disableprimaryondown", resource.Disableprimaryondown)
	d.Set("dns64", resource.Dns64)
	d.Set("dnsprofilename", resource.Dnsprofilename)
	d.Set("downstateflush", resource.Downstateflush)
	d.Set("hashlength", resource.Hashlength)
	d.Set("healththreshold", resource.Healththreshold)
	d.Set("httpprofilename", resource.Httpprofilename)
	d.Set("httpsredirecturl", resource.Httpsredirecturl)
	d.Set("icmpvsrresponse", resource.Icmpvsrresponse)
	d.Set("insertvserveripport", resource.Insertvserveripport)
	d.Set("ipmask", resource.Ipmask)
	d.Set("ippattern", resource.Ippattern)
	d.Set("ipv46", resource.Ipv46)
	d.Set("l2conn", resource.L2conn)
	d.Set("lbmethod", resource.Lbmethod)
	d.Set("lbprofilename", resource.Lbprofilename)
	d.Set("listenpolicy", resource.Listenpolicy)
	d.Set("listenpriority", resource.Listenpriority)
	d.Set("m", resource.M)
	d.Set("macmoderetainvlan", resource.Macmoderetainvlan)
	d.Set("maxautoscalemembers", resource.Maxautoscalemembers)
	d.Set("minautoscalemembers", resource.Minautoscalemembers)
	d.Set("mssqlserverversion", resource.Mssqlserverversion)
	d.Set("mysqlcharacterset", resource.Mysqlcharacterset)
	d.Set("mysqlprotocolversion", resource.Mysqlprotocolversion)
	d.Set("mysqlservercapabilities", resource.Mysqlservercapabilities)
	d.Set("mysqlserverversion", resource.Mysqlserverversion)
	d.Set("name", resource.Name)
	d.Set("netmask", resource.Netmask)
	d.Set("netprofile", resource.Netprofile)
	d.Set("newservicerequest", resource.Newservicerequest)
	d.Set("newservicerequestincrementinterval", resource.Newservicerequestincrementinterval)
	d.Set("newservicerequestunit", resource.Newservicerequestunit)
	d.Set("oracleserverversion", resource.Oracleserverversion)
	d.Set("persistavpno", resource.Persistavpno)
	d.Set("persistencebackup", resource.Persistencebackup)
	d.Set("persistencetype", resource.Persistencetype)
	d.Set("persistmask", resource.Persistmask)
	d.Set("port", resource.Port)
	d.Set("pq", resource.Pq)
	d.Set("processlocal", resource.Processlocal)
	d.Set("push", resource.Push)
	d.Set("pushlabel", resource.Pushlabel)
	d.Set("pushmulticlients", resource.Pushmulticlients)
	d.Set("pushvserver", resource.Pushvserver)
	d.Set("range", resource.Range)
	d.Set("recursionavailable", resource.Recursionavailable)
	d.Set("redirectfromport", resource.Redirectfromport)
	d.Set("redirectportrewrite", resource.Redirectportrewrite)
	d.Set("redirurl", resource.Redirurl)
	d.Set("resrule", resource.Resrule)
	d.Set("retainconnectionsoncluster", resource.Retainconnectionsoncluster)
	d.Set("rhistate", resource.Rhistate)
	d.Set("rtspnat", resource.Rtspnat)
	d.Set("rule", resource.Rule)
	d.Set("sc", resource.Sc)
	d.Set("servicetype", resource.Servicetype)
	d.Set("sessionless", resource.Sessionless)
	d.Set("skippersistency", resource.Skippersistency)
	d.Set("sobackupaction", resource.Sobackupaction)
	d.Set("somethod", resource.Somethod)
	d.Set("sopersistence", resource.Sopersistence)
	d.Set("sopersistencetimeout", resource.Sopersistencetimeout)
	d.Set("sothreshold", resource.Sothreshold)
	d.Set("tcpprofilename", resource.Tcpprofilename)
	d.Set("td", resource.Td)
	d.Set("timeout", resource.Timeout)
	d.Set("tosid", resource.Tosid)
	d.Set("trofspersistence", resource.Trofspersistence)
	d.Set("v6netmasklen", resource.V6netmasklen)
	d.Set("v6persistmasklen", resource.V6persistmasklen)
	d.Set("vipheader", resource.Vipheader)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_lbvserver_key(d *schema.ResourceData) nitro.LbvserverKey {

	key := nitro.LbvserverKey{
		d.Get("name").(string),
	}
	return key
}

func create_lbvserver(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbvserver")

	client := meta.(*nitro.NitroClient)

	resource := get_lbvserver(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbvserver(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbvserver(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbvserver(d, resource)
	} else {
		err := client.AddLbvserver(get_lbvserver(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbvserver(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbvserver(d, resource)
	}

	return nil
}

func read_lbvserver(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbvserver")

	client := meta.(*nitro.NitroClient)

	resource := get_lbvserver(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbvserver(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbvserver(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbvserver(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_lbvserver(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbvserver")

	client := meta.(*nitro.NitroClient)

	update := nitro.LbvserverUpdate{}
	unset := nitro.LbvserverUnset{}

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
	if d.HasChange("persistencetype") {
		updateFlag = true

		value := d.Get("persistencetype").(string)
		update.Persistencetype = value

		if value == "" {
			unsetFlag = true

			unset.Persistencetype = true
		}

	}
	if d.HasChange("timeout") {
		updateFlag = true

		value := d.Get("timeout").(int)
		update.Timeout = value

		if value == 0 {
			unsetFlag = true

			unset.Timeout = true
		}

	}
	if d.HasChange("persistencebackup") {
		updateFlag = true

		value := d.Get("persistencebackup").(string)
		update.Persistencebackup = value

		if value == "" {
			unsetFlag = true

			unset.Persistencebackup = true
		}

	}
	if d.HasChange("backuppersistencetimeout") {
		updateFlag = true

		value := d.Get("backuppersistencetimeout").(int)
		update.Backuppersistencetimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Backuppersistencetimeout = true
		}

	}
	if d.HasChange("lbmethod") {
		updateFlag = true

		value := d.Get("lbmethod").(string)
		update.Lbmethod = value

		if value == "" {
			unsetFlag = true

			unset.Lbmethod = true
		}

	}
	if d.HasChange("hashlength") {
		updateFlag = true

		value := d.Get("hashlength").(int)
		update.Hashlength = value

		if value == 0 {
			unsetFlag = true

			unset.Hashlength = true
		}

	}
	if d.HasChange("netmask") {
		updateFlag = true

		value := d.Get("netmask").(string)
		update.Netmask = value

		if value == "" {
			unsetFlag = true

			unset.Netmask = true
		}

	}
	if d.HasChange("v6netmasklen") {
		updateFlag = true

		value := d.Get("v6netmasklen").(int)
		update.V6netmasklen = value

		if value == 0 {
			unsetFlag = true

			unset.V6netmasklen = true
		}

	}
	if d.HasChange("backuplbmethod") {
		updateFlag = true

		value := d.Get("backuplbmethod").(string)
		update.Backuplbmethod = value

		if value == "" {
			unsetFlag = true

			unset.Backuplbmethod = true
		}

	}
	if d.HasChange("rule") {
		updateFlag = true

		value := d.Get("rule").(string)
		update.Rule = value

		if value == "" {
			unsetFlag = true

			unset.Rule = true
		}

	}
	if d.HasChange("cookiename") {
		updateFlag = true

		value := d.Get("cookiename").(string)
		update.Cookiename = value

		if value == "" {
			unsetFlag = true

			unset.Cookiename = true
		}

	}
	if d.HasChange("resrule") {
		updateFlag = true

		value := d.Get("resrule").(string)
		update.Resrule = value

		if value == "" {
			unsetFlag = true

			unset.Resrule = true
		}

	}
	if d.HasChange("persistmask") {
		updateFlag = true

		value := d.Get("persistmask").(string)
		update.Persistmask = value

		if value == "" {
			unsetFlag = true

			unset.Persistmask = true
		}

	}
	if d.HasChange("v6persistmasklen") {
		updateFlag = true

		value := d.Get("v6persistmasklen").(int)
		update.V6persistmasklen = value

		if value == 0 {
			unsetFlag = true

			unset.V6persistmasklen = true
		}

	}
	if d.HasChange("pq") {
		updateFlag = true

		value := d.Get("pq").(string)
		update.Pq = value

		if value == "" {
			unsetFlag = true

			unset.Pq = true
		}

	}
	if d.HasChange("sc") {
		updateFlag = true

		value := d.Get("sc").(string)
		update.Sc = value

		if value == "" {
			unsetFlag = true

			unset.Sc = true
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
	if d.HasChange("m") {
		updateFlag = true

		value := d.Get("m").(string)
		update.M = value

		if value == "" {
			unsetFlag = true

			unset.M = true
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
	if d.HasChange("datalength") {
		updateFlag = true

		value := d.Get("datalength").(int)
		update.Datalength = value

		if value == 0 {
			unsetFlag = true

			unset.Datalength = true
		}

	}
	if d.HasChange("dataoffset") {
		updateFlag = true

		value := d.Get("dataoffset").(int)
		update.Dataoffset = value

		if value == 0 {
			unsetFlag = true

			unset.Dataoffset = true
		}

	}
	if d.HasChange("sessionless") {
		updateFlag = true

		value := d.Get("sessionless").(string)
		update.Sessionless = value

		if value == "" {
			unsetFlag = true

			unset.Sessionless = true
		}

	}
	if d.HasChange("trofspersistence") {
		updateFlag = true

		value := d.Get("trofspersistence").(string)
		update.Trofspersistence = value

		if value == "" {
			unsetFlag = true

			unset.Trofspersistence = true
		}

	}
	if d.HasChange("connfailover") {
		updateFlag = true

		value := d.Get("connfailover").(string)
		update.Connfailover = value

		if value == "" {
			unsetFlag = true

			unset.Connfailover = true
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
	if d.HasChange("redirurl") {
		updateFlag = true

		value := d.Get("redirurl").(string)
		update.Redirurl = value

		if value == "" {
			unsetFlag = true

			unset.Redirurl = true
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
	if d.HasChange("sothreshold") {
		updateFlag = true

		value := d.Get("sothreshold").(int)
		update.Sothreshold = value

		if value == 0 {
			unsetFlag = true

			unset.Sothreshold = true
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
	if d.HasChange("healththreshold") {
		updateFlag = true

		value := d.Get("healththreshold").(int)
		update.Healththreshold = value

		if value == 0 {
			unsetFlag = true

			unset.Healththreshold = true
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
	if d.HasChange("disableprimaryondown") {
		updateFlag = true

		value := d.Get("disableprimaryondown").(string)
		update.Disableprimaryondown = value

		if value == "" {
			unsetFlag = true

			unset.Disableprimaryondown = true
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
	if d.HasChange("oracleserverversion") {
		updateFlag = true

		value := d.Get("oracleserverversion").(string)
		update.Oracleserverversion = value

		if value == "" {
			unsetFlag = true

			unset.Oracleserverversion = true
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
	if d.HasChange("newservicerequest") {
		updateFlag = true

		value := d.Get("newservicerequest").(int)
		update.Newservicerequest = value

		if value == 0 {
			unsetFlag = true

			unset.Newservicerequest = true
		}

	}
	if d.HasChange("newservicerequestunit") {
		updateFlag = true

		value := d.Get("newservicerequestunit").(string)
		update.Newservicerequestunit = value

		if value == "" {
			unsetFlag = true

			unset.Newservicerequestunit = true
		}

	}
	if d.HasChange("newservicerequestincrementinterval") {
		updateFlag = true

		value := d.Get("newservicerequestincrementinterval").(int)
		update.Newservicerequestincrementinterval = value

		if value == 0 {
			unsetFlag = true

			unset.Newservicerequestincrementinterval = true
		}

	}
	if d.HasChange("minautoscalemembers") {
		updateFlag = true

		value := d.Get("minautoscalemembers").(int)
		update.Minautoscalemembers = value

		if value == 0 {
			unsetFlag = true

			unset.Minautoscalemembers = true
		}

	}
	if d.HasChange("maxautoscalemembers") {
		updateFlag = true

		value := d.Get("maxautoscalemembers").(int)
		update.Maxautoscalemembers = value

		if value == 0 {
			unsetFlag = true

			unset.Maxautoscalemembers = true
		}

	}
	if d.HasChange("persistavpno") {
		updateFlag = true

		value := utils.Convert_set_to_int_array(d.Get("persistavpno").(*schema.Set))
		update.Persistavpno = value

	}
	if d.HasChange("skippersistency") {
		updateFlag = true

		value := d.Get("skippersistency").(string)
		update.Skippersistency = value

		if value == "" {
			unsetFlag = true

			unset.Skippersistency = true
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
	if d.HasChange("macmoderetainvlan") {
		updateFlag = true

		value := d.Get("macmoderetainvlan").(string)
		update.Macmoderetainvlan = value

		if value == "" {
			unsetFlag = true

			unset.Macmoderetainvlan = true
		}

	}
	if d.HasChange("dbslb") {
		updateFlag = true

		value := d.Get("dbslb").(string)
		update.Dbslb = value

		if value == "" {
			unsetFlag = true

			unset.Dbslb = true
		}

	}
	if d.HasChange("dns64") {
		updateFlag = true

		value := d.Get("dns64").(string)
		update.Dns64 = value

		if value == "" {
			unsetFlag = true

			unset.Dns64 = true
		}

	}
	if d.HasChange("bypassaaaa") {
		updateFlag = true

		value := d.Get("bypassaaaa").(string)
		update.Bypassaaaa = value

		if value == "" {
			unsetFlag = true

			unset.Bypassaaaa = true
		}

	}
	if d.HasChange("recursionavailable") {
		updateFlag = true

		value := d.Get("recursionavailable").(string)
		update.Recursionavailable = value

		if value == "" {
			unsetFlag = true

			unset.Recursionavailable = true
		}

	}
	if d.HasChange("processlocal") {
		updateFlag = true

		value := d.Get("processlocal").(string)
		update.Processlocal = value

		if value == "" {
			unsetFlag = true

			unset.Processlocal = true
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
	if d.HasChange("lbprofilename") {
		updateFlag = true

		value := d.Get("lbprofilename").(string)
		update.Lbprofilename = value

		if value == "" {
			unsetFlag = true

			unset.Lbprofilename = true
		}

	}
	if d.HasChange("redirectfromport") {
		updateFlag = true

		value := d.Get("redirectfromport").(int)
		update.Redirectfromport = value

		if value == 0 {
			unsetFlag = true

			unset.Redirectfromport = true
		}

	}
	if d.HasChange("httpsredirecturl") {
		updateFlag = true

		value := d.Get("httpsredirecturl").(string)
		update.Httpsredirecturl = value

		if value == "" {
			unsetFlag = true

			unset.Httpsredirecturl = true
		}

	}
	if d.HasChange("retainconnectionsoncluster") {
		updateFlag = true

		value := d.Get("retainconnectionsoncluster").(string)
		update.Retainconnectionsoncluster = value

		if value == "" {
			unsetFlag = true

			unset.Retainconnectionsoncluster = true
		}

	}
	key := get_lbvserver_key(d)

	if updateFlag {
		if err := client.UpdateLbvserver(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetLbvserver(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetLbvserver(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_lbvserver(d, resource)
	}

	return nil
}

func delete_lbvserver(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbvserver")

	client := meta.(*nitro.NitroClient)

	resource := get_lbvserver(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbvserver(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbvserver(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
