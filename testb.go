package uwhQuad

func FindQuadB(x, y int) []string {
	quadb := []rune(nil)
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quadb = append(quadb, '/')
					} else if j == x {
						quadb = append(quadb, '\\')
					} else {
						quadb = append(quadb, '*')
					}
				}
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quadb = append(quadb, '\\')
					} else if j == x {
						quadb = append(quadb, '/')
					} else {
						quadb = append(quadb, '*')
					}
				}
			} else {
				for k := 1; k <= x; k++ {
					if k == 1 || k == x {
						quadb = append(quadb, '*')
					} else {
						quadb = append(quadb, ' ')
					}
				}
			}
			quadb = append(quadb, '\n')
		}
	}
	return []string{string(quadb), "B"}
}
