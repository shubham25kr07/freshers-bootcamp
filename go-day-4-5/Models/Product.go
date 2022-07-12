package Models

import "go-day-4-5/Config"

func GetAllProducts(products *[]Product) (err error) {
	if err = Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(product *Product, id string) (err error) {

	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product) (err error) {

	if err = Config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}
