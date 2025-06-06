 <h1 style="text-align: center">☀️ Sun</h1> 

![GitHub last commit](https://img.shields.io/github/last-commit/SunDeveloppments/sun)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/SunDeveloppments/sun/go-tests.yml)
![GitHub License](https://img.shields.io/github/license/SunDeveloppments/sun)
![GitHub top language](https://img.shields.io/github/languages/top/SunDeveloppments/sun)

The core of the Sun system. Sun is a programming project manger for manage information ( like language, hosting platform...). 

> [!IMPORTANT]
> Yet on developpment. Some features are no implemented.

## 🛠️ Functioning

#### Init sun environnment

When you use `sun init yourapp` , sun create a .sunenv.yaml on your working directory. 

###### *.sunenv.yaml* syntax :  

Example : sun

```yaml  
name: "sun"  
author: "Jellyfish"  
maintener: "Jellyfish"  
language: "go"  
hosting:  
   platform: "github"  
   repo: "github.com/SunDeveloppments/sun"   
```  

> [!NOTE]
> The reader for .sunenv.yaml is not implemented.

### 🚀 Roadmap

- [x] init feature for initialize a .sunenv.yaml on working directory. ( just init name )
- [x] Read the .sunenv.yaml file ( read name, author and language)
- [x] Makefile for a better installation

## 👋 Contributing 

Welcome, nerds, coders, and others! Ready to contribute? You can fork this repo and open a pull request. We use the [gitmoji](https://gitmoji.dev) code for commits, if you want, you can do use it

## 📜 License

Sun is licensed under the MIT license.
