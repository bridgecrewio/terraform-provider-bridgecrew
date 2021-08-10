package bridgecrew

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestFlattenRepositoryData(t *testing.T) {
	cases := []struct {
		ownerID      *string
		repositories []map[string]interface{}
		expected     []*GroupIdentifier
	}{
		// simple, no user id included
		{
			ownerId: aws.String("user1234"),
			repositories: []*ec2.UserIdGroupPair{
				&ec2.UserIdGroupPair{
					GroupId: aws.String("sg-12345"),
				},
			},
			expected: []*GroupIdentifier{
				&GroupIdentifier{
					GroupId: aws.String("sg-12345"),
				},
			},
		},
	}

	for _, c := range cases {
		out := flattenRepositoryData(c.pairs, c.ownerId)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}
