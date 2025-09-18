# Coffee Shop Platform - Multi-Tenant Frontend

A modern React frontend for the multi-tenant coffee shop management platform, integrated with a Go backend API.

## 🚀 Features

- ✅ **API Integration**: Fully integrated with Go backend APIs
- ✅ **Multi-Tenant Support**: Works with subdomain-based tenant routing
- ✅ **Category Management**: Dynamic categories from backend
- ✅ **Menu Management**: Full CRUD operations for menu items
- ✅ **Authentication**: JWT-based authentication for both admin types
- ✅ **Responsive Design**: Mobile-first responsive design
- ✅ **Real-time Updates**: Live data from backend APIs
- ✅ **Error Handling**: Comprehensive error handling and user feedback

## 🏗️ Architecture

### Frontend Structure
```
src/
├── components/          # Reusable UI components
├── entities/           # Data models and API integration
│   ├── MenuItem.js     # Menu item entity with API methods
│   ├── CafeSettings.js # Cafe settings entity
│   ├── Category.js     # Category entity
│   └── index.js        # Entity exports
├── pages/              # Page components
│   ├── Menu.jsx        # Public menu page
│   ├── AdminLogin.jsx  # Admin login page
│   └── AdminDashboard.jsx # Admin dashboard
├── services/           # API service layer
│   └── api.js          # Centralized API service
└── utils/              # Utility functions
```

### API Integration
- **Centralized API Service**: Single point for all API calls
- **Entity Classes**: Object-oriented data models with API methods
- **Error Handling**: Consistent error handling across the app
- **Token Management**: Automatic JWT token handling

## 🚀 Quick Start

### Prerequisites
- Node.js 16+
- npm or yarn
- Go backend running on port 8080

### Installation

1. **Install dependencies**:
   ```bash
   npm install
   # or
   yarn install
   ```

2. **Configure environment**:
   ```bash
   cp .env.example .env
   # Edit .env with your API URL
   ```

3. **Start the backend** (in another terminal):
   ```bash
   cd backend
   make setup  # Migrate and seed database
   make dev    # Start backend server
   ```

4. **Start the frontend**:
   ```bash
   npm run dev
   # or
   yarn dev
   ```

5. **Open your browser**:
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080

## 🔧 Configuration

### Environment Variables
Create a `.env` file in the root directory:

```env
# API Configuration
VITE_API_BASE_URL=http://localhost:8080

# Development Configuration
VITE_APP_NAME=Coffee Shop Platform
VITE_APP_VERSION=1.0.0
```

### API Endpoints
The frontend automatically connects to these backend endpoints:

- **Public**: `/api/public/*` - No authentication required
- **Admin**: `/api/admin/*` - Requires authentication
- **Auth**: `/api/auth/*` - Authentication endpoints

## 📱 Pages and Features

### 1. Public Menu Page (`/`)
- **Dynamic Categories**: Loaded from backend API
- **Menu Items**: Real-time data from backend
- **Category Filtering**: Filter items by category
- **Responsive Design**: Works on all devices
- **Cafe Information**: Display cafe settings from backend

### 2. Admin Login Page (`/admin-login`)
- **Dual Admin Support**: Main admin and shop admin login
- **JWT Authentication**: Secure token-based authentication
- **Error Handling**: User-friendly error messages
- **Form Validation**: Client-side validation

### 3. Admin Dashboard (`/admin-dashboard`)
- **Menu Management**: Full CRUD operations for menu items
- **Category Selection**: Choose from backend categories
- **Real-time Updates**: Live data synchronization
- **Role-based Access**: Different features for different admin types

## 🔐 Authentication

### Login Flow
1. User selects admin type (Main Admin or Shop Admin)
2. Enters credentials
3. Frontend calls appropriate API endpoint
4. JWT token stored in localStorage
5. Token included in all subsequent API calls

### Token Management
- **Automatic Storage**: Tokens stored in localStorage
- **Auto-include**: Tokens automatically added to API requests
- **Logout**: Tokens cleared on logout
- **Expiration**: Handled by backend

## 🎨 UI Components

### Design System
- **Tailwind CSS**: Utility-first CSS framework
- **Framer Motion**: Smooth animations and transitions
- **Lucide React**: Beautiful icon library
- **Responsive**: Mobile-first design approach

### Key Components
- **MenuItemCard**: Display menu items with pricing
- **CategorySection**: Group items by category
- **CafeHeader**: Display cafe information
- **AdminForm**: Form for adding/editing items

## 🔄 Data Flow

### 1. Data Loading
```
Page Load → API Service → Backend API → Database → Response → Entity Classes → React State
```

### 2. Data Updates
```
User Action → Form Submission → API Service → Backend API → Database → Response → State Update → UI Update
```

### 3. Error Handling
```
API Error → Error Response → User-friendly Message → UI Display
```

## 🛠️ Development

### Available Scripts
```bash
# Development
npm run dev          # Start development server
npm run build        # Build for production
npm run preview      # Preview production build

# Linting
npm run lint         # Run ESLint
npm run lint:fix     # Fix ESLint issues
```

### Code Structure
- **Components**: Reusable UI components
- **Pages**: Full page components
- **Entities**: Data models with API integration
- **Services**: API communication layer
- **Utils**: Helper functions

## 🚀 Deployment

### Production Build
```bash
npm run build
```

### Environment Configuration
Update `.env` for production:
```env
VITE_API_BASE_URL=https://your-api-domain.com
```

### Static Hosting
The built files in `dist/` can be deployed to any static hosting service:
- Vercel
- Netlify
- GitHub Pages
- AWS S3 + CloudFront

## 🔧 API Integration

### Entity Classes
Each entity class provides methods for API operations:

```javascript
// MenuItem operations
const items = await MenuItem.list();           // Get all items
const item = await MenuItem.getById(1);        // Get single item
const newItem = await MenuItem.create(data);   // Create new item
await item.update(data);                       // Update item
await item.delete();                           // Delete item

// Category operations
const categories = await Category.getPublicCategories();
const allCategories = await Category.getAllCategories();

// Cafe settings
const settings = await CafeSettings.getPublicSettings();
```

### API Service
Centralized API service handles all HTTP requests:

```javascript
import apiService from './services/api';

// Authentication
await apiService.loginMainAdmin(credentials);
await apiService.loginShopAdmin(credentials);

// Data operations
await apiService.getPublicMenu();
await apiService.getMenuItems();
await apiService.createMenuItem(data);
```

## 🐛 Troubleshooting

### Common Issues

1. **API Connection Failed**
   - Check if backend is running on port 8080
   - Verify `VITE_API_BASE_URL` in `.env`
   - Check browser console for CORS errors

2. **Authentication Issues**
   - Clear localStorage and try logging in again
   - Check if JWT token is valid
   - Verify admin credentials

3. **Data Not Loading**
   - Check browser network tab for failed requests
   - Verify backend API endpoints
   - Check console for error messages

### Debug Mode
Enable debug logging by adding to `.env`:
```env
VITE_DEBUG=true
```

## 📚 API Documentation

### Backend API Endpoints

#### Public Endpoints
- `GET /api/public/categories` - Get active categories
- `GET /api/public/menu` - Get public menu
- `GET /api/public/shop` - Get shop settings

#### Authentication
- `POST /api/auth/main-admin/login` - Main admin login
- `POST /api/auth/shop-admin/login` - Shop admin login

#### Admin Endpoints
- `GET /api/admin/menu` - Get menu items
- `POST /api/admin/menu` - Create menu item
- `PUT /api/admin/menu/:id` - Update menu item
- `DELETE /api/admin/menu/:id` - Delete menu item
- `GET /api/admin/categories` - Get categories
- `GET /api/admin/settings` - Get shop settings
- `PUT /api/admin/settings` - Update shop settings

## 🎉 Features Summary

- ✅ **Full API Integration**: Complete backend integration
- ✅ **Multi-Tenant Support**: Subdomain-based routing
- ✅ **Dynamic Categories**: Backend-managed categories
- ✅ **Real-time Data**: Live updates from backend
- ✅ **JWT Authentication**: Secure token-based auth
- ✅ **Responsive Design**: Mobile-first approach
- ✅ **Error Handling**: User-friendly error messages
- ✅ **Type Safety**: Consistent data models
- ✅ **Performance**: Optimized API calls
- ✅ **Maintainable**: Clean, organized code structure

The frontend is now fully integrated with the Go backend and ready for production use! 🚀
