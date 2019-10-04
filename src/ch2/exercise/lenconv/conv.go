package lenconv

func MtoDM(m M) DM {
	return DM(m * 10)
}

func DMtoCM(dm DM) CM {
	return CM(dm * 10)
}
