package uwhQuad

func FindQuadC(x, y int) string {
	quadc := []rune(nil)
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quadc = append(quadc, 'A')
					} else if j == x {
						quadc = append(quadc, 'A')
					} else {
						quadc = append(quadc, 'B')
					}
				}
			} else if i == y {
				for j := 1; j <= x; j++ {
					if j == 1 {
						quadc = append(quadc, 'C')
					} else if j == x {
						quadc = append(quadc, 'C')
					} else {
						quadc = append(quadc, 'B')
					}
				}
			} else {
				for k := 1; k <= x; k++ {
					if k == 1 || k == x {
						quadc = append(quadc, 'B')
					} else {
						quadc = append(quadc, ' ')
					}
				}
			}
			quadc = append(quadc, '\n')
		}
	}
	return string(quadc)
}
