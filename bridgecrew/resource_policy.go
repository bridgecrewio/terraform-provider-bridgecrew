package bridgecrew

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePolicyCreate,
		ReadContext:   resourcePolicyRead,
		UpdateContext: resourePolicyUpdate,
		DeleteContext: resourcePolicyDelete,
		Schema: map[string]*schema.Schema{
			//"policies": {
			//	Type:     schema.TypeSet,
			//	Computed: false,
			//	Required: true,
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			"cloud_provider": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			//"id": {
			//  Type:     schema.TypeString,
			//  Computed: true,
			//},
			"title": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"descriptive_title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"constructive_title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Required: true,
				//todo
				// needs to one of the severity classes not a free string field
			},
			"category": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_types": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"accountsdata": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Required: true,
							Type:     schema.TypeString,
						},
						"amounts": {
							Required: true,
							Type:     schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"lastupdatedate": {
							Computed: true,
							Type:     schema.TypeString,
						},
					},
				},
			},
			"guideline": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"iscustom": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"condition_query": {
				Type:     schema.TypeString,
				Required: true,
			},
			"benchmarks": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"benchmark": {
							Type:     schema.TypeString,
							Required: true,
						},
						"version": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"createdby": {
				Type:     schema.TypeString,
				Computed: true,
				//todo
				// should this not be estimated automatically from the TOKEN?
			},
			"code": {
				Type:     schema.TypeString,
				Required: true,
			},
			//	},
			//	},
			//	},
		},
	}
}

func resourcePolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	policies := d.Get("policies").([]interface{})

	log.Print(policies)

	//d.SetId(strconv.Itoa(o.ID))

	return diags
}

func resourcePolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourePolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePolicyRead(ctx, d, m)
}

func resourcePolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

// CreateOrder - Create new order
//func (c *Client) CreatePolicy(PolicyItems []PolicyItem) (*Policy, error) {
//	rb, err := json.Marshal(PolicyItem)
//	if err != nil {
//		return nil, err
//	}
//
//	req, err := http.NewRequest("POST", fmt.Sprintf("%s/orders", c.HostURL), strings.NewReader(string(rb)))
//	if err != nil {
//		return nil, err
//	}
//
//	body, err := c.doRequest(req)
//	if err != nil {
//		return nil, err
//	}
//
//	policy := Policy{}
//	err = json.Unmarshal(body, &policy)
//	if err != nil {
//		return nil, err
//	}
//
//	return &policy, nil
//}
