# Build Jar stage
FROM eclipse-temurin:21-alpine AS jar-build-stage
WORKDIR /app
COPY ./gradlew ./build.gradle.kts settings.gradle.kts ./
RUN ["./gradlew", "--refresh-dependencies"]
COPY ./src ./
RUN ["./gradlew", "bootJar"]

# Second stage
FROM eclipse-temurin:21-alpine AS prod
COPY --from=jar-build-stage /app/build/libs/api.jar /app/api.jar
CMD [ "java" , "-jar" , "/app/api.jar" ]

