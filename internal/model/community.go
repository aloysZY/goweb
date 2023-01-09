package model

import "time"

type Community struct {
	Id   int    `db:"community_id"`
	Name string `db:"community_name"`
}

type CommunityDetail struct {
	Id         int       `db:"community_id"`
	Name       string    `db:"community_name"`
	Intr       string    `db:"introduction"`
	CreateTime time.Time `db:"create_time"`
}
