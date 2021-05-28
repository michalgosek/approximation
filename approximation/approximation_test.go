package approximation_test

import (
	"lab3/approximation"

	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_CreateMatrixS(t *testing.T) {
	x := []float64{-1.8, -1, 0.2, 1.2, 2}

	want := [][]float64{
		{5, 0.6000000000000003, 9.72},
		{0.6000000000000003, 9.72, 2.903999999999999},
		{9.72, 2.903999999999999, 29.5728},
	}

	got := approximation.CreateMatrixS(x, 2)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatal(diff)
	}
}

func Test_CreateMatrixT(t *testing.T) {
	x := []float64{-1.8, -1, 0.2, 1.2, 2}
	y := []float64{5.68, 1.99, 1.24, 9.81, 14.5}

	want := [][]float64{
		{33.22},
		{28.805999999999997},
	}

	got := approximation.CreateTMatrix(x, y, 1)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatal(diff)
	}
}

func Test_CreateMinor(t *testing.T) {
	tests := []struct {
		name  string
		input [][]float64
		r     int
		c     int
		want  [][]float64
	}{
		{
			name: "positive should return 2x2 row1: [3 4]; row2 [2 8] for row 1; col 0; to elimnate",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: [][]float64{
				{3, 4},
				{3, 8},
			},
			r: 1,
			c: 0,
		},
		{
			name: "positive should return 2x2 row1: [2 2]; row2 [3 8] for row 0; col 0; to elimnate",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: [][]float64{
				{2, 2},
				{3, 8},
			},
			r: 0,
			c: 0,
		},
		{
			name: "positive should return 2x2 row1: [2 2]; row2 [1 8] for row 1; col 2; to elimnate",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: [][]float64{
				{2, 2},
				{1, 8},
			},
			r: 0,
			c: 1,
		},
		{
			name: "positive should return 2x2 row1: [2 3]; row2 [1 3] for row 2; col 3; to elimnate",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: [][]float64{
				{2, 3},
				{1, 3},
			},
			r: 1,
			c: 2,
		},
		{
			name: "positive should return 2x2 row1: [2 3]; row2 [1 3] for row 3; col 3; to elimnate",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: [][]float64{
				{2, 3},
				{2, 2},
			},
			r: 2,
			c: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := approximation.CreateMinor(tt.input, tt.r, tt.c)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func Test_CalculateMatrixDet(t *testing.T) {
	tests := []struct {
		name  string
		input [][]float64
		want  float64
	}{
		{
			name: "should return 10",
			input: [][]float64{
				{2, 2},
				{3, 8},
			},
			want: 10,
		},
		{
			name: "should return -6",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: -6,
		},
		{
			name: "should return -100",
			input: [][]float64{
				{5, 2, 0, 0, -2},
				{0, 1, 4, 3, 2},
				{0, 0, 2, 6, 3},
				{0, 0, 3, 4, 1},
				{0, 0, 0, 0, 2},
			},
			want: -100,
		},
		{
			name: "should return 400",
			input: [][]float64{
				{1, 2, 3, 4},
				{0, 5, 6, 7},
				{0, 0, 8, 9},
				{0, 0, 0, 10},
			},
			want: 400,
		},
		{
			name: "should return 501.34955200000013",
			input: [][]float64{
				{5, 0.6, 9.72},
				{0.6, 9.72, 2.9},
				{9.72, 2.9, 29.6},
			},
			want: 501.34955200000013,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := approximation.CalculateMatrixDet(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func Test_TransposeMatrix(t *testing.T) {
	tests := []struct {
		name  string
		input [][]float64
		want  [][]float64
	}{
		{
			name: "should transpose matrix {2,3,4}{2,2,2}{1,3,8} to {2,3,4}{3,2,3}{4,2,8}",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: [][]float64{
				{2, 2, 1},
				{3, 2, 3},
				{4, 2, 8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := approximation.TransposeMatrix(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func Test_CalculateCofactor(t *testing.T) {
	tests := []struct {
		name  string
		i, j  int
		want  float64
		input [][]float64
	}{
		{
			name: "should return {9.72}",
			input: [][]float64{
				{5, 0.6},
				{0.6, 9.72},
			},
			want: 9.72,
		},
		{
			name: "should return 10 for a11",
			i:    0, // arrays index starts from 0
			j:    0,
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: 10,
		},
		{
			name: "should return 10 for a32",
			i:    2, // arrays index starts from 0
			j:    1,
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := approximation.CalculateCofactor(tt.input, tt.i, tt.j)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func Test_CreateCofactorMatrix(t *testing.T) {
	tests := []struct {
		name  string
		input [][]float64
		want  [][]float64
	}{
		{
			name: "should return {10,-14,4}{-12,12,-3}{-2,4,-2}",
			input: [][]float64{
				{2, 3, 4},
				{2, 2, 2},
				{1, 3, 8},
			},
			want: [][]float64{
				{10, -14, 4},
				{-12, 12, -3},
				{-2, 4, -2},
			},
		},
		{
			name: "should return {279,10.5,−93}{10.5,53.4,−8,7}{−93,−8.7,48.2}",
			input: [][]float64{
				{5, 0.6, 9.72},
				{0.6, 9.72, 2.9},
				{9.72, 2.9, 29.6},
			},
			want: [][]float64{
				{279.302, 10.428, -92.73840000000001},
				{10.428, 53.52159999999999, -8.668},
				{-92.73840000000001, -8.668, 48.24},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := approximation.CreateCofactorMatrix(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func Test_CreateInverseMatrix(t *testing.T) {
	tests := []struct {
		name  string
		input [][]float64
		want  [][]float64
		det   float64
	}{
		{
			name: "should return {0.55710032827555, 0.020799859017326894, -0.1849775264185336}{0.020799859017326894, 0.10675505699863472, -0.017289334288664124}{-0.1849775264185336, -0.017289334288664124, 0.09622029142652946}",
			input: [][]float64{
				{279.302, 10.428, -92.73840000000001},
				{10.428, 53.52159999999999, -8.668},
				{-92.73840000000001, -8.668, 48.24},
			},
			want: [][]float64{
				{0.55710032827555, 0.020799859017326894, -0.1849775264185336},
				{0.020799859017326894, 0.10675505699863472, -0.017289334288664124},
				{-0.1849775264185336, -0.017289334288664124, 0.09622029142652946},
			},
			det: 501.34955200000013,
		},
		{
			name: "should return {0.2016597510373444, -0.012448132780082987}{-0.012448132780082987, 0.10373443983402489}",
			input: [][]float64{
				{9.72, -0.6},
				{-0.6, 5},
			},
			det: 48.2,
			want: [][]float64{
				{0.2016597510373444, -0.012448132780082987},
				{-0.012448132780082987, 0.10373443983402489},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := approximation.CreateInverseMatrix(tt.input, tt.det)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func Test_FindPolynomialValues(t *testing.T) {
	tests := []struct {
		name          string
		matrixT       [][]float64
		inverseMatrix [][]float64
		want          []float64
	}{
		{
			name: "should return {6.329,2.5869999999999997}",
			matrixT: [][]float64{
				{33.2},
				{28.7},
			},
			inverseMatrix: [][]float64{
				{0.201, -0.012},
				{-0.012, 0.104},
			},
			want: []float64{6.329, 2.5869999999999997},
		},
		{
			name: "should return {1.9849999999999994, 2.1659999999999995, 2.266}",
			matrixT: [][]float64{
				{33.22},
				{28.805999999999997},
				{92.5692},
			},
			inverseMatrix: [][]float64{
				{0.5574890219560878, 0.02081437125748503, -0.18510658682634734},
				{0.02081437125748503, 0.10682954091816366, -0.01730139720558882},
				{-0.18510658682634734, -0.01730139720558882, 0.09628742514970061},
			},
			want: []float64{1.9849999999999994, 2.1659999999999995, 2.266},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := approximation.FindPolynomialValues(tt.inverseMatrix, tt.matrixT)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}

}

func Example_FindApproximationFunc() {
	x := []float64{-1.8, -1, 0.2, 1.2, 2}
	y := []float64{5.68, 1.99, 1.24, 9.81, 14.5}

	degree := 1
	approximation.FindApproximationFunc(x, y, degree)

	degree = 2
	approximation.FindApproximationFunc(x, y, degree)
	// Output:
	// F(x) = 6.336000+2.573000x^1
	// F(x) = 1.973000+2.164000x^1+2.269000x^2
}
