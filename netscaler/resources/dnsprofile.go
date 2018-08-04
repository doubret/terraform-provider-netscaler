package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerDnsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_dnsprofile_func,
		//                Read:          read_dnsprofile_func,
		//                Update:        update_dnsprofile_func,
		//                Delete:        delete_dnsprofile_func,
		Schema: map[string]*schema.Schema{
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cacheecsresponses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cachenegativeresponses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cacherecords": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsanswerseclogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnserrorlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsextendedlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsquerylogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropmultiqueryrequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
