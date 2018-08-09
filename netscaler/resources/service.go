package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerService() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_service,
		Read:          read_service,
		Update:        update_service,
		Delete:        delete_service,
		Schema: map[string]*schema.Schema{
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

func get_service(d *schema.ResourceData) nitro.Service {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Service{
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
		Name:               d.Get("name").(string),
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
	var _ = strconv.Itoa

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
	d.Set("name", resource.Name)
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

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_service_key(d *schema.ResourceData) nitro.ServiceKey {

	key := nitro.ServiceKey{
		d.Get("name").(string),
	}
	return key
}

func create_service(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_service")

	client := meta.(*nitro.NitroClient)

	resource := get_service(d)
	key := resource.ToKey()

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

	resource := get_service(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.ServiceUpdate{}
	unset := nitro.ServiceUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("maxclient") {
		updateFlag = true

		value := d.Get("maxclient").(int)
		update.Maxclient = value

		if value == 0 {
			unsetFlag = true

			unset.Maxclient = true
		}

	}
	if d.HasChange("maxreq") {
		updateFlag = true

		value := d.Get("maxreq").(int)
		update.Maxreq = value

		if value == 0 {
			unsetFlag = true

			unset.Maxreq = true
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
	if d.HasChange("cip") {
		updateFlag = true

		value := d.Get("cip").(string)
		update.Cip = value

		if value == "" {
			unsetFlag = true

			unset.Cip = true
		}

	}
	if d.HasChange("cipheader") {
		updateFlag = true

		value := d.Get("cipheader").(string)
		update.Cipheader = value

		if value == "" {
			unsetFlag = true

			unset.Cipheader = true
		}

	}
	if d.HasChange("usip") {
		updateFlag = true

		value := d.Get("usip").(string)
		update.Usip = value

		if value == "" {
			unsetFlag = true

			unset.Usip = true
		}

	}
	if d.HasChange("pathmonitor") {
		updateFlag = true

		value := d.Get("pathmonitor").(string)
		update.Pathmonitor = value

		if value == "" {
			unsetFlag = true

			unset.Pathmonitor = true
		}

	}
	if d.HasChange("pathmonitorindv") {
		updateFlag = true

		value := d.Get("pathmonitorindv").(string)
		update.Pathmonitorindv = value

		if value == "" {
			unsetFlag = true

			unset.Pathmonitorindv = true
		}

	}
	if d.HasChange("useproxyport") {
		updateFlag = true

		value := d.Get("useproxyport").(string)
		update.Useproxyport = value

		if value == "" {
			unsetFlag = true

			unset.Useproxyport = true
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
	if d.HasChange("sp") {
		updateFlag = true

		value := d.Get("sp").(string)
		update.Sp = value

		if value == "" {
			unsetFlag = true

			unset.Sp = true
		}

	}
	if d.HasChange("rtspsessionidremap") {
		updateFlag = true

		value := d.Get("rtspsessionidremap").(string)
		update.Rtspsessionidremap = value

		if value == "" {
			unsetFlag = true

			unset.Rtspsessionidremap = true
		}

	}
	if d.HasChange("healthmonitor") {
		updateFlag = true

		value := d.Get("healthmonitor").(string)
		update.Healthmonitor = value

		if value == "" {
			unsetFlag = true

			unset.Healthmonitor = true
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
	if d.HasChange("svrtimeout") {
		updateFlag = true

		value := d.Get("svrtimeout").(int)
		update.Svrtimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Svrtimeout = true
		}

	}
	if d.HasChange("customserverid") {
		updateFlag = true

		value := d.Get("customserverid").(string)
		update.Customserverid = value

		if value == "" {
			unsetFlag = true

			unset.Customserverid = true
		}

	}
	if d.HasChange("serverid") {
		updateFlag = true

		value := d.Get("serverid").(int)
		update.Serverid = value

		if value == 0 {
			unsetFlag = true

			unset.Serverid = true
		}

	}
	if d.HasChange("cka") {
		updateFlag = true

		value := d.Get("cka").(string)
		update.Cka = value

		if value == "" {
			unsetFlag = true

			unset.Cka = true
		}

	}
	if d.HasChange("tcpb") {
		updateFlag = true

		value := d.Get("tcpb").(string)
		update.Tcpb = value

		if value == "" {
			unsetFlag = true

			unset.Tcpb = true
		}

	}
	if d.HasChange("cmp") {
		updateFlag = true

		value := d.Get("cmp").(string)
		update.Cmp = value

		if value == "" {
			unsetFlag = true

			unset.Cmp = true
		}

	}
	if d.HasChange("maxbandwidth") {
		updateFlag = true

		value := d.Get("maxbandwidth").(int)
		update.Maxbandwidth = value

		if value == 0 {
			unsetFlag = true

			unset.Maxbandwidth = true
		}

	}
	if d.HasChange("accessdown") {
		updateFlag = true

		value := d.Get("accessdown").(string)
		update.Accessdown = value

		if value == "" {
			unsetFlag = true

			unset.Accessdown = true
		}

	}
	if d.HasChange("monthreshold") {
		updateFlag = true

		value := d.Get("monthreshold").(int)
		update.Monthreshold = value

		if value == 0 {
			unsetFlag = true

			unset.Monthreshold = true
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
	if d.HasChange("hashid") {
		updateFlag = true

		value := d.Get("hashid").(int)
		update.Hashid = value

		if value == 0 {
			unsetFlag = true

			unset.Hashid = true
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
	if d.HasChange("monconnectionclose") {
		updateFlag = true

		value := d.Get("monconnectionclose").(string)
		update.Monconnectionclose = value

		if value == "" {
			unsetFlag = true

			unset.Monconnectionclose = true
		}

	}
	key := get_service_key(d)

	if updateFlag {
		if err := client.UpdateService(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetService(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetService(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_service(d, resource)
	}

	return nil
}

func delete_service(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_service")

	client := meta.(*nitro.NitroClient)

	resource := get_service(d)
	key := resource.ToKey()

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
