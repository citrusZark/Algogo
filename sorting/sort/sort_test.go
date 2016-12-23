package sort_test

import (
	"algorityhm/sorting/sort"
	"math/rand"
	s "sort"
	"testing"
	"time"

	"github.com/cheekybits/is"
)

type person struct {
	Name string
	Age  int
	job
}
type job struct {
	Title  string
	Salary float32
}
type aray []string

func (a aray) Len() int {
	return len(a)
}
func (a aray) Less(i, j int) bool {
	return a[i] < a[j]
}
func (x aray) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func TestSelectionSort(t *testing.T) {
	is := is.New(t)
	var sort sort.SelectionSort
	ss := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	si := []int{3, 1, 4, 5, 6, 7, 1, 4, 8, 9, 0, 2}
	sf := []float32{12.2, 34.5, 67.8, 1.2}
	ns := 3
	us := []bool{true, false, true}
	sstruct := []person{
		{
			Name: "agung",
			Age:  34,
			job: job{
				Title:  "Dev",
				Salary: 1000.00,
			},
		},
		{
			Name: "budi",
			Age:  12,
			job: job{
				Title:  "AM",
				Salary: 2500.00,
			},
		},
	}
	sort.Sort(sstruct, "Title")
	sort.Show(sstruct)
	is.True(sort.IsSorted(sstruct, "Title"))
	sort.Sort(ss, "")
	sort.Show(ss)
	is.True(sort.IsSorted(ss, ""))
	sort.Sort(si, "")
	sort.Show(si)
	is.True(sort.IsSorted(si, ""))
	sort.Sort(sf, "")
	sort.Show(sf)
	is.True(sort.IsSorted(sf, ""))
	is.PanicWith("error: only take slice as argument", func() { sort.Sort(ns, "") })
	is.PanicWith("Unsupported data or struct's field not found", func() { sort.Sort(us, "") })
}

func TestInsertionSort(t *testing.T) {
	is := is.New(t)
	var sort sort.InsertionSort
	ss := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	si := []int{3, 1, 4, 5, 6, 7, 1, 4, 8, 9, 0, 2}
	sf := []float32{12.2, 34.5, 67.8, 1.2}
	ns := 3
	us := []bool{true, false, true}
	sstruct := []person{
		{
			Name: "agung",
			Age:  34,
			job: job{
				Title:  "Dev",
				Salary: 1000.00,
			},
		},
		{
			Name: "budi",
			Age:  12,
			job: job{
				Title:  "AM",
				Salary: 2500.00,
			},
		},
	}
	sort.Sort(sstruct, "Title")
	sort.Show(sstruct)
	is.True(sort.IsSorted(sstruct, "Title"))
	sort.Sort(ss, "")
	sort.Show(ss)
	is.True(sort.IsSorted(ss, ""))
	sort.Sort(si, "")
	sort.Show(si)
	is.True(sort.IsSorted(si, ""))
	sort.Sort(sf, "")
	sort.Show(sf)
	is.True(sort.IsSorted(sf, ""))
	is.PanicWith("error: only take slice as argument", func() { sort.Sort(ns, "") })
	is.PanicWith("Unsupported data or struct's field not found", func() { sort.Sort(us, "") })
}

func TestMergeSortUpBottom(t *testing.T) {
	is := is.New(t)
	var sort sort.MergeSortUpBottom
	ss := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	si := []int{3, 1, 4, 5, 6, 7, 1, 4, 8, 9, 0, 2}
	sf := []float32{12.2, 34.5, 67.8, 1.2}
	ns := 3
	us := []bool{true, false, true}
	sstruct := []person{
		{
			Name: "agung",
			Age:  34,
			job: job{
				Title:  "Dev",
				Salary: 1000.00,
			},
		},
		{
			Name: "budi",
			Age:  12,
			job: job{
				Title:  "AM",
				Salary: 2500.00,
			},
		},
	}
	sort.Sort(sstruct, "Title")
	sort.Show(sstruct)
	is.True(sort.IsSorted(sstruct, "Title"))
	sort.Sort(ss, "")
	sort.Show(ss)
	is.True(sort.IsSorted(ss, ""))
	sort.Sort(si, "")
	sort.Show(si)
	is.True(sort.IsSorted(si, ""))
	sort.Sort(sf, "")
	sort.Show(sf)
	is.True(sort.IsSorted(sf, ""))
	is.PanicWith("error: only take slice as argument", func() { sort.Sort(ns, "") })
	is.PanicWith("Unsupported data or struct's field not found", func() { sort.Sort(us, "") })
}

func TestMergeSortBottomUp(t *testing.T) {
	is := is.New(t)
	var sort sort.MergeSortBottomUp
	ss := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	si := []int{3, 1, 4, 5, 6, 7, 1, 4, 8, 9, 0, 2}
	sf := []float32{12.2, 34.5, 67.8, 1.2}
	ns := 3
	us := []bool{true, false, true}
	sstruct := []person{
		{
			Name: "agung",
			Age:  34,
			job: job{
				Title:  "Dev",
				Salary: 1000.00,
			},
		},
		{
			Name: "budi",
			Age:  12,
			job: job{
				Title:  "AM",
				Salary: 2500.00,
			},
		},
	}
	sort.Sort(sstruct, "Title")
	sort.Show(sstruct)
	is.True(sort.IsSorted(sstruct, "Title"))
	sort.Sort(ss, "")
	sort.Show(ss)
	is.True(sort.IsSorted(ss, ""))
	sort.Sort(si, "")
	sort.Show(si)
	is.True(sort.IsSorted(si, ""))
	sort.Sort(sf, "")
	sort.Show(sf)
	is.True(sort.IsSorted(sf, ""))
	is.PanicWith("error: only take slice as argument", func() { sort.Sort(ns, "") })
	is.PanicWith("Unsupported data or struct's field not found", func() { sort.Sort(us, "") })
}

func TestQuickSort(t *testing.T) {
	is := is.New(t)
	var sort sort.QuickSort
	ss := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	si := []int{3, 1, 4, 5, 6, 7, 1, 4, 8, 9, 0, 2}
	sf := []float32{12.2, 34.5, 67.8, 1.2}
	ns := 3
	us := []bool{true, false, true}
	sstruct := []person{
		{
			Name: "agung",
			Age:  34,
			job: job{
				Title:  "Dev",
				Salary: 1000.00,
			},
		},
		{
			Name: "budi",
			Age:  12,
			job: job{
				Title:  "AM",
				Salary: 2500.00,
			},
		},
	}
	sort.Sort(sstruct, "Title")
	sort.Show(sstruct)
	is.True(sort.IsSorted(sstruct, "Title"))
	sort.Sort(ss, "")
	sort.Show(ss)
	is.True(sort.IsSorted(ss, ""))
	sort.Sort(si, "")
	sort.Show(si)
	is.True(sort.IsSorted(si, ""))
	sort.Sort(sf, "")
	sort.Show(sf)
	is.True(sort.IsSorted(sf, ""))
	is.PanicWith("error: only take slice as argument", func() { sort.Sort(ns, "") })
	is.PanicWith("Unsupported data or struct's field not found", func() { sort.Sort(us, "") })
}
func TestHeapSort(t *testing.T) {
	is := is.New(t)
	var sort sort.HeapSort
	ss := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}
	si := []int{3, 1, 4, 5, 6, 7, 1, 4, 8, 9, 0, 2}
	sf := []float32{12.2, 34.5, 67.8, 1.2}
	ns := 3
	us := []bool{true, false, true}
	sstruct := []person{
		{
			Name: "agung",
			Age:  34,
			job: job{
				Title:  "Dev",
				Salary: 1000.00,
			},
		},
		{
			Name: "budi",
			Age:  12,
			job: job{
				Title:  "AM",
				Salary: 2500.00,
			},
		},
	}
	sort.Sort(sstruct, "Title")
	sort.Show(sstruct)
	is.True(sort.IsSorted(sstruct, "Title"))
	sort.Sort(ss, "")
	sort.Show(ss)
	is.True(sort.IsSorted(ss, ""))
	sort.Sort(si, "")
	sort.Show(si)
	is.True(sort.IsSorted(si, ""))
	sort.Sort(sf, "")
	sort.Show(sf)
	is.True(sort.IsSorted(sf, ""))
	is.PanicWith("error: only take slice as argument", func() { sort.Sort(ns, "") })
	is.PanicWith("Unsupported data or struct's field not found", func() { sort.Sort(us, "") })
}

//func BenchmarkSelectionSort(b *testing.B) {
//	rand.Seed(time.Now().UTC().UnixNano())
//	var randomStringSlice []string
//	for i := 0; i < 1000000; i++ {
//		randomStringSlice = append(randomStringSlice, randomString(1))
//	}
//	var sort sort.SelectionSort
//	sort.Sort(randomStringSlice, "")
//}
func BenchmarkInsertionSort(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Now().UTC().UnixNano())
	var randomStringSlice []string
	for i := 0; i < 10000000; i++ {
		randomStringSlice = append(randomStringSlice, randomString(1))
	}

	b.StartTimer()
	var sort sort.InsertionSort
	sort.Sort(randomStringSlice, "")
	b.StopTimer()
}

func BenchmarkMergeSortUpBottom(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Now().UTC().UnixNano())
	var randomStringSlice []string
	for i := 0; i < 10000000; i++ {
		randomStringSlice = append(randomStringSlice, randomString(1))
	}

	b.StartTimer()
	var sort sort.MergeSortUpBottom
	sort.Sort(randomStringSlice, "")
	b.StopTimer()
}
func BenchmarkMergeSortBottomUp(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Now().UTC().UnixNano())
	var randomStringSlice []string
	for i := 0; i < 10000000; i++ {
		randomStringSlice = append(randomStringSlice, randomString(1))
	}

	b.StartTimer()
	var sort sort.MergeSortBottomUp
	sort.Sort(randomStringSlice, "")
	b.StopTimer()
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Now().UTC().UnixNano())
	var randomStringSlice []string
	for i := 0; i < 10000000; i++ {
		randomStringSlice = append(randomStringSlice, randomString(1))
	}
	s.Sort(aray(randomStringSlice))
	b.StartTimer()
	b.StopTimer()
}

/*func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Now().UTC().UnixNano())
	var randomStringSlice []string
	for i := 0; i < 1000; i++ {
		randomStringSlice = append(randomStringSlice, randomString(1))
	}

	b.StartTimer()
	var sort sort.QuickSort
	sort.Sort(randomStringSlice, "")
	b.StopTimer()
}*/

/*func BenchmarkHeapSort(b *testing.B) {
	b.StopTimer()
	rand.Seed(time.Now().UTC().UnixNano())
	var randomStringSlice []string
	for i := 0; i < 10000000; i++ {
		randomStringSlice = append(randomStringSlice, randomString(1))
	}

	b.StartTimer()
	var sort sort.HeapSort
	sort.Sort(randomStringSlice, "")
	b.StopTimer()
}*/

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
