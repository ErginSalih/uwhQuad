package uwhQuad

func FindQuadD(x, y int) []string {
	quadd := []rune(nil)
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quadd = append(quadd, 'A')
					} else if j == x {
						quadd = append(quadd, 'C')
					} else {
						quadd = append(quadd, 'B')
					}
				}
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quadd = append(quadd, 'A')
					} else if j == x {
						quadd = append(quadd, 'C')
					} else {
						quadd = append(quadd, 'B')
					}
				}
			} else {
				for k := 1; k <= x; k++ {
					if k == 1 || k == x {
						quadd = append(quadd, 'B')
					} else {
						quadd = append(quadd, ' ')
					}
				}
			}
			quadd = append(quadd, '\n')
		}
	}
	return []string{string(quadd), "D"}
}
