package bridgecrew

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_dataSourceApitokens(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Resource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataSourceApitokens(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataSourceApitokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataSourceApitokensRead(t *testing.T) {
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
			if got := dataSourceApitokensRead(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataSourceApitokensRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flattenApitokensData(t *testing.T) {
	type args struct {
		Apitokens *[]map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenApitokensData(tt.args.Apitokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenApitokensData() = %v, want %v", got, tt.want)
			}
		})
	}
}
