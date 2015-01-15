package dataFormat

import (
	"time"
)

type SideStats struct {
	Wins uint32
	Losses uint32
}

type DetailedNumberOf struct {
	Wins uint32
	Losses uint32
	TimePlayed uint32 // Seconds
	Kills uint32
	Assists uint32
	Deaths uint32
	DoubleKills uint32
	TripleKills uint32
	QuadraKills uint32
	PentaKills uint32
	GoldEarned uint32
	MinionsKilled uint32
	MonstersKilled uint32
	WardsPlaced uint32
	WardsKilled uint32
	Blue SideStats
	Red SideStats
}

type BasicNumberOf struct {
	Wins uint32
	Losses uint32
	TimePlayed uint32 // Seconds
	Kills uint32
	Assists uint32
	Deaths uint32
	// DoubleKills uint32
	// TripleKills uint32
	// QuadraKills uint32
	// PentaKills uint32
	MinionsKilled uint32
	MonstersKilled uint32
	WardsPlaced uint32
}

type Stats struct {
	All DetailedNumberOf
	GameTypeStats map[string]DetailedNumberOf
	Champions map[string]BasicNumberOf
}

type Player struct {
	Region string // Region code. Ex: oce
	Tier string
	Id string
	RecordStart string // Date of first ever game recorded
	NextUpdate time.Time
	NextLongUpdate time.Time
	ProcessedGames []string
	All Stats
	MonthlyStats map[string]Stats
}

type BasicPlayer struct {
	Region string
	Id string
	Tier string
	RecordStart time.Time
	NextUpdate time.Time
	NextLongUpdate time.Time
}

type BrowserPlayer struct {
	Name string
	Region string
}

type Champion struct {
	Name string
	Title string
	SquareURL string
	SplashURL string
	PortraitURL string
}

type Game struct {
	DidWin bool
	IsOnBlue bool
	IsNormals bool // AKA Not Custom
	ChampionId string
	Duration uint32
	Id string
	Type string
	Kills uint32
	Assists uint32
	Deaths uint32
	DoubleKills uint32
	TripleKills uint32
	QuadraKills uint32
	PentaKills uint32
	GoldEarned uint32
	MinionsKilled uint32
	MonstersKilled uint32
	WardsPlaced uint32
	WardsKilled uint32
	YearMonth string
	Date time.Time
}
