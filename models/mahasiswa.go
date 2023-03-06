package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mahasiswa struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nama      string             `json:"nama" bson:"nama" validate:"required"`
	Npm       string             `json:"npm" bson:"npm" validate:"required"`
	Ipk       string             `json:"ipk" bson:"ipk" validate:"required"`
	Kelas     string             `json:"kelas" bson:"kelas" validate:"required"`
	Kelulusan string             `json:"kelulusan" bson:"kelulusan" validate:"required"`
}
