package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// seperate is
// BaseModel to be emmbered to other struct as audit trail perpurse
type BaseModel struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
	CreatedBy string        `bson:"created_by,omitempty" json:"created_by"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
	UpdatedBy string        `bson:"updated_by,omitempty" json:"updated_by"`
	IsRemoved bool          `bson:"is_removed,omitempty" json:"is_removed"`
	RemovedAt time.Time     `bson:"removed_at,omitempty" json:"removed_at"`
	RemovedBy string        `bson:"removed_by,omitempty" json:"removed_by"`
}

// Camel case
// BaseModel to be emmbered to other struct as audit trail perpurse
type BaseMod struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time     `bson:"CreatedAt,omitempty" json:"created_at"`
	CreatedBy string        `bson:"CreatedBy,omitempty" json:"created_by"`
	UpdatedAt time.Time     `bson:"UpdatedAt,omitempty" json:"updated_at"`
	UpdatedBy string        `bson:"UpdatedBy,omitempty" json:"updated_by"`
	IsRemoved bool          `bson:"IsRemoved,omitempty" json:"is_removed"`
	RemovedAt time.Time     `bson:"RemovedAt,omitempty" json:"removed_at"`
	RemovedBy string        `bson:"RemovedBy,omitempty" json:"removed_by"`
}

//ChangeLog
type ChangeLog struct {
	BaseMod      `bson:",inline"`
	ModelObjId   bson.ObjectId `bson:"ModelObjId,omitempty" json:"model_obj_id"`
	ModelName    string        `bson:"ModelName,omitempty" json:"model_name"`
	ModelValue   interface{}   `bson:"ModelValue,omitempty" json:"model_value"`
	ChangeReason string        `bson:"ChangeReason,omitempty" json:"change_reason"`
}
