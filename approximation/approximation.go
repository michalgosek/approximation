package approximation

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func CreateMatrixS(x []float64, m int) [][]float64 {
	s := 2 * m

	var sk []float64

	for k := 0; k <= s; k++ {
		var rec float64
		for i := 0; i < len(x); i++ {
			rec += math.Pow(x[i], float64(k))
		}
		sk = append(sk, rec)
	}

	sMatrix := make([][]float64, m+1)

	for i := 0; i <= m; i++ {
		var row []float64
		for j := 0; j <= m; j++ {
			row = append(row, sk[j+i])
		}
		sMatrix[i] = row
	}

	return sMatrix
}

func CreateTMatrix(x, y []float64, m int) [][]float64 {
	var tk []float64
	for k := 0; k <= m; k++ {
		var result float64
		for i := 0; i < len(x); i++ {
			result += math.Pow(x[i], float64(k)) * float64(y[i])
		}
		tk = append(tk, result)
	}

	tMatrix := make([][]float64, len(tk))
	for k := 0; k < len(tk); k++ {
		tMatrix[k] = []float64{tk[k]}
	}

	return tMatrix
}

func CreateMinor(s [][]float64, r, c int) [][]float64 {
	minor := make([][]float64, len(s)-1)

	var idx int

	for i := 0; i < len(s); i++ {
		if i == r {
			continue
		}
		var row []float64
		for j := 0; j < len(s); j++ {
			if j == c {
				continue
			}
			row = append(row, s[i][j])
		}
		minor[idx] = row
		idx++
	}

	return minor
}

func CalculateMatrixDet(matrix [][]float64) float64 {
	var determinant float64

	for i := 0; i < len(matrix); i++ {
		if len(matrix) > 2 {
			minor := CreateMinor(matrix, i, 0)
			e := matrix[i][0]
			p := math.Pow(-1, float64(i))
			determinant += e * p * CalculateMatrixDet(minor)
		} else if len(matrix) == 2 {
			determinant = matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
		}
	}
	return determinant
}

func CalculateCofactor(matrix [][]float64, i, j int) float64 {
	min := CreateMinor(matrix, i, j)

	if len(min) > 1 {
		det := CalculateMatrixDet(min)
		return math.Pow(-1, float64(i+j)) * det
	}

	return math.Pow(-1, float64(i+j)) * min[0][0]
}

func TransposeMatrix(matrix [][]float64) [][]float64 {
	t := make([][]float64, len(matrix))

	for i := 0; i < len(matrix); i++ {
		t[i] = make([]float64, len(matrix))
	}

	var c int
	for i := 0; i < len(matrix); i++ {
		var row []float64

		r := 0

		for j := 0; j < len(matrix); j++ {
			row = append(row, matrix[i][j])
		}

		for k := 0; k < len(row); k++ {
			t[r][c] = row[k]
			r++
		}
		c++

	}

	return t
}

func CreateCofactorMatrix(matrix [][]float64) [][]float64 {
	cofactor := make([][]float64, len(matrix))

	for i := 0; i < len(cofactor); i++ {
		cofactor[i] = make([]float64, len(matrix))
	}

	var r int
	for i := 0; i < len(matrix); i++ {
		var cofactors []float64
		for j := 0; j < len(matrix); j++ {
			a := CalculateCofactor(matrix, r, j)
			cofactors = append(cofactors, a)
		}

		cofactor[i] = cofactors
		r++
	}

	return cofactor
}

func CreateInverseMatrix(matrix [][]float64, det float64) [][]float64 {
	// macierz dołączona :)
	adjugateMatrix := TransposeMatrix(matrix)

	inverseMatrix := make([][]float64, len(matrix))

	for i := 0; i < len(inverseMatrix); i++ {
		inverseMatrix[i] = make([]float64, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {

			v := adjugateMatrix[i][j] / det
			adjugateMatrix[i][j] = v
		}
	}
	return adjugateMatrix
}

func FindPolynomialValues(inverseMatrix, tMatrix [][]float64) []float64 {
	var values []float64

	for i := 0; i < len(inverseMatrix); i++ {
		var p float64
		for j := 0; j < len(tMatrix); j++ {
			v := inverseMatrix[i][j] * tMatrix[j][0]
			s := fmt.Sprintf("%.3f", v)

			parsed, err := strconv.ParseFloat(s, 10)
			if err != nil {
				panic(err)
			}

			p += parsed

		}
		values = append(values, p)

	}

	return values
}

func FindApproximationFunc(x, y []float64, m int) {
	sMatirx := CreateMatrixS(x, m)
	tMatrix := CreateTMatrix(x, y, m)

	cofactorMatrixS := CreateCofactorMatrix(sMatirx)
	transposedMatrixS := TransposeMatrix(cofactorMatrixS)

	detS := CalculateMatrixDet(sMatirx)

	inverseMatrix := CreateInverseMatrix(transposedMatrixS, detS)

	values := FindPolynomialValues(inverseMatrix, tMatrix)

	var sb strings.Builder

	for i := 0; i < len(values); i++ {
		var s string
		if i == 0 {
			s = fmt.Sprintf("%f", values[i])
		} else {
			s = fmt.Sprintf("%fx^%d", values[i], i)
		}

		if i != len(values) && i > 0 {
			_, err := sb.WriteString("+")
			if err != nil {
				panic(err)
			}
		}

		_, err := sb.WriteString(s)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("F(x) =", sb.String())
}
