package repository

import (
	"context"
	"database/sql"
	"fmt"
	"inventory-management/model/dto"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) TransactionRepositoryInterface {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) GetTransactions(ctx context.Context, params dto.GetTransactionRequest) (response []dto.TransactionData, err error) {

	query := `SELECT
    id,
    customer_id,
    paid,
    change,
    created_at
    FROM transactions t WHERE 1 = 1`

	if params.CustomerId != "" {
		query += fmt.Sprintf(" AND t.customer_id = '%s'", params.CustomerId)
	}

	if params.CreatedAt == "asc" {
		query += " ORDER BY created_at ASC"
	} else if params.CreatedAt == "desc" {
		query += " ORDER BY created_at DESC"
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, params.Offset)

	
	rows, err := r.db.QueryContext(context.Background(), query)

    if err != nil { 
		fmt.Println("Error in tranasction_repo > GetTransactions > in QueryContext: ", err)
		return nil, err
	}
    if(rows == nil){
        fmt.Println("No transacion result")
        return []dto.TransactionData{}, err
    }
	
    defer rows.Close()

	var transactions []dto.TransactionData

	for rows.Next() {
        var transaction dto.TransactionData
        // Scan the row data into TransactionData fields
        err := rows.Scan(
            &transaction.TransactionID,
            &transaction.CustomerID,
            &transaction.Paid,
            &transaction.Change,
            &transaction.CreatedAt,
        )
        if err != nil {
            fmt.Println("Error in San transaction rows", err)
            panic(err)
        }

        // Query product details and populate ProductDetails field
        productDetails, err := getProductDetails(r.db, transaction.TransactionID)
        if err != nil {
            fmt.Println("Error in getProductDetails", err)
            panic(err)
        }
        transaction.ProductDetails = productDetails

        // Append transaction to the result slice
        transactions = append(transactions, transaction)
    }
	
    // Print the result
    fmt.Println(transactions)
	return transactions, err
}

func getProductDetails(db *sql.DB, transactionID string) ([]dto.ProductDetail, error) {
    rows, err := db.QueryContext(context.Background(), "SELECT product_id, quantity FROM transaction_products WHERE transaction_id = $1", transactionID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var productDetails []dto.ProductDetail
    for rows.Next() {
        var productDetail dto.ProductDetail
        if err := rows.Scan(&productDetail.ProductID, &productDetail.Quantity); err != nil {
            return nil, err
        }
        productDetails = append(productDetails, productDetail)
    }
    return productDetails, nil
}
