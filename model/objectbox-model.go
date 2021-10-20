// Code generated by ObjectBox; DO NOT EDIT.

package model

import (
	"github.com/objectbox/objectbox-go/objectbox"
)

// ObjectBoxModel declares and builds the model from all the entities in the package.
// It is usually used when setting-up ObjectBox as an argument to the Builder.Model() function.
func ObjectBoxModel() *objectbox.Model {
	model := objectbox.NewModel()
	model.GeneratorVersion(6)

	model.RegisterBinding(TimeseriesBinding)
	model.LastEntityId(1, 8153158015204459380)

	return model
}
