package sublist

func Sublist(l1, l2 []int) Relation {
	super := contains(l1, l2)
	sub := contains(l2, l1)
	switch {
	case super && sub:
		return RelationEqual
	case super:
		return RelationSuperlist
	case sub:
		return RelationSublist
	default:
		return RelationUnequal
	}
}

func contains(l1, l2 []int) bool {
	if len(l2) == 0 {
		return true
	}

	if len(l2) > len(l1) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		isEqual := true
		for j := 0; j < len(l2); j++ {
			var p1 int
			if i+j < len(l1) {
				p1 = i + j
			} else {
				p1 = len(l1) - 1
			}
			var p2 = j
			if l1[p1] != l2[p2] {
				isEqual = false
				break
			}
		}

		if isEqual {
			return true
		}
	}

	return false
}
