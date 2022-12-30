package model

import "ciel-admin/internal/model/entity"

type PracticeWordLevel struct {
	Level uint   `json:"level"`
	Class string `json:"class"`
	Desc  string `json:"desc"`
	Msg   string `json:"msg"`
}

type EnglishDocument struct {
	Info  *entity.EnDocument
	Level *entity.EnDocumentUser
}
type EnglishDocumentParagraph struct {
	Info  *entity.EnDocumentParagraph
	Level *entity.EnDocumentParagraphUser
}
type EnglishParagraphWord struct {
	Info  *entity.EnDocumentParagraphWord
	Level *entity.EnDocumentParagraphWordUser
}
