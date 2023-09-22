package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type PermainanDadu struct {
	Dadu             []string
	MendapatkanDadu1 int
}
type AfterEvaluasi struct {
	DaduOld  []string
	Dadu1    []string
	DaduNot1 []string
}

func main() {
	var pemain, dadu int
	fmt.Print("jumlah pemain : ")
	fmt.Scanln(&pemain)
	fmt.Print("jumlah dadu : ")
	fmt.Scanln(&dadu)

	pemainDadu := make(map[int]PermainanDadu)

	for i := 0; i < pemain; i++ {
		pemainDadu[i+1] = PermainanDadu{}
	}

	forI := 0
	stop := len(pemainDadu)
	var skor int
	var pemenang int
	for {
		fmt.Println("giliran", forI+1, "lempar dadu")
		daduMapping := make(map[int]AfterEvaluasi)
		for i1 := 0; i1 < len(pemainDadu); i1++ {
			daduRandom := []string{}
			if forI == 0 {
				for i3 := 0; i3 < dadu; i3++ {
					stringRandom := strconv.Itoa(rand.Intn(6) + 1)
					daduRandom = append(daduRandom, stringRandom)
				}
			} else {
				for i2 := 0; i2 < len(pemainDadu[i1+1].Dadu); i2++ {
					stringRandom := strconv.Itoa(rand.Intn(6) + 1)
					daduRandom = append(daduRandom, stringRandom)
				}
			}
			pemainDadu[i1+1] = PermainanDadu{
				Dadu:             append(daduRandom),
				MendapatkanDadu1: pemainDadu[i1+1].MendapatkanDadu1,
			}
			fmt.Println(fmt.Sprintf("   pemain #%d (%d) : %s ", i1+1, pemainDadu[i1+1].MendapatkanDadu1, strings.Join(pemainDadu[i1+1].Dadu, ", ")))

			dadu1, daduNot1 := contains(daduRandom, "1")
			daduMapping[i1+1] = AfterEvaluasi{
				Dadu1:    append(dadu1),
				DaduOld:  append(daduRandom),
				DaduNot1: append(daduNot1),
			}
		}
		fmt.Println("setelah evaluasi")
		for i4 := 0; i4 < len(daduMapping); i4++ {
			var appendDadu []string
			if len(daduMapping[i4+1].DaduOld) > 0 {
				if i4 < len(daduMapping) && i4 > 0 {
					appendDadu = append(appendDadu, daduMapping[i4].Dadu1...)
					appendDadu = append(appendDadu, daduMapping[i4+1].DaduNot1...)
					pemainDadu[i4+1] = PermainanDadu{
						Dadu:             append(appendDadu),
						MendapatkanDadu1: pemainDadu[i4+1].MendapatkanDadu1 + len(daduMapping[i4+1].Dadu1),
					}
				} else {
					if i4 == 0 {
						appendDadu = append(appendDadu, daduMapping[len(daduMapping)].Dadu1...)
						appendDadu = append(appendDadu, daduMapping[1].DaduNot1...)
						pemainDadu[1] = PermainanDadu{
							Dadu:             append(appendDadu),
							MendapatkanDadu1: pemainDadu[1].MendapatkanDadu1 + len(daduMapping[1].Dadu1),
						}
					} else {
						appendDadu = append(appendDadu, daduMapping[i4+1].DaduNot1...)
						pemainDadu[i4+1] = PermainanDadu{
							Dadu:             append(appendDadu),
							MendapatkanDadu1: pemainDadu[i4+1].MendapatkanDadu1 + len(daduMapping[i4+1].Dadu1),
						}
					}
				}
			}

		}

		for i5 := 0; i5 < len(pemainDadu); i5++ {
			var daduNot6 []string
			// if len(daduMapping[i+1].DaduOld) <= 0 {
			for _, v := range pemainDadu[i5+1].Dadu {
				if v != "6" {
					daduNot6 = append(daduNot6, v)
				}
			}
			pemainDadu[i5+1] = PermainanDadu{
				Dadu:             append(daduNot6),
				MendapatkanDadu1: pemainDadu[i5+1].MendapatkanDadu1,
			}
			fmt.Println(fmt.Sprintf("   pemain #%d (%d) : %s ", i5+1, pemainDadu[i5+1].MendapatkanDadu1, strings.Join(pemainDadu[i5+1].Dadu, ", ")))
			// }
		}
		fmt.Println("=================")
		forI += 1

		for i6 := 0; i6 < len(pemainDadu); i6++ {
			if i6 == 0 {
				skor = 0
				pemenang = i6 + 1
			}
			if skor < pemainDadu[i6+1].MendapatkanDadu1 {
				skor = pemainDadu[i6+1].MendapatkanDadu1
				pemenang = i6 + 1
			}
			if len(pemainDadu[i6+1].Dadu) <= 0 {
				stop = stop - 1
			}
			if len(pemainDadu[i6+1].Dadu) > 0 {
				pemain = i6 + 1
			}
		}
		if stop == 1 || stop == 0 {
			fmt.Println(fmt.Sprintf("game berakhir karena hanya pemain #%d yang memiliki dadu", pemain))
			fmt.Println(fmt.Sprintf("game dimenangkan oleh pemain #%d karena memiliki point yang lebih banyak skor : %d", pemenang, skor))
			break
		}
		stop = len(pemainDadu)
	}
}

func contains(s []string, str string) ([]string, []string) {
	var new []string
	var notI []string
	for _, v := range s {
		if v == str {
			new = append(new, v)
		} else {
			notI = append(notI, v)
		}
	}

	return new, notI
}
