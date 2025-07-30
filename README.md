# 🚀🌠 Pulsarship

<div align="center">
  
<!-- Badges -->
[![License: GPL3](https://img.shields.io/github/license/xeyossr/pulsarship?style=for-the-badge&logo=opensourceinitiative&logoColor=white)](https://github.com/xeyossr/pulsarship/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.24+-blue?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/dl/)
[![Release](https://img.shields.io/github/v/release/xeyossr/pulsarship?style=for-the-badge&logo=github&logoColor=white&v=1)](https://github.com/xeyossr/pulsarship/releases/latest)
[![Platform](https://img.shields.io/badge/platform-Linux-lightgrey?style=for-the-badge&logo=linux&logoColor=white)](https://www.kernel.org/)
[![Development](https://img.shields.io/badge/development-active-brightgreen?style=for-the-badge&logo=github&logoColor=white)](https://github.com/xeyossr/pulsarship/graphs/commit-activity)

---

**🚀 Pulsarship** is a minimal, fast and customizable shell prompt tool written in Go.

[✨ Features](#-features) • [📦 Installation](#-installation) • [🔧 Shell Config](#-add-to-your-shell-config) • [🛠 Configuration](#-configuration) • [📜 License](#-license)

</div>

## ✨ Features

<table>
  <tr>
    <td valign="top">


- ⚙️ <b>Modular architecture</b> – Easily customize each part of the prompt  
- 🎨 <b>Color palette support</b> – Define theme colors using `palette`  
- 🧩 <b>Components</b> – Includes:
  - `cwd`, `username`, `hostname`, `git`, `character`, etc  
- 🧰 <b>Custom components</b> – Define your own modules  
- ⚡ <b>Blazing fast</b> – Lightweight Go binary  
- 🧪 <b>Extensible</b> – Right prompt, async, etc  

</td>
    <td valign="top">
      <img src="https://github.com/user-attachments/assets/e9ac28f5-a464-4a63-b74a-95968314ff0e" width="400"/>
    </td>
  </tr>
</table>


---

## 📦 Installation
### 🔁 AUR (Arch Linux / Manjaro / EndeavourOS)

If you're using an Arch-based distribution, you can install `pulsarship` from the AUR using an AUR helper like [`yay`](https://github.com/Jguer/yay) or [`paru`](https://github.com/Morganamilo/paru):
```bash
yay -S pulsarship
```
or
```bash
paru -S pulsarship
```

### 🛠️ Install via Script
You can install `pulsarship` with a single command:

```bash
curl -sS https://raw.githubusercontent.com/xeyossr/pulsarship/main/install.sh | bash
```

This script will clone the repository, build the binary, and install it for you.

### 🔧 Add to your shell config
Add the following to your `~/.config/fish/config.fish`:
**Fish:**
```bash
set -Ux PULSARSHIP_CONFIG ~/.config/pulsarship/pulsarship.toml
pulsarship init fish | source
```

**Zsh:**
Add the following to your `~/.zshrc`:
```zsh
export PULSARSHIP_CONFIG="$HOME/.config/pulsarship/pulsarship.toml"
eval "$(pulsarship init zsh)"
```

**Bash:**
Add the following to your `~/.bashrc`:
```bash
export PULSARSHIP_CONFIG="$HOME/.config/pulsarship/pulsarship.toml"
eval "$(pulsarship init bash)"
```

> Make sure to restart your shell or source the config file after editing:   
> `source ~/.config/fish/config.fish` or `source ~/.bashrc` or `source ~/.zshrc`

## 🛠 Configuration
The default config file location is `~/.config/pulsarship/pulsarship.toml`

Pulsarship uses TOML-based configuration:
```toml
prompt = '''
{cwd} 
{character} '''

add_newline = true

[cwd]
format = "^(lavender)[{cwd}]^"
max_length = 5

[character]
icon = "❯"
format = "^(peach){character}^"

[palette]
lavender = "#b4befe"
peach = "#fab387"
```

For more information visit the [Wiki](https://github.com/xeyossr/pulsarship/wiki).

## 🤝 Contributing
Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## 🐞 Reporting Issues
Found a bug or have a suggestion? [Open an issue](https://github.com/xeyossr/pulsarship/issues). Be concise and include any relevant output or screenshots.

## 📜 License
This project is licensed under the **GNU General Public License v3.0.**
See the [LICENSE](LICENSE) file for details.
