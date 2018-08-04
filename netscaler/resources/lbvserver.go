package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerLbvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbvserver,
		Read:          read_lbvserver,
		Update:        update_lbvserver,
		Delete:        delete_lbvserver,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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

func key_lbvserver(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_lbvserver(d *schema.ResourceData) nitro.Lbvserver {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbvserver{
		Name:                               d.Get("name").(string),
		State:                              d.Get("state").(string),
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
	d.Set("name", resource.Name)
	d.Set("state", resource.State)
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
	d.SetId(resource.Name)
}

func create_lbvserver(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbvserver")

	client := meta.(*nitro.NitroClient)

	key := key_lbvserver(d)

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

	key := key_lbvserver(d)

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

	return nil
}

func delete_lbvserver(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbvserver")

	client := meta.(*nitro.NitroClient)

	key := key_lbvserver(d)

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
