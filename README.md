# **GoLang Web Server with SQLite**

## **Overview**
This project demonstrates how to build a simple, yet well-structured, web server in GoLang with support for CRUD (Create, Read, Update, Delete) operations. It uses SQLite for data persistence and follows best practices for code organization.

This README is designed to help beginners understand the process of building and running the project step by step.

---

## **Features**
- **CRUD Operations**:
  - Create new resources (`POST`)
  - Retrieve all resources (`GET`)
  - Update a resource (`PUT`)
  - Delete a resource (`DELETE`)
- SQLite database integration for persistent data storage.
- Well-organized codebase with clear abstraction layers.

---

## **Project Structure**
Here’s how the project is organized:

```
project/
│
├── main.go               # Entry point of the application
│
├── db/
│   ├── db.go             # Database initialization and utility functions
│
├── handlers/
│   ├── resource.go       # HTTP handlers for resource-related endpoints
│
├── models/
│   ├── resource.go       # Resource data model
│
├── data/
│   ├── resources.db      # SQLite database file
│
├── static/
│   ├── index.html        # Static file example (if applicable)
│
├── go.mod                # Go module file
└── go.sum                # Dependency lock file
```

### **Explanation**
- **`main.go`**: Sets up the server, initializes the database, and defines the HTTP routes.
- **`db/`**: Contains database-related logic like connecting to SQLite and managing the schema.
- **`handlers/`**: Contains HTTP handlers for processing requests and interacting with the database.
- **`models/`**: Defines the `Resource` model used throughout the application.
- **`data/`**: Stores the SQLite database file (`resources.db`).

---

## **Setup Instructions**

### **1. Prerequisites**
- GoLang installed (version 1.17+)
- SQLite installed

---

### **2. Clone the Repository**
Clone this project to your local machine:
```bash
git clone https://github.com/your-repo-name/project
cd project
```

---

### **3. Initialize the Project**
Run the following commands to set up the Go module and install dependencies:
```bash
go mod tidy
```

---

### **4. Set Up the Database**
Ensure the `data/resources.db` file exists. If it doesn’t, the application will create it automatically when it runs for the first time.

Alternatively, you can create it manually with:
```bash
sqlite3 data/resources.db < script.sql
```

The `script.sql` file (if present) contains the SQL schema:
```sql
CREATE TABLE IF NOT EXISTS resources (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);
```

---

### **5. Run the Server**
Start the server:
```bash
go run main.go
```

The server will listen on `http://localhost:8080`.

---

### **6. Test the Endpoints**
You can test the endpoints using `curl` or Postman.

#### **Retrieve All Resources (`GET`)**
```bash
curl -X GET http://localhost:8080/resources
```

#### **Add a New Resource (`POST`)**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"New Resource"}' http://localhost:8080/resources/add
```

#### **Update a Resource (`PUT`)**
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Resource"}' "http://localhost:8080/resources/update?id=1"
```

#### **Delete a Resource (`DELETE`)**
```bash
curl -X DELETE "http://localhost:8080/resources/delete?id=1"
```

---

## **Code Walkthrough**

### **1. Main Components**
1. **Database Initialization (`db/db.go`):**
   - Connects to SQLite and ensures the `resources` table exists.
   - Uses the `database/sql` package for interaction.

2. **HTTP Handlers (`handlers/resource.go`):**
   - Defines the logic for each HTTP method (`GET`, `POST`, `PUT`, `DELETE`).
   - Interacts with the database to perform CRUD operations.

3. **Resource Model (`models/resource.go`):**
   - Defines the `Resource` struct used to represent resources in the application.

---

### **2. Key Concepts**
- **Separation of Concerns:**
  - Database logic, HTTP handlers, and data models are separated into different packages for clarity.
- **Dependency Management:**
  - Dependencies are defined in `go.mod` and managed with Go modules.
- **SQLite Integration:**
  - The project uses SQLite for a lightweight and self-contained database.

---

## **Customization**
- Modify the database schema in `db/db.go` or `script.sql` to add more fields.
- Add authentication middleware for more complex use cases.
- Replace SQLite with PostgreSQL or MySQL for a more scalable solution.

---

## **Next Steps**
Once you’re comfortable with this project, consider:
1. Adding middleware for logging or authentication.
2. Writing unit tests using `httptest` to verify the handlers.
3. Deploying the application to a cloud platform like Heroku, AWS, or Google Cloud.
