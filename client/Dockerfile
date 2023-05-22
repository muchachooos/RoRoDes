FROM node:latest
WORKDIR /app
COPY . .
RUN npm install
RUN npm install xdg-utils --fix-missing
RUN npm run build
ENTRYPOINT npm run start