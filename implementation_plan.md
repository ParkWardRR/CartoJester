# Implementation Plan: CartoJester

Based on `spec.md` and the user's preference for **Go** and **Swift** (targeting Apple Silicon M3/M4), this plan outlines the architecture and execution steps.

## Architecture

We will split the stack into three distinct layers:

1.  **Presentation (Web)**: SvelteKit + D3 + TailwindCSS v4.
    *   *Why*: Best-in-class interactive graph visualization, required by spec.
    *   *Language*: TypeScript/Svelte.
2.  **Ingestion (ML/AI)**: Swift CLI (`tools/ingest`).
    *   *Why*: Access to Apple's Core ML and Natural Language frameworks for on-device embedding and deduplication.
    *   *Language*: Swift.
3.  **Data Pipeline & Orchestration**: Go CLI (`tools/bunt`).
    *   *Why*: Robust, fast data processing for validation, normalization, and merging of large JSON datasets (`seed.json` + `auto.json`). Replaces Node.js scripts.
    *   *Language*: Go.

## Data Flow

1.  **Seed Data**: Manually curated in `src/lib/data/seed.json`.
2.  **Ingestion**: `tools/ingest` (Swift) fetches from Wikidata, runs ML dedupe, and outputs `src/lib/data/auto.json`.
3.  **Processing**: `tools/bunt` (Go) validates both files, merges them deterministically, and produces the final specific `src/lib/data/graph.json` for the frontend.
4.  **Build**: SvelteKit builds the static site using `src/lib/data/graph.json`.

## Phase 1: Foundation Setup
- [ ] Initialize Git repository.
- [ ] Create `tools/ingest` (Swift package) structure.
- [ ] Create `tools/bunt` (Go module) structure.
- [ ] Scaffold SvelteKit app in `src/` with Tailwind v4.

## Phase 2: Ingestion Engine (Swift)
- [ ] Implement `WikidataSource` to fetch comedian data.
- [ ] Implement `NaturalLanguage` embedding generation (using `NLEmbedding`).
- [ ] Implement `CoreML` or heuristic deduplication logic.
- [ ] Output `auto.json` with source metadata.

## Phase 3: Data Processing (Go)
- [ ] Define Go structs matching the Data Model (Node/Edge).
- [ ] Implement strict validation logic (replacing Zod usage for build-time checks).
- [ ] Implement deterministic merge logic (Seed > Auto).
- [ ] Create CLI commands: `bunt validate`, `bunt merge`.

## Phase 4: Frontend Implementation (SvelteKit + D3)
- [ ] Implement D3 force-directed graph.
- [ ] Create UI panels (Filters, Details, Time Slider).
- [ ] Integrate `graph.json` loading.
- [ ] Polish UI (Dark mode, Mobile responsive).

## Phase 5: Deployment & Documentation
- [ ] Verify `docs/` generation for GitHub Pages.
- [ ] Write `DEPLOY.md`.
- [ ] Finalize `README.md`.

## Constraints Checklist
- [x] No CI (Manual build).
- [x] Static Output (`adapter-static`).
- [x] Apple Silicon Optimized (Core ML usage).
- [x] "Go and Swift" utilized for tooling.
