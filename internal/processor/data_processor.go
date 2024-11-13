package processor

func ProcessData(data []byte) ([]byte, error) {
	// Simulate data transformation
	return append(data, []byte(" processed")...), nil
}
