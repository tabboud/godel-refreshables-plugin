// Generated by godel-refreshable-plugin: do not edit.

package test3

import refreshable "github.com/palantir/pkg/refreshable"

type RefreshableMainStruct interface {
	refreshable.Refreshable
	CurrentMainStruct() MainStruct
	MapMainStruct(func(MainStruct) interface{}) refreshable.Refreshable
	SubscribeToMainStruct(func(MainStruct)) (unsubscribe func())

	IncludedString() refreshable.String
	Sub() RefreshableSubStruct
}

type RefreshingMainStruct struct {
	refreshable.Refreshable
}

func NewRefreshingMainStruct(in refreshable.Refreshable) RefreshingMainStruct {
	return RefreshingMainStruct{Refreshable: in}
}

func (r RefreshingMainStruct) CurrentMainStruct() MainStruct {
	return r.Current().(MainStruct)
}

func (r RefreshingMainStruct) MapMainStruct(mapFn func(MainStruct) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(MainStruct))
	})
}

func (r RefreshingMainStruct) SubscribeToMainStruct(consumer func(MainStruct)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(MainStruct))
	})
}

func (r RefreshingMainStruct) IncludedString() refreshable.String {
	return refreshable.NewString(r.MapMainStruct(func(i MainStruct) interface{} {
		return i.IncludedString
	}))
}

func (r RefreshingMainStruct) Sub() RefreshableSubStruct {
	return NewRefreshingSubStruct(r.MapMainStruct(func(i MainStruct) interface{} {
		return i.Sub
	}))
}

type RefreshableSubStruct interface {
	refreshable.Refreshable
	CurrentSubStruct() SubStruct
	MapSubStruct(func(SubStruct) interface{}) refreshable.Refreshable
	SubscribeToSubStruct(func(SubStruct)) (unsubscribe func())

	IncludedInt() refreshable.Int
}

type RefreshingSubStruct struct {
	refreshable.Refreshable
}

func NewRefreshingSubStruct(in refreshable.Refreshable) RefreshingSubStruct {
	return RefreshingSubStruct{Refreshable: in}
}

func (r RefreshingSubStruct) CurrentSubStruct() SubStruct {
	return r.Current().(SubStruct)
}

func (r RefreshingSubStruct) MapSubStruct(mapFn func(SubStruct) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(SubStruct))
	})
}

func (r RefreshingSubStruct) SubscribeToSubStruct(consumer func(SubStruct)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(SubStruct))
	})
}

func (r RefreshingSubStruct) IncludedInt() refreshable.Int {
	return refreshable.NewInt(r.MapSubStruct(func(i SubStruct) interface{} {
		return i.IncludedInt
	}))
}