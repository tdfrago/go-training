# GO PROGRAMMING PROJECT
---
# NAME

Movies To Watch (Keeps record of movies watched)

# REQUIREMENTS

1. http API
2. http client (webapp or CLI)
3. Storage - file-based (csv, json) or SQL database
4. Logging capability, multiple users
5. Deployed using Docker
6. Full documentation (README.md, sufficiently commented source code)
7. Video prezo (min 1 minute, max 3 minutes)

# DESCRIPTION

-Ideation phase

*USERS TABLE*

| Field | Data Type | Description |
|-------|-----------|-------------|
| LastName  | string    | User's Last name   |
| FirstName | string    | User's First name  |
| UserName | string  | User's Username |
| Password | string | User's Password |
| Role | string | User's role |

*MOVIES TABLE*

| Field | Data Type | Description |
|-------|-----------|-------------|
| Title  | string    | Movie title   |
| Year | integer    | Release year  |
| Director | string  | Director Name |
| Genre | string | Movie genre |
| Language | string | Language used in the movie |
| Country | string | Origin of the movie |

# ERROR HANDLING

Always report errors encountered (e.g., incorrect url)