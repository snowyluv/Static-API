# API for Static services

### Installation
```bash
git clone https://github.com/snowyluv/Static-API.git 
cd static-api
cp .env.example .env
go build
```

### Configure the .env
```
DATABASE_URL="postgresql://johndoe:randompassword@localhost:5432/mydb?schema=public
JWT_SECRET=""
```

### Routes
```
http://localhost:9000/register
http://localhost:9000/login
http://localhost:9000/msg/sendmessage
```
