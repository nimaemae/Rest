# Coffee Shop Platform - Multi-Tenant Backend

A comprehensive multi-tenant coffee shop management platform built with Go, Echo, and PostgreSQL.

## 🏗️ Architecture

### Multi-Tenant Structure
- **Main Admin**: Platform owner who manages tenants, coffee shops, and categories
- **Tenants**: Organizations that can have multiple coffee shops
- **Coffee Shops**: Individual shops under tenants with their own admins and menus
- **Categories**: Centralized category management shared across all shops

### Key Features
- ✅ **Centralized Category Management**: Main admin controls all categories
- ✅ **Multi-Tenant Support**: Each tenant can have multiple coffee shops
- ✅ **Subdomain Routing**: Each tenant gets their own subdomain
- ✅ **JWT Authentication**: Secure authentication for both admin types
- ✅ **Database Migrations**: Automated schema management
- ✅ **Sample Data Seeding**: Pre-populated with realistic data
- ✅ **RESTful APIs**: Complete CRUD operations for all entities

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Make (optional, for convenience commands)

### Installation

1. **Clone and setup**:
   ```bash
   git clone <repository>
   cd backend
   cp .env.example .env
   # Edit .env with your database credentials
   ```

2. **Install dependencies**:
   ```bash
   make deps
   # or
   go mod tidy
   ```

3. **Setup database**:
   ```bash
   make setup
   # or
   go run cmd/main.go -migrate
   go run cmd/main.go -seed
   ```

4. **Start the server**:
   ```bash
   make dev
   # or
   go run cmd/main.go
   ```

## 📊 Database Schema

### Core Tables
- `main_admins` - Platform administrators
- `tenants` - Multi-tenant organizations
- `coffee_shops` - Individual coffee shops
- `shop_admins` - Coffee shop administrators
- `categories` - **Centralized category management**
- `menu_items` - Menu items linked to categories

### Category Management
Categories are managed centrally by the main admin and shared across all coffee shops:

```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    emoji VARCHAR(10),
    color VARCHAR(50),
    order_index INTEGER DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);
```

## 🔧 API Endpoints

### Public Endpoints
- `GET /api/public/categories` - Get active categories
- `GET /api/public/menu` - Get public menu (requires tenant resolution)
- `GET /api/public/shop` - Get shop settings (requires tenant resolution)

### Main Admin Endpoints
- `POST /api/auth/main-admin/login` - Main admin login
- `GET /api/admin/categories` - Get all categories
- `POST /api/admin/categories` - Create category
- `PUT /api/admin/categories/:id` - Update category
- `DELETE /api/admin/categories/:id` - Delete category
- `GET /api/admin/tenants` - Manage tenants
- `GET /api/admin/tenants/:id/shops` - Manage coffee shops

### Shop Admin Endpoints
- `POST /api/auth/shop-admin/login` - Shop admin login
- `GET /api/admin/categories` - Get active categories (read-only)
- `GET /api/admin/menu` - Manage menu items
- `POST /api/admin/menu` - Create menu item
- `PUT /api/admin/menu/:id` - Update menu item
- `DELETE /api/admin/menu/:id` - Delete menu item

## 🎯 Category Management

### Creating Categories
Only main admins can create and manage categories:

```bash
curl -X POST http://localhost:8080/api/admin/categories \
  -H "Authorization: Bearer <main_admin_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "coffee",
    "display_name": "قهوه",
    "emoji": "☕",
    "color": "from-amber-400 to-orange-500",
    "order_index": 1
  }'
```

### Using Categories in Menu Items
Shop admins can only select from existing categories when creating menu items:

```bash
curl -X POST http://localhost:8080/api/admin/menu \
  -H "Authorization: Bearer <shop_admin_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "اسپرسو",
    "category_id": 1,
    "price": 45000,
    "price_premium": 55000,
    "has_dual_pricing": true,
    "is_available": true
  }'
```

## 🏪 Multi-Tenant Setup

### Tenant Configuration
1. **Create Tenant**:
   ```bash
   curl -X POST http://localhost:8080/api/admin/tenants \
     -H "Authorization: Bearer <main_admin_token>" \
     -H "Content-Type: application/json" \
     -d '{
       "subdomain": "mycoffee",
       "name": "My Coffee Company"
     }'
   ```

2. **Create Coffee Shop**:
   ```bash
   curl -X POST http://localhost:8080/api/admin/tenants/1/shops \
     -H "Authorization: Bearer <main_admin_token>" \
     -H "Content-Type: application/json" \
     -d '{
       "name": "Downtown Location",
       "location": "123 Main St",
       "phone": "+1-555-0123"
     }'
   ```

3. **Create Shop Admin**:
   ```bash
   curl -X POST http://localhost:8080/api/admin/shops/1/admins \
     -H "Authorization: Bearer <main_admin_token>" \
     -H "Content-Type: application/json" \
     -d '{
       "username": "shopadmin",
       "password": "securepassword"
     }'
   ```

### Subdomain Access
- Main platform: `http://localhost:8080`
- Tenant subdomain: `http://mycoffee.localhost:8080`
- Public menu: `http://mycoffee.localhost:8080/api/public/menu`

## 🗄️ Database Management

### Migration Commands
```bash
# Run migrations only
make migrate
# or
go run cmd/main.go -migrate

# Seed with sample data
make seed
# or
go run cmd/main.go -seed

# Full setup (migrate + seed)
make setup
```

### Sample Data
The seeding process creates:
- 1 main admin (username: `admin`, password: `admin123`)
- 8 predefined categories (Coffee, Shake, Cold Bar, etc.)
- 1 sample tenant (`demo` subdomain)
- 1 sample coffee shop
- 1 shop admin (username: `shopadmin`, password: `shop123`)
- 51 sample menu items with proper category assignments

## 🔐 Authentication

### Main Admin Login
```bash
curl -X POST http://localhost:8080/api/auth/main-admin/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

### Shop Admin Login
```bash
curl -X POST http://localhost:8080/api/auth/shop-admin/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "shopadmin",
    "password": "shop123"
  }'
```

## 🌐 Environment Variables

Create a `.env` file with:

```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=coffee_shop_platform

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRE_HOURS=24
```

## 📁 Project Structure

```
backend/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go          # Configuration management
│   ├── database/
│   │   ├── database.go        # Database connection
│   │   └── migrate.go         # Migration functions
│   ├── handlers/
│   │   ├── auth.go            # Authentication handlers
│   │   ├── category.go        # Category management handlers
│   │   ├── coffee_shop.go     # Coffee shop handlers
│   │   ├── menu.go            # Menu item handlers
│   │   └── tenant.go          # Tenant handlers
│   ├── middleware/
│   │   └── auth.go            # Authentication middleware
│   ├── models/
│   │   ├── category.go        # Category model
│   │   └── models.go          # All other models
│   ├── routes/
│   │   └── routes.go          # Route definitions
│   └── utils/
│       ├── jwt.go             # JWT utilities
│       └── password.go        # Password hashing
├── scripts/
│   └── seed.go                # Database seeding
├── .env.example               # Environment template
├── go.mod                     # Go modules
├── Makefile                   # Build commands
└── README.md                  # This file
```

## 🚀 Deployment

### Production Setup
1. Set up PostgreSQL database
2. Configure environment variables
3. Run migrations: `./bin/server -migrate`
4. Seed initial data: `./bin/server -seed`
5. Start server: `./bin/server`

### Docker Support
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o bin/server cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin/server .
CMD ["./server"]
```

## 🔄 Development Workflow

1. **Make changes** to models, handlers, or routes
2. **Run migrations** to update database schema
3. **Test APIs** using curl or Postman
4. **Seed data** if needed for testing
5. **Deploy** to production

## 📝 API Examples

### Get All Categories
```bash
curl http://localhost:8080/api/public/categories
```

### Create Menu Item with Category
```bash
curl -X POST http://localhost:8080/api/admin/menu \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cappuccino",
    "category_id": 1,
    "price": 40000,
    "price_premium": 50000,
    "has_dual_pricing": true,
    "is_available": true
  }'
```

## 🎉 Features Summary

- ✅ **Centralized Category Management**: All shops use the same category list
- ✅ **Multi-Tenant Architecture**: Support for multiple organizations
- ✅ **Subdomain Routing**: Each tenant gets their own subdomain
- ✅ **JWT Authentication**: Secure token-based authentication
- ✅ **Database Migrations**: Automated schema management
- ✅ **Sample Data**: Pre-populated with realistic coffee shop data
- ✅ **RESTful APIs**: Complete CRUD operations
- ✅ **Environment Configuration**: All settings via environment variables
- ✅ **Command Line Tools**: Migration and seeding flags
- ✅ **Comprehensive Documentation**: Complete API documentation

The platform is now ready for production use with centralized category management! 🚀
