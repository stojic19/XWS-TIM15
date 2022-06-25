ARG ASPNET_VERSION=5.0
ARG SDK_VERSION=5.0

FROM mcr.microsoft.com/dotnet/aspnet:${ASPNET_VERSION} AS base
WORKDIR /app
ENV ASPNETCORE_URLS=http://*:8000
EXPOSE 8000

FROM mcr.microsoft.com/dotnet/sdk:${SDK_VERSION} AS build
ENV PROJECT=Chat_microservice
WORKDIR /app
COPY ./Chat_microservice .
RUN dotnet restore "${PROJECT}/${PROJECT}.csproj" && \
    dotnet build "${PROJECT}/${PROJECT}.csproj" -c Release

FROM build AS publish
ENV PROJECT=Chat_microservice
ENV PATH $PATH:/root/.dotnet/tools
RUN dotnet tool install -g dotnet-ef --version 5.0.11 && \
    dotnet publish "${PROJECT}/${PROJECT}.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app .
WORKDIR /app/publish
ENTRYPOINT ["dotnet", "Chat_microservice.dll"]