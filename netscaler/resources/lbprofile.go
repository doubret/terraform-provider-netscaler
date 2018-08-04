package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerLbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_lbprofile_func,
		//                Read:          read_lbprofile_func,
		//                Update:        update_lbprofile_func,
		//                Delete:        delete_lbprofile_func,
		Schema: map[string]*schema.Schema{
			"lbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cookiepassphrase": &schema.Schema{
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
			"httponlycookieflag": &schema.Schema{
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
			"useencryptedpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"usesecuredpersistencecookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
