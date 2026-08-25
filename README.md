# RecipeApp

A full-stack recipe management application.

## Architecture

- **Backend**: Go + Fiber + PostgreSQL (`/backend`)
- **Frontend**: React + Vite + Tailwind CSS (`/frontend`)

## Quick Start

### Backend

```bash
cd backend
cp .env.example .env
go run main.go
# API läuft auf http://localhost:8080
```

### Frontend

```bash
cd frontend
npm install
npm run dev
# App läuft auf http://localhost:5173
```

### Production

```bash
docker-compose up
```

## Project Structure

```
recipeApp/
├── backend/    - Go REST API
├── frontend/   - React Web App
└── docker-compose.yml
```