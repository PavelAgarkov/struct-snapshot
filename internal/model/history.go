package model

type Entity interface {
	GetTargetStream() string
	MakeNewSnapshot(tag string)
}

type History interface {
	InitStructs()
	PrepareFieldsData() Changes
	ToFillEntitiesFields(oldEntity Entity, updatedEntity Entity)
	FieldIsEqual(old any, new any) bool
	Reset()
}

type Field struct {
	OldValue any
	NewValue any
}
