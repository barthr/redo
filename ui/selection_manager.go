package ui

var (
	selectionManager = &SelectionManager{items: []*HistoryItem{}}
)

type SelectionManager struct {
	items []*HistoryItem
}

func (r *SelectionManager) Add(item *HistoryItem) {
	if r.Contains(item) {
		return
	}
	r.items = append(r.items, item)
}

func (r *SelectionManager) Remove(item *HistoryItem) {
	index := r.IndexOf(item)
	if index == -1 {
		return
	}
	r.items = append(r.items[:index], r.items[index+1:]...)
}

func (r *SelectionManager) Contains(item *HistoryItem) bool {
	return r.IndexOf(item) != -1
}

func (r *SelectionManager) IndexOf(item *HistoryItem) int {
	for i, v := range r.items {
		if v == item {
			return i
		}
	}
	return -1
}
