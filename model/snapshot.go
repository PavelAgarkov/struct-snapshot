package model

type Changes map[string]*Field

type Snapshot struct {
	tag     string
	changes Changes
}

func (sht *Snapshot) GetChanges() Changes {
	return sht.changes
}

func (sht *Snapshot) GetTag() string {
	return sht.tag
}

type Snap struct {
	current   Entity
	history   History
	snapshots []*Snapshot
}

func NewSnapshot(entity Entity, history History) *Snap {
	return &Snap{
		current:   entity,
		history:   history,
		snapshots: make([]*Snapshot, 0),
	}
}

func (s *Snap) makeNewSnapshot(updatedEntity Entity, tag string) {
	s.history.InitStructs()
	s.history.ToFillEntitiesFields(s.current, updatedEntity)

	s.current = updatedEntity
	if len(s.history.PrepareFieldsData()) == 0 {
		s.snapshots = append(s.snapshots, &Snapshot{
			tag: tag,
		})
	} else {
		s.snapshots = append(s.snapshots, &Snapshot{
			tag:     tag,
			changes: s.history.PrepareFieldsData(),
		})
	}
	s.history.Reset()
}

func (s *Snap) GetSnapshots() []*Snapshot {
	return s.snapshots
}

func (s *Snap) GetSnapshotByIndex(key int64) *Snapshot {
	if key < int64(len(s.snapshots)) {
		return s.snapshots[key]
	}
	return nil
}

func (s *Snap) GetSnapshotByTag(tag string) *Snapshot {
	for _, v := range s.snapshots {
		if v.GetTag() == tag {
			return v
		}
	}
	return nil
}
