<div align="center">

# 🃏 CartoJester

### **The Global Comedy Atlas**

**[🌐 Explore the Live Atlas →](https://parkwardrr.github.io/CartoJester/)**

An interactive network graph mapping comedian alliances — collaborations, troupes, influences, mentorships, and rivalries — spanning 130 years across 6 continents.

---

<!-- Tech Stack -->
[![SvelteKit](https://img.shields.io/badge/SvelteKit-5-FF3E00?style=for-the-badge&logo=svelte&logoColor=white)](https://svelte.dev)
[![D3.js](https://img.shields.io/badge/D3.js-7-F9A03C?style=for-the-badge&logo=d3.js&logoColor=white)](https://d3js.org)
[![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://typescriptlang.org)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![Swift](https://img.shields.io/badge/Swift-6-F05138?style=for-the-badge&logo=swift&logoColor=white)](https://swift.org)
[![Canvas2D](https://img.shields.io/badge/Canvas2D-GPU_Accelerated-orange?style=for-the-badge&logo=html5&logoColor=white)](#)

<!-- Status -->
[![Live Site](https://img.shields.io/badge/🌐_Live_Site-GitHub_Pages-blueviolet?style=for-the-badge&logo=github)](https://parkwardrr.github.io/CartoJester/)
[![Static Site](https://img.shields.io/badge/Deploy-No_CI_Required-success?style=for-the-badge&logo=rocket)](#-build--deploy)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)
[![Apple Silicon](https://img.shields.io/badge/Apple_Silicon-M3%2FM4-333?style=for-the-badge&logo=apple&logoColor=white)](https://apple.com)

<!-- Data Stats -->
[![Comedians](https://img.shields.io/badge/Comedians-753-8b5cf6?style=flat-square&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PGNpcmNsZSBjeD0iMTIiIGN5PSIxMiIgcj0iMTAiIGZpbGw9IiNmZmYiLz48L3N2Zz4=)](#-dataset)
[![Alliances](https://img.shields.io/badge/Alliances-546-ec4899?style=flat-square)](#-dataset)
[![Alliance Types](https://img.shields.io/badge/Types-7-06b6d4?style=flat-square)](#-alliance-types)
[![Time Span](https://img.shields.io/badge/Years-1895–2025-f59e0b?style=flat-square)](#-dataset)
[![Regions](https://img.shields.io/badge/Regions-6_Continents-22c55e?style=flat-square)](#-global-coverage)

<!-- Scanners -->
[![Reddit Scanner](https://img.shields.io/badge/Reddit_Scanner-135_edges-FF4500?style=flat-square&logo=reddit&logoColor=white)](#-data-scanners)
[![News Scanner](https://img.shields.io/badge/News_Scanner-7_edges-1a73e8?style=flat-square&logo=googlenews&logoColor=white)](#-data-scanners)
[![Wikipedia](https://img.shields.io/badge/Wikipedia-Enriched-999?style=flat-square&logo=wikipedia&logoColor=white)](#-data-scanners)

<!-- Quality -->
[![Mobile Ready](https://img.shields.io/badge/Mobile-Responsive-4ade80?style=flat-square&logo=iphone&logoColor=white)](#-mobile)
[![Dark Mode](https://img.shields.io/badge/Theme-Dark_%2F_Light-818cf8?style=flat-square&logo=adjustments&logoColor=white)](#)
[![a11y](https://img.shields.io/badge/a11y-WCAG_2.1-blue?style=flat-square)](#)
[![Performance](https://img.shields.io/badge/Rendering-60fps_Canvas-brightgreen?style=flat-square)](#)

---

</div>

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🕸️ **Smart Network Graph** | GPU-accelerated Canvas2D with zoom-dependent level-of-detail, viewport culling, and density-aware rendering |
| 🧠 **Density Navigation** | Automatically reduces visual clutter in dense clusters — edges fade, labels intelligently show/hide based on zoom & importance |
| ⏱️ **Time Travel** | Dual-range slider with play/pause animation — scrub through comedy history decade by decade |
| 🔍 **Command Palette** | `⌘K` fuzzy search to instantly jump to any of 753 comedians |
| 🎨 **7 Alliance Types** | Collaboration, Troupe, Influence, Mentor, Rivalry, Studio, Family — each color-coded |
| 🏷️ **Tag Filtering** | Filter by medium (Silent, Standup, Sketch, Film, TV, SNL, Improv) and region (India, Philippines, Indonesia, Malaysia, Singapore) |
| 🌙 **Day/Dark Mode** | One-click toggle with persistence |
| 👥 **Groups & Movements** | Browse 19 comedy movements from Vaudeville to modern Indian stand-up |
| 📊 **Data Explorer** | Searchable table of all comedians and alliances with JSON download |
| 📱 **Mobile First** | Collapsible filter panel, bottom-sheet details, floating FAB toggle |
| ♿ **Accessible** | Keyboard navigation, focus management, reduced motion, WCAG contrast |

## 🌍 Global Coverage

| Region | Comedians | Key Names |
|--------|-----------|-----------|
| 🇺🇸 **USA** | 400+ | Chappelle, Seinfeld, Pryor, Rock, Martin |
| 🇬🇧 **UK** | 80+ | Gervais, Atkinson, Milligan, Cleese |
| 🇮🇳 **India** | 22 | Kapil Sharma, Zakir Khan, Vir Das, Tanmay Bhat |
| 🇵🇭🇮🇩🇲🇾🇸🇬 **SE Asia** | 18 | Dolphy, Harith Iskander, Nigel Ng, Ronny Chieng |
| 🇨🇦 **Canada** | 25+ | Jim Carrey, Mike Myers, Russell Peters |
| 🇦🇺 **Australia** | 15+ | Hannah Gadsby, Jim Jefferies |
| 🌎 **World Legends** | 10 | Chaplin, Lenny Bruce, Robin Williams |

## 🎭 Alliance Types

| Type | Color | Count | Example |
|------|-------|-------|---------|
| 🤝 **Collaboration** | ![#a855f7](https://via.placeholder.com/12/a855f7/a855f7.png) Purple | 297 | Tina Fey ↔ Amy Poehler |
| 🎪 **Troupe** | ![#f97316](https://via.placeholder.com/12/f97316/f97316.png) Orange | 198 | Monty Python members |
| 🎓 **Influence** | ![#3b82f6](https://via.placeholder.com/12/3b82f6/3b82f6.png) Blue | 22 | Richard Pryor → Eddie Murphy |
| 👨‍🏫 **Mentor** | ![#22c55e](https://via.placeholder.com/12/22c55e/22c55e.png) Green | 5 | Garry Shandling → Judd Apatow |
| ⚔️ **Rivalry** | ![#ef4444](https://via.placeholder.com/12/ef4444/ef4444.png) Red | 9 | Letterman vs. Leno |
| 🏢 **Studio** | ![#8b5cf6](https://via.placeholder.com/12/8b5cf6/8b5cf6.png) Violet | 6 | UCB founding members |
| 👨‍👩‍👧 **Family** | ![#ec4899](https://via.placeholder.com/12/ec4899/ec4899.png) Pink | 9 | Jerry Stiller → Ben Stiller |

## 🏗️ Architecture

```
CartoJester/
├── app/                          # SvelteKit frontend
│   ├── src/
│   │   ├── lib/
│   │   │   ├── data/            # seed.json (753 nodes, 546 edges) + types
│   │   │   ├── components/      # Graph, Panels, Search, Legend, TimeSlider
│   │   │   └── stores.ts       # Reactive state management
│   │   └── routes/              # Pages: /, /about, /data, /groups
│   └── package.json
├── tools/
│   ├── ingest/                  # Swift CLI — Wikidata ingestion + Apple ML dedupe
│   ├── bunt/                    # Go CLI — validation, merge, stats
│   ├── globalmerge/             # Go — KDL scaffold merger (India + SE Asia)
│   ├── redditscan/              # Go — Reddit co-mention scanner
│   └── newsscan/                # Go — Comedy news site scraper
├── docs/                        # Static output → GitHub Pages
├── DEPLOY.md
├── spec.md
└── README.md
```

## 🚀 Quick Start

```bash
git clone https://github.com/ParkWardRR/CartoJester.git
cd CartoJester/app
npm install
npm run dev
```

Open [http://localhost:5173](http://localhost:5173) and explore.

## 🔬 Data Scanners

### Reddit Scanner (`tools/redditscan/`)

[![Reddit](https://img.shields.io/badge/Subreddits_Scanned-4-FF4500?style=flat-square&logo=reddit)](tools/redditscan/)
[![Edges Found](https://img.shields.io/badge/Edges_Added-135-success?style=flat-square)](tools/redditscan/)

Scans r/StandUpComedy, r/Standup, r/comedy, r/comedians for comedian co-mentions. Only adds edges when comedians appear together in 2+ posts or in posts scoring 100+.

```bash
cd tools/redditscan && go run main.go
```

### News Scanner (`tools/newsscan/`)

[![Sources](https://img.shields.io/badge/News_Sources-5-1a73e8?style=flat-square&logo=googlenews)](tools/newsscan/)
[![Edges Found](https://img.shields.io/badge/Edges_Added-7-success?style=flat-square)](tools/newsscan/)

Scrapes Chortle, Vulture, Paste Magazine, Comedy Bureau, and The Laugh Button using proximity-window analysis for comedian co-occurrences.

```bash
cd tools/newsscan && go run main.go
```

### Global Merge (`tools/globalmerge/`)

[![Comedians Added](https://img.shields.io/badge/New_Comedians-40-8b5cf6?style=flat-square)](tools/globalmerge/)
[![Wikipedia](https://img.shields.io/badge/Enrichment-Wikipedia_API-999?style=flat-square&logo=wikipedia)](tools/globalmerge/)

Merges KDL scaffold data for India, SE Asia, and World cohorts. Enriches with Wikipedia birth years.

```bash
cd tools/globalmerge && go run main.go
```

## 🛠️ Tools

### Go: `bunt` — Data Pipeline

```bash
cd tools/bunt && go build -o bunt .
./bunt validate ../../app/src/lib/data/seed.json
./bunt stats ../../app/src/lib/data/seed.json
./bunt merge seed.json auto.json merged.json
```

### Swift: `ingest` — AI-Powered Ingestion

Requires macOS 15+ on Apple Silicon (M3/M4).

```bash
cd tools/ingest
swift run ingest fetch --top 75 --since 1890
swift run ingest merge --seed seed.json --auto auto.json
```

## 📦 Build & Deploy

[![No CI](https://img.shields.io/badge/CI-Not_Required-success?style=flat-square)](#)
[![GitHub Pages](https://img.shields.io/badge/Hosting-GitHub_Pages-blueviolet?style=flat-square&logo=github)](#)

```bash
cd app
npm run build        # dev build
# or
PUBLIC_BASE_PATH=/CartoJester npm run build   # GitHub Pages build
cd ..
git add docs/
git commit -m "build: update static site"
git push origin main
```

See [DEPLOY.md](DEPLOY.md) for detailed setup.

## 📊 Dataset

| Metric | Value |
|--------|-------|
| **Total Comedians** | 753 |
| **Total Alliances** | 546 |
| **Alliance Types** | 7 |
| **Comedy Groups** | 19 |
| **Time Span** | 1895–2025 |
| **Regions** | USA, UK, Canada, Australia, India, SE Asia, World |
| **Data Sources** | Hand-curated, Wikipedia, Reddit, Comedy news sites |

### Edge Breakdown

| Type | Count | Share |
|------|-------|-------|
| Collaboration | 297 | 54.4% |
| Troupe | 198 | 36.3% |
| Influence | 22 | 4.0% |
| Family | 9 | 1.6% |
| Rivalry | 9 | 1.6% |
| Studio | 6 | 1.1% |
| Mentor | 5 | 0.9% |

## 📱 Mobile

[![Responsive](https://img.shields.io/badge/Responsive-Yes-4ade80?style=flat-square)](#)
[![Tested On](https://img.shields.io/badge/Tested-iPhone_14_Pro-333?style=flat-square&logo=apple)](#)

- **Collapsible filter panel** — floating FAB with comedian count badge
- **Bottom-sheet details** — tapping a comedian shows info from bottom
- **Full-width graph** — no sidebar blocking on small screens
- **Touch-optimized** — pinch-to-zoom, tap-to-select

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

Built with ❤️ using SvelteKit, D3.js, Canvas2D, Swift, and Go

[![Stars](https://img.shields.io/github/stars/ParkWardRR/CartoJester?style=social)](https://github.com/ParkWardRR/CartoJester)
[![Forks](https://img.shields.io/github/forks/ParkWardRR/CartoJester?style=social)](https://github.com/ParkWardRR/CartoJester)

</div>
