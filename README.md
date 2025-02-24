# URL Shortener Project

This project is a simple URL shortener that allows users to shorten long URLs into more manageable links. It generates a unique ID for each long URL, creates a short link using that ID, and stores both the long and short URLs in a database. When the short URL is accessed, the system retrieves the corresponding long URL from the database and redirects the user to it.

## Features

- **URL Shortening:** Takes a long URL as input and generates a shorter, unique URL.
- **Unique ID Generation:** Employs a robust mechanism to create unique IDs for each shortened URL.
- **Database Storage:** Stores the original long URL and its corresponding shortened URL in a database.
- **Redirection:** Upon accessing the shortened URL, the system retrieves the original URL from the database and redirects the user to it.
- **Easy to Use:** Simple interface for entering the long URL and getting the shortened link.

## How It Works

1.  **Input:** The user enters a long URL into the designated input field.
2.  **ID Generation:** The system generates a unique ID for the URL.
3.  **Short Link Creation:** A short link is constructed using the base URL of the service and the unique ID (e.g., `short.url/uniqueID`).
4.  **Database Storage:** The long URL and its associated short link (including the unique ID) are stored in the database.
5.  **Redirection:** When a user accesses the short link, the system:
    - Extracts the unique ID from the short URL.
    - Queries the database for the long URL associated with that ID.
    - Redirects the user to the found long URL.

## Technologies Used

- _(Placeholder - Add the specific technologies you used here. For example:)_
  - Golang
  - Gin
  - SQL (SQLite)
  - _HTML_
  - _CSS_
  - _Bootstrap_

## Setup and Installation

_(Placeholder - Provide detailed instructions on how to set up and run the project. For example:)_

1.  **Clone the repository:**
    ```bash
    git clone [repository-url]
    ```
2.  **Install dependencies:**
    ```bash
    pip install -r requirements.txt
    ```
3.  **database:**
    - Since we're using SQLite the database is create in the root project
      automatically
4.  **Run the application:**
    ```bash
    go run main.go
    ```
5.  Access on `http://127.0.0.1:8080` or specified link. You'll need to setup the .env file with the BASE_URL set to the port you're using when running locally.

## Usage

1.  Open the application in your web browser.
2.  Enter the long URL you wish to shorten in the provided field.
3.  Click the "Shorten" button.
4.  The shortened URL will be displayed.
5.  Copy the shortened URL and share it.

## Future Updates

This project is continuously being improved. Here are some planned features for future updates:

- **Favorites:**
  - Users will be able to mark their shortened links as "favorites."
  - A dedicated section will be provided to view and manage favorite links.
  - Each favorite URL will have an associated name.
- **My Links:**
  - A "My Links" section will allow users to view a history of all the URLs they have shortened.
  - User can login/signup to have a personalised experience
  - Links will be displayed in a list, optionally with the original long URL and creation date/time.
  - Optionally, users might be able to delete their created links.
- **Custom Short URLs:**
  - Allow users to specify a custom ending for their short URLs, if the chosen ID is available.
- **Analytics:**
  - Add how many times a link has been clicked.
  <!-- - **API:**
  - Develop a RESTful API to allow other applications to use the URL shortening service. -->
- **Improve UI:**
  - Make the user experience more user-friendly.

## Contributing

_(Placeholder - If you're open to contributions, add guidelines here. For example:)_

Contributions to this project are welcome! If you find a bug or have an idea for an improvement, please open an issue or submit a pull request.

## License

_(Placeholder - Specify the license under which the project is released. For example:)_

This project is licensed under the [MIT License](LICENSE).
