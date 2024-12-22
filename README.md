### **Order and Product Management System with gRPC and JWT Authentication**  

This project demonstrates a **microservices-based architecture** for managing **orders** and **products** using **gRPC** and **REST** APIs. It features two independent services—**Order Service** and **Product Service**—that communicate via **gRPC** with secure **JWT-based authentication**. The system supports placing orders, checking product availability, and updating stock in real-time.  

Key highlights include:  
- **gRPC APIs** for efficient service-to-service communication.  
- **REST Gateway** for seamless HTTP integration.  
- **JWT Authentication** for secure API access and token propagation.  
- **Dynamic Metadata Handling** to pass authentication headers across services.  
- **Scalable Design** with modular code structure and dependency management using Go modules.  

This project is ideal for demonstrating **secure inter-service communication** and **scalable microservices design** in **Go** using **gRPC**.

### **Endpoints**  

**Order Service**  
- **POST /api/v1/orders** - Place a new order  

**Product Service**  
- **GET /api/v1/products/{product_id}** - Get product details  
- **POST /api/v1/products/{product_id}/stock** - Update product stock  
