// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ciel-admin/internal/dao/internal"
)

// internalEnDocumentParagraphWordUserDao is internal type for wrapping internal DAO implements.
type internalEnDocumentParagraphWordUserDao = *internal.EnDocumentParagraphWordUserDao

// enDocumentParagraphWordUserDao is the data access object for table e_en_document_paragraph_word_user.
// You can define custom methods on it to extend its functionality as you wish.
type enDocumentParagraphWordUserDao struct {
	internalEnDocumentParagraphWordUserDao
}

var (
	// EnDocumentParagraphWordUser is globally public accessible object for table e_en_document_paragraph_word_user operations.
	EnDocumentParagraphWordUser = enDocumentParagraphWordUserDao{
		internal.NewEnDocumentParagraphWordUserDao(),
	}
)

// Fill with you ideas below.
