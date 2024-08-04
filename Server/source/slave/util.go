package slave

func xorbytes(chars []byte, key int) []byte {

	var xoredBytes []byte

	for _, character := range chars {
		xoredBytes = append(xoredBytes, character^byte(key))
	}

	return xoredBytes

}
