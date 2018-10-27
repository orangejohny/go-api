package model

import (
	"time"
)

// AdItem struct describes ad that users supposed to add
type AdItem struct {
	ID            int64     `db:"id" json:"id" schema:"id,optional" valid:"-"`
	Title         string    `db:"title" json:"title" schema:"title,required" valid:"utfletternum,required"`
	Price         int       `db:"price" json:"price,omitempty" schema:"price,optional" valid:"-"`
	Country       string    `db:"country" json:"country,omitempty" schema:"country,optional" valid:"alpha,optional"`
	City          string    `db:"city" json:"city,omitempty" schema:"city,optional" valid:"alpha,optional"`
	SubwayStation string    `db:"subway_station" json:"subway_station,omitempty" schema:"subway_station,optional" valid:"alpha,optional"`
	ImagesFolder  string    `db:"images_folder" json:"images_folder,omitempty" schema:"-" valid:"-"`
	Owner         User      `db:"owner_ad" json:"owner_ad" schema:"-" valid:"-"`
	Description   string    `db:"description_ad" json:"description_ad" schema:"description_ad,required" valid:"utfletternum,required"`
	CreationTime  time.Time `db:"creation_time" json:"creation_time" schema:"-" valid:"-"`
}
