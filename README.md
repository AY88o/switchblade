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

<details>
<summary><strong>Troubleshooting: "Command Not Found"</strong></summary>

If the `switchblade` command doesn't work after running the installer:

1. Search Windows for **"Edit the system environment variables"**.
2. Click **Environment Variables** > **Path** > **Edit**.
3. Click **New** and paste: `C:\Program Files\Switchblade`
4. Click **OK** and restart your terminal.

</details>

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

## 🛠️ Troubleshooting

<details>
<summary><strong>⚠️ "Virus Detected" or Installation Blocked</strong></summary>

Because Switchblade is a new open-source tool, Windows Defender acts cautiously.
1. Go to **Windows Security** > **Virus & threat protection**.
2. Click **Protection history**.
3. Find the "Threat blocked" item (Switchblade), click **Actions**, and select **Allow on device**.
4. Redownload the zip file.
</details>

<details>
<summary><strong>🚫 "Command Not Found" (Manual Install)</strong></summary>

If the installer script fails to set the path automatically:
1. Create a folder at `C:\Program Files\Switchblade`.
2. Manually move `switchblade.exe` into that folder.
3. Add `C:\Program Files\Switchblade` to your **System Environment Variables (Path)**.
4. Restart your terminal.
</details>

<details>
<summary><strong>📂 `ls` command shows too many files?</strong></summary>

Switchblade scans the current directory for profile files. If you run it from your Desktop, it might see *everything*.
* **Tip:** Create a dedicated folder (e.g., `Documents\Switchblade`) and save your profiles there to keep your `ls` output clean.
</details>

## 📄 License
MIT License. Built with ❤️ in Go.


