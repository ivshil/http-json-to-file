# HTTP JSON Post File Saver

        This project is a small practice application designed for learning and practicing how to handle HTTP POST requests with JSON bodies and save the received JSON data as files.
        The primary objective of this project is to gain hands-on experience with processing JSON data from incoming HTTP requests and persisting it to files.

#### Features:

        Accepts HTTP POST requests with JSON bodies.
        Extracts JSON data from the request body.
        Buffers the data until it reaches a certain size.
        Saves the received JSON data as files in the local file system.
        (-/Provides a simple interface for testing and exploring JSON data handling./-)
        Getting Started:

#### How to try it:
        Clone the repository to your local machine 

        Run the application using the go run command:

        go run main.go
        - Usage:

        Send an HTTP POST request to the application with a JSON body containing the data you wish to save. For example:

        bash
        Copy code
        POST http://localhost:8080/save
        Content-Type: application/json

        {
          "id": 1,
          "name": "John Doe",
          "email": "johndoe@example.com",
          "address": {"street": "123 Main St", "city": "New York", "state": "NY", "postalCode": "10001"},
          "phoneNumbers": ["+1 123-456-7890", "+1 987-654-3210", "+1 555-555-5555"],
          "isActive": true,
          "registeredAt": "2022-07-01T10:00:00Z",
          "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsias": ["basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia1", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia2", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia3", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia4", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia5", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia6", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia7", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia8", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia9", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia10", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia11", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia12", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia13", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia14", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia15", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia16", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia17", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia18", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia19", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia20", "basgasgahdokghaokasjfi;aejgad;da;jgaegjiaegina;adsia21"]
        }
        when the specified buffer size is reached, the application will extract the JSON data from the request body/ies and save it as a file in the data folder.

#### Project Structure:

        main.go: Contains the main application logic for handling HTTP requests and saving JSON data as files.
        data/: The folder where the received JSON data will be saved as files.
        Contributing:

- Happy learning and coding! 🚀