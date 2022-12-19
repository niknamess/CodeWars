package main

func Tribonacci(signature [3]float64, n int) []float64 {
	// your code here

	switch n {
	case 0:
		sam := make([]float64, 0)
		return sam
	case 1:
		var sam []float64

		sam = append(sam, signature[0])

		return sam
	case 2:
		var sam []float64

		sam = append(sam, signature[0])
		sam = append(sam, signature[1])
		return sam
	default:
		var sam []float64

		sam = append(sam, signature[0])
		sam = append(sam, signature[1])
		sam = append(sam, signature[2])
		for i := 3; i < n; i++ {
			sam = append(sam[:], (sam[i-1] + sam[i-2] + sam[i-3]))

		}
		return sam
	}
}
