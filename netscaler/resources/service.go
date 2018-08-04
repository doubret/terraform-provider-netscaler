package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerService() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_service,
		Read:          read_service,
		Update:        update_service,
		Delete:        delete_service,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"accessdown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"appflowlog": &schema.Schema{
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
			"cachetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cka": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cleartextport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cmp": &schema.Schema{
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
			"customserverid": &schema.Schema{
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
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"healthmonitor": &schema.Schema{
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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxclient": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxreq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"monconnectionclose": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"monthreshold": &schema.Schema{
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
			"pathmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pathmonitorindv": &schema.Schema{
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
			"processlocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rtspsessionidremap": &schema.Schema{
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
			"serverid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"svrtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcpb": &schema.Schema{
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
			"useproxyport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"usip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_service(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_service(d *schema.ResourceData) nitro.Service {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Service{
		Name:               d.Get("name").(string),
		Accessdown:         d.Get("accessdown").(string),
		Appflowlog:         d.Get("appflowlog").(string),
		Cacheable:          d.Get("cacheable").(string),
		Cachetype:          d.Get("cachetype").(string),
		Cip:                d.Get("cip").(string),
		Cipheader:          d.Get("cipheader").(string),
		Cka:                d.Get("cka").(string),
		Cleartextport:      d.Get("cleartextport").(int),
		Clttimeout:         d.Get("clttimeout").(int),
		Cmp:                d.Get("cmp").(string),
		Comment:            d.Get("comment").(string),
		Customserverid:     d.Get("customserverid").(string),
		Dnsprofilename:     d.Get("dnsprofilename").(string),
		Downstateflush:     d.Get("downstateflush").(string),
		Hashid:             d.Get("hashid").(int),
		Healthmonitor:      d.Get("healthmonitor").(string),
		Httpprofilename:    d.Get("httpprofilename").(string),
		Ip:                 d.Get("ip").(string),
		Maxbandwidth:       d.Get("maxbandwidth").(int),
		Maxclient:          d.Get("maxclient").(int),
		Maxreq:             d.Get("maxreq").(int),
		Monconnectionclose: d.Get("monconnectionclose").(string),
		Monthreshold:       d.Get("monthreshold").(int),
		Netprofile:         d.Get("netprofile").(string),
		Pathmonitor:        d.Get("pathmonitor").(string),
		Pathmonitorindv:    d.Get("pathmonitorindv").(string),
		Port:               d.Get("port").(int),
		Processlocal:       d.Get("processlocal").(string),
		Rtspsessionidremap: d.Get("rtspsessionidremap").(string),
		Sc:                 d.Get("sc").(string),
		Serverid:           d.Get("serverid").(int),
		Servername:         d.Get("servername").(string),
		Servicetype:        d.Get("servicetype").(string),
		Sp:                 d.Get("sp").(string),
		Svrtimeout:         d.Get("svrtimeout").(int),
		Tcpb:               d.Get("tcpb").(string),
		Tcpprofilename:     d.Get("tcpprofilename").(string),
		Td:                 d.Get("td").(int),
		Useproxyport:       d.Get("useproxyport").(string),
		Usip:               d.Get("usip").(string),
	}

	return resource
}

func set_service(d *schema.ResourceData, resource *nitro.Service) {
	d.Set("name", resource.Name)
	d.Set("accessdown", resource.Accessdown)
	d.Set("appflowlog", resource.Appflowlog)
	d.Set("cacheable", resource.Cacheable)
	d.Set("cachetype", resource.Cachetype)
	d.Set("cip", resource.Cip)
	d.Set("cipheader", resource.Cipheader)
	d.Set("cka", resource.Cka)
	d.Set("cleartextport", resource.Cleartextport)
	d.Set("clttimeout", resource.Clttimeout)
	d.Set("cmp", resource.Cmp)
	d.Set("comment", resource.Comment)
	d.Set("customserverid", resource.Customserverid)
	d.Set("dnsprofilename", resource.Dnsprofilename)
	d.Set("downstateflush", resource.Downstateflush)
	d.Set("hashid", resource.Hashid)
	d.Set("healthmonitor", resource.Healthmonitor)
	d.Set("httpprofilename", resource.Httpprofilename)
	d.Set("ip", resource.Ip)
	d.Set("maxbandwidth", resource.Maxbandwidth)
	d.Set("maxclient", resource.Maxclient)
	d.Set("maxreq", resource.Maxreq)
	d.Set("monconnectionclose", resource.Monconnectionclose)
	d.Set("monthreshold", resource.Monthreshold)
	d.Set("netprofile", resource.Netprofile)
	d.Set("pathmonitor", resource.Pathmonitor)
	d.Set("pathmonitorindv", resource.Pathmonitorindv)
	d.Set("port", resource.Port)
	d.Set("processlocal", resource.Processlocal)
	d.Set("rtspsessionidremap", resource.Rtspsessionidremap)
	d.Set("sc", resource.Sc)
	d.Set("serverid", resource.Serverid)
	d.Set("servername", resource.Servername)
	d.Set("servicetype", resource.Servicetype)
	d.Set("sp", resource.Sp)
	d.Set("svrtimeout", resource.Svrtimeout)
	d.Set("tcpb", resource.Tcpb)
	d.Set("tcpprofilename", resource.Tcpprofilename)
	d.Set("td", resource.Td)
	d.Set("useproxyport", resource.Useproxyport)
	d.Set("usip", resource.Usip)
	d.SetId(resource.Name)
}

func create_service(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_service")

	client := meta.(*nitro.NitroClient)

	key := key_service(d)

	exists, err := client.ExistsService(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetService(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_service(d, resource)
	} else {
		err := client.AddService(get_service(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetService(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_service(d, resource)
	}

	return nil
}

func read_service(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_service")

	client := meta.(*nitro.NitroClient)

	key := key_service(d)

	exists, err := client.ExistsService(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetService(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_service(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_service(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_service")

	return nil
}

func delete_service(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_service")

	client := meta.(*nitro.NitroClient)

	key := key_service(d)

	exists, err := client.ExistsService(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteService(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
