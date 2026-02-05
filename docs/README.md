# CLI Documentation

This directory contains the documentation website for the CLI framework, built with [VitePress](https://vitepress.dev/).

## Development

### Prerequisites

- Node.js 18+ 
- pnpm 9+

### Setup

```bash
# Install dependencies
pnpm install

# Start development server
pnpm dev

# Build for production
pnpm build

# Preview production build
pnpm preview
```

## Project Structure

```
docs/
├── .vitepress/          # VitePress configuration
│   ├── config.ts        # Main config file
│   └── theme/           # Custom theme
├── guide/               # Guide documentation
├── examples/            # Code examples
├── api/                 # API reference
├── index.md             # Homepage
├── package.json         # Dependencies
└── tsconfig.json        # TypeScript config
```

## Deployment

The documentation is automatically deployed to GitHub Pages via GitHub Actions when changes are pushed to the `master` branch in the `docs/` directory.

## Configuration

- Base path: `/cli/` (configured in `.vitepress/config.ts`)
- Build output: `.vitepress/dist/`
- GitHub Pages: Automatically deployed from `docs/` directory
