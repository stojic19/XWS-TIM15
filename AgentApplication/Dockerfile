ARG ASPNET_VERSION=5.0
ARG SDK_VERSION=5.0

FROM mcr.microsoft.com/dotnet/aspnet:${ASPNET_VERSION} AS base
WORKDIR /app
ENV ASPNETCORE_URLS=http://*:9000
EXPOSE 9000

FROM mcr.microsoft.com/dotnet/sdk:${SDK_VERSION} AS build
WORKDIR /app
COPY ./AgentApplication .
RUN dotnet restore "AgentApplication.ClassLib/AgentApplication.ClassLib.csproj" && \
    dotnet build "AgentApplication.ClassLib/AgentApplication.ClassLib.csproj" -c Release

RUN dotnet restore "AgentApplication.API/AgentApplication.API.csproj" && \
    dotnet build "AgentApplication.API/AgentApplication.API.csproj" -c Release

FROM build AS publish
ENV PATH $PATH:/root/.dotnet/tools
RUN dotnet tool install -g dotnet-ef --version 5.0.11 && \
    dotnet publish "AgentApplication.API/AgentApplication.API.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app .
WORKDIR /app/publish
ENTRYPOINT ["dotnet", "AgentApplication.API.dll"]