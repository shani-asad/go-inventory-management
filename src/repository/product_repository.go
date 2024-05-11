package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-management/model/dto"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepositoryInterface {
	return &ProductRepository{db}
}

func (r *ProductRepository) SearchSku(ctx context.Context, params dto.SearchSkuParams) (response []dto.SearchSkuResponse, err error) {
	
	query := constructQuery(params)
	rows, err := r.db.QueryContext(ctx, query)

	for rows.Next() {
		var sku dto.SearchSkuResponse
		err := rows.Scan(&sku)
		fmt.Println(err)
		if err != nil {
				return nil, err
		}
		response = append(response, sku)
}
	return response, err
}

func constructQuery(params dto.SearchSkuParams) string {
	query := "SELECT * FROM products WHERE 1=1"
	if params.Name != "" {
		query += fmt.Sprintf(" AND LOWER(name) LIKE LOWER('%%%s%%')", params.Name)
	}
	if params.Category != "" {
		query += fmt.Sprintf(" AND category = '%s'", params.Category)
	}
	if params.Sku != "" {
		query += fmt.Sprintf(" AND sku = '%s'", params.Sku)
	}
	if(params.IsInstockValid){
		if params.InStock {
			query += " AND stock > 0"
		} else {
			query += " AND stock < 1"
		}
	}
	if params.Price == "asc" {
		query += " ORDER BY price ASC"
	} else if params.Price == "desc" {
		query += " ORDER BY price DESC"
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, params.Offset)
	return query
}
