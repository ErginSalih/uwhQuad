package uwhQuad

func FindQuadA(x, y int) string {
	quada := []rune(nil)
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			if i == 1 || i == y {
				for j := x; j > 0; j-- {
					if j == 1 || j == x {
						quada = append(quada, 'o')
					} else {
						quada = append(quada, '-')
					}
				}
			} else {
				for k := 1; k <= x; k++ {
					if k == 1 || k == x {
						quada = append(quada, '|')
					} else {
						quada = append(quada, ' ')
					}
				}
			}
			quada = append(quada, '\n')
		}
	}
	return []string{string(quada),"A"}
}
