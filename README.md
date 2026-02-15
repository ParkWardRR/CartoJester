<div align="center">

# 🃏 CartoJester

### **Comedian Alliances Atlas**

**[🌐 Live Site →](https://parkwardrr.github.io/CartoJester/)**

An interactive network graph exploring comedian alliances — collaborations, troupes, influences, mentorships, and rivalries — from the silent era to today.

[![Live Site](https://img.shields.io/badge/Live_Site-GitHub_Pages-blueviolet?style=for-the-badge&logo=github)](https://parkwardrr.github.io/CartoJester/)
[![SvelteKit](https://img.shields.io/badge/SvelteKit-5-FF3E00?style=for-the-badge&logo=svelte&logoColor=white)](https://svelte.dev)
[![D3.js](https://img.shields.io/badge/D3.js-7-F9A03C?style=for-the-badge&logo=d3.js&logoColor=white)](https://d3js.org)
[![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://typescriptlang.org)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind-v4-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)](https://tailwindcss.com)
[![Swift](https://img.shields.io/badge/Swift-6-F05138?style=for-the-badge&logo=swift&logoColor=white)](https://swift.org)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Apple Silicon](https://img.shields.io/badge/Apple_Silicon-M3%2FM4-333?style=for-the-badge&logo=apple&logoColor=white)](https://apple.com)
[![Static Site](https://img.shields.io/badge/Static-No_CI_Required-success?style=for-the-badge)](DEPLOY.md)

---

<img width="900" alt="CartoJester Explorer showing comedian network graph" src="https://img.shields.io/badge/Preview-See_Live_Site-blueviolet?style=flat-square">

</div>

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🕸️ **Interactive Graph** | GPU-accelerated Canvas2D force-directed network with zoom, pan, drag, and glow effects |
| ⏱️ **Time Travel** | Dual-range slider with play/pause animation to scrub through comedy history decade by decade |
| 🔍 **Command Palette** | `⌘K` fuzzy search to instantly jump to any comedian |
| 🎨 **7 Alliance Types** | Collaboration, Troupe, Influence, Mentor, Rivalry, Studio, Family — each color-coded |
| 🏷️ **Tag Filtering** | Filter by medium: Silent, Standup, Sketch, Film, TV, SNL, Improv, Vaudeville |
| 🌙 **Day/Dark Mode** | Light mode default with one-click dark mode toggle (persisted) |
| 👥 **Groups & Movements** | Compare comedy schools, movements, and ensembles across eras |
| 📊 **Data Explorer** | Searchable table of all comedians and alliances with JSON download |
| ♿ **Accessible** | Keyboard navigation, focus management, reduced motion support, WCAG contrast |
| 📱 **Responsive** | Adapts to any screen size with collapsible panels |

## 🏗️ Architecture

```
CartoJester/
├── app/                      # SvelteKit frontend
│   ├── src/
│   │   ├── lib/
│   │   │   ├── data/         # Seed JSON + types
│   │   │   ├── components/   # Graph, Panels, Search, etc.
│   │   │   └── stores.ts     # Svelte stores (state management)
│   │   └── routes/           # Pages: /, /about, /data, /groups
│   └── package.json
├── tools/
│   ├── ingest/               # Swift CLI — Wikidata ingestion + Apple ML dedupe
│   │   ├── Package.swift
│   │   └── Sources/
│   └── bunt/                 # Go CLI — validation, merge, stats
│       ├── go.mod
│       └── main.go
├── docs/                     # Static output for GitHub Pages
├── DEPLOY.md                 # Manual deployment guide
├── spec.md                   # Full specification
└── README.md
```

## 🚀 Quick Start

```bash
# Clone
git clone https://github.com/ParkWardRR/CartoJester.git
cd CartoJester/app

# Install & run
npm install
npm run dev
```

Open [http://localhost:5173](http://localhost:5173) and explore.

## 🛠️ Tools

### Go: `bunt` — Data Pipeline

```bash
cd tools/bunt
go build -o bunt .

# Validate dataset
./bunt validate ../../app/src/lib/data/seed.json

# Show statistics
./bunt stats ../../app/src/lib/data/seed.json

# Merge seed + auto datasets
./bunt merge seed.json auto.json merged.json
```

### Swift: `ingest` — AI-Powered Ingestion

Requires macOS 15+ on Apple Silicon (M3/M4).

```bash
cd tools/ingest

# Fetch top comedians from Wikidata
swift run ingest fetch --top 75 --since 1890

# Merge datasets
swift run ingest merge --seed seed.json --auto auto.json

# Audit deduplication decisions
swift run ingest audit --auto auto.json
```

## 📦 Build & Deploy

No CI required. Build locally and commit.

```bash
cd app
npm run build:pages
cd ..
git add docs/
git commit -m "build: update static site"
git push origin main
```

See [DEPLOY.md](DEPLOY.md) for detailed GitHub Pages setup instructions.

## 📊 Dataset

| Metric | Count |
|--------|-------|
| Comedians | 130+ |
| Alliances | 200+ |
| Alliance Types | 7 |
| Time Span | 1895–2025 |
| Eras Covered | Silent, Golden Age, Counterculture, Comedy Boom, Modern |

All data is best-effort with evidence links. Seed entries are hand-curated; auto-ingested entries are flagged for verification.

## 🤝 Contributing

1. Fork the repository
2. Edit `app/src/lib/data/seed.json`
3. Run `tools/bunt/bunt validate` to check your changes
4. Run `npm run dev` in `app/` to preview
5. Submit a pull request

## 📄 License

MIT — see [LICENSE](LICENSE) for details.

---

<div align="center">

**[🌐 Visit the Live Atlas →](https://parkwardrr.github.io/CartoJester/)**

Built with ❤️ using SvelteKit, D3.js, Swift, and Go

</div>
