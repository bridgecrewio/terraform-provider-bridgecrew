package bridgecrew

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_dataSourceUsers(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Resource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataSourceUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataSourceUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataSourceUsersRead(t *testing.T) {
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
			if got := dataSourceUsersRead(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dataSourceUsersRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flattenUserData(t *testing.T) {
	type args struct {
		Users *[]map[string]interface{}
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
			if got := flattenUserData(tt.args.Users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}
