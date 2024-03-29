<a name="readme-top"></a>



[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/khalidrianda/ArticleApp">
  </a>

<h3 align="center">Belajar bikin api reza</h3>

  <p align="center">
    Project reza belajar golang.
    <br />
    <a href="https://github.com/khalidrianda/ArticleApp"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/khalidrianda/ArticleApp">View Demo</a>
    ·
    <a href="https://github.com/khalidrianda/ArticleApp/issues">Report Bug</a>
    ·
    <a href="https://github.com/khalidrianda/ArticleApp/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#features">Features</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## Features

- Article
    - Users can make articles
    - Users can read articles sorted by the newest first
    - Users can search article based on title/body and author

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

<div>
    <a href="https://go.dev/">
    <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" title="Golang" alt="Golang" width="60"/>
    </a>&nbsp;
    <a href="https://code.visualstudio.com/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/9/9a/Visual_Studio_Code_1.35_icon.svg" title="Visual Studio Code" alt="Visual Studio Code" width="40"/>
    </a>&nbsp;
    <a href="https://aws.amazon.com/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/9/93/Amazon_Web_Services_Logo.svg" title="AWS" alt="AWS" width="50"/></a>&nbsp;
    <a href="https://www.mysql.com/">
    <img src="https://1000logos.net/wp-content/uploads/2020/08/MySQL-Logo.png" title="MySQL" alt="MySQL" width="60"/></a>&nbsp;
    <a href="https://hub.docker.com/">
    <img src="https://www.docker.com/wp-content/uploads/2022/03/vertical-logo-monochromatic.png" title="Docker" alt="Docker" width="50"/></a>&nbsp;
    <a href="https://swagger.io/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/a/ab/Swagger-logo.png" title="Swagger" alt="Swagger" width="40"/></a>&nbsp;
    <a href="https://www.postman.com/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/c/c2/Postman_%28software%29.png" title="Postman" alt="Postman" width="120"/></a>&nbsp;
</div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/khalidrianda/ArticleApp
   ```
2. Create file [config.env]
3. Enter your config in `config.env`
   ```bash
   SERVER_PORT = "YOUR SERVER PORT"
   DB_PORT     = "YOUR DB PORT"
   DB_USERNAME = "YOUR DB USERNAME" 
   DB_HOST     = "YOUR DB HOST"
   DB_NAME     = "YOUR DB NAME"
   JWT_SECRET  = "YOUR JWT SECRET"
   ```
4. Run project
   ```bash
   cd ArticleApp
   create your config.env
   go run .
   ```

### BUILD API

1. bikin folder config dan bikin config.go didalamnya
2. Create file [config.env]
3. Enter your config in `config.env`
   ```bash
   SERVER_PORT = "YOUR SERVER PORT"
   DB_PORT     = "YOUR DB PORT"
   DB_USERNAME = "YOUR DB USERNAME" 
   DB_HOST     = "YOUR DB HOST"
   DB_NAME     = "YOUR DB NAME"
   JWT_SECRET  = "YOUR JWT SECRET"
   ```
4. Setelah itu bikin folder utils/db dan create db.go
    ```
    didalamnya kita buka koneksi ke database dan mengembalikannya ke main.go
    ```
5. Bikin kerangka internal yang didalamnya terdapat handler, repository, dan usecase
    ```
    pola pengerjaaannya dari repository, lalu baru ke usecase, terus ke handler, baru kita bikin si routes untuk akses endpoint
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Documentation

_For more examples, please refer to the [OPEN API](https://app.swaggerhub.com/apis-docs/KHALIDRIANDA22_1/ArticleApp/1.0.0#/)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Article
    - [x] Make an Article
    - [x] Search articles

See the [open issues](https://github.com/khalidrianda/ArticleApp/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

[![Khalid][khalid]][khalid-url] [![LinkedIn-Khalid][linkedin-shield]][linkedin-url-1] [![Email-Khalid][email-shield]][email-1]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/khalidrianda/ArticleApp.svg?style=for-the-badge
[contributors-url]: https://github.com/khalidrianda/ArticleApp/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/khalidrianda/ArticleApp.svg?style=for-the-badge
[forks-url]: https://github.com/khalidrianda/ArticleApp/network/members
[stars-shield]: https://img.shields.io/github/stars/khalidrianda/ArticleApp.svg?style=for-the-badge
[stars-url]: https://github.com/khalidrianda/ArticleApp/stargazers
[issues-shield]: https://img.shields.io/github/issues/khalidrianda/ArticleApp.svg?style=for-the-badge
[issues-url]: https://github.com/khalidrianda/ArticleApp/issues
[license-shield]: https://img.shields.io/github/license/khalidrianda/ArticleApp.svg?style=for-the-badge
[license-url]: https://github.com/khalidrianda/ArticleApp/blob/main/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url-1]: https://linkedin.com/in/khalidrianda
[Go]: https://img.shields.io/github/go-mod/go-version/khalidrianda/ArticleApp
[go-url]: https://go.dev/
[Echo]: https://img.shields.io/badge/Echo-v4-9cf
[echo-url]: https://echo.labstack.com/
[Oauth]: https://img.shields.io/badge/OAuth-Google-informational
[oauth-url]: https://developers.google.com/identity/protocols/oauth2
[Gmail]: https://img.shields.io/badge/Gmail-Google-informational
[mail-url]: https://github.com/googleapis/google-api-go-client
[Calendar]: https://img.shields.io/badge/Calender-Google-informational
[calendar-url]: https://github.com/googleapis/google-api-go-client
[AWS]: https://img.shields.io/badge/AWS-EC2-orange
[aws-url]: https://aws.amazon.com/
[khalid]: https://img.shields.io/badge/-Khalid-black.svg?style=for-the-badge&logo=Khalid&colorB=555
[khalid-url]: https://github.com/khalidrianda
[email-shield]: https://img.shields.io/badge/gmail-DD0031?style=for-the-badge&logo=gmail&logoColor=white
[email-1]: khalidrianda12@gmail.com