<svelte:head>
    <title>About — CartoJester</title>
    <meta
        name="description"
        content="How the Comedian Alliances Atlas works: alliance types, time filtering, and seed notability."
    />
</svelte:head>

<div class="about-page">
    <div class="about-content">
        <h1>About CartoJester</h1>
        <p class="subtitle">Understanding the Comedian Alliances Atlas</p>

        <section>
            <h2>What is an "Alliance"?</h2>
            <p>
                An alliance represents a meaningful connection between two
                comedians. We track seven distinct types:
            </p>
            <div class="type-grid">
                <div class="type-card" style="--tc:#a855f7">
                    <span class="dot"></span>
                    <h3>Collaboration</h3>
                    <p>Co-starred in films, shows, or tours together</p>
                </div>
                <div class="type-card" style="--tc:#f59e0b">
                    <span class="dot"></span>
                    <h3>Troupe</h3>
                    <p>Members of the same comedy group, cast, or ensemble</p>
                </div>
                <div class="type-card" style="--tc:#3b82f6">
                    <span class="dot"></span>
                    <h3>Influence</h3>
                    <p>
                        Documented comedic influence on another performer's
                        style
                    </p>
                </div>
                <div class="type-card" style="--tc:#10b981">
                    <span class="dot"></span>
                    <h3>Mentor</h3>
                    <p>
                        Direct mentorship or professional guidance relationship
                    </p>
                </div>
                <div class="type-card" style="--tc:#ef4444">
                    <span class="dot"></span>
                    <h3>Rivalry</h3>
                    <p>
                        Professional competition or creative rivalry between
                        performers
                    </p>
                </div>
                <div class="type-card" style="--tc:#8b5cf6">
                    <span class="dot"></span>
                    <h3>Studio</h3>
                    <p>
                        Connected through the same studio, production company,
                        or label
                    </p>
                </div>
                <div class="type-card" style="--tc:#ec4899">
                    <span class="dot"></span>
                    <h3>Family</h3>
                    <p>
                        Blood relatives who performed together or in the same
                        field
                    </p>
                </div>
            </div>
        </section>

        <section>
            <h2>Time Filtering</h2>
            <p>
                Every node and edge has optional time bounds. When you set a
                year range:
            </p>
            <ul>
                <li>
                    <strong>Nodes</strong> appear if their active period overlaps
                    the filter range
                </li>
                <li>
                    <strong>Edges</strong> appear if their relationship period overlaps
                    the range
                </li>
                <li>
                    Use the <strong>Play</strong> button to animate through comedy
                    history
                </li>
                <li>Quick-jump buttons let you focus on specific decades</li>
            </ul>
        </section>

        <section>
            <h2>Seed Notability</h2>
            <p>
                Each comedian has a notability score from 1–5. This is a
                best-effort heuristic reflecting general public recognition and
                historical impact. It is <strong>not</strong> a quality judgment.
            </p>
            <p>
                Notability affects node size on the graph. Higher-notability
                comedians appear larger. This score is a placeholder for a more
                rigorous metric and should be extended with community input.
            </p>
        </section>

        <section>
            <h2>Data Sources</h2>
            <p>
                The seed dataset is hand-curated with evidence links (primarily
                Wikipedia). Auto-ingested data comes from Wikidata queries
                processed through Apple ML deduplication on-device. All
                auto-ingested entries are flagged for verification.
            </p>
        </section>
    </div>
</div>

<style>
    .about-page {
        flex: 1;
        overflow-y: auto;
        display: flex;
        justify-content: center;
        padding: 40px 20px;
    }
    .about-content {
        max-width: 720px;
        width: 100%;
    }
    h1 {
        font-size: 32px;
        font-weight: 900;
        letter-spacing: -0.04em;
        background: linear-gradient(
            135deg,
            var(--color-brand-500),
            var(--color-accent-500)
        );
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
        margin-bottom: 4px;
    }
    .subtitle {
        font-size: 16px;
        color: var(--text-secondary);
        margin-bottom: 36px;
    }
    section {
        margin-bottom: 32px;
    }
    h2 {
        font-size: 20px;
        font-weight: 700;
        letter-spacing: -0.02em;
        color: var(--text-primary);
        margin-bottom: 12px;
        padding-bottom: 8px;
        border-bottom: 1px solid var(--border);
    }
    p {
        font-size: 15px;
        line-height: 1.7;
        color: var(--text-secondary);
        margin-bottom: 12px;
    }
    ul {
        padding-left: 20px;
        margin-bottom: 12px;
    }
    li {
        font-size: 15px;
        line-height: 1.7;
        color: var(--text-secondary);
        margin-bottom: 4px;
    }
    li strong {
        color: var(--text-primary);
    }
    .type-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
        gap: 10px;
        margin-top: 12px;
    }
    .type-card {
        padding: 14px;
        border-radius: 10px;
        border: 1px solid var(--border);
        background: var(--bg-card);
    }
    .type-card .dot {
        display: inline-block;
        width: 10px;
        height: 10px;
        border-radius: 50%;
        background: var(--tc);
        margin-bottom: 6px;
    }
    .type-card h3 {
        font-size: 14px;
        font-weight: 700;
        color: var(--text-primary);
        text-transform: capitalize;
        margin-bottom: 4px;
    }
    .type-card p {
        font-size: 13px;
        margin: 0;
    }
</style>
