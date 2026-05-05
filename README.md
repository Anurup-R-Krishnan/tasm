# TASM - Asset Management System

TASM is an enterprise-grade Asset Management System designed for modern organizational requirements. It provides a comprehensive suite of tools for tracking asset lifecycles, financial analytics, and maintenance scheduling through a high-performance web interface.

## Table of Contents
1. [Features](#features)
2. [Architecture](#architecture)
3. [Technology Stack](#technology-stack)
4. [Deployment](#deployment)
5. [Configuration](#configuration)
6. [Initial Setup](#initial-setup)
7. [License](#license)

## Features

### System Onboarding
- Guided 5-step initialization wizard for organizational settings.
- Dynamic configuration of physical sites, asset categories, and custom fields.

### Asset Lifecycle Management
- Real-time QR code scanning for mobile and desktop browsers.
- Full custody tracking and automated check-out/check-in workflows.
- Detailed asset profiles with historical activity logs.

### Financial and Compliance
- Automated depreciation calculation (Straight-line and custom methods).
- Procurement pipeline management with approval workflows.
- Physical inventory audit tools with automated discrepancy reporting.

### Maintenance and Operations
- Scheduled maintenance tracking and automated work order generation.
- Service history logging and maintenance contract management.
- Software license registry with seat tracking.

## Architecture

The system follows a decoupled micro-architecture:
- **Client**: Single Page Application (SPA) providing the management interface.
- **Server**: RESTful API handling business logic and database transactions.
- **Database**: Relational storage for persistent system data.
- **Monitoring**: Integrated metrics collection and visualization.

## Technology Stack

### Frontend
- Framework: Vue.js 3 (Composition API)
- Build Tool: Vite
- Styling: Tailwind CSS 4
- UI Components: PrimeVue 4
- Language: TypeScript

### Backend
- Language: Go (Golang)
- Web Framework: Gin Gonic
- ORM: GORM
- Communication: RESTful JSON API

### Infrastructure
- Database: PostgreSQL 15
- Observability: Prometheus & Grafana
- Containerization: Docker & Docker Compose

## Deployment

### Prerequisites
- Docker Engine 20.10+
- Docker Compose v2.0+

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/anuruprkris/tasm.git
   cd tasm
   ```

2. Launch the services:
   ```bash
   docker compose up -d
   ```

### Service Access
- Management Interface: http://localhost
- API Backend: http://localhost:8080
- Metrics Dashboard: http://localhost:9090

## License

This project is licensed under the MIT License.
