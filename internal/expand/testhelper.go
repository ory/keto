package expand

import (
	"testing"

	"github.com/ory/keto/internal/relationtuple"

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

func AssertInternalTreesAreEqual(t *testing.T, expected, actual *relationtuple.Tree) bool {
	if !assert.ObjectsAreEqual(expected.Type, actual.Type) {
		t.Logf("expected type %+v, but got %+v", expected.Type, actual.Type)
		return false
	}
	if !assert.ObjectsAreEqual(expected.Subject, actual.Subject) {
		t.Logf("expected subject %+v, but got %+v", expected.Subject, actual.Subject)
		return false
	}
	if len(expected.Children) != len(actual.Children) {
		t.Logf("expected %d children, but got %d", len(expected.Children), len(actual.Children))
		return false
	}

outer:
	for _, child := range expected.Children {
		for _, actualChild := range actual.Children {
			if AssertInternalTreesAreEqual(t, child, actualChild) {
				continue outer
			}
		}
		assert.Truef(t, false, "could not find %+v", child)
		return false
	}
	return true
}
