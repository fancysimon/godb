// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package comparator

// Comparator provides a total order across slices that are
// used as keys in an sstable or a database.
type Comparator interface {
	// Compare between |first| and |second|.
	// Returns value:
	//   < 0 iff |first| < |second|,
	//   == 0 iff |first| == |second|,
	//   > 0 iff |first| > |second|
	Compare(first, second string) int
	// The name of the comparator.  Used to check for comparator
    // mismatches (i.e., a DB created with one comparator is
    // accessed using a different comparator.
	Name() string
}

type BytewiseComparator struct {}

func (BytewiseComparator) Compare(first, second string) int {
	if first < second {
		return -1
	} else if first > second {
		return 1
	}
	return 0
}

func (BytewiseComparator) Name() string {
	return "godb.BytewiseComparator"
}