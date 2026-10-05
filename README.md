# Nigeria Wards API

A RESTful API for fetching Nigerian ward boundary data from the GRID3 dataset. This API provides endpoints to retrieve ward information by state and Local Government Area (LGA).

## Features

- **RESTful API** built with Go and Gin framework
- **PostgreSQL database** for efficient data storage and retrieval
- **Docker support** for easy deployment and consistent environments
- **Environment-based configuration** using .env files
- **Flexible querying** by state and LGA
- **4044 ward records** covering 15 Nigerian states

## Dataset Information

This API uses the GRID3 NGA - Operational Wards v2.0 dataset, which contains operational ward boundaries for the following Nigerian states:

- Adamawa
- Bauchi
- Bayelsa
- Borno
- Delta
- Gombe
- Jigawa
- Kano
- Katsina
- Kwara
- Niger
- Ogun
- Osun
- Oyo
- Yobe

**Data Citation:** Center for Integrated Earth System Information (CIESIN), Columbia University 2026. GRID3 NGA - Operational Wards v2.0. New York: GRID3. https://doi.org/10.7916/gpv6-dq34

**License:** Creative Commons Attribution-Share Alike 4.0 International (CC BY-SA 4.0)

## Prerequisites

### For Local Development
- Go 1.25 or higher
- PostgreSQL 16 or higher
- Python 3.x (for data import)

### For Docker Deployment
- Docker 20.10 or higher
- Docker Compose 2.0 or higher

## Installation

### Option 1: Docker (Recommended)

This is the easiest way to run the application with all dependencies.

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd api
   ```

2. **Configure environment variables**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` with your configuration (optional - defaults are provided).

3. **Start the application**
   ```bash
   docker-compose up -d
   ```

4. **Verify the application is running**
   ```bash
   curl http://localhost:8080/api/v1/wards
   ```

### Option 2: Local Development

1. **Install Go dependencies**
   ```bash
   cd api
   go mod download
   ```

2. **Set up PostgreSQL**
   - Create a database named `grid3_wards`
   - Import the data using the provided Python script:
   ```bash
   cd ..
   python3 -m venv venv
   source venv/bin/activate  # On Windows: venv\Scripts\activate
   pip install pandas sqlalchemy psycopg2-binary
   python3 import_data.py
   ```

3. **Configure environment variables**
   ```bash
   cd api
   cp .env.example .env
   ```
   Edit `.env` with your database credentials.

4. **Run the application**
   ```bash
   go run main.go
   ```

## Configuration

The application uses environment variables for configuration. Create a `.env` file based on `.env.example`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=grid3_wards
DB_SSLMODE=disable

# Server Configuration
SERVER_PORT=8080
GIN_MODE=debug  # Use 'release' for production
```

## API Endpoints

### Get All Wards
```http
GET /api/v1/wards
```

**Response:**
```json
[
  {
    "objectid": 1,
    "country": "Nigeria",
    "iso3": "NGA",
    "state": "Adamawa",
    "statecode": "AD",
    "lga": "Demsa",
    "lga_alt_names": "",
    "ward": "Bille",
    "ward_alt_names": "Gengle",
    "multipart_count": 1,
    "source": "CIESIN",
    "date": "2026-03-31",
    "area_sqkm": 209
  }
]
```

### Get Wards by State
```http
GET /api/v1/wards/state/:state
```

**Example:**
```bash
curl http://localhost:8080/api/v1/wards/state/Adamawa
```

### Get Wards by LGA
```http
GET /api/v1/wards/lga/:lga
```

**Example:**
```bash
curl http://localhost:8080/api/v1/wards/lga/Demsa
```

## Usage Examples

### Using curl
```bash
# Get all wards
curl http://localhost:8080/api/v1/wards

# Get wards by state
curl http://localhost:8080/api/v1/wards/state/Kano

# Get wards by LGA
curl http://localhost:8080/api/v1/wards/lga/Ikeja
```

### Using Python
```python
import requests

# Get all wards
response = requests.get('http://localhost:8080/api/v1/wards')
wards = response.json()

# Get wards by state
response = requests.get('http://localhost:8080/api/v1/wards/state/Adamawa')
adamawa_wards = response.json()
```

### Using JavaScript/Node.js
```javascript
// Get all wards
fetch('http://localhost:8080/api/v1/wards')
  .then(response => response.json())
  .then(data => console.log(data));

// Get wards by state
fetch('http://localhost:8080/api/v1/wards/state/Kano')
  .then(response => response.json())
  .then(data => console.log(data));
```

## Docker Commands

```bash
# Build and start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Stop services and remove volumes
docker-compose down -v

# Restart services
docker-compose restart

# Execute commands in the API container
docker-compose exec api sh
```

## Data Import

If you need to re-import the data or use a different dataset:

1. Place your GeoPackage (.gpkg) file in the project root
2. Update the `gpkg_path` in `import_data.py`
3. Run the import script:
   ```bash
   python3 import_data.py
   ```

The script will:
- Extract data from the GeoPackage file
- Save it as CSV (backup)
- Import it directly into PostgreSQL

## Troubleshooting

### Database Connection Issues

**Error:** `connection refused` or `could not connect to server`

**Solution:**
- Verify PostgreSQL is running: `docker-compose ps`
- Check database credentials in `.env`
- Ensure the database exists: `docker-compose exec postgres psql -U postgres -l`

### Port Already in Use

**Error:** `bind: address already in use`

**Solution:**
- Change the port in `.env` (SERVER_PORT)
- Or kill the process using port 8080:
  ```bash
  lsof -ti:8080 | xargs kill -9
  ```

### Docker Build Issues

**Error:** Build fails during dependency download

**Solution:**
- Clear Docker cache: `docker system prune -a`
- Rebuild: `docker-compose build --no-cache`

## Project Structure

```
api/
├── main.go              # Main application file
├── go.mod              # Go module dependencies
├── go.sum              # Go dependency checksums
├── Dockerfile          # Docker image configuration
├── docker-compose.yml  # Docker Compose configuration
├── .env.example        # Example environment variables
├── .gitignore          # Git ignore rules
└── README.md           # This file

../
├── import_data.py      # Data import script
└── venv/               # Python virtual environment
```

## API Response Format

All endpoints return JSON with the following structure:

```json
{
  "objectid": 1,
  "country": "Nigeria",
  "iso3": "NGA",
  "state": "Adamawa",
  "statecode": "AD",
  "lga": "Demsa",
  "lga_alt_names": "",
  "ward": "Bille",
  "ward_alt_names": "Gengle",
  "multipart_count": 1,
  "source": "CIESIN",
  "date": "2026-03-31",
  "area_sqkm": 209
}
```

### Field Descriptions

- `objectid`: Unique identifier for the ward
- `country`: Country name (Nigeria)
- `iso3`: ISO 3-letter country code
- `state`: State name
- `statecode`: 2-letter state code
- `lga`: Local Government Area name
- `lga_alt_names`: Alternative LGA names
- `ward`: Ward name
- `ward_alt_names`: Alternative ward names
- `multipart_count`: Number of polygon parts
- `source`: Data source (CIESIN)
- `date`: Data creation date
- `area_sqkm`: Area in square kilometers

## Performance Considerations

- The API uses PostgreSQL for efficient querying
- Indexes are recommended on frequently queried columns (state, lga)
- For large datasets, consider implementing pagination
- Response caching can be added for frequently accessed endpoints

## Security Notes

- Never commit `.env` files version control
- Use strong passwords in production
- Enable SSL/TLS for database connections in production
- Implement rate limiting for production deployments
- Use `GIN_MODE=release` in production

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

This project uses the GRID3 dataset licensed under Creative Commons Attribution-Share Alike 4.0 International (CC BY-SA 4.0).

## Support

For issues or questions:
- Open an issue on GitHub
- Contact: omolanweshly@gmail.com
- X:https://x.com/oluwaloseye1?s=11
- LinkedIn: oluwaseyearemu32

## Acknowledgments

- GRID3 (Global Rural Geospatial data for Development)
- Center for Integrated Earth System Information (CIESIN), Columbia University
- All contributing institutions mentioned in the dataset documentation
