package relation

type IRelation interface {
	Create(t *Relation) (*Relation, bool, error)
}
