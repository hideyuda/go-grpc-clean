package entity

type Article struct {
	Id        uint   `db:"id" json:"id"`
	Uuid      string `db:"uuid" json:"uuid"`
	Title     string `db:"title" json:"title"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`

	// Paragraphs []Paragraph `db:"paragraphs" json:"paragraphs"`
	// Keywords   []Keyword   `db:"keywords" json:"keywords"`
}
