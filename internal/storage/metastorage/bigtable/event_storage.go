package bigtable

import (
	"context"

	"cloud.google.com/go/bigtable"

	"github.com/coinbase/chainstorage/internal/storage/metastorage/internal"
	"github.com/coinbase/chainstorage/internal/storage/metastorage/model"
	"github.com/coinbase/chainstorage/internal/utils/instrument"
)

type (
	eventStorageImpl struct {
		adminClient                            *bigtable.AdminClient
		client                                 *bigtable.Client
		latestEventTag                         uint32
		instrumentAddEvents                    instrument.Call
		instrumentGetEventByEventId            instrument.Call
		instrumentGetEventsAfterEventId        instrument.Call
		instrumentGetEventsByEventIdRange      instrument.Call
		instrumentGetMaxEventId                instrument.Call
		instrumentSetMaxEventId                instrument.Call
		instrumentGetFirstEventIdByBlockHeight instrument.Call
		instrumentGetEventsByBlockHeight       instrument.Call
	}
)

func newEventStorage(params Params, adminClient *bigtable.AdminClient, client *bigtable.Client) (internal.EventStorage, error) {
	metrics := params.Metrics.SubScope("event_storage_bigtable")
	storage := eventStorageImpl{
		adminClient:                            adminClient,
		client:                                 client,
		latestEventTag:                         params.Config.GetLatestEventTag(),
		instrumentAddEvents:                    instrument.NewCall(metrics, "add_events"),
		instrumentGetEventByEventId:            instrument.NewCall(metrics, "get_event_by_event_id"),
		instrumentGetEventsAfterEventId:        instrument.NewCall(metrics, "get_events_after_event_id"),
		instrumentGetEventsByEventIdRange:      instrument.NewCall(metrics, "get_events_by_event_id_range"),
		instrumentGetMaxEventId:                instrument.NewCall(metrics, "get_max_event_id"),
		instrumentSetMaxEventId:                instrument.NewCall(metrics, "set_max_event_id"),
		instrumentGetFirstEventIdByBlockHeight: instrument.NewCall(metrics, "get_first_event_id_by_block_height"),
		instrumentGetEventsByBlockHeight:       instrument.NewCall(metrics, "get_events_by_block_height"),
	}
	return &storage, nil
}

// AddEvents implements internal.EventStorage.
func (*eventStorageImpl) AddEvents(ctx context.Context, eventTag uint32, events []*model.BlockEvent) error {
	panic("unimplemented")
}

// GetEventByEventId implements internal.EventStorage.
func (*eventStorageImpl) GetEventByEventId(ctx context.Context, eventTag uint32, eventId int64) (*model.EventEntry, error) {
	panic("unimplemented")
}

// GetEventsAfterEventId implements internal.EventStorage.
func (*eventStorageImpl) GetEventsAfterEventId(ctx context.Context, eventTag uint32, eventId int64, maxEvents uint64) ([]*model.EventEntry, error) {
	panic("unimplemented")
}

// GetEventsByBlockHeight implements internal.EventStorage.
func (*eventStorageImpl) GetEventsByBlockHeight(ctx context.Context, eventTag uint32, blockHeight uint64) ([]*model.EventEntry, error) {
	panic("unimplemented")
}

// GetEventsByEventIdRange implements internal.EventStorage.
func (*eventStorageImpl) GetEventsByEventIdRange(ctx context.Context, eventTag uint32, minEventId int64, maxEventId int64) ([]*model.EventEntry, error) {
	panic("unimplemented")
}

// GetFirstEventIdByBlockHeight implements internal.EventStorage.
func (*eventStorageImpl) GetFirstEventIdByBlockHeight(ctx context.Context, eventTag uint32, blockHeight uint64) (int64, error) {
	panic("unimplemented")
}

// GetMaxEventId implements internal.EventStorage.
func (*eventStorageImpl) GetMaxEventId(ctx context.Context, eventTag uint32) (int64, error) {
	panic("unimplemented")
}

// SetMaxEventId implements internal.EventStorage.
func (*eventStorageImpl) SetMaxEventId(ctx context.Context, eventTag uint32, maxEventId int64) error {
	panic("unimplemented")
}
