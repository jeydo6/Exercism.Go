package strand

func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	for i := 0; i < len(dna); i++ {
		rna[i] = toRna(dna[i])
	}

	return string(rna)
}

func toRna(dna uint8) rune {
	switch dna {
	case 'A':
		return 'U'
	case 'T':
		return 'A'
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	default:
		return -1
	}
}
