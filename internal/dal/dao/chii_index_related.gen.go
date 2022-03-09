// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameIndexSubject = "chii_index_related"

// IndexSubject mapped from table <chii_index_related>
type IndexSubject struct {
	ID       uint32  `gorm:"column:idx_rlt_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"idx_rlt_id"`
	Cat      int8    `gorm:"column:idx_rlt_cat;type:tinyint(3);not null;index:idx_rlt_cat,priority:1;index:idx_order,priority:2" json:"idx_rlt_cat"`
	Rid      uint32  `gorm:"column:idx_rlt_rid;type:mediumint(8) unsigned;not null;index:idx_rlt_rid,priority:1;index:idx_rlt_sid,priority:1;index:idx_order,priority:1" json:"idx_rlt_rid"`   // 关联目录
	Type     uint8   `gorm:"column:idx_rlt_type;type:smallint(6) unsigned;not null;index:idx_rlt_rid,priority:2" json:"idx_rlt_type"`                                                          // 关联条目类型
	Sid      uint32  `gorm:"column:idx_rlt_sid;type:mediumint(8) unsigned;not null;index:idx_rlt_sid,priority:2;index:idx_rlt_sid_2,priority:1;index:idx_order,priority:4" json:"idx_rlt_sid"` // 关联条目ID
	Order    uint32  `gorm:"column:idx_rlt_order;type:mediumint(8) unsigned;not null;index:idx_order,priority:3;default:0" json:"idx_rlt_order"`
	Comment  string  `gorm:"column:idx_rlt_comment;type:mediumtext;not null" json:"idx_rlt_comment"`
	Dateline int32   `gorm:"column:idx_rlt_dateline;type:int(10) unsigned;not null" json:"idx_rlt_dateline"`
	Subject  Subject `gorm:"foreignKey:idx_rlt_sid;references:subject_id" json:"subject"`
}

// TableName IndexSubject's table name
func (*IndexSubject) TableName() string {
	return TableNameIndexSubject
}
