# ProdCheck

**ProdCheck** is an advanced product monitoring and enforcement tool designed to help businesses protect their brand integrity and ensure compliance with online marketplace regulations. ProdCheck automates the tracking of product listings, detects violations, and ensures that authorized resellers adhere to brand policies.

## Note

Make sure to configure the necessary environment variables in the `.env` file and set up Firebase configuration in the internal configuration engine. Firebase is used for key services within the system, and proper configuration is crucial for the system to function correctly. Please check [wiki](https://github.com/Moxirboy/ProdCheck/wiki) too for doc.

## Key Features

- **Automated Product Monitoring**: Monitor product listings across multiple e-commerce platforms, detecting unauthorized sellers, price violations, and counterfeit products.
- **Infringement Detection**: Identify intellectual property violations, ensuring that brand assets are not misused.
- **Compliance Tracking**: Monitor and enforce Minimum Advertised Price (MAP) policies, along with other compliance standards.
- **Screenshot Capturing**: Automatically capture screenshots of non-compliant listings for documentation and enforcement.
- **Real-Time Alerts**: Receive instant notifications on detected violations to enable timely corrective actions.
- **Analytics & Reporting**: Get actionable insights into reseller behavior, marketplace trends, and compliance patterns via integrated analytics.

## System Architecture

ProdCheck adopts a microservice architecture to ensure scalability, high performance, and integration with various online platforms.

### Core Components:

1. **Web Crawlers**: Golang and PHP-based scrapers gather product data in real-time from e-commerce platforms.
2. **Screenshot Service**: A dedicated microservice for capturing and processing screenshots of violating product listings.
3. **Data Queue**: Uses RabbitMQ/Kafka for asynchronous task management, ensuring smooth data processing and load distribution.
4. **User Interface**: Built with Next.js and React, providing users with an intuitive and real-time product monitoring experience.
5. **Backend API**: A Golang-based API for managing product data, scraping instructions, and integrating third-party services.
6. **Authorization**: Casbin is used for role-based access control (RBAC). It manages user permissions by enforcing policies defined for different user roles, ensuring that only authorized users can access or modify specific system features.
7. **Authentication**: Firebase Authentication manages user accounts and session handling, integrating seamlessly with Casbin for secure access control.

### Authorization Workflow (Casbin)

1. **User Roles**: Define different roles such as Admin, Manager, and Viewer within Casbin.
2. **Policy Definitions**: Set policies specifying which actions each role can perform (e.g., access dashboard, initiate scrapers, generate reports).
3. **Request Evaluation**: When a user requests access to a resource, Casbin checks if their role is permitted to perform the requested action based on the policy configuration.
4. **Enforcement**: Access is granted or denied based on the policy rules, ensuring a secure and compliant system environment.

Casbin provides flexibility and security, allowing fine-grained control over system resources and user permissions.

## Tech Stack

- **Frontend**: Next.js, React
- **Backend**: Golang, PHP, Redis, RabbitMQ/Kafka
- **Database**: PostgreSQL, Redis (for caching)
- **Web Scraping**: Colly (for scraping), chromedp (for HTML cache extraction)
- **Authentication & Cloud Services**: Firebase
- **Authorization**: Casbin, JWT
- **Deployment**: Docker, Kubernetes

## Prerequisites

Ensure that you have the following installed:

- Docker & Docker Compose
- Node.js (for running the frontend)
- Go (for backend services)
- Firebase credentials for authentication and services
- RabbitMQ or Kafka setup for message queues

## Installation and Setup

Follow the steps below to set up the ProdCheck system:

1. **Clone the repository**:

    ```bash
    git clone https://github.com/Moxirboy/ProdCheck.git
    ```

2. **Navigate to the project directory**:

    ```bash
    cd ProdCheck
    ```

3. **Set up environment variables**:

    Copy the sample `.env.example` file to `.env`:

    ```bash
    cp .env.example .env
    ```

    Open the `.env` file and add the necessary configurations, including:

    - Firebase credentials (API keys, authentication, etc.)
    - PostgreSQL database credentials
    - Redis connection details
    - RabbitMQ/Kafka settings for the queue system

4. **Install Node.js dependencies** (for the frontend):

    Navigate to the frontend directory and install the required packages:

    ```bash
    cd frontend
    npm install
    ```

5. **Configure Firebase**:

    Make sure to configure Firebase with the correct credentials. This includes setting up Firebase Authentication and any required Firebase services.

6. **Run Docker Compose**:

    To spin up all services (API, web crawlers, Redis, RabbitMQ/Kafka, etc.), use Docker Compose:

    ```bash
    docker-compose up --build
    ```

    Docker will automatically start all microservices, including backend APIs, web scrapers, and other components.

7. **Run Migrations**:

    Ensure that the database schema is up to date by running migrations. This can vary depending on your setup, but typically this would involve:

    ```bash
    docker exec -it prodcheck-api bash -c "go run migrations/migrate.go"
    ```

8. **Access the Application**:

    Once all services are up and running, you can access the ProdCheck UI through your browser at:

    ```bash
    http://localhost:3000
    ```

## Running the Application

### Using the UI

1. **Define Monitoring Parameters**: Use the user-friendly interface to define which products, resellers, or marketplaces to monitor.
2. **Schedule Crawlers**: Automate the crawling process by scheduling tasks or manually triggering them from the dashboard.
3. **Monitor Alerts**: View real-time alerts and notifications on the dashboard when violations are detected.
4. **Capture Evidence**: The system automatically captures screenshots of non-compliant product listings for documentation.
5. **Generate Reports**: Export detailed reports about compliance trends, violations, and enforcement actions.

### Using the API

ProdCheck provides a robust API to interact with its services programmatically. You can define products to monitor, manage scrapers, and access reporting via API endpoints. Refer to the API documentation (provided separately) for endpoint details and usage.
