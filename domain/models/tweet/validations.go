package tweet

func (t *Tweet) MessagePresent() bool {
	return len(t.Message) > 0
}

func (t *Tweet) LimitValidMessage(limit int) bool {
	return len(t.Message) <= limit
}

func (t *Tweet) UserIdPresent() bool {
	return len(t.UserID) > 0
}
