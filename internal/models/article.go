package models

// Article represents a knowledge base article
type Article struct {
	BaseModel

	Title   string `gorm:"size:500;not null;index" json:"title"`
	Slug    string `gorm:"size:500;uniqueIndex;not null" json:"slug"`
	Content string `gorm:"type:text;not null" json:"content"`
	Summary string `gorm:"type:text" json:"summary"`
	Status  string `gorm:"size:50;default:'draft'" json:"status"` // draft, published, archived

	CategoryID *uint `gorm:"index" json:"category_id,omitempty"`
	AuthorID   uint  `gorm:"index;not null" json:"author_id"`

	ViewCount int    `gorm:"default:0" json:"view_count"`
	Tags      string `gorm:"type:text" json:"tags"` // Comma-separated tags
}

// Category represents an article category
type Category struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	Slug        string `gorm:"size:255;uniqueIndex;not null" json:"slug"`
	Description string `gorm:"type:text" json:"description"`
	ParentID    *uint  `gorm:"index" json:"parent_id,omitempty"`
}

// TableName overrides
func (Article) TableName() string {
	return "articles"
}

func (Category) TableName() string {
	return "categories"
}
