package infrastructure

import (
	"encoding/json"
	"hrm/pkg/candidate/domain"
)

type dbEventDispatcher struct {
	repo domain.MessageRepository
}

func (d *dbEventDispatcher) Dispatch(event domain.Event) error {
	serialized, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return d.repo.Save(domain.Message{Msg: string(serialized)})
}

type dbEventDispatcherFactory struct {
}

func (d dbEventDispatcherFactory) Create(mr domain.MessageRepository) domain.EventDispatcher {
	return &dbEventDispatcher{repo: mr}
}
