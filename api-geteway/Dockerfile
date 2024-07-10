# Dockerfile for API Gateway (e.g., Kong)
FROM kong:latest
COPY kong.yml /usr/local/kong/declarative/
ENV KONG_DATABASE=off
ENV KONG_DECLARATIVE_CONFIG=/usr/local/kong/declarative/kong.yml
