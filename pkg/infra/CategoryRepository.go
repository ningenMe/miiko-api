package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryRepository struct{}

type CategoryDto struct {
	CategoryId          string `db:"category_id"`
	CategoryDisplayName string `db:"category_display_name"`
	CategorySystemName  string `db:"category_system_name"`
	CategoryOrder       int32  `db:"category_order"`
}

func (CategoryRepository) GetList() []*CategoryDto {
	rows, err := ComproMysql.Queryx(`SELECT category_id, category_display_name, category_system_name, category_order FROM category ORDER BY category_order ASC`)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var list []*CategoryDto
	for rows.Next() {
		c := &CategoryDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}

	return list
}

func (CategoryRepository) Upsert(category *CategoryDto) {

	_, err := ComproMysql.NamedExec(`REPLACE INTO category (category_id, category_display_name, category_system_name, category_order) 
                                 VALUES (:category_id, :category_display_name, :category_system_name, :category_order)`, category)
	if err != nil {
		fmt.Println(err)
	}
}

func (CategoryRepository) Delete(categoryId string) {
	_, err := ComproMysql.NamedExec(`DELETE FROM category WHERE category_id = :categoryId`,
		map[string]interface{}{
			"categoryId": categoryId,
		})
	if err != nil {
		fmt.Println(err)
	}
}