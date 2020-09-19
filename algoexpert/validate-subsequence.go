package main

func IsValidSubsequence(array []int, sequence []int) bool {
	arrIdx, seqIdx := 0, 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}
	return seqIdx == len(sequence)
}
