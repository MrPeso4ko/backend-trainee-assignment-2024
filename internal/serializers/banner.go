package serializers

type Banner struct {
	Title string `db:"title" json:"title"`
	Text  string `db:"text" json:"text"`
	Url   string `db:"url" json:"url"`
}
