package bridgecrew

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceComplexPolicy(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Resource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourceComplexPolicy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceComplexPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourceComplexPolicyCreate(t *testing.T) {
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
			if got := resourceComplexPolicyCreate(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceComplexPolicyCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setComplexPolicy(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		want    complexPolicy
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setComplexPolicy(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("setComplexPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setComplexPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setComplexConditions(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		want    ConditionQuery
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setComplexConditions(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("setComplexConditions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setComplexConditions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourceComplexPolicyRead(t *testing.T) {
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
			if got := resourceComplexPolicyRead(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceComplexPolicyRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourceComplexPolicyUpdate(t *testing.T) {
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
			if got := resourceComplexPolicyUpdate(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceComplexPolicyUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complexPolicyChange(t *testing.T) {
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
			if got := complexPolicyChange(tt.args.d); got != tt.want {
				t.Errorf("complexPolicyChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
