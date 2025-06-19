 <h1 style="text-align: center">☀️ Sun</h1> 

![GitHub last commit](https://img.shields.io/github/last-commit/SunDeveloppments/sun)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/SunDeveloppments/sun/go-tests.yml)
![GitHub License](https://img.shields.io/github/license/SunDeveloppments/sun)
![GitHub top language](https://img.shields.io/github/languages/top/SunDeveloppments/sun)
[![gitmoji-changelog](https://img.shields.io/badge/Changelog-gitmoji-brightgreen.svg)](https://github.com/frinyvonnick/gitmoji-changelog)
<a href="https://gitmoji.dev">
  <img
    src="https://img.shields.io/badge/gitmoji-%20😜%20😍-FFDD67.svg?style=flat-square"
    alt="Gitmoji"
  />
</a>
![GitHub contributors](https://img.shields.io/github/contributors/SunDeveloppments/sun)

The core of the Sun system. Sun is a project information manager ( like language, hosting platform...). 

> [!IMPORTANT]
> Still on developpment. Some features are no implemented.

## 🛠️ Functioning

#### Init sun environnment

When you use `sun init ` , sun create a .sunenv.yaml in your working directory. 

###### *.sunenv.yaml* syntax :  

Example : sun

```yaml  
name: "sun"
author: "Jellyfish"
author-email: ""
maintener: "Jellyfish"
maintener-email: ""
language: "go"
hosting:
   platform: "github"
   repo: "github.com/SunDeveloppments/sun" 
```

> [!NOTE]
> Now the reader for .sunenv.yaml is implemented !

### 🚀 Roadmap

#### Done

- [x] init feature for initialize a .sunenv.yaml in working directory. ( just init name )
- [x] Read the .sunenv.yaml file ( read name, author and language)
- [x] Makefile for a better installation
- [X] Uninstall target in the Makefile if the user wants uninstall the project
- [X] Questions in init.go if the information flags are not provided
- [X] A function to do purcentages of langages from the current repertory!
- [X] Another function to detect used frameworks.
- [X] Automatic use of these stats to init .sunenv.yaml.

#### To do

- [ ] Automatically read config files of frameworks.
- [ ] Use these informations.
- [ ] Real-time monitoring.
- [ ] Auto-initialize frameworks.
- [ ] Detect librairies.
- [ ] Git integration.
- [ ] Automatic git pull / git push / git fetch

## 💻 Installing

### Dependencies  
For build Sun, you need

- 🐧 A linux computer  
- 🐹 Golang   
```bash  
sudo dnf install go # for Fedora, RHEL…
sudo apt update && sudo apt install go # For Ubuntu, Debian…  
sudo zypper install go # For OpenSUSE and SUSE  
sudo pacman -S go or sudo yay -S go # For ArchLinux…  
doas apk add go # For Alpine Linux  
```  
- 🏗️ Make
- 📚 Pandoc ( optionnal, doc feature )


### Install
Clone the repo:
```bash
git clone https://github.com/SunDeveloppments/sun
cd sun
```
Compile or compile and install:
```bash
make
```
```bash
sudo make install
```
Start !
```bash
sun
```
```bash
./sun
```
### Man page

Using :

```sh
man sun
```

### Uninstall
```
sudo make uninstall
```


## 👋 Contributing 

Welcome, nerds, coders, and others! Ready to contribute? You can fork this repo and open a pull request. We use the [gitmoji](https://gitmoji.dev) code for commits, if you want, you can do use it

## 📜 License

Sun is licensed under the MIT license.
