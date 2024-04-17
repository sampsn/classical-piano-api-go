package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, _ = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

func main() {
	app := fiber.New()

	db.AutoMigrate(&Composer{}, &Piece{})

	// jsonFile, err := os.Open("composers.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Successfully opened file.json")
	// defer jsonFile.Close()
	//
	// byteValue, _ := io.ReadAll(jsonFile)
	//
	// var composers []Composer
	//
	// err = json.Unmarshal(byteValue, &composers)
	// if err != nil {
	// 	fmt.Println("Decoding error: ", err)
	// }
	//
	// for _, composer := range composers {
	// 	db.Create(&composer)
	// }
	//
	// jsonFile, err = os.Open("pieces.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Successfully opened file.json")
	// defer jsonFile.Close()
	//
	// byteValue, _ = io.ReadAll(jsonFile)
	//
	// var pieces []Piece
	//
	// err = json.Unmarshal(byteValue, &pieces)
	// if err != nil {
	// 	fmt.Println("Decoding error: ", err)
	// }
	//
	// for _, piece := range pieces {
	// 	db.Create(&piece)
	// }

	app.Get("/composers", getComposers)
	app.Get("/pieces", getPieces)
	app.Post("/composers", createComposer)
	app.Post("/pieces", createPiece)
	app.Put("/composers/:composer_id", updateComposer)
	app.Put("/pieces/:piece_name", updatePiece)
	app.Delete("/composers/:composer_id", deleteComposer)
	app.Delete("/pieces/:piece_name", deletePiece)

	app.Listen(":8001")
}

func getComposers(c *fiber.Ctx) error {
	var composers []Composer
	db.Preload("Pieces").Find(&composers)
	return c.JSON(composers)
}

func getPieces(c *fiber.Ctx) error {
	id := c.Query("composer_id")
	if id == "" {
		var pieces []Piece
		db.Find(&pieces)
		return c.JSON(pieces)
	} else {
		var pieces []Piece
		db.Where("composer_id = ?", id).Find(&pieces)
		return c.JSON(pieces)
	}
}

func createComposer(c *fiber.Ctx) error {
	new_composer := new(Composer)
	c.BodyParser(&new_composer)
	db.Create(&new_composer)
	return nil
}

func createPiece(c *fiber.Ctx) error {
	new_piece := new(Piece)
	c.BodyParser(&new_piece)
	db.Create(&new_piece)
	return nil
}

func updateComposer(c *fiber.Ctx) error {
	id := c.Params("composer_id")
	updated_composer := new(Composer)
	c.BodyParser(&updated_composer)
	composer := new(Composer)
	db.Where("id = ?", id).First(&composer)
	composer.Name = updated_composer.Name
	composer.Home_Country = updated_composer.Home_Country
	db.Save(&composer)
	return nil
}

func updatePiece(c *fiber.Ctx) error {
	pieceName := c.Params("piece_name")
	updated_piece := new(Piece)
	c.BodyParser(&updated_piece)
	piece := new(Piece)
	db.Where("name = ?", pieceName).First(&piece)
	piece.Name = updated_piece.Name
	piece.Alt_Name = updated_piece.Alt_Name
	piece.Difficulty = updated_piece.Difficulty
	piece.Composer_ID = updated_piece.Composer_ID
	db.Save(&piece)
	return nil
}

func deleteComposer(c *fiber.Ctx) error {
	id := c.Params("composer_id")
	db.Delete(&Composer{}, id)
	return nil
}

func deletePiece(c *fiber.Ctx) error {
	pieceName := c.Params("piece_name")
	db.Delete(&Piece{}, "name = ?", pieceName)
	return nil
}
