# Deploying CartoJester to GitHub Pages

This guide covers publishing the site manually (no CI/Actions required).

## Prerequisites

- Node.js 18+ and npm
- macOS with Apple Silicon (M3/M4) for ingestion tools

## Build for GitHub Pages

Install dependencies:

```bash
cd app
npm install
```

Build for GitHub Pages (with base path for project pages):

```bash
npm run build:pages
```

This outputs the static site to `/docs` at the repo root.

## Commit and Push

```bash
cd ..
git add docs/
git commit -m "build: update static site"
git push origin main
```

## Enable GitHub Pages

1. Go to **Settings → Pages** in your repository
2. Set **Source** to `Deploy from a branch`
3. Set **Branch** to `main` and **Folder** to `/docs`
4. Click **Save**

The site will be live at `https://ParkWardRR.github.io/CartoJester/` within a few minutes.

## Custom Domain (Optional)

If using a custom domain, update the base path:

```bash
PUBLIC_BASE_PATH="" npm run build
```

Then add a `CNAME` file to `docs/` with your domain.

## Updating the Site

After any content or code changes:

```bash
cd app
npm run build:pages
cd ..
git add docs/
git commit -m "build: update site"
git push origin main
```
