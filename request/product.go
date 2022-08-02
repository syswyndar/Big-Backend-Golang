package request

type ProductRequest struct {
	Product_Name        string `json:"product_name"`
	Product_Description string `json:"product_description"`
	Product_Price       int    `json:"product_price"`
	Product_Stock       int    `json:"product_stock"`
	Product_Category_Id int    `json:"product_category_id"`
}
