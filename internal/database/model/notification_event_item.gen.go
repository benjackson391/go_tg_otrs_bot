// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameNotificationEventItem = "notification_event_item"

// NotificationEventItem mapped from table <notification_event_item>
type NotificationEventItem struct {
	NotificationID int32  `gorm:"column:notification_id;not null" json:"notification_id"`
	EventKey       string `gorm:"column:event_key;not null" json:"event_key"`
	EventValue     string `gorm:"column:event_value;not null" json:"event_value"`
}

// TableName NotificationEventItem's table name
func (*NotificationEventItem) TableName() string {
	return TableNameNotificationEventItem
}
