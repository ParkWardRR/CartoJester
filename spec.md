# spec.md — Comedian Alliances Atlas (GitHub Pages, no CI)

## Objective
Generate a beautiful, fast, static GitHub Pages site that lets users explore a time-filtered network graph of comedians and their alliances (collaborations, troupes, influences, mentorships, rivalries) from the silent era to today, starting with Charlie Chaplin and expanding through the most widely-known comedians across decades.

This repo must be publishable to GitHub Pages without GitHub Actions or any CI: the user will build locally and commit the static output.

## Non-negotiables
- Do not add any GitHub Actions workflows or CI configs.
- Must be 100% static output suitable for GitHub Pages.
- Must be interactive: graph pan/zoom, hover tooltips, click-to-focus, time slider (year or range), filters, search.
- Must be “production-beautiful”: typography, spacing, motion, dark/light mode, accessibility, keyboard navigation.
- Seed with “mega-known” comedians across time, beginning with Chaplin, and include a starter dataset that is clearly marked as “seed; verify + extend.”

## Tech stack (bleeding-edge but practical)
Implement with:
- SvelteKit + TypeScript, using Svelte 5 style (runes where it helps) and `@sveltejs/adapter-static` for static prerender.
- Tailwind CSS v4 for styling.
- shadcn-svelte components (or equivalent high-quality Svelte component primitives) for UI shell: dialogs, sheets, tabs, toasts, command palette.
- D3 for force-directed graph layout + zoom/pan + transitions.
- Vite tooling (default via SvelteKit).
- Prefer native browser APIs (View Transitions API where safe, IntersectionObserver, ResizeObserver) and keep bundles lean.

## Deliverables (repo contents)
Create a complete repository with:
- `README.md` (what it is, how to run, how to add data)
- `DEPLOY.md` (manual GitHub Pages publishing instructions; no CI)
- `src/` app implementing the site
- `src/lib/data/seed.json` (initial dataset)
- `src/lib/data/schema.ts` (zod schema + TS types)
- `src/lib/data/normalize.ts` (utilities to validate/normalize seed data)
- `src/lib/graph/` (D3 graph engine, layout, rendering)
- `src/routes/` pages:
  - `/` main explorer
  - `/about` methodology + definitions
  - `/data` dataset viewer + “how to contribute data”
- `static/` assets (favicon, og image placeholder)
- `package.json` scripts that support local build and GitHub Pages output to `/docs`

## Data model (must support time)
Use JSON with explicit time bounds and source links per claim.

### Node
Each comedian node:
- `id` (string, stable)
- `name` (string)
- `aka` (string[])
- `bornYear` (number | null)
- `diedYear` (number | null)
- `activeStartYear` (number | null)
- `activeEndYear` (number | null)
- `tags` (string[]) e.g. `silent`, `standup`, `sketch`, `film`, `tv`, `snl`, `keystone`
- `notability` (number 1–5, heuristic for default prominence; seed only)
- `links` ({ label, url }[]) official or canonical references

### Edge (Alliance)
Each alliance edge:
- `id`
- `sourceId`
- `targetId`
- `type` (enum): `collaboration`, `troupe`, `influence`, `mentor`, `rivalry`, `studio`, `family`
- `startYear` (number | null)
- `endYear` (number | null)
- `weight` (1–5)
- `summary` (short, original text; no quotes)
- `evidence` ({ url, label }[]) links only (no copied text)

### Rules
- Treat all seed entries as “best-effort” and keep `evidence` URLs so users can verify.
- Graph filtering by year Y: include nodes that are active at Y (if bounds exist) and edges whose [start,end] overlaps Y.
- Also support range filtering [Y0,Y1].

## UX requirements (beauty + usability)
Implement the UI with a three-panel layout:
- Center: graph canvas (SVG or Canvas; choose what gives best perf and crispness; include retina scaling)
- Left: filters (type toggles, era presets, tags, min weight)
- Right: details panel for selected node/edge (bio-ish metadata + alliances list + evidence links)

Interaction details:
- Scroll/trackpad zoom, drag pan, double-click (or button) to re-center selection.
- Hover tooltip with name + era + top alliances.
- Click a node: focus, freeze highlight, show alliances list sorted by weight and overlap with time filter.
- Time control: a bottom slider with play/pause “scrub through years” animation; include quick jumps to 1910, 1920, … 2020s.
- Search: command palette (⌘K) to jump to a comedian; fuzzy search.
- Accessibility: all controls keyboard reachable; visible focus states; respect reduced motion; color contrast safe in both themes.

## Visualization requirements (graph)
- Force layout that converges quickly and remains stable when the year changes (avoid “exploding” layout).
- Use edge bundling or subtle curve routing if it improves readability, but only if it stays performant.
- Provide “cluster by” modes: none, era, medium (film/tv/standup/sketch), organization (e.g., SNL).
- Provide legend and clear encoding:
  - Node size = notability (seed) *and/or* degree in filtered graph
  - Node color = primary tag/medium
  - Edge color = type
  - Edge opacity/width = weight

## Seed dataset (must start with Chaplin)
Create `src/lib/data/seed.json` with at least 30 nodes and 50 edges spanning decades.

Must include (at minimum) these “mega-known” anchors:
- Charlie Chaplin (first, earliest anchor)
- Buster Keaton
- Harold Lloyd
- Laurel & Hardy (both nodes + strong collaboration edge)
- The Marx Brothers (at least Groucho + “troupe/family” edges)
- Lucille Ball
- Bob Hope
- Lenny Bruce
- Richard Pryor
- George Carlin
- Eddie Murphy
- Robin Williams
- Steve Martin
- Bill Murray
- Whoopi Goldberg
- Jerry Seinfeld
- Chris Rock
- Adam Sandler
- Tina Fey
- Amy Poehler
- Dave Chappelle
- Kevin Hart
- Ali Wong
- John Mulaney

Edge guidance (keep seed defensible and avoid over-claiming):
- Prefer edges for direct collaborations (films, shows, tours), shared troupe/cast membership, documented influence/mentorship, or well-known partnerships.
- When uncertain, either omit the edge or set low weight and include evidence URLs clearly.
- Put 2–5 evidence links per edge (Wikipedia pages are acceptable as a starting point; official sites where available are better).

## Pages
### /
Primary explorer with the graph and panels.

### /about
Explain:
- What “alliance” means in this project
- How time filtering works
- What “seed notability” is and why it’s a placeholder

### /data
Show:
- schema summary (auto-generated from zod)
- the loaded dataset in a searchable table-like UI (no huge wall of JSON on screen)
- a “Download JSON” button
- a “How to contribute” section (edit seed.json, run validation)

## Build + deployment (manual, no CI)
Use adapter-static and prerender all pages.

Output strategy for GitHub Pages without Actions:
- Configure the SvelteKit build to emit into `/docs` (repo root).
- The user will run build locally and commit `/docs` to main.
- `DEPLOY.md` must give exact steps and include a one-liner build command.

Add npm scripts:
- `dev`
- `build`
- `preview`
- `build:pages` (ensures base path is correct for project pages and outputs to `/docs`)

Implement base-path handling:
- Support both user/organization pages and project pages.
- Provide a config switch via environment variable (e.g. `PUBLIC_BASE_PATH`) and document it.

In `DEPLOY.md`, include these command examples as standalone code blocks (not in tables):
- install deps
- build for pages
- commit docs
- enable Pages in repo settings (source: main, folder: /docs)

## Quality bar
- Lighthouse: aim for 90+ Performance/Accessibility/Best Practices (best-effort for an interactive graph).
- No console errors.
- Fast first load; progressive enhancement.
- Clean, modern, minimal aesthetic; tasteful motion; great typography.

## Implementation plan (what you should generate)
1) Scaffold SvelteKit + TS project.
2) Add Tailwind v4 setup and shadcn-svelte (or comparable) component foundation.
3) Implement zod schema + seed loader + validation with readable errors.
4) Build D3 graph engine with stable layout and time filtering.
5) Build the UI shell and interactions (panels, slider, search).
6) Add `/about` and `/data`.
7) Add `DEPLOY.md` and ensure `/docs` output works with GitHub Pages manually.
8) Final pass: accessibility, responsive layout, polish, and documentation.

## Constraints / safety
- Do not copy-paste copyrighted text (bios, long descriptions). Keep summaries original and short.
- If you can’t confidently assert an alliance, either omit it or mark it low-weight with “seed; verify” and include evidence links.
## NEW: Offline ingestion + Apple-ML dedupe

### Why offline
The site stays fully static for GitHub Pages; ingestion runs locally (no CI) and produces `src/lib/data/auto.json` which is committed. [web:72][web:65]

### New repo deliverables
Add:
- `tools/ingest/` (Swift Package, macOS CLI)
- `tools/ingest/Sources/` (pluggable “source adapters”)
- `tools/ingest/ML/` (Core ML + NaturalLanguage embedding utilities)
- `src/lib/data/auto.json` (generated candidates + edges)
- `src/lib/data/merge.ts` (deterministic merge of `seed.json` + `auto.json`, with conflict rules)
- `src/routes/data/ingestion` (read-only UI to inspect auto-ingested items and dedupe decisions)

### Ingestion scope (no copyright copying)
Only ingest structured metadata needed for the map: names, aliases, years, tags, link URLs, and relationship edges. Never copy long biographical text; keep `summary` fields short and original. (If a source doesn’t allow automated access, skip it.)

### Source 1: Wikidata (required)
Implement `WikidataSource` that queries Wikidata Query Service SPARQL endpoint `https://query.wikidata.org/sparql`. [web:72][web:65]
It must be able to:
- Fetch “comedians” (and adjacent categories like comedic actors, stand-up comedians) with fields: QID, label, aliases, birth/death years (if present), sitelinks count (if available), and Wikipedia URL.
- Fetch relationship-style statements usable as edges: influences, collaborations, troupe membership, etc., when available as structured statements; attach evidence URLs back to the Wikidata item or statement.

### “Top comedians” ranking (seed expansion)
Because “top” is ambiguous, define a transparent heuristic score for candidates:
- `score = a*sitelinks + b*statementCount + c*recency + d*networkPotential`
- Keep weights configurable in a small JSON config committed to repo.
- Mark results as “auto-ranked; verify”.

### Apple ML “to the max” (even if overkill)
Implement an on-device ML-assisted deduper + ranker:

1) Embeddings for similarity:
- Use Apple Natural Language embeddings (word/sentence/contextual as available) to embed `name + aliases + short descriptor`.
- Compute cosine similarity / distance to detect near-duplicates. [web:71][web:74]

2) Core ML dedupe classifier:
- Build a simple pairwise duplicate classifier (Create ML or trained externally, delivered as `.mlmodel`) that takes features such as:
  - normalizedNameExactMatch
  - editDistance / Jaro-Winkler
  - embeddingSimilarity
  - birthYearDelta (if both present)
  - sharedAliasesCount
  - sharedExternalLinksCount
  - sameWikidataQID (hard positive)
- Run inference via Core ML on-device (CPU/GPU/Neural Engine). [web:76][web:70]

3) Human-friendly explainability:
- Every dedupe merge must record “why” (top features that drove the decision) in a `dedupeLog` object stored alongside generated output.

### Dedupe rules (must prevent duplicates)
Hard keys:
- If Wikidata QID matches an existing node’s `externalIds.wikidata`, it is the same entity.

Soft keys:
- If (embeddingSimilarity >= threshold AND nameSimilarity >= threshold) OR Core ML classifier probability >= threshold, merge.
- If uncertain, do NOT merge automatically; output as a “review needed” candidate.

Canonicalization:
- Maintain a stable `id` scheme: prefer existing node id; otherwise `wd:Q####`.
- Maintain an alias index (case-folded, punctuation-stripped) to map variants.

### Output format
`tools/ingest` writes:
- `src/lib/data/auto.json` containing:
  - `generatedAt`
  - `sources` (list of source adapters used + endpoint URLs)
  - `nodes` (new + updated nodes, each with `externalIds` and `links`)
  - `edges` (new edges with evidence URLs)
  - `dedupe` (merge decisions + review-needed list)

### CLI UX (beautiful, fast)
Create a Swift CLI with polished output:
- Progress bars, timing, cache hit ratio
- A “diff-like” preview of merges (before/after)
- A `--dry-run` mode

Commands (examples):
- `swift run ingest --top 75 --since 1890 --out src/lib/data/auto.json`
- `swift run ingest --merge --seed src/lib/data/seed.json --auto src/lib/data/auto.json --out src/lib/data/merged.json`
- `swift run ingest --dedupe-audit --auto src/lib/data/auto.json`

### Caching + rate limiting
- Cache raw query results under `tools/ingest/.cache/` keyed by hash(query).
- Respect endpoint constraints; throttle requests and retry with backoff. [web:72][web:65]

### App integration
The site should load `merged.json` (or seed+auto merged at build time) and visually distinguish:
- Seed nodes/edges (hand-curated)
- Auto-ingested nodes/edges (needs verification badge)
Add a filter toggle: “Show only verified”.

### Cost
Running ingestion locally is $0 (your Mac does the compute). Core ML and NaturalLanguage run on-device. [web:76][web:70][web:71]
