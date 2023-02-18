package kata

func DNAStrand(dna string) string {
	var res string
	for _, n := range dna {
		switch n {
		case 'A':
			res += "T"
		case 'T':
			res += "A"
		case 'G':
			res += "C"
		default:
			res += "G"
		}
	}
	return res
}
