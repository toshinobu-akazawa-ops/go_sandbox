package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "fmt"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

func main() {
  db, err := gorm.Open("postgres", "host=db port=5432 user=root dbname=postgres password=root sslmode=disable")
  if err != nil {
    fmt.Println(err)
    panic("データベースへの接続に失敗しました")
  }
  defer db.Close()

  // スキーマのマイグレーション
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "L1212", Price: 1000})

  // Read
  var product Product
  db.First(&product, 1) // idが1の製品を探します
  fmt.Println(product.Price)
  db.First(&product, "code = ?", "L1212") // codeがL1212の製品を探します

  // Update - 製品価格を2,000に更新します
  db.Model(&product).Update("Price", 2000)

  // Delete - 製品を削除します
  //db.Delete(&product)
}