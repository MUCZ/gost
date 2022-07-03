FROM ubuntu
COPY . .
CMD ["./gost", "s"]
EXPOSE 8080