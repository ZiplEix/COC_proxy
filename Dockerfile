# Utiliser l'image de base golang avec la version 1.16
FROM golang:latest

# Copier les fichiers nécessaires dans le conteneur
WORKDIR /app
COPY . .

# Télécharger les dépendances
RUN go mod download

# Compiler l'application
RUN go build -o main .

# Exposer le port 8080 utilisé par l'API
EXPOSE 8080

# Commande à exécuter lors du démarrage du conteneur
CMD ["./main"]
