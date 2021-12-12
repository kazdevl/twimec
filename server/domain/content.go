package domain

type ContentInfo struct {
	ID         string `json:"id"`
	AuthorName string `json:"author_name"`
	Header     string `json:"header"`
	Cover      string `json:"cover"`
	Title      string `json:"title"`
}
type Chapter struct {
	ContentID string `json:"content_id"`
	Index     int    `json:"name"`
	Icon      string `json:"icon"`
	Pages     string `json:"pages"` // great way: use list in one column
}
