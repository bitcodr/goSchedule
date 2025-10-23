# goSchedule

A lightweight, MongoDB-backed task scheduler for sending scheduled emails in Go.

## Overview

goSchedule is a Go-based scheduled email task system that fetches pending tasks from MongoDB, sends emails when they're due, and tracks task completion status. It's designed to be simple, efficient, and easy to integrate into existing systems.

## Features

- **MongoDB Integration**: Store and retrieve tasks from MongoDB
- **Scheduled Email Delivery**: Automatically send emails when tasks reach their execution date
- **Status Tracking**: Track task status (Pending, Done, Failed)
- **Configurable Limits**: Control how many tasks to process per execution
- **SMTP Support**: Send emails via any SMTP server
- **Repository Pattern**: Clean architecture with interfaces for easy testing and extensibility

## Architecture

The project follows a clean architecture pattern with clear separation of concerns:

```
goSchedule/
├── config/          # Configuration management (database, environment)
├── handler/         # Business logic handlers
├── helper/          # Utility functions (email, context)
├── model/           # Data models
├── provider/        # Service providers and dependency injection
└── repository/      # Data access layer with interfaces
```

## Requirements

- Go 1.13+
- MongoDB
- SMTP Server (for sending emails)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/amiraliio/goSchedule.git
cd goSchedule
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file based on `.env_example`:
```bash
cp .env_example .env
```

4. Configure your environment variables in `.env`:
```env
###### MongoDB ######
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_DATABASE=your_database
MONGO_USERNAME=your_username
MONGO_PASSWORD=your_password

#### Mail Server #####
MAIL_SERVER=smtp.gmail.com
MAIL_PORT=587
MAIL_USERNAME=your_email@gmail.com
MAIL_PASSWORD=your_password
MAIL_SENDER=your_email@gmail.com
```

## Usage

### Building the Project

```bash
go build -o goSchedule
```

### Running the Scheduler

```bash
./goSchedule
```

The scheduler will:
1. Load environment configuration
2. Connect to MongoDB
3. Fetch pending tasks with `executeDate <= current time`
4. Send emails for each task
5. Update task status based on email delivery success

### Task Model

Tasks in MongoDB should follow this structure:

```json
{
  "_id": ObjectId,
  "executeDate": 1234567890,  // Unix timestamp
  "reference": "task-ref-123",
  "status": "p",              // p=pending, d=done, f=failed
  "email": {
    "attachments": [],
    "body": "Email body content",
    "receiver": "recipient@example.com",
    "subject": "Email subject"
  }
}
```

### Status Values

- `p` - Pending: Task is waiting to be executed
- `d` - Done: Task completed successfully
- `f` - Failed: Task failed to complete

## Configuration

### MongoDB Connection

The application supports both authenticated and non-authenticated MongoDB connections:

- **With Authentication**: Provide `MONGO_USERNAME` and `MONGO_PASSWORD`
- **Without Authentication**: Leave `MONGO_USERNAME` and `MONGO_PASSWORD` empty

### Task Limit

By default, the application processes 3 tasks per execution. You can modify this in `handler/task.go`:

```go
filter := &model.Filter{
    Limit: 3,  // Change this value
}
```

## Project Structure Details

### Config Package
- `app.go`: Environment variable loading
- `database.go`: MongoDB connection management

### Handler Package
- `task.go`: Task processing business logic

### Helper Package
- `helpers.go`: Context timeout utilities
- `mailhelper.go`: Email sending functionality
- `mongohelper.go`: MongoDB helper functions

### Model Package
- `task.go`: Task data model and constants
- `mail.go`: Email data model
- `filter.go`: Query filter model

### Provider Package
- `app.go`: Service initialization and dependency injection

### Repository Package
- `task.go`: Task data access implementation
- `interfaces/`: Repository interfaces for dependency inversion

## Development

### Adding New Features

1. Define interfaces in `repository/interfaces/`
2. Implement business logic in `handler/`
3. Add data models in `model/`
4. Implement repository in `repository/`
5. Wire up dependencies in `provider/`

### Running as a Cron Job

To run the scheduler periodically, set up a cron job:

```bash
# Run every 5 minutes
*/5 * * * * /path/to/goSchedule >> /path/to/logs/goSchedule.log 2>&1
```

## Error Handling

The application handles errors gracefully:
- Failed email sends are logged and tasks are marked as "Failed"
- MongoDB connection errors are logged with fatal status
- Context timeouts prevent long-running operations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.

## Author

Developed by [amiraliio](https://github.com/amiraliio)

## Support

For issues, questions, or contributions, please visit the [GitHub repository](https://github.com/amiraliio/goSchedule).
