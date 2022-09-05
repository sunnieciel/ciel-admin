// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoldStatisticsLog is the golang structure for table gold_statistics_log.
type GoldStatisticsLog struct {
	Id          uint64      `json:"id"          description:""`
	Uid         uint64      `json:"uid"         description:""`
	T1          float64     `json:"t1"          description:""`
	T2          float64     `json:"t2"          description:""`
	T3          float64     `json:"t3"          description:""`
	T4          float64     `json:"t4"          description:""`
	T5          float64     `json:"t5"          description:""`
	T6          float64     `json:"t6"          description:""`
	T7          float64     `json:"t7"          description:""`
	T8          float64     `json:"t8"          description:""`
	T9          float64     `json:"t9"          description:""`
	T10         float64     `json:"t10"         description:""`
	T11         float64     `json:"t11"         description:""`
	T12         float64     `json:"t12"         description:""`
	T13         float64     `json:"t13"         description:""`
	CreatedDate *gtime.Time `json:"createdDate" description:""`
}
