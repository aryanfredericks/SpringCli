# 🌱 SpringCli

**SpringCli** is a fast, interactive CLI tool that lets you generate Spring Boot starter projects directly from your terminal — no need to visit Spring Initializr every time!

Built with Go, SpringCli helps developers create new Spring Boot applications in seconds with all the required metadata and dependencies configured.

---

## 🚀 Features

- Interactive prompts for project setup
- Fetches official Spring Initializr metadata
- Downloads and extracts Spring Boot project as a ZIP
- Supports custom group ID, artifact ID, Java version, and more
- Includes popular dependencies like `web`, `jpa`, `security`, `lombok`, etc.

---

## 📦 Installation

### 🐧 Linux / 🍎 macOS

1. [Download the binary](https://github.com/yourusername/springcli/releases) for your OS
2. Make it executable and move it to a folder in your `$PATH`:

```bash
chmod +x springcli-linux
sudo mv springcli-linux /usr/local/bin/springcli
```

3. Run it from anywhere:
```bash
springcli
```


### 🐧 Windows
1. [Download the binary](https://github.com/yourusername/springcli/releases) from the Releases page.
2. Move it to a folder like `C:\Tools\SpringCli`
3. Add that folder to your system variable `Path`:
4. Use from any terminal

```bash
springcli.exe
```

## Example
```bash
$ springcli
? Select Project Type:       maven-project
? Select Project Language:   java
? Select Boot Version:       3.2.5
? Group ID:                  com.example
? Artifact ID:               myapp
? Project Name:              myapp
? Package Name:              com.example.myapp
? Select Packaging:          jar
? Select Java Version:       17
? Select dependencies:       web, data-jpa, lombok, security

✅ Metadata fetched successfully!
📦 Downloading project...
✅ Project downloaded successfully!
🚀 Happy coding!
```

## Author
Made with ❤️ by me
