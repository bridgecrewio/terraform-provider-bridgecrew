package bridgecrew

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourcePolicy(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Resource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourcePolicy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourcePolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourcePolicyCreate(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *schema.ResourceData
		m   interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourcePolicyCreate(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourcePolicyCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPolicy(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		want    Policy
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setPolicy(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("setPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCastToStringList(t *testing.T) {
	type args struct {
		temp []interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CastToStringList(tt.args.temp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CastToStringList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourcePolicyRead(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *schema.ResourceData
		m   interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourcePolicyRead(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourcePolicyRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_highlight(t *testing.T) {
	type args struct {
		myPolicy interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			highlight(tt.args.myPolicy)
		})
	}
}

func Test_resourcePolicyUpdate(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *schema.ResourceData
		m   interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourcePolicyUpdate(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourcePolicyUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_policyChange(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := policyChange(tt.args.d); got != tt.want {
				t.Errorf("policyChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourcePolicyDelete(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *schema.ResourceData
		m   interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourcePolicyDelete(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourcePolicyDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
