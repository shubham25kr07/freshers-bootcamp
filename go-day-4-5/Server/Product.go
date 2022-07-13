package Server

import (
	"go-day-4-5/Config"
	"go-day-4-5/Models"
)

func GetAllProducts(products *[]Models.Product) (err error) {
	if err = Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(product *Models.Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(product *Models.Product, id string) (err error) {

	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Models.Product) (err error) {

	if err = Config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}
