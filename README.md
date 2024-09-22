# ProdCheck

ProdCheck is a comprehensive product monitoring and enforcement tool designed to help businesses safeguard their brand integrity and maintain compliance in online marketplaces. Similar to Trackstreet's system, ProdCheck automates the process of tracking product listings, monitoring potential violations, and ensuring that authorized resellers comply with brand policies.

## Note

Make sure to configure the necessary environment variables in the `.env` file and set up Firebase configuration in the internal configuration engine. Firebase is used for key services within the system, and proper configuration is crucial for the system to function correctly.

## Features

- **Automated Product Monitoring**: Continuously track product listings across multiple online platforms to detect unauthorized sellers, price violations, and counterfeit products.
- **Infringement Detection**: Identify potential intellectual property violations with a robust system designed to detect unauthorized usage of your brand's assets.
- **Compliance Tracking**: Ensure resellers are adhering to the company's Minimum Advertised Price (MAP) policy and other relevant compliance measures.
- **Screenshot Capturing**: Capture visual evidence of non-compliance for reporting and enforcement purposes.
- **Custom Alerts**: Receive real-time notifications on violations to take immediate action.
- **Data Analytics**: Gain insights into reseller behavior and marketplace trends with built-in analytics and reporting tools.

## Architecture

ProdCheck uses a microservice architecture to ensure scalability, performance, and ease of integration with various e-commerce platforms. 

Key components include:

- **Web Crawlers**: Deployed using Golang and PHP for scraping product information and detecting violations in real-time.
- **Screenshot Service**: A dedicated microservice that captures and processes screenshots of infringing products.
- **Data Queue**: Implements RabbitMQ/Kafka for managing data processing tasks asynchronously.
- **User Interface**: Built with a React-based front end for easy user interaction and real-time updates on product compliance.
- **Backend API**: A robust API for managing product data, integrating external services, and storing monitoring results in the database.
  
## Technologies

- **Frontend**: Next.js, React
- **Backend**: Golang, PHP, Redis, RabbitMQ/Kafka
- **Database**: SQLite, Redis for caching
- **Web Scraping**: Colly (for crawling), chromedp (for HTML cache extraction)
- **Firebase**: For authentication and cloud services
- **Deployment**: Docker, Kubernetes

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/Moxirboy/ProdCheck.git
    ```

2. **Navigate to the project directory:**

    ```bash
    cd ProdCheck
    ```

3. **Set up the environment variables**:

    ```bash
    cp .env.example .env
    ```

    Fill in the necessary details in the `.env` file.

4. **Configure Firebase**:

    Ensure that your Firebase configuration is set up correctly in the internal configuration engine. This includes Firebase credentials, API keys, and any additional services that Firebase provides.

5. **Build and run the Docker containers**:

    ```bash
    docker-compose up --build
    ```

    This will set up the required services (API, web scraper, etc.) using Docker.

6. **Access the application**:

    Once the services are running, you can access the UI via:

    ```bash
    http://localhost:3000
    ```

## Usage

1. **Configure Scraping Parameters**: Define which products or marketplaces to monitor by using the UI or via API.
2. **Schedule Crawling**: Use the scheduling feature to automate the crawling process. Crawlers can be triggered manually or via scheduled tasks.
3. **Monitor Alerts**: Check the dashboard for alerts on product violations or non-compliance issues.
4. **Generate Reports**: Export reports detailing compliance trends, violations, and enforcement actions taken.

## Contributing

We welcome contributions to the project. Please submit a pull request or open an issue for discussion.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
