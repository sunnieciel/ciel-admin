// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Read is the golang structure for table read.
type Read struct {
	Id          uint64      `json:"id"          description:""`
	Category1   string      `json:"category1"   description:""`
	Category2   string      `json:"category2"   description:""`
	Tag         string      `json:"tag"         description:""`
	Title       string      `json:"title"       description:""`
	Icon        string      `json:"icon"        description:""`
	Poster      string      `json:"poster"      description:""`
	Visit       uint64      `json:"visit"       description:""`
	Keywords    string      `json:"keywords"    description:""`
	Writer      string      `json:"writer"      description:""`
	Description string      `json:"description" description:""`
	Content     string      `json:"content"     description:""`
	Status      uint        `json:"status"      description:""`
	Sort        uint64      `json:"sort"        description:""`
	Recommend   uint        `json:"recommend"   description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
}
