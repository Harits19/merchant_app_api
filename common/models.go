package common

type ModifiedModel struct {
	CreatedAt string `json:"created_at" bson:"created_at"`
	CreatedBy int    `json:"created_by" bson:"created_by"`
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	UpdatedBy int    `json:"updated_by" bson:"updated_by"`
}
