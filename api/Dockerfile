# Stage 1: Build the application
FROM maven:3.9.9 AS builder

WORKDIR /workdir/server
COPY pom.xml /workdir/server/pom.xml

RUN mvn dependency:go-offline

COPY src /workdir/server/src

RUN mvn clean package -DskipTests

FROM eclipse-temurin:21-jre-alpine AS prod

RUN adduser -D runner

WORKDIR /app

COPY --from=builder /workdir/server/target/*.jar /app/app.jar

RUN rm -rf /workdir

RUN chown runner:runner /app/app.jar

USER runner

EXPOSE 9000
ENTRYPOINT ["java", "-jar", "/app/app.jar"]
