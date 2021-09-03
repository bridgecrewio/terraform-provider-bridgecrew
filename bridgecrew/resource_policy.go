package bridgecrew

import (
	"context"
	"fmt"
	"log"
	"strconv"

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
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					switch val.(string) {
					case
						"CRITICAL",
						"HIGH",
						"LOW",
						"MEDIUM":
						return
					}
					errs = append(errs, fmt.Errorf("%q Must be one of CRITICAL, HIGH, MEDIUM or LOW", val))
					return
				},
			},
			"category": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					switch val.(string) {
					case
						"LOGGING",
						"ELASTICSEARCH",
						"GENERAL",
						"STORAGE",
						"ENCRYPTION",
						"NETWORKING",
						"MONITORING",
						"KUBERNETES",
						"SERVERLESS",
						"BACKUP_AND_RECOVERY",
						"IAM",
						"SECRETS",
						"PUBLIC",
						"GENERAL_SECURITY":
						return
					}
					errs = append(errs,
						fmt.Errorf("%q Must be one of LOGGING, ELASTICSEARCH, GENERAL, STORAGE, ENCRYPTION,"+
							" NETWORKING, MONITORING, KUBERNETES, SERVERLESS, BACKUP_AND_RECOVERY, SECRETS, PUBLIC,"+
							" GENERAL_SECURITY or IAM", val))
					return
				},
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
				Type:     schema.TypeList,
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
		},
	}
}

func resourcePolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	myPolicy := Policy{}

	//if d.Get("accountsdata") != nil {
	//	myAccounts := d.Get("accountsdata").([]interface{})
	//	log.Print(myAccounts)
	//	//log.Print(myAccounts["repository"])
	//	//for i, Account := range myAccounts {
	//	//	myAccount:=Account.(map[string]interface{})
	//	//
	//	//	myPolicy.AccountsData[i].repository = myAccount["repository"].(string)
	//	//}
	//}

	myBenchmarks := d.Get("benchmarks").([]interface{})

	var myItems []Benchmark
	if len(myBenchmarks) != 0 {
		for _, myBenchmark := range myBenchmarks {
			s := myBenchmark.(map[string]interface{})
			var Item Benchmark
			versions := CastToStringList(s["version"].([]interface{}))
			Item.version = versions
			Item.benchmark = s["benchmark"].(string)
			myItems = append(myItems, Item)
		}
	}

	myPolicy.Benchmarks = myItems
	myPolicy.Category = d.Get("category").(string)
	myPolicy.Code = d.Get("title").(string)
	myPolicy.Code = d.Get("code").(string)
	myPolicy.Constructivetitle = d.Get("constructive_title").(string)
	myPolicy.Descriptivetitle = d.Get("descriptive_title").(string)
	myPolicy.Provider = d.Get("cloud_provider").(string)
	myPolicy.Severity = d.Get("severity").(string)
	myPolicy.Title = d.Get("title").(string)
	myPolicy.Conditionquery = d.Get("condition_query").(string)
	myPolicy.Createdby = d.Get("createdby").(string)
	myPolicy.Guideline = d.Get("guideline").(string)
	myPolicy.Iscustom = d.Get("iscustom").(bool)

	Types := d.Get("resource_types").([]interface{})

	if len(Types) != 0 {
		for _, Type := range Types {
			myPolicy.Resourcetypes = append(myPolicy.Resourcetypes, Type.(string))
		}
	}

	highlight(myPolicy)

	//jspolicy, err := json.Marshal(myPolicy)
	//if err != nil {
	//	log.Fatal("json could no be written")
	//}
	//log.Print(strings.NewReader(string(jspolicy)))

	//configure := m.(ProviderConfig)
	//url := configure.URL+"/policies"
	//
	//req, _ := http.NewRequest("POST", url, strings.NewReader(string(jspolicy)))
	//
	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("authorization", configure.Token)
	//
	//res, _ := client.Do(req)
	//
	//defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//
	//if err != nil {
	//	log.Fatal("json could no be written")
	//}
	//
	//Policy := Policy{}
	//err = json.Unmarshal(body, &Policy)
	//
	//if err != nil {
	//	log.Fatal("json could no be unmarshalled")
	//}

	d.SetId(strconv.Itoa(myPolicy.ID))

	return diags
}

// CastToStringList is a helper to work with coversion of types
// If theres a better way (most likely)?
func CastToStringList(temp []interface{}) []string {
	var versions []string
	for _, version := range temp {
		versions = append(versions, version.(string))
	}
	return versions
}

// highlight is just to help with manual debugging so you can find the lines
func highlight(myPolicy interface{}) {
	log.Print("XXXXXXXXXXX")
	log.Print(myPolicy)
	log.Print("XXXXXXXXXXX")
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
