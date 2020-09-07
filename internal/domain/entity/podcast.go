package entity

// Podcast Podcasts entity
type Podcast struct {
	RssID       string `json:"rss_id"`
	DurationSec int64  `json:"duration_sec"`
	Duration    string `json:"duration"`
	Episode     int32  `json:"episode"`
	EpisodeType string `json:"episode_type"`
	Transcript  string `json:"transcript"`
}
