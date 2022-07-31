package request

type CategoryRequest struct {
	Category_Name string `json:"category_name" binding:"required"`
}
