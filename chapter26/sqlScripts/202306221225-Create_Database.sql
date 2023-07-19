DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS Products;

CREATE TABLE IF NOT EXISTS Categories
(
    Id INTEGER NOT NULL PRIMARY KEY,
    Name TEXT
);

CREATE TABLE IF NOT EXISTS Products
(
    Id INTEGER NOT NULL PRIMARY KEY,
    Name TEXT,
    CategoryId INTEGER,
    Price decimal(8,2),
    CONSTRAINT CatRef FOREIGN KEY(CategoryId) REFERENCES Categories(Id)
);

INSERT INTO Categories (Id, Name)
VALUES (1, "Food"),
       (2, "Clothes");

INSERT INTO Products (Id, Name, Category, Price)
VALUES (1, "Milk", 1, 23.34),
       (2, "Bread", 1, 45.56),
       (3, "Beer", 1, 67.78),
       (4, "Jacket", 2, 123.34);

