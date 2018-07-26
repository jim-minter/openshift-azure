package addons

import (
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestClean(t *testing.T) {
	tests := []struct {
		name    string
		testNum int
		input   *unstructured.Unstructured
		want    *unstructured.Unstructured
	}{
		{
			name:    "Status clean",
			testNum, 1
		},
	}

	for _, tt := range tests {

		_, filename, _, ok := runtime.Caller(1)
		log.Print(filename)
	}

}
