package addons

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var obj *unstructured.Unstructured

func newTestObj() unstructured.Unstructured {

	trueVar := true
	ten := int64(10)
	obj := unstructured.Unstructured{
		Object: map[string]interface{}{
			"kind":       "test_kind",
			"apiVersion": "test_version",
			"metadata": map[string]interface{}{
				// keep fields
				"generateName": "test_generateName",
				"labels": map[string]interface{}{
					"test_label": "test_value",
				},
				// clean fields
				"resourceVersion":            "test_resourceVersion",
				"uid":                        "test_uid",
				"generation":                 ten,
				"deletionGracePeriodSeconds": ten,
				"selfLink":                   "test_selfLink",
				"creationTimestamp":          "2009-11-10T23:00:00Z",
				"deletionTimestamp":          "2010-11-10T23:00:00Z",
				"annotations": map[string]interface{}{
					// keep annotation
					"test_annotation": "test_value",
					// clean annotations
					"kubectl.kubernetes.io/last-applied-configuration": "test",
					"openshift.io/generated-by":                        "test_generator",
				},
				"ownerReferences": []interface{}{
					map[string]interface{}{
						"kind":       "Pod",
						"name":       "poda",
						"apiVersion": "v1",
						"uid":        "1",
					},
				},
			},
			"spec": map[string]interface{}{
				"initContainers": []interface{}{
					map[string]interface{}{
						"name":            "init1",
						"image":           "image1",
						"imagePullPolicy": "Always",
					},
					map[string]interface{}{
						"name":            "init2",
						"image":           "image2",
						"imagePullPolicy": "Always",
					},
				},
				"containers": []interface{}{
					map[string]interface{}{
						"name":            "init1",
						"image":           "image1",
						"imagePullPolicy": "Always",
					},
					map[string]interface{}{
						"name":            "init2",
						"image":           "image2",
						"imagePullPolicy": "Always",
					},
				},
			},
		},
	}

	return obj
}
