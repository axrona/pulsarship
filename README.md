# 🚀🌠 Pulsarship

<div align="center">

<!-- Badges -->
[![License: MIT](https://img.shields.io/github/license/xeyossr/pulsarship?style=for-the-badge&logo=opensourceinitiative&logoColor=white)](https://github.com/xeyossr/pulsarship/blob/main/LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.24+-blue?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/dl/)
[![Project Status](https://img.shields.io/badge/status-pre--release-orange?style=for-the-badge&logo=git&logoColor=white)](https://github.com/xeyossr/pulsarship/issues)
[![Platform](https://img.shields.io/badge/platform-Linux-lightgrey?style=for-the-badge&logo=linux&logoColor=white)](https://www.kernel.org/)
[![Development](https://img.shields.io/badge/development-active-brightgreen?style=for-the-badge&logo=github&logoColor=white)](https://github.com/xeyossr/pulsarship/graphs/commit-activity)


---

**🚀 Pulsarship** is a minimal, fast and customizable shell prompt tool written in Go.

[✨ Features](#-features) • [📦 Installation](#-installation) • [🔧 Shell Config](#-add-to-your-shell-config) • [🛠 Configuration](#-configuration) • [🚧 Roadmap](#-roadmap) • [📜 License](#-license)

</div>

---
> ⚠️ Note: Pulsarship is still under development. Until the first stable release, it is not recommended for daily use.

## ✨ Features

- ⚙️ **Modular architecture** – Easily customize each part of the prompt
- 🎨 **Color palette support** – Define theme colors using `palette` and reference them in components
- 🧩 **Components** – Includes built-in components like:
  - `cwd`, `username`, `hostname`, `character`, `time` 
  - with more planned in the future.
- ⚡ **Blazing fast** – Lightweight Go binary with minimal memory usage
- 🧪 **Extensible** – Future support planned for right prompt, custom modules, async updates

---

## 📦 Installation

```bash
git clone https://github.com/xeyossr/pulsarship
cd pulsarship
make install
```

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
Pulsarship uses a TOML-based configuration file:
```ini
prompt = '''

{username}@{hostname}: {cwd}  {git}
{character} '''

[character]
icon = ">"
format = "^(peach){character}^"

[hostname]
color = "mauve"
format = "^(mauve){hostname}^"

[username]
format = "^(sky){username}^"

[cwd]
format = "^(lavender){cwd}^"
max_length = 3

[git]
format = "^(yellow) {git}^"
dirty_suffix = " ^(red)[*]^"
clean_suffix = ""

[palette]
yellow = "#f9e2af"
red = "#f38ba8"
peach = "#fab387"
lavender = "#b4befe"
sky = "#89dceb"
mauve = "#cba6f7"
```

➡️ For more, check out the [Wiki](Wiki) (coming soon)

## 🚧 Roadmap
- [x] Basic prompt rendering
- [x] Component system
- [x] Color palette mapping
- [x] Async component rendering
- [ ] Right prompt support
- [ ] AUR packaging and support for Arch-based systems
- [ ] Support for writing custom user-defined modules
- [ ] Performance improvements and further optimizations

## 🤝 Contributing
Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## 🐞 Reporting Issues
Found a bug or have a suggestion? [Open an issue](https://github.com/xeyossr/pulsarship/issues). Be concise and include any relevant output or screenshots.

## 📜 License
This project is licensed under the **GNU General Public License v3.0.**
See the **[LICENSE](LICENSE)** file for details.