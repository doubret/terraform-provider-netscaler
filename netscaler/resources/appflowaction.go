package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerAppflowaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowaction,
		Read:          read_appflowaction,
		Update:        update_appflowaction,
		Delete:        delete_appflowaction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"clientsidemeasurements": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"collectors": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"distributionalgorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"metricslog": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pagetracking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"securityinsight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"transactionlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"videoanalytics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"webinsight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_appflowaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_appflowaction(d *schema.ResourceData) nitro.Appflowaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appflowaction{
		Name: d.Get("name").(string),
		Clientsidemeasurements: d.Get("clientsidemeasurements").(string),
		Collectors:             utils.Convert_set_to_string_array(d.Get("collectors").(*schema.Set)),
		Comment:                d.Get("comment").(string),
		Distributionalgorithm:  d.Get("distributionalgorithm").(string),
		Metricslog:             d.Get("metricslog").(bool),
		Pagetracking:           d.Get("pagetracking").(string),
		Securityinsight:        d.Get("securityinsight").(string),
		Transactionlog:         d.Get("transactionlog").(string),
		Videoanalytics:         d.Get("videoanalytics").(string),
		Webinsight:             d.Get("webinsight").(string),
	}

	return resource
}

func set_appflowaction(d *schema.ResourceData, resource *nitro.Appflowaction) {
	d.Set("name", resource.Name)
	d.Set("clientsidemeasurements", resource.Clientsidemeasurements)
	d.Set("collectors", resource.Collectors)
	d.Set("comment", resource.Comment)
	d.Set("distributionalgorithm", resource.Distributionalgorithm)
	d.Set("metricslog", resource.Metricslog)
	d.Set("pagetracking", resource.Pagetracking)
	d.Set("securityinsight", resource.Securityinsight)
	d.Set("transactionlog", resource.Transactionlog)
	d.Set("videoanalytics", resource.Videoanalytics)
	d.Set("webinsight", resource.Webinsight)
	d.SetId(resource.Name)
}

func create_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowaction")

	client := meta.(*nitro.NitroClient)

	key := key_appflowaction(d)

	exists, err := client.ExistsAppflowaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appflowaction(d, resource)
	} else {
		err := client.AddAppflowaction(get_appflowaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppflowaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appflowaction(d, resource)
	}

	return nil
}

func read_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowaction")

	client := meta.(*nitro.NitroClient)

	key := key_appflowaction(d)

	exists, err := client.ExistsAppflowaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appflowaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appflowaction")

	return nil
}

func delete_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowaction")

	client := meta.(*nitro.NitroClient)

	key := key_appflowaction(d)

	exists, err := client.ExistsAppflowaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppflowaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
