package inventory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualTo(t *testing.T) {
	po := &PartialObject{
		TypeMeta: TypeMeta{
			Kind:         "Deployment",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "deployments",
		},
		ObjectMeta: ObjectMeta{
			Name:      "deployment1",
			Namespace: "default",
		},
	}

	tests := []struct {
		other  *PartialObject
		result bool
	}{
		{
			other: &PartialObject{
				TypeMeta: TypeMeta{
					Kind:         "Deployment",
					APIGroup:     "apps",
					APIVersion:   "v1",
					ResourceType: "deployments",
				},
				ObjectMeta: ObjectMeta{
					Name:      "deployment1",
					Namespace: "default",
				},
			},
			result: true,
		},
		{
			other: &PartialObject{
				TypeMeta: TypeMeta{
					Kind:         "ReplicaSet",
					APIGroup:     "apps",
					APIVersion:   "v1",
					ResourceType: "deployments",
				},
				ObjectMeta: ObjectMeta{
					Name:      "deployment1",
					Namespace: "default",
				},
			},
			result: false,
		},
		{
			other: &PartialObject{
				TypeMeta: TypeMeta{
					Kind:         "Deployment",
					APIGroup:     "apps",
					APIVersion:   "v1",
					ResourceType: "deployments",
				},
				ObjectMeta: ObjectMeta{
					Name:      "non-equal",
					Namespace: "default",
				},
			},
			result: false,
		},
	}

	for _, tt := range tests {
		result := po.EqualTo(tt.other)
		assert.Equal(t, tt.result, result)
	}
}
