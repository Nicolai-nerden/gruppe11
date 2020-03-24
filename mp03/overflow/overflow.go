package overflow

//Funksjoner som sjekker overflows for ulike inttyper.
// Kan evt hardcode inn maxLimit for vartypes for å korte ned prosessering betraketlig. Tall med flere enn 8 siffer bruker svært lang tid på å regne ut overflowgrense.

func Int8Overflow(sum int) bool {
	var a int8 = 1
	var b int8 = 2

	for a < b {
		if b <= 0 {
			break
		} else {
			a = b
			b++
		}
	}
	if sum > int(a) {
		return true
	}
	return false
}

func Int32Overflow(sum int) bool {
	var a int32 = 1
	var b int32 = 2

	for a < b {
		if b <= 0 {
			break
		} else {
			a = b
			b++
		}
	}

	if sum > int(a) {
		return true
	}
	return false
}

func Uint32Overflow(sum int) bool {
	var a uint32 = 1
	var b uint32 = 2

	for a < b {
		if b == 0 {
			break
		} else {
			a = b
			b++
		}
	}

	if sum > int(a) {
		return true
	}
	return false
}

func Int64Overflow(sum int) bool {
	var a int64 = 1
	var b int64 = 2

	for a < b {
		if b <= 0 {
			break
		} else {
			a = b
			b++
		}
	}

	if sum > int(a) {
		return true
	}
	return false
}
