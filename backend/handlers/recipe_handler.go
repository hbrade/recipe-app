package handlers

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/dada/recipe-api/config"
	"github.com/dada/recipe-api/models"
	"github.com/gofiber/fiber/v2"
)

// GET /api/recipes
func GetRecipes(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(),
		"SELECT id, title, description, ingredients, image_url, tags, created_at, updated_at FROM recipes ORDER BY created_at DESC")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		var ingredientsJSON []byte
		var tagsArray []string

		err := rows.Scan(
			&recipe.ID,
			&recipe.Title,
			&recipe.Description,
			&ingredientsJSON,
			&recipe.ImageURL,
			&tagsArray,
			&recipe.CreatedAt,
			&recipe.UpdatedAt,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		json.Unmarshal(ingredientsJSON, &recipe.Ingredients)
		recipe.Tags = tagsArray

		recipes = append(recipes, recipe)
	}

	return c.JSON(recipes)
}

// GET /api/recipes/:id
func GetRecipe(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var recipe models.Recipe
	var ingredientsJSON []byte
	var tagsArray []string

	err = config.DB.QueryRow(context.Background(),
		"SELECT id, title, description, ingredients, image_url, tags, created_at, updated_at FROM recipes WHERE id = $1",
		id,
	).Scan(
		&recipe.ID,
		&recipe.Title,
		&recipe.Description,
		&ingredientsJSON,
		&recipe.ImageURL,
		&tagsArray,
		&recipe.CreatedAt,
		&recipe.UpdatedAt,
	)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Recipe not found"})
	}

	json.Unmarshal(ingredientsJSON, &recipe.Ingredients)
	recipe.Tags = tagsArray

	return c.JSON(recipe)
}

// POST /api/recipes
func CreateRecipe(c *fiber.Ctx) error {
	req := new(models.CreateRecipeRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validation
	if req.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Title is required"})
	}

	ingredientsJSON, _ := json.Marshal(req.Ingredients)

	var id int
	err := config.DB.QueryRow(context.Background(),
		`INSERT INTO recipes (title, description, ingredients, image_url, tags)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
		req.Title, req.Description, ingredientsJSON, req.ImageURL, req.Tags,
	).Scan(&id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"id": id, "message": "Recipe created"})
}

// PUT /api/recipes/:id
func UpdateRecipe(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	req := new(models.CreateRecipeRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	ingredientsJSON, _ := json.Marshal(req.Ingredients)

	result, err := config.DB.Exec(context.Background(),
		`UPDATE recipes 
		 SET title=$1, description=$2, ingredients=$3, image_url=$4, tags=$5, updated_at=NOW()
		 WHERE id=$6`,
		req.Title, req.Description, ingredientsJSON, req.ImageURL, req.Tags, id,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if result.RowsAffected() == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Recipe not found"})
	}

	return c.JSON(fiber.Map{"message": "Recipe updated"})
}

// DELETE /api/recipes/:id
func DeleteRecipe(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	result, err := config.DB.Exec(context.Background(), "DELETE FROM recipes WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if result.RowsAffected() == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Recipe not found"})
	}

	return c.JSON(fiber.Map{"message": "Recipe deleted"})
}
