Search API in Go
   - A simple Go API that searches users by name using phonetic ranking with PostgreSQL.
   - It accepts a `name` query parameter.
   - It returns a ranked list of users matching the query based on phonetic similarity.
   - It includes metadata like total matches and time taken for the search.

Setup Instructions:

1. Clone the Repository
   - git clone https://github.com/Gowrisankar-GO/search-api.git

2. Install Dependencies
   - go mod tidy

3. Env Configuration
   - DB_PORT     =  [add your db port]
   - DB_HOST     =  [add your db host]
   - DB_NAME     =  [add your db name]
   - DB_USER     =  [add your db username]
   - DB_PASSWORD =  [add your db password]

4. Run the API
   - go run main.go

5. Test the API
   - curl "http://localhost:8080/search?name=john"

6.Features
   - Phonetic Matching Algorithm
   - Soundex: Transforms names into phonetic codes for easier matching.

7. Performance
   - GIN Index for fast name matching.
   - Pagination with LIMIT 100 to handle large datasets.
