package lazypdf

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveToPNGOK(t *testing.T) {
	file, err := os.Open("testdata/sample4.pdf")
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	require.NoError(t, err)
	defer func() { require.NoError(t, file.Close()) }()
	for i := uint16(0); i < 13; i++ {
		require.NoError(t, err)

		buf := bytes.NewBuffer([]byte{})
		err = SaveToPNG(context.Background(), i, buffer.Bytes(), buf)
		require.NoError(t, err)

		// expectedPage, err := ioutil.ReadFile(fmt.Sprintf("testdata/test2%d.png", i))
		// require.NoError(t, err)
		out, err := os.Create(fmt.Sprintf("testdata/sample123_%d.png", i))
		require.NoError(t, err)
		_, err = out.Write(buf.Bytes())
		require.NoError(t, err)
		// resultPage, err := io.ReadAll(buf)
		// require.NoError(t, err)
		// require.Equal(t, expectedPage, resultPage)
	}
}

func TestSaveToPNGFail(t *testing.T) {
	file, err := os.Open("testdata/sample4.pdf")
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	require.NoError(t, err)
	defer func() { require.NoError(t, file.Close()) }()

	err = SaveToPNG(context.Background(), 0, buffer.Bytes(), bytes.NewBuffer([]byte{}))
	require.Error(t, err)
	require.Equal(t, "failure at the C/MuPDF layer: no objects found", err.Error())
}

func TestPageCount(t *testing.T) {
	file, err := os.Open("testdata/sample4.pdf")
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	require.NoError(t, err)
	defer func() { require.NoError(t, file.Close()) }()

	count, err := PageCount(context.Background(), buffer.Bytes())
	require.NoError(t, err)
	require.Equal(t, 13, count)
}

func TestPageCountFail(t *testing.T) {
	file, err := os.Open("testdata/sample4.pdf")
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	require.NoError(t, err)
	defer func() { require.NoError(t, file.Close()) }()

	_, err = PageCount(context.Background(), buffer.Bytes())
	require.Error(t, err)
	require.Equal(t, "failure at the C/MuPDF layer: no objects found", err.Error())
}

func BenchmarkSaveToPNGPage0(b *testing.B)  { benchmarkSaveToPNGRunner(0, b) }
func BenchmarkSaveToPNGPage1(b *testing.B)  { benchmarkSaveToPNGRunner(1, b) }
func BenchmarkSaveToPNGPage2(b *testing.B)  { benchmarkSaveToPNGRunner(2, b) }
func BenchmarkSaveToPNGPage3(b *testing.B)  { benchmarkSaveToPNGRunner(3, b) }
func BenchmarkSaveToPNGPage4(b *testing.B)  { benchmarkSaveToPNGRunner(4, b) }
func BenchmarkSaveToPNGPage5(b *testing.B)  { benchmarkSaveToPNGRunner(5, b) }
func BenchmarkSaveToPNGPage6(b *testing.B)  { benchmarkSaveToPNGRunner(6, b) }
func BenchmarkSaveToPNGPage7(b *testing.B)  { benchmarkSaveToPNGRunner(7, b) }
func BenchmarkSaveToPNGPage8(b *testing.B)  { benchmarkSaveToPNGRunner(8, b) }
func BenchmarkSaveToPNGPage9(b *testing.B)  { benchmarkSaveToPNGRunner(9, b) }
func BenchmarkSaveToPNGPage10(b *testing.B) { benchmarkSaveToPNGRunner(10, b) }
func BenchmarkSaveToPNGPage11(b *testing.B) { benchmarkSaveToPNGRunner(11, b) }
func BenchmarkSaveToPNGPage12(b *testing.B) { benchmarkSaveToPNGRunner(12, b) }

func benchmarkSaveToPNGRunner(page uint16, b *testing.B) {
	file, err := os.Open("testdata/sample4.pdf")
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	require.NoError(b, err)
	defer func() { require.NoError(b, file.Close()) }()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		output := bytes.NewBuffer([]byte{})
		err := SaveToPNG(context.Background(), page, buffer.Bytes(), output)
		require.NoError(b, err)
	}
}
