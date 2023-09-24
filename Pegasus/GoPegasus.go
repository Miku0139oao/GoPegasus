package GoPegasus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RawDataMAp map[string]PegasusItems

type RawApiData struct {
	Items    RawDataMAp `json:"items"`
	MinPrice int        `json:"minPrice"`
	MaxPrice int        `json:"maxPrice"`
	Tags     []string   `json:"tags"`
}

type PegasusItems struct {
	BrandId            string   `json:"brandId"`
	Categories         []string `json:"categories"`
	DeliveryBySupplier bool     `json:"deliveryBySupplier"`
	Ean                string   `json:"ean"`
	Id                 int      `json:"id"`
	Image              string   `json:"image"`
	MenuOrder          int      `json:"menuOrder"`
	Name               string   `json:"name"`
	NeedSets           int      `json:"needSets"`
	OrderingTag        int      `json:"orderingTag"`
	Price              int      `json:"price"`
	Sku                string   `json:"sku"`
	Status             string   `json:"status"`
	TagVisibility      int      `json:"tagVisibility"`
	Tags               []string `json:"tags"`
	Visibility         string   `json:"visibility"`
	Tag                int      `json:"tag"`
}

type ItemList []PegasusItems

var (
	SoftWare            = []int{13, 16}
	PowerSupplier       = []int{88}
	MotherBoard         = []int{5}
	CPU                 = []int{2}
	Case                = []int{12}
	Storage             = []int{7, 9}
	Cooler              = []int{68, 86, 87, 33}
	RAM                 = []int{8}
	DisplayCard         = []int{10}
	DisplayCard_Add_Ons = []int{117}
)

func GetItems(Categories []int) (Items ItemList) {
	url := fmt.Sprintf("https://api.pegasus.hk/api/product/query?searchKey=&limit=10&categories=")
	for i, category := range Categories {
		if i != len(Categories)-1 {
			url += fmt.Sprintf("%v,", category)
		} else {
			url += fmt.Sprintf("%v", category)
		}
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return Items
	}
	req.Header.Set("Authorization", "Bearer")
	req.Header.Set("User-Agent", "Pegasus Status Bot")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var Map RawApiData
	if err = json.Unmarshal(body, &Map); err != nil {
		fmt.Println(err)
		return
	}
	for _, product := range Map.Items {
		Items = append(Items, PegasusItems{
			BrandId:            product.BrandId,
			Categories:         product.Categories,
			DeliveryBySupplier: product.DeliveryBySupplier,
			Ean:                product.Ean,
			Id:                 product.Id,
			Image:              product.Image,
			MenuOrder:          product.MenuOrder,
			Name:               product.Name,
			NeedSets:           product.NeedSets,
			OrderingTag:        product.OrderingTag,
			Price:              product.Price,
			Sku:                product.Sku,
			Status:             product.Status,
			TagVisibility:      product.TagVisibility,
			Tags:               product.Tags,
			Visibility:         product.Visibility,
			Tag:                product.Tag,
		})
	}
	return Items
}
