package ui

var (
	selectionManager = &SelectionManager{items: []*HistoryItem{}}
)

type SelectionManager struct {
	items []*HistoryItem
}

func (r *SelectionManager) Add(item *HistoryItem) {
	r.items = append(r.items, item)
}

func (r *SelectionManager) Remove(item *HistoryItem) {
	for i, v := range r.items {
		if v == item {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return
		}
	}
}

func (r *SelectionManager) Contains(item *HistoryItem) bool {
	for _, v := range r.items {
		if v == item {
			return true
		}
	}
	return false
}

func (r *SelectionManager) IndexOf(item *HistoryItem) int {
	for i, v := range r.items {
		if v == item {
			return i
		}
	}
	return -1
}
