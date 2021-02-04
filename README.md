# Crypto Faucet
Incognito Faucet - A fullstack Go + ReactJS + Nginx  App deployed with Docker

## What did I do ?
- Built a simple Golang + ReactJS web application running on docker compose. The application uses the Gin framework and stores documents in MySQL.

### Pre-installation
* Git
* Docker
* Docker-compose
### Installation
* Clone this repo

  ```
    git clone https://github.com/hoangnguyen-1312/crypto-faucet..git <project-name>
    cd <project-name>
  ```
### Quick Start
1) `sudo docker-compose up --build`
2) Visit `https://localhost:3000` (*note **https***)
3) Make changes to either golang or react code, and re-run step 1 to apply changes.

### Preview of the app
![Screenshot of the app](docs/demo.png?raw=true "Screenshot")

### Description
- Crypto faucet - a really easy way to earn crypto for free
- This web application was developed to give a reward of 30 PRV - a native coin of Incognito - for users who own Incognito Wallet.  
