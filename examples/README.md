# Examples
In this section, there is examples and tutorial to learn how to use sun.

## Part 1: Basics
### The .sunenv.yaml file
#### Write .sunenv.yaml
The structure of .sunenv.yaml seems like this:
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
Here values are the values of this repos.
You can run ```sun init``` to start configuration procedure of this file.
Add `-y` to set fields to heir default values.

#### Read .sunenv.yaml
To read .sunenv.yaml, simply run this:
```bash
sun read
```
Sun will read the .sunenv.yaml file in the current directory, and shows it.
