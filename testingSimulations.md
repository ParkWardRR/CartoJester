# testingSimulations.md — High-signal test runs

## How to use this file
Each simulation is a repeatable “script” you run on local dev + built preview, and you log: pass/fail, notes, screenshots (optional), and any regressions.

## Simulation matrix (run these before “ship”)

| ID | Simulation | Setup | Steps | Expected result |
|---|---|---|---|---|
| S1 | First-time wow | Fresh browser profile, default dataset | Load `/` and do nothing for 5 seconds | Fast load, no layout jumpiness, graph readable, clear CTA to search/time |
| S2 | Chaplin anchor | Default dataset | Open ⌘K, type “Chaplin”, select | Smooth zoom-to-focus, details panel shows timeline + top alliances + evidence links |
| S3 | Time scrubbing stability | Default dataset | Drag year slider across decades, then play/pause | Graph transitions feel stable; nodes don’t “explode”; UI stays responsive |
| S4 | Range filter sanity | Default dataset | Set range 1930–1950; filter to “collaboration” | Only edges overlapping range render; counts update; legend stays accurate |
| S5 | Filter stress test | Default dataset | Toggle many edge types quickly; change cluster mode | No console errors; frame rate remains acceptable; controls never overlap content |
| S6 | Search disambiguation | Dataset includes duplicate-like names (seed a few) | Search for “Steve” or “Williams” | Results show disambiguators (years/tags); selecting one focuses correct node |
| S7 | Dedupe audit UX | After running ingestion | Visit `/data` ingestion view | Shows merge decisions + “review needed”; explains why merges occurred |
| S8 | Offline ingestion pipeline | macOS, network on | Run ingest → merge → open app | New `auto.json` loads; auto nodes flagged; no duplicates in graph |
| S9 | Deterministic build | Same inputs twice | Build twice and diff outputs | Output is stable (no random IDs, stable ordering); only timestamps differ |
| S10 | Static-host constraints | Preview built `/docs` output | Open built site with a simple static server | Works without any server APIs; no reliance on headers or backend |

## Accessibility simulations (must-pass)

| ID | Simulation | Setup | Steps | Expected result |
|---|---|---|---|---|
| A1 | Keyboard-only navigation | No mouse | Tab through filters, open details, close panels, use slider | Every interactive element reachable; focus never disappears; visible focus indicator always present (WCAG Focus Visible) [web:80] |
| A2 | Focus not obscured | Small viewport, sticky headers/panels enabled | Tab into elements near edges and within drawers | Focused element is not hidden behind sticky UI (align with WCAG 2.2 focus-not-obscured intent) [web:84] |
| A3 | Reduced motion | OS “Reduce motion” ON | Interact with graph + transitions | Animations reduce/disable; still understandable; no motion-heavy effects |
| A4 | Contrast sanity | Dark + light mode | Check legend, selected node/edge, badges | Text and key UI states remain legible; no “selected state” ambiguity |
| A5 | Screen reader smoke test | VoiceOver (macOS) | Navigate main controls + details panel | Controls have sensible names/roles; major sections announced; no traps |

## Performance simulations (practical)

| ID | Simulation | Setup | Steps | Expected result |
|---|---|---|---|---|
| P1 | Big graph preview | Inflate dataset to 300 nodes / 800 edges | Load and interact | App remains usable; progressive rendering or level-of-detail kicks in |
| P2 | Memory/leak check | Open for 10 minutes | Scrub time, open/close panels repeatedly | No runaway memory growth; no accumulating event listeners |
| P3 | Mobile throttle | DevTools CPU throttle | Pan/zoom + filter | Degrades gracefully; offers “simplify rendering” mode if needed |

## Reporting template (copy/paste per run)
| Run date | Build hash | Simulation IDs | Pass/fail summary | Notes | Follow-ups |
|---|---|---|---|---|---|
|  |  |  |  |  |  |
