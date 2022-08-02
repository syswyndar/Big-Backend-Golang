package request

type ProductRequest struct {
	Product_Name        string `json:"productName" binding:"required"`
	Product_Description string `json:"productDescription" binding:"required"`
	Product_Price       int64  `json:"productPrice" binding:"required,number"`
	Product_Stock       int    `json:"productStock" binding:"required,number"`
	Product_Category_Id int    `json:"productCategoryId" binding:"required,number"`
}
