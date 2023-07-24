package e2e_test

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/cespare/permute"
)

var daten = [][]float64{
	{13.02, 27.73, 62.70, 105.90, 128.97, 132.70, 129.40, 106.26, 76.37, 42.73, 17.04, 9.39},  // 0
	{15.54, 31.44, 67.19, 110.52, 131.61, 134.21, 131.34, 109.39, 81.01, 47.50, 20.15, 11.60}, // 5,1
	{18.03, 34.92, 71.24, 114.48, 133.58, 135.12, 132.66, 111.95, 85.14, 51.93, 23.16, 13.84}, // 10,2
	{20.40, 38.14, 74.83, 117.77, 134.95, 135.40, 133.38, 113.96, 88.76, 55.98, 25.99, 16.01}, // 15,3
	{22.60, 41.07, 77.97, 120.42, 135.74, 135.12, 133.53, 115.42, 91.85, 59.64, 28.60, 18.04}, // 20,4
	{24.62, 43.70, 80.65, 122.39, 135.85, 134.45, 133.15, 116.28, 94.42, 62.89, 30.98, 19.91}, // 25,5
	{26.44, 46.03, 82.84, 123.70, 135.24, 133.04, 132.12, 116.54, 96.45, 65.72, 33.13, 21.61}, // 30,6
	{28.06, 48.03, 84.55, 124.33, 133.96, 130.95, 130.40, 116.21, 97.92, 68.13, 35.02, 23.14}, // 35,7
	{29.47, 49.71, 85.77, 124.28, 132.01, 128.23, 128.05, 115.29, 98.84, 70.10, 36.66, 24.49}, // 40,8
	{30.67, 51.07, 86.48, 123.55, 129.41, 124.90, 125.09, 113.77, 99.19, 71.62, 38.04, 25.65}, // 45,9
	{31.66, 52.09, 86.69, 122.14, 126.18, 121.02, 121.52, 111.66, 98.97, 72.69, 39.16, 26.63}, // 50,10
	{32.43, 52.79, 86.40, 120.01, 122.30, 116.51, 117.38, 108.95, 98.18, 73.31, 40.02, 27.43}, // 55,11
	{32.99, 53.15, 85.60, 117.17, 117.71, 111.47, 112.56, 105.61, 96.81, 73.48, 40.62, 28.03}, // 60,12
	{33.33, 53.18, 84.30, 113.68, 112.54, 105.84, 107.23, 101.75, 94.87, 73.19, 40.95, 28.44},
	{33.45, 52.88, 82.49, 109.52, 106.74, 99.65, 101.32, 97.32, 92.37, 72.45, 41.01, 28.65},
	{33.35, 52.25, 80.20, 104.69, 100.35, 92.95, 94.88, 92.34, 89.31, 71.25, 40.81, 28.68},
	{33.03, 51.30, 77.42, 99.21, 93.41, 85.82, 87.97, 86.84, 85.70, 69.61, 40.35, 28.51},
	{32.50, 50.02, 74.16, 93.11, 85.77, 77.71, 80.59, 80.84, 81.56, 67.52, 39.62, 28.15},
	{31.76, 48.42, 70.44, 86.41, 77.63, 69.30, 71.98, 74.41, 76.90, 65.00, 38.64, 27.60},
}

func MaxSum(t *testing.T, numPairs int) {
	t.Helper()

	monate := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	maxSum := 0.0
	maxMonat := 0
	maxWinkel := []int{}

	for offset := 0; offset < numPairs; offset++ {
		summe := 0.0
		winkel := make([]int, 12)

		if offset > 0 {
			monate = append(monate[1:], monate[0])
		}

		for i := 0; i < 12; i += numPairs {
			maxWinkel := 0
			max := 0.0

			for winkel := 0; winkel < 19; winkel++ {
				ertrag := 0.0
				for o := 0; o < numPairs; o++ {
					ertrag += daten[winkel][monate[i+o]]
				}

				if ertrag > max {
					max = ertrag
					maxWinkel = winkel * 5
				}
			}

			for j := i; j < i+numPairs; j++ {
				winkel[monate[j]] = maxWinkel
			}
			summe += max
		}

		if summe > maxSum {
			maxSum = summe
			maxMonat = monate[0] + 1
			maxWinkel = winkel
		}
	}
	t.Logf("### Max Summe (%.2f) Monate (%d) Start in (%d) Winkel (%v) ###", maxSum, numPairs, maxMonat, maxWinkel)
}

func MaxPerm(t *testing.T, perm []int) {
	t.Helper()

	monate := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	maxSum := 0.0
	maxWinkel := []int{}
	uMaxWinkel := 0

	for rotateMonths := 0; rotateMonths < 12; rotateMonths++ {
		summe := 0.0
		winkel := make([]int, 12)
		uWinkel := make(map[int]bool)

		if rotateMonths > 0 {
			monate = append(monate[1:], monate[0])
		}

		for i, offset := 0, 0; i < len(perm); i, offset = i+1, offset+perm[i] {
			maxWinkel := 0
			max := 0.0

			// for winkel := 0; winkel < 19; winkel++ {
			for winkel := 4; winkel <= 12; winkel += 4 {
				ertrag := 0.0

				for o := 0; o < perm[i]; o++ {
					ertrag += daten[winkel][monate[offset+o]]
				}

				if ertrag > max {
					max = ertrag
					maxWinkel = winkel * 5
				}
			}
			summe += max

			for o := 0; o < perm[i]; o++ {
				winkel[monate[offset+o]] = maxWinkel
				uWinkel[maxWinkel] = true
			}
		}

		if summe > maxSum {
			maxSum = summe
			maxWinkel = winkel
			uMaxWinkel = len(uWinkel)
		}
	}

	fmt.Printf("%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%s;;%v\n",
		uMaxWinkel, maxWinkel[0], maxWinkel[1], maxWinkel[2], maxWinkel[3], maxWinkel[4], maxWinkel[5],
		maxWinkel[6], maxWinkel[7], maxWinkel[8], maxWinkel[9], maxWinkel[10], maxWinkel[11],
		strings.Replace(strconv.FormatFloat(maxSum, 'f', 2, 32), ".", ",", 1), perm)
}

func TestPVMaxSum(t *testing.T) {
	MaxSum(t, 1)
	MaxSum(t, 2)
	MaxSum(t, 3)
	MaxSum(t, 4)
	MaxSum(t, 4)
	MaxSum(t, 6)
	MaxSum(t, 12)
}

func TestPVPerm(t *testing.T) {
	perms := [][]int{
		{6, 4, 2},
		{6, 2, 4},
		{6, 3, 3},
		{6, 2, 2, 2},
		{6, 3, 3},
		{5, 5, 2},
		{5, 2, 5},
		{5, 4, 3},
		{5, 3, 4},
		{5, 3, 2, 2},
		{5, 2, 3, 2},
		{5, 2, 2, 3},
		{4, 3, 3, 2},
		{4, 3, 2, 3},
		{4, 2, 3, 3},
		{4, 2, 2, 2, 2},
		{3, 3, 3, 3},
		{3, 3, 2, 2, 2},
		{3, 2, 3, 2, 2},
		{3, 2, 2, 3, 2},
		{3, 2, 2, 2, 3},
		{2, 2, 2, 2, 2, 2},
	}

	for _, perm := range perms {
		MaxPerm(t, perm)
	}
}

func partition(val int, minValue int) [][]int {
	set := make(map[string]bool)

	for i := minValue; i <= val; i++ {
		for _, result := range worker([]int{i}, val-i, minValue) {
			set[result] = true
		}
	}

	out := make([][]int, len(set))
	pos := 0

	for key := range set {
		pp := strings.Split(key, ",")
		pi := make([]int, len(pp))

		for i, p := range pp {
			pi[i], _ = strconv.Atoi(p)
		}
		out[pos] = pi
		pos++
	}

	return out
}

func worker(already []int, rest int, minValue int) []string {
	if rest == 0 {
		sort.Slice(already, func(i, j int) bool {
			return already[i] > already[j]
		})

		return []string{strings.Trim(strings.Join(strings.Fields(fmt.Sprint(already)), ","), "[]")}
	}

	out := make([]string, 0)
	for i := minValue; i <= rest; i++ {
		cp := make([]int, len(already))
		copy(cp, already)

		out = append(out, worker(append(cp, i), rest-i, minValue)...)
	}

	return out
}

func TestIterPerm(t *testing.T) {
	partitions := partition(12, 1)

	unique := make(map[string]bool, 0)

	for _, part := range partitions {
		first := part[0]
		fmt.Println("part", part)

		if first == 1 {
			continue
		}

		p := permute.Ints(part)
		for p.Permute() {
			if part[0] == first {
				pp := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(part)), ","), "[]")

				if _, ok := unique[pp]; !ok {
					unique[pp] = true
				}
			}
		}
	}

	// clean
	for perm := range unique {
		pp := strings.Split(perm, ",")
		pi := make([]int, len(pp))

		for i, p := range pp {
			pi[i], _ = strconv.Atoi(p)
		}

		MaxPerm(t, pi)
	}
}
