package bridgecrew

// The Policy record
type Policy struct {
	Provider string `json:"provider"`
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title"`
	//descriptivetitle  string
	//constructivetitle string
	Severity string `json:"severity"`
	Category string `json:"category"`
	//Resourcetypes []string `json:"resourcetypes"`
	//accountsData      []Account
	Guidelines string     `json:"guidelines"`
	Conditions Conditions `json:"conditions"`
	Benchmarks Benchmark  `json:"benchmarks,omitempty"`
	Code       string     `json:"code,omitempty"`
}

// Benchmark is child object to Policy
type Benchmark struct {
	Cisawsv12 []string `json:"CIS AWS V1.2,omitempty"`
	Cisawsv13 []string `json:"CIS AWS V1.3,omitempty"`
}

// Account is a child object to Policy
type Account struct {
	repository     string
	amounts        Amount
	lastupdatedate string
}

// Amount is a sub-object of Account
type Amount struct {
	CLOSED     int
	DELETED    int
	OPEN       int
	REMEDIATED int
	SUPPRESSED int
}

//Conditions is part of the simple query
type Conditions struct {
	Attribute     string   `json:"attribute"`
	CondType      string   `json:"cond_type"`
	Operator      string   `json:"operator"`
	ResourceTypes []string `json:"resource_types"`
	Value         string   `json:"value"`
}
