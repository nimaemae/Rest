# Coffee Shop Platform - Complete Setup Guide

This guide will help you set up the complete Coffee Shop Platform with both frontend and backend.

## 🏗️ Architecture Overview

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   React Frontend │    │   Go Backend    │    │   PostgreSQL    │
│   (Port 5173)   │◄──►│   (Port 8080)   │◄──►│   (Port 5432)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 📋 Prerequisites

### Required Software
- **Node.js 16+** - For React frontend
- **Go 1.21+** - For backend API
- **PostgreSQL 12+** - For database
- **Git** - For version control

### Optional Tools
- **Docker** - For containerized deployment
- **Postman** - For API testing
- **VS Code** - Recommended IDE

## 🚀 Quick Start (Automated)

### 1. Clone and Setup
```bash
git clone <repository-url>
cd coffee-shop-platform
```

### 2. Start Everything
```bash
./start-dev.sh
```

This script will:
- ✅ Build the Go backend
- ✅ Run database migrations
- ✅ Seed the database with sample data
- ✅ Start the backend server (port 8080)
- ✅ Start the frontend dev server (port 5173)

### 3. Access the Application
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080
- **Health Check**: http://localhost:8080/health

## 🔧 Manual Setup

### Backend Setup

1. **Navigate to backend directory**:
   ```bash
   cd backend
   ```

2. **Install Go dependencies**:
   ```bash
   go mod tidy
   ```

3. **Configure environment**:
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Setup database**:
   ```bash
   # Create PostgreSQL database
   createdb coffee_shop_platform
   
   # Run migrations and seed data
   make setup
   # or
   go run cmd/main.go -migrate
   go run cmd/main.go -seed
   ```

5. **Start backend**:
   ```bash
   make dev
   # or
   go run cmd/main.go
   ```

### Frontend Setup

1. **Navigate to root directory**:
   ```bash
   cd ..
   ```

2. **Install dependencies**:
   ```bash
   npm install
   ```

3. **Configure environment**:
   ```bash
   cp .env.example .env
   # Edit .env with your API URL
   ```

4. **Start frontend**:
   ```bash
   npm run dev
   ```

## 🗄️ Database Configuration

### PostgreSQL Setup

1. **Install PostgreSQL**:
   ```bash
   # Ubuntu/Debian
   sudo apt-get install postgresql postgresql-contrib
   
   # macOS (with Homebrew)
   brew install postgresql
   
   # Windows
   # Download from https://www.postgresql.org/download/
   ```

2. **Start PostgreSQL**:
   ```bash
   # Ubuntu/Debian
   sudo systemctl start postgresql
   
   # macOS
   brew services start postgresql
   ```

3. **Create database**:
   ```bash
   sudo -u postgres createdb coffee_shop_platform
   ```

4. **Configure user** (optional):
   ```bash
   sudo -u postgres psql
   CREATE USER your_user WITH PASSWORD 'your_password';
   GRANT ALL PRIVILEGES ON DATABASE coffee_shop_platform TO your_user;
   \q
   ```

### Environment Variables

Create `backend/.env`:
```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=coffee_shop_platform

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRE_HOURS=24
```

Create `.env` (root directory):
```env
# API Configuration
VITE_API_BASE_URL=http://localhost:8080

# Development Configuration
VITE_APP_NAME=Coffee Shop Platform
VITE_APP_VERSION=1.0.0
```

## 🧪 Testing the Setup

### 1. Backend Health Check
```bash
curl http://localhost:8080/health
```
Expected response:
```json
{"status":"ok"}
```

### 2. API Endpoints Test
```bash
# Get public categories
curl http://localhost:8080/api/public/categories

# Get public menu
curl http://localhost:8080/api/public/menu
```

### 3. Frontend Test
- Open http://localhost:5173
- You should see the coffee shop menu
- Try logging in as admin (username: `admin`, password: `admin123`)

## 🔐 Default Credentials

### Main Admin
- **Username**: `admin`
- **Password**: `admin123`

### Shop Admin
- **Username**: `shopadmin`
- **Password**: `shop123`

## 📱 Features Overview

### Public Features
- ✅ View coffee shop menu
- ✅ Filter by categories
- ✅ View cafe information
- ✅ Responsive design

### Admin Features
- ✅ Login system (Main Admin / Shop Admin)
- ✅ Menu item management (CRUD)
- ✅ Category management (Main Admin only)
- ✅ Shop settings management
- ✅ Real-time data updates

## 🚀 Production Deployment

### Backend Deployment

1. **Build for production**:
   ```bash
   cd backend
   go build -o bin/server cmd/main.go
   ```

2. **Set production environment**:
   ```env
   SERVER_HOST=0.0.0.0
   SERVER_PORT=8080
   DB_HOST=your-production-db-host
   DB_USER=your-production-db-user
   DB_PASSWORD=your-production-db-password
   JWT_SECRET=your-production-jwt-secret
   ```

3. **Run migrations**:
   ```bash
   ./bin/server -migrate
   ```

4. **Start server**:
   ```bash
   ./bin/server
   ```

### Frontend Deployment

1. **Build for production**:
   ```bash
   npm run build
   ```

2. **Deploy dist/ folder** to your hosting service:
   - Vercel
   - Netlify
   - AWS S3 + CloudFront
   - GitHub Pages

## 🐛 Troubleshooting

### Common Issues

1. **Backend won't start**:
   - Check if PostgreSQL is running
   - Verify database credentials in `.env`
   - Check if port 8080 is available

2. **Frontend can't connect to API**:
   - Verify `VITE_API_BASE_URL` in `.env`
   - Check if backend is running on port 8080
   - Check browser console for CORS errors

3. **Database connection failed**:
   - Ensure PostgreSQL is running
   - Check database credentials
   - Verify database exists

4. **Build errors**:
   - Run `npm install` to install dependencies
   - Check Node.js version (16+ required)
   - Clear node_modules and reinstall

### Debug Mode

Enable debug logging:
```bash
# Backend
export DEBUG=true
go run cmd/main.go

# Frontend
VITE_DEBUG=true npm run dev
```

## 📚 API Documentation

### Authentication
- `POST /api/auth/main-admin/login` - Main admin login
- `POST /api/auth/shop-admin/login` - Shop admin login

### Public Endpoints
- `GET /api/public/categories` - Get active categories
- `GET /api/public/menu` - Get public menu
- `GET /api/public/shop` - Get shop settings

### Admin Endpoints
- `GET /api/admin/menu` - Get menu items
- `POST /api/admin/menu` - Create menu item
- `PUT /api/admin/menu/:id` - Update menu item
- `DELETE /api/admin/menu/:id` - Delete menu item
- `GET /api/admin/categories` - Get categories
- `GET /api/admin/settings` - Get shop settings
- `PUT /api/admin/settings` - Update shop settings

## 🎉 Success!

If everything is working correctly, you should have:

- ✅ A running React frontend on port 5173
- ✅ A running Go backend on port 8080
- ✅ A PostgreSQL database with sample data
- ✅ Full CRUD functionality for menu items
- ✅ Category management system
- ✅ Authentication system
- ✅ Responsive design

## 📞 Support

If you encounter any issues:

1. Check the troubleshooting section above
2. Review the logs in your terminal
3. Check the browser console for errors
4. Verify all prerequisites are installed
5. Ensure all services are running

Happy coding! 🚀
