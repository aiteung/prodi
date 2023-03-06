package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Heregistrasi struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nisn        string             `json:"nisn" bson:"nisn" validate:"required"`
	Nama        string             `json:"nama" bson:"nama" validate:"required"`
	AsalSekolah string             `json:"asal_sekolah" bson:"asal_sekolah" validate:"required"`
	Alamat      string             `json:"alamat" bson:"alamat" validate:"required"`
	Provinsi    string             `json:"provinsi" bson:"provinsi" validate:"required"`
}

type NilaiRaport struct {
	Matematika      string    `json:"matematika" bson:"matematika" validate:"required"`
	BahasaIndonesia string    `json:"b_indo" bson:"b_indo" validate:"required"`
	BahasaInggris   string    `json:"b_ing" bson:"b_ing" validate:"required"`
	Peminatan       Peminatan `bson:"peminatan,omitempty"`
}

type Peminatan struct {
	Ipa Ipa `bson:"ipa,omitempty"`
	Ips Ips `bson:"ips,omitempty"`
}

type Ipa struct {
	Biologi string `json:"biologi" bson:"biologi" validate:"required"`
	Fisika  string `json:"fisika" bson:"fisika" validate:"required"`
	Kimia   string `json:"kimia" bson:"kimia" validate:"required"`
}

type Ips struct {
	Ekonomi   string `json:"ekonomi" bson:"ekonomi" validate:"required"`
	Geografi  string `json:"geografi" bson:"geografi" validate:"required"`
	Sejarah   string `json:"sejarah" bson:"sejarah" validate:"required"`
	Sosiologi string `json:"sosiologi" bson:"sosiologi" validate:"required"`
}
