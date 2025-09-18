#!/bin/bash

# Coffee Shop Platform - Development Startup Script

echo "🚀 Starting Coffee Shop Platform Development Environment"
echo "=================================================="

# Check if backend directory exists
if [ ! -d "backend" ]; then
    echo "❌ Backend directory not found!"
    exit 1
fi

# Function to cleanup background processes
cleanup() {
    echo ""
    echo "🛑 Shutting down services..."
    kill $BACKEND_PID $FRONTEND_PID 2>/dev/null
    exit 0
}

# Set up signal handlers
trap cleanup SIGINT SIGTERM

# Start backend
echo "🔧 Starting Go backend..."
cd backend
if [ ! -f "bin/server" ]; then
    echo "📦 Building backend..."
    go build -o bin/server cmd/main.go
fi

# Check if database is set up
echo "🗄️ Checking database setup..."
if ! ./bin/server -migrate > /dev/null 2>&1; then
    echo "⚠️ Database migration failed. Please check your database connection."
    echo "Make sure PostgreSQL is running and configured correctly."
    exit 1
fi

# Seed database if needed
echo "�� Seeding database..."
./bin/server -seed > /dev/null 2>&1

# Start backend server
echo "🚀 Starting backend server on port 8080..."
./bin/server &
BACKEND_PID=$!

# Wait a moment for backend to start
sleep 2

# Go back to root directory
cd ..

# Start frontend
echo "🎨 Starting React frontend..."
npm run dev &
FRONTEND_PID=$!

echo ""
echo "✅ Services started successfully!"
echo "=================================================="
echo "🌐 Frontend: http://localhost:5173"
echo "🔧 Backend API: http://localhost:8080"
echo "📚 API Docs: http://localhost:8080/health"
echo ""
echo "Press Ctrl+C to stop all services"
echo "=================================================="

# Wait for processes
wait
