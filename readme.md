<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<a name="readme-top"></a>

<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->

<!-- PROJECT LOGO -->
<br />
<div align="center">

<h3 align="center">TEMPLATE-BACKEND-ULaMM-GO
</h3>
  <p align="center">
    Template untuk backend dari microservices untuk project ULaMM
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#project-structure">Project Structure</a>
      <ul>
        <li><a href="#go-graphql-architecture">Go-REST-API Architecture</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

-   Proyek ini merupakan template untuk project backend yang digunakan oleh tim ULaMM dengan menggunakan bahasa premrograman Go
-   Terdapat dua branch utama pada repository ini, yaitu:
    - master : branch utama dengan berbagai contoh use case (penggunaan lebih dari dua database, HTTP request, etc.)
    - blank : branch yang bisa kalian gunakan untuk memulai project baru. Silahkan rename module dan beberapa resources sesuai kebutuhan

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

-   [![Go][go.dev]][Go-Lang-url]
-   [![Docker][docker.com]][Docker-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Project Structure

Project structure overview

```sh
|-- api
|   |-- controller
|   |-- middleware
|   |-- registry
|   `-- route
|-- docs
|-- domain
|   `-- entity
|-- pkg
|   |-- datasource
|   |-- repository
|   `-- usecase
|-- test
|-- tools
|-- tzinit
`-- utils
    |-- config
    |-- constantvar
    `-- httputility
```

-   **API** berisi `controller`,`middleware`,`registry`,`route`
    **Controller** merupakan layer terluar yang berfungsi untuk menerima request dan memanipulasi response
    **Middleware** berisikan middleware
    **Registry** digunakan untuk instantiate objek (dependency injection)
    **Route** berisikan daftar seluruh endpoint
-   **docs** berisikan hasil generate dari swagger
-   **Domain** berisikan seluruh class/struct seperti entity, request, dan response untuk kebutuhan http request.
-   **pkg** berisikan layer usecase dan repository
    - **usecase** digunakan untuk mengisi logika bisnis pada request yang berasal dari controller
    - **repository** layer untuk menarik data yang berasal dari database, external API, dll.
    - **datasource** sebagai sumber data
-   **utils** berisikan fungsi-fungsi bantuan

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Go-REST-API Architecture

![Go-REST-API Architecture](docs/go-rest-api-architecture.webp)

Go-REST-API consist of:

`Controller`\
For handling request from Graphql Resolvers. Make sure Resolver doesn't have any logic to convert data.

`UseCases`\
For app logic handling. Use usecase for logic handling (remapping data, etc..)

`Repository`\
For data storing and fetching from datasource

<!-- GETTING STARTED -->

## Getting Started

Berikut dependency yang dibutuhkan serta cara instalasi dan deployment dari proyek ini.

### Prerequisites

-   gorm
    ```sh
    go get -u gorm.io/gorm
    ```
-   gin
    ```sh
    go get -u github.com/gin-gonic/gin
    ```
-   viper
    ```sh
    go get github.com/spf13/viper
    ```
-   swaggo
    ``` sh
    go install github.com/swaggo/swag/cmd/swag@latest
    ```    
-   docker
-   make

### Installation

1. go mod tidy
2. go run .

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[Go-Lang-url]: https://go.dev/
[go.dev]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[AWS-url]: https://aws.amazon.com/
[AWS.amazon.com]: https://img.shields.io/badge/Amazon_AWS-FF9900?style=for-the-badge&logo=amazonaws&logoColor=white
[Docker-url]: https://docker.com
[docker.com]: https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=fff&style=for-the-badge