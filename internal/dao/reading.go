// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"blog/internal/dao/internal"
)

// readingDao is the data access object for the table reading.
// You can define custom methods on it to extend its functionality as needed.
type readingDao struct {
	*internal.ReadingDao
}

var (
	// Reading is a globally accessible object for table reading operations.
	Reading = readingDao{internal.NewReadingDao()}
)

// Add your custom methods and functionality below.
