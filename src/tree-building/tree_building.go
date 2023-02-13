package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

const rootRecordId = 0

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := map[int]*Node{}
	seen := map[int]bool{}

	var previousRecordId = -1
	for _, record := range sortRecords(records) {
		if isValid, _ := validateRecord(record, seen, previousRecordId); !isValid {
			return nil, fmt.Errorf("invalid record: %d", record.ID)
		}

		nodes[record.ID] = &Node{ID: record.ID}

		if record.ID != rootRecordId {
			parent := nodes[record.Parent]
			node := nodes[record.ID]
			parent.Children = append(parent.Children, node)
		}

		previousRecordId++
	}

	return nodes[rootRecordId], nil
}

func sortRecords(records []Record) []Record {
	var result = make([]Record, len(records))
	copy(result, records)
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func validateRecord(record Record, seen map[int]bool, previousRecordId int) (bool, error) {
	if seen[record.ID] {
		return false, fmt.Errorf("record with id %d occurs multiple times", record.ID)
	}
	seen[record.ID] = true

	if record.ID == rootRecordId {
		if record.Parent != rootRecordId {
			return false, fmt.Errorf("root node has non-0 parent %d", record.Parent)
		}
	} else {
		if record.ID <= record.Parent {
			return false, fmt.Errorf("record %d has self or later parent %d", record.ID, record.Parent)
		} else if record.ID != previousRecordId+1 {
			return false, fmt.Errorf("too high id %d", record.ID)
		}
	}

	return true, nil
}
