package main

func toStringUnit(in int) string {
	switch in {
	case 1:
		return "satu"
	case 2:
		return "dua"
	case 3:
		return "tiga"
	case 4:
		return "empat"
	case 5:
		return "lima"
	case 6:
		return "enam"
	case 7:
		return "tujuh"
	case 8:
		return "delapan"
	case 9:
		return "sembilan"
	case 10:
		return "sepuluh"
	case 11:
		return "sebelas"
	default:
		return "null"
	}
}

func toStringUtils(in int, temps string) string {
	if in < 12 {
		temps = temps + toStringUnit(in)
	} else if in < 20 {
		temps = temps + toStringUnit(in%10) + " belas"
	} else if in < 100 {
		temps = temps + toStringUnit(in/10) + " puluh " + toStringUnit(in%10)
	} else if in/100 == 1 {
		if in == 100 {
			temps = "seratus"
		} else {
			temps = temps + "seratus " + toStringUtils(in%100, temps)
		}
	} else {
		temps = temps + toStringUnit(in/100) + " ratus " + toStringUtils(in%100, temps)
	}
	return temps
}

func spellUtils(in int) string {
	var temps = ""
	var res = ""
	for in > 0 {
		if in/1000000000 > 0 {
			res = res + toStringUtils(in/1000000000, temps) + " milyar "
			in = in % 1000000000
		} else if in/1000000 > 0 {
			res = res + toStringUtils(in/1000000, temps) + " juta "
			in = in % 1000000
		} else if in/1000 == 1 {
			res = res + "seribu "
			in = in % 1000
		} else if in/1000 > 0 {
			res = res + toStringUtils(in/1000, temps) + " ribu "
			in = in % 1000
		} else {
			res = res + toStringUtils(in, temps)
			in = in / 1000
		}
	}
	return res
}
