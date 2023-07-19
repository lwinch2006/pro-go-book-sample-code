package main

import (
	"chapter26/models"
	"chapter26/utils"
	"database/sql"
	_ "modernc.org/sqlite"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		utils.Printfln("Driver: %v", driver)
	}
}

func openDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "./data/chapter26.db")

	if err == nil {
		utils.Printfln("Database open success")
	}

	return
}

func Database1() {
	utils.Printfln("Database1()")
	listDrivers()
}

func Database2() {
	utils.Printfln("Database2()")

	db, err := openDatabase()

	if err == nil {
		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database3() {
	utils.Printfln("Database3()")

	sqlQuery := "SELECT * FROM Products"

	db, err := openDatabase()
	if err == nil {
		rows, err := db.Query(sqlQuery)
		if err == nil {
			var (
				id, categoryId int
				name           string
				price          float64
			)

			for rows.Next() {
				err = rows.Scan(&id, &categoryId, &name, &price)
				if err == nil {
					utils.Printfln("Product #%v %v of category %v with price %.2f", id, name, categoryId, price)
				} else {
					utils.Printfln("Row scan error: %v", err.Error())
				}
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database4() {
	utils.Printfln("Database4()")

	sqlQuery := "SELECT * FROM Products"

	db, err := openDatabase()
	if err == nil {
		rows, err := db.Query(sqlQuery)
		if err == nil {
			var (
				id, categoryId int
				name           string
				price          float64
			)

			for rows.Next() {
				err = rows.Scan(&id, &name, &categoryId, &price)
				if err == nil {
					utils.Printfln("Product #%v %v of category %v with price %.2f", id, name, categoryId, price)
				} else {
					utils.Printfln("Row scan error: %v", err.Error())
				}
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database5() {
	utils.Printfln("Database5()")

	sqlQuery := "SELECT * FROM Products"

	db, err := openDatabase()
	if err == nil {
		rows, err := db.Query(sqlQuery)
		if err == nil {
			product := &models.Product{}

			for rows.Next() {
				err = rows.Scan(&product.Id, &product.Name, &product.CategoryId, &product.Price)
				if err == nil {
					utils.Printfln("Product %v ", *product)
				} else {
					utils.Printfln("Row scan error: %v", err.Error())
				}
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database6() {
	utils.Printfln("Database6()")

	sqlQuery := "SELECT Products.*, Categories.* FROM Products, Categories WHERE Products.CategoryId == Categories.Id"

	db, err := openDatabase()
	if err == nil {
		rows, err := db.Query(sqlQuery)
		if err == nil {
			for rows.Next() {
				product := &models.Product{
					Category: &models.Category{},
				}

				err = rows.Scan(&product.Id, &product.Name, &product.CategoryId, &product.Price, &product.Category.Id, &product.Category.Name)
				if err == nil {
					utils.Printfln("Product %v with category reference %v", *product, *product.Category)
				} else {
					utils.Printfln("Row scan error: %v", err.Error())
				}
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database7() {
	utils.Printfln("Database7()")

	sqlQuery := `
		SELECT 
		    Products.*, 
		    Categories.* 
		FROM 
			Products, Categories 
		WHERE Products.CategoryId == Categories.Id
		AND Categories.Name = ?`

	categoryName := "Food"

	db, err := openDatabase()
	if err == nil {
		rows, err := db.Query(sqlQuery, categoryName)
		if err == nil {
			for rows.Next() {
				product := &models.Product{
					Category: &models.Category{},
				}

				err = rows.Scan(&product.Id, &product.Name, &product.CategoryId, &product.Price, &product.Category.Id, &product.Category.Name)
				if err == nil {
					utils.Printfln("Product %v with category reference %v", *product, *product.Category)
				} else {
					utils.Printfln("Row scan error: %v", err.Error())
				}
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database8() {
	utils.Printfln("Database8()")

	sqlQuery := `
		SELECT 
		    Products.*, 
		    Categories.* 
		FROM 
			Products, Categories 
		WHERE Products.CategoryId == Categories.Id
		AND Products.Id = ?`

	productId := 1

	db, err := openDatabase()
	if err == nil {
		row := db.QueryRow(sqlQuery, productId)
		if row.Err() == nil {
			product := &models.Product{
				Category: &models.Category{},
			}

			err = row.Scan(&product.Id, &product.Name, &product.CategoryId, &product.Price, &product.Category.Id, &product.Category.Name)
			if err == nil {
				utils.Printfln("Product %v with category reference %v", *product, *product.Category)
			} else {
				utils.Printfln("Row scan error: %v", err.Error())
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database9() {
	utils.Printfln("Database9()")

	sqlSelectQuery := `
		SELECT [Id]
		FROM [Products]
		ORDER BY [Id] DESC
	`
	sqlInsertQuery := `
		INSERT INTO [Products] ([Id], [Name], [CategoryId], [Price])
		VALUES (?, ?, ?, ?)
		`

	db, err := openDatabase()
	if err == nil {
		row := db.QueryRow(sqlSelectQuery)
		if row.Err() == nil {
			latestProductId := 0
			err = row.Scan(&latestProductId)
			if err == nil {
				latestProductId++
				newProduct := &models.Product{
					Id:         latestProductId,
					Name:       "Skirt",
					CategoryId: 2,
					Price:      56.67,
				}

				result, err := db.Exec(sqlInsertQuery, &newProduct.Id, &newProduct.Name, &newProduct.CategoryId, &newProduct.Price)
				if err == nil {
					rowsAffected, _ := result.RowsAffected()
					if rowsAffected == 1 {
						lastInsertedId, _ := result.LastInsertId()
						utils.Printfln("Product successfully inserted with ID=%v", lastInsertedId)
					} else {
						utils.Printfln("Product not inserted")
					}
				} else {
					utils.Printfln("Database query error: %v", err.Error())
				}
			} else {
				utils.Printfln("Row scan error: %v", err.Error())
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database10() {
	utils.Printfln("Database10()")

	var (
		newCategoryStatement           *sql.Stmt
		changeProductCategoryStatement *sql.Stmt
	)

	newCategorySqlQuery := `
			INSERT INTO [Categories] ([Name])
			VALUES (?)
		`
	changeProductCategorySqlQuery := `
			UPDATE [Products]
			SET [CategoryId] = ?
			WHERE [Id] = ?
		`

	db, err := openDatabase()
	if err == nil {
		newCategoryStatement, _ = db.Prepare(newCategorySqlQuery)
		changeProductCategoryStatement, _ = db.Prepare(changeProductCategorySqlQuery)

		newCategory := &models.Category{
			Name: "NewFood",
		}

		result, err := newCategoryStatement.Exec(&newCategory.Name)
		if err == nil {
			rowsAffected, _ := result.RowsAffected()
			if rowsAffected == 1 {
				lastInsertedId, _ := result.LastInsertId()
				utils.Printfln("Category successfully inserted with ID=%v", lastInsertedId)

				_, _ = changeProductCategoryStatement.Exec(&lastInsertedId, 1)

			} else {
				utils.Printfln("Category not inserted")
			}
		} else {
			utils.Printfln("Database query error: %v", err.Error())
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}

func Database11() {
	utils.Printfln("Database11()")

	var (
		lastInsertedId                 int64
		rowsAffected                   int64
		sqlExecFailed                  bool
		newCategoryStatement           *sql.Stmt
		changeProductCategoryStatement *sql.Stmt
	)

	newCategorySqlQuery := `
			INSERT INTO [Categories] ([Name])
			SELECT ?
			WHERE NOT EXISTS (SELECT 1 FROM [Categories] WHERE [Name] == ?)
		`
	changeProductCategorySqlQuery := `
			UPDATE [Products]
			SET [CategoryId] = ?
			WHERE [Id] = ?
		`

	db, err := openDatabase()
	if err == nil {
		transaction, _ := db.Begin()
		newCategoryStatement, _ = transaction.Prepare(newCategorySqlQuery)
		changeProductCategoryStatement, _ = transaction.Prepare(changeProductCategorySqlQuery)

		newCategory := &models.Category{
			Name: "NewFood2",
		}

		result, err := newCategoryStatement.Exec(&newCategory.Name, &newCategory.Name)
		if err == nil {
			rowsAffected, err = result.RowsAffected()
			if err == nil {
				if rowsAffected == 1 {
					lastInsertedId, err = result.LastInsertId()
					if err == nil {
						utils.Printfln("Category successfully inserted with ID=%v", lastInsertedId)
						result, err = changeProductCategoryStatement.Exec(&lastInsertedId, 1)
						if err == nil {
							if rowsAffected, err = result.RowsAffected(); err == nil && rowsAffected == 0 {
								sqlExecFailed = true
							}
						}
					}
				} else {
					sqlExecFailed = true
				}
			}
		}

		if err != nil || sqlExecFailed {
			errorMessage := "empty"
			if err != nil {
				errorMessage = err.Error()
			}

			utils.Printfln("Aborting transaction")
			utils.Printfln("Error: %v", errorMessage)
			utils.Printfln("SQL Exec failed: %v", sqlExecFailed)
			transaction.Rollback()
		} else {
			transaction.Commit()
		}

		db.Close()
	} else {
		utils.Printfln("Database open error: %v", err.Error())
	}
}
