FROM scratch

# add timezone data
# this allows the user to specify a timezone as an environment variable: TZ
# RUN apk add --no-cache tzdata

# copy go binary into container
ADD digital_trainer_backend /app/digital_trainer_backend

# copy db schema into container
ADD resources/schema.sql /app/schema.sql

# expose port for API access
EXPOSE 8080

# set go binary as entrypoint
CMD ["/app/digital_trainer_backend"]