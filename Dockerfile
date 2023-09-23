FROM golang


# Set variabel lingkungan untuk Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Buat direktori kerja di dalam kontainer
WORKDIR /app

# Salin file kode Go dan berkas proyek lainnya ke direktori kerja di dalam kontainer
COPY . .

# Kompilasi kode Go
RUN go build -o JWT_Login app/main.go

# Port yang akan digunakan aplikasi Go
EXPOSE 8080

# Perintah yang akan dijalankan saat kontainer berjalan
CMD ["./JWT_Login", "rest"]