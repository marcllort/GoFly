# GoFly

GoFly is a chatbot designed to answer and suggest your travel questions. It can recommend cities, and places in them to visit (parks, restaurants, hotels, public attractions...). You can ask the bot for more information of the recommendation, and even see a history of the places it recommended before.

Installation

1. Download latest version of Go (https://golang.org/)
2. Clone the project to your local
3. In terminal/cmd, go to the folder of the project, and run: 
    - (MacOS only) chmod +x ./GoFlyMac.app
    - ./GoFlyWin.exe // ./GoFlyMac.app

To try it out, you can use the GoFly app or Postman. Both should connect to localhost:4444

The structure of the message body is:

{
    "User":"your_username",
    "Message":"your_request_message"
}

GoFly app can be found here: https://github.com/alexalmansa/fligthbot-app

To be able to use the app, use it with android studio emulator, or ask me to open the server, where I have public IP and open ports for the app.

For a more detailed explanation of how the chatbot is built, and its capabilities, refer to "info.pdf".
