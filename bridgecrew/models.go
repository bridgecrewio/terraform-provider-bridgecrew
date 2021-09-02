package bridgecrew

// The Policy record
type Policy struct {
	Provider          string
	ID                int
	Title             string
	Descriptivetitle  string
	Constructivetitle string
	Severity          string
	Category          string
	Resourcetypes     []string
	AccountsData      []Account
	Guideline         string
	Iscustom          bool
	Conditionquery    string
	Benchmarks        []Benchmark
	Createdby         string
	Code              string
}

// Benchmark is child object to Policy
type Benchmark struct {
	benchmark string
	version   []string
}

// Account is a child object to Policy
type Account struct {
	repository     string
	Amounts        Amount
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
