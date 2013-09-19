parallel_sort
=============

Parallel sorting library for Google Go.

Usage
=====

This library implements several sequential and parallel sorting algorithms for integer slices, arbitrary type slices (with the reflect library), and for any collection in the same style as the sort library of the Go language.

To use it just create an instance of the algorithm and call the Sort method, like in the following example:

  s := []int{5, 1, 9, 10}
	qs := integers.QuickSortParallel {}
	qs.SetNumCPU(4)
	qs.Sort(s)
	
