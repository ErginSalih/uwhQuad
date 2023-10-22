package uwhQuad

func FindQuadE(x, y int) []string {
	quade := []rune(nil)
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quade = append(quade, 'A')
					} else if j == x {
						quade = append(quade, 'C')
					} else {
						quade = append(quade, 'B')
					}
				}
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quade = append(quade, 'C')
					} else if j == x {
						quade = append(quade, 'A')
					} else {
						quade = append(quade, 'B')
					}
				}
			} else {
				for k := 1; k <= x; k++ {
					if k == 1 || k == x {
						quade = append(quade, 'B')
					} else {
						quade = append(quade, ' ')
					}
				}
			}
			quade = append(quade, '\n')
		}
	}
	return []string{string(quadE), "e"}
}
