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

The webapp implements REST API for a database of movies. Data is stored in a MySQL database. The webapp is deployed using Docker.

This webapp allows a user to make his/her own movie watchlist. A user can signup to the webapp, login to create, edit, and view the movie watchlist, and logout when done. In creating a movie entry in the user's list, user must specify the movie title, genre, release year, director, language, country, and the status whether user plans to watch it, has watched it, or currently watching it. All user activies are then saved into a log text file.

| HTTP Verb | User signup /signup  | User login /login | User logout /logout |
|-----------|------------|----------------|----------------|----------------|
| POST      | 201 (Created), creates user; username already taken  | 200 (OK); username/password incorect | 200 (OK); not logged in |
| GET       | 405 (Not allowed) | 405 (Not allowed) | 405 (Not allowed) |
| PUT      | 405 (Not allowed) | 405 (Not allowed) | 405 (Not allowed) |
| DELETE   | 405 (Not allowed) | 405 (Not allowed) | 405 (Not allowed) | 

| HTTP Verb | Movie watchlist /movies  | Specific movie /movies/?Id={id} |
|-----------|------------|----------------|
| POST      | 201 (Created), creates movie entry; movie already created  | 405 (Not allowed) |
| GET       | 200 (OK), retrieves movie watchlist | 200 (OK), retrieves movie; 404 (Not found) |
| PUT      | 405 (Not allowed) | 200 (OK), updates movie entry; 404 (Not found) |
| DELETE   | 405 (Not allowed) | 200 (OK), removes movie entry; 404 (Not found) | 

*USERS TABLE*

| Field | Data Type | Description |
|-------|-----------|-------------|
| Id  | int    | User Id number   |
| LastName  | string    | Last name   |
| FirstName | string    | First name  |
| UserName | string  | Username |
| Password | string | Password (hash)|

*MOVIES TABLE*

| Field | Data Type | Description |
|-------|-----------|-------------|
| Id  | int    | Movie Id on user's list   |
| Title  | string    | Movie title   |
| Genre | string | Movie genre |
| Year | integer    | Release year  |
| Director | string  | Director Name |
| Language | string | Language used in the movie |
| Country | string | Origin of the movie |
| Status | string | If watched, on-going, or to be watched |
| UserName | string | owner of list |

# ERROR HANDLING

Always report errors encountered (e.g., incorrect url)