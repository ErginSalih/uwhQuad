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
	return string(quada)
}

func FindQuadB(x, y int) string {
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
	return string(quadb)
}

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

func FindQuadD(x, y int) string {
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
	return string(quadd)
}

func FindQuadE(x, y int) string {
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
	return string(quade)
}



