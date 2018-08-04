package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAuthorizationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_authorizationpolicylabel_func,
		//                Read:          read_authorizationpolicylabel_func,
		//                Update:        update_authorizationpolicylabel_func,
		//                Delete:        delete_authorizationpolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
