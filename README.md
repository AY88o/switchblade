# ⚡ Switchblade
> **Context Switching for Windows, solved.**

Switchblade is a lightweight CLI tool written in Go that lets you snapshot your running applications and switch between different "work modes" (e.g., Coding, Gaming, Work) instantly.

---

## 🆕 What's New in v1.4
- **One-Click Installer:** No more manual PATH editing. Just run `install.bat`.
- **Delete Command:** Finally, you can remove old profiles with `switchblade delete`.

---

## 🚀 Installation (The Easy Way)

1. Go to the [Releases Page](https://github.com/AY88o/switchblade/releases).
2. Download **`Switchblade-v1.4-Installer.zip`**.
3. **Extract** the folder.
4. Double-click **`install.bat`**.
   *(If Windows asks, click "More Info" -> "Run Anyway". It's safe!)*

That's it! Open a **new** terminal and type `switchblade help` to verify.

---

## ⚡ Getting Started (Critical Step!)

Before you start saving profiles, you must teach Switchblade what "Normal" looks like on your PC.
**1. Calibrate (Do this once)**
Close all your heavy apps (Games, IDEs, Spotify). Leave only Windows running. Then run:

```powershell
switchblade calibrate
```
This creates a Noise.json file so Switchblade knows which system processes to ignore.

## 🎮 Usage Guide

### 1. Save a Profile
Open the apps you want (e.g., VS Code + Chrome + Spotify).
```powershell
switchblade save coding
```
### 2. Switch Contexts

* **Launch Only:** Opens the apps in the profile.
    ```powershell
    switchblade go coding
    ```

* **Interactive Switch:** Asks to kill current apps before opening new ones.
    ```powershell
    switchblade go -k gaming
    ```

* **Force Switch:** Instantly kills current apps and opens the new profile.
    ```powershell
    switchblade go -fk work
    ```
### 3. Manage Profiles

* **List all saved profiles:**
    ```powershell
    switchblade ls
    ```

* **Delete a profile:**
    ```powershell
    switchblade delete coding
    ```

## 📦 Roadmap

- [x] **v1.0:** Core Logic (Save/Load)
- [x] **v1.3:** Interactive Kill Mode (`-k`)
- [x] **v1.4:** Installer Script & Delete Command
- [ ] **v2.0:** Window Layouts (Future)

---

## 📄 License
MIT License. Built with ❤️ in Go.


