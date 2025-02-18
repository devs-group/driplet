# 🌊 Driplet

A comprehensive monorepo containing all Driplet services and components.

## 🌟 Overview

Driplet is a complete ecosystem for data capture and processing with these key components:

- **API Service**: Authentication backbone and data ingestion service
- **Scheduler**: Cloud Run-based recurring job processor
- **Chrome Extension**: User-facing browser extension
- **Landing Page**: Public-facing website built with Next.js

## 🏗️ Project Structure

```
driplet/
├── api/                # Go-based REST API service
├── scheduler/          # Go-based Cloud Run scheduler service
├── extension/          # Chrome extension (TypeScript/Vue)
├── landing/            # Next.js website
└── pkg/                # Shared Go packages
```

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Make (optional, but recommended)
- Node.js and pnpm (for extension and landing page development)
- Go (for direct API/scheduler development)

### Environment Setup

1. Clone the repository
   ```bash
   git clone https://github.com/devs-group/driplet.git
   cd driplet
   ```

2. Copy a `.env.example` file in the root directory and replace required variables

## 🛠️ Development

### Running the Backend Services

Start all services using Docker Compose:

```bash
docker compose up -d
```

Stop all services:

```bash
docker compose down
```

### Database Management

Run migrations to set up the database schema:

```bash
make migrate
```

Create a new migration:

```bash
make migration name=add_new_table
```

### Extension Development

Navigate to the extension directory:

```bash
cd extension
```

Install dependencies:

```bash
pnpm install
```

Start development server:

```bash
pnpm dev
```

Load the extension in Chrome:
1. Open Chrome and navigate to `chrome://extensions`
2. Enable "Developer mode"
3. Click "Load unpacked" and select the `extension/` folder

Build the extension for production:

```bash
pnpm build
```

### Landing Page Development

Navigate to the landing page directory:

```bash
cd landing
```

Install dependencies:

```bash
pnpm install
```

Start development server:

```bash
pnpm dev
```

Build for production:

```bash
pnpm build
```

## 🧩 Architecture

### API Service

The API service handles:
- User authentication (Google OAuth)
- Authorization
- Data ingestion into PubSub topics
- Database access and management

### Scheduler

The scheduler service:
- Runs in Cloud Run environment
- Executes recurring cron jobs
- Processes scheduled tasks

### Chrome Extension

The extension includes:
- Background scripts
- Content scripts
- Popup UI
- Options page

### Data Flow

1. User interactions captured by the Chrome extension
2. Data sent to API service
3. API publishes events to PubSub
4. Scheduler processes events according to defined schedules

## 🧪 Testing

Run backend tests:

```bash
docker compose run --rm api go test ./...
docker compose run --rm scheduler go test ./...
```

## 📦 Deployment

### API and Scheduler

The services are containerized and can be deployed to any container orchestration platform:
- Google Cloud Run

### Extension

1. Build the extension:
   ```bash
   cd extension
   pnpm build
   ```
2. Package files under `extension/`
3. Publish to Chrome Web Store
