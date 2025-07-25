// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"blog/internal/dao/internal"
)

// sentenceTagDao is the data access object for the table sentence_tag.
// You can define custom methods on it to extend its functionality as needed.
type sentenceTagDao struct {
	*internal.SentenceTagDao
}

var (
	// SentenceTag is a globally accessible object for table sentence_tag operations.
	SentenceTag = sentenceTagDao{internal.NewSentenceTagDao()}
)

// Add your custom methods and functionality below.
