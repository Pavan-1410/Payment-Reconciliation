# माझ्या Docker image साठी Go 1.25 असलेली ready-made image वापर. "Docker Hub वर Go ची official image available असते
FROM golang:1.27        
# create/app work directry in docker image
WORKDIR /app
# या दोन files Docker container मध्ये copy करते.
COPY go.mod go.sum ./
# आता Docker container मध्ये Go dependencies download
RUN go mod download
# Current project मधील सर्व files container च्या /app मध्ये copy कर.
COPY . . 
# source code compile करून executable तयार करणे.
RUN go build -o payment-reconciliation .
# application container मध्ये port 8080 वर listen करेल अशी माहिती Docker ला सांग.
EXPOSE 8080
# ही container start झाल्यावर execute होणारी command आहे.
CMD [ "./payment-reconciliation" ]
