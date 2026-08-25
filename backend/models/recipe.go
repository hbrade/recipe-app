package models

import "time"

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

type Recipe struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
	ImageURL    string       `json:"image_url"`
	Tags        []string     `json:"tags"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type CreateRecipeRequest struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
	ImageURL    string       `json:"image_url"`
	Tags        []string     `json:"tags"`
}
