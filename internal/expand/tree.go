package expand

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/ketoapi"
)

func AssertExternalTreesAreEqual(t *testing.T, expected, actual *ketoapi.ExpandTree) {
	t.Helper()
	assert.Truef(t, treesAreEqual(t, expected, actual),
		"expected:\n%+v\n\nactual:\n%+v", expected, actual)
}

func treesAreEqual(t *testing.T, expected, actual *ketoapi.ExpandTree) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if expected.Type != actual.Type {
		t.Logf("expected type %q, actual type %q", expected.Type, actual.Type)
		return false
	}
	if !assert.ObjectsAreEqual(expected.SubjectID, actual.SubjectID) || !assert.ObjectsAreEqual(expected.SubjectSet, actual.SubjectSet) {
		t.Logf("expected subject: %+v %+v, actual subject: %+v %+v", expected.SubjectID, expected.SubjectSet, actual.SubjectID, actual.SubjectSet)
		return false
	}
	if len(expected.Children) != len(actual.Children) {
		t.Logf("expected len(children)=%d, actual len(children)=%d", len(expected.Children), len(actual.Children))
		return false
	}

	// For children, we check for equality disregarding the order
outer:
	for _, expectedChild := range expected.Children {
		for _, actualChild := range actual.Children {
			if treesAreEqual(t, expectedChild, actualChild) {
				continue outer
			}
		}
		t.Logf("expected child:\n%+v\n\nactual child:\n%+v", expectedChild, actual)
		return false
	}
	return true
}
