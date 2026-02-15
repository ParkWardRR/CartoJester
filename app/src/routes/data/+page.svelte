<script lang="ts">
    import { nodes, edges } from "$lib/stores";
    import {
        EDGE_COLORS,
        TAG_COLORS,
        TAG_TO_CATEGORY,
        TAG_CATEGORIES,
    } from "$lib/data/types";
    import seedData from "$lib/data/seed.json";

    let searchTerm = "";
    let activeTab: "nodes" | "edges" = "nodes";

    $: filteredNodeList = $nodes.filter(
        (n) =>
            !searchTerm ||
            n.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
            n.tags.some((t) => t.includes(searchTerm.toLowerCase())),
    );

    $: filteredEdgeList = $edges.filter(
        (e) =>
            !searchTerm ||
            e.summary.toLowerCase().includes(searchTerm.toLowerCase()) ||
            e.type.includes(searchTerm.toLowerCase()),
    );

    function downloadJSON() {
        const blob = new Blob([JSON.stringify(seedData, null, 2)], {
            type: "application/json",
        });
        const url = URL.createObjectURL(blob);
        const a = document.createElement("a");
        a.href = url;
        a.download = "seed.json";
        a.click();
        URL.revokeObjectURL(url);
    }
</script>

<svelte:head>
    <title>Data — CartoJester</title>
    <meta
        name="description"
        content="Explore and download the CartoJester comedian dataset."
    />
</svelte:head>

<div class="data-page">
    <div class="data-content">
        <h1>Dataset</h1>
        <p class="subtitle">
            {$nodes.length} comedians · {$edges.length} alliances
        </p>

        <div class="toolbar">
            <div class="tabs">
                <button
                    class:active={activeTab === "nodes"}
                    on:click={() => (activeTab = "nodes")}>Comedians</button
                >
                <button
                    class:active={activeTab === "edges"}
                    on:click={() => (activeTab = "edges")}>Alliances</button
                >
            </div>
            <input
                type="text"
                bind:value={searchTerm}
                placeholder="Filter..."
                class="filter-input"
            />
            <button class="dl-btn" on:click={downloadJSON}>↓ JSON</button>
        </div>

        {#if activeTab === "nodes"}
            <div class="table-wrap">
                <table>
                    <thead
                        ><tr
                            ><th>Name</th><th>Years</th><th>Tags</th><th>★</th
                            ></tr
                        ></thead
                    >
                    <tbody>
                        {#each filteredNodeList as n}
                            <tr>
                                <td class="name-cell"
                                    ><strong>{n.name}</strong
                                    >{#if n.aka.length > 0}<br /><span
                                            class="aka">{n.aka.join(", ")}</span
                                        >{/if}</td
                                >
                                <td class="year-cell"
                                    >{n.bornYear ?? "?"}–{n.diedYear ??
                                        "now"}</td
                                >
                                <td class="tags-cell"
                                    >{#each n.tags as t}{@const catId =
                                            TAG_TO_CATEGORY[t]}{@const cat =
                                            TAG_CATEGORIES.find(
                                                (c) => c.id === catId,
                                            )}<span
                                            class="minitag"
                                            style="color:{TAG_COLORS[t] ||
                                                '#94a3b8'}"
                                            >{cat?.emoji ?? ""} {t}</span
                                        >{/each}</td
                                >
                                <td>{n.notability}</td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
        {:else}
            <div class="table-wrap">
                <table>
                    <thead
                        ><tr
                            ><th>Type</th><th>From → To</th><th>Years</th><th
                                >W</th
                            ><th>Summary</th></tr
                        ></thead
                    >
                    <tbody>
                        {#each filteredEdgeList as e}
                            <tr>
                                <td
                                    ><span
                                        class="etype"
                                        style="color:{EDGE_COLORS[e.type]}"
                                        >{e.type}</span
                                    ></td
                                >
                                <td class="name-cell"
                                    >{e.sourceId} → {e.targetId}</td
                                >
                                <td class="year-cell"
                                    >{e.startYear ?? "?"}–{e.endYear ??
                                        "now"}</td
                                >
                                <td>{e.weight}</td>
                                <td class="sum-cell">{e.summary}</td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
        {/if}

        <section class="contribute">
            <h2>How to Contribute</h2>
            <ol>
                <li>
                    Fork the <a
                        href="https://github.com/ParkWardRR/CartoJester"
                        target="_blank">repository</a
                    >
                </li>
                <li>Edit <code>app/src/lib/data/seed.json</code></li>
                <li>Add nodes and edges following the schema</li>
                <li>Run <code>npm run dev</code> to verify your changes</li>
                <li>Submit a pull request</li>
            </ol>
        </section>
    </div>
</div>

<style>
    .data-page {
        flex: 1;
        overflow-y: auto;
        display: flex;
        justify-content: center;
        padding: 40px 20px;
    }
    .data-content {
        max-width: 960px;
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
        margin-bottom: 24px;
    }
    .toolbar {
        display: flex;
        gap: 8px;
        align-items: center;
        margin-bottom: 16px;
        flex-wrap: wrap;
    }
    .tabs {
        display: flex;
        gap: 2px;
        background: var(--bg-elevated);
        border-radius: 8px;
        padding: 2px;
    }
    .tabs button {
        padding: 6px 14px;
        border: none;
        border-radius: 6px;
        background: none;
        color: var(--text-secondary);
        font-size: 13px;
        font-weight: 500;
        cursor: pointer;
        transition: all 0.15s;
    }
    .tabs button.active {
        background: var(--bg-card);
        color: var(--text-primary);
        font-weight: 600;
        box-shadow: var(--shadow-sm);
    }
    .filter-input {
        flex: 1;
        min-width: 120px;
        padding: 6px 12px;
        border: 1px solid var(--border);
        border-radius: 6px;
        background: var(--bg-card);
        color: var(--text-primary);
        font-size: 13px;
        outline: none;
        font-family: var(--font-sans);
    }
    .filter-input:focus {
        border-color: var(--border-active);
    }
    .dl-btn {
        padding: 6px 14px;
        border: 1px solid var(--color-brand-500);
        border-radius: 6px;
        background: var(--color-brand-500);
        color: white;
        font-size: 13px;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.15s;
    }
    .dl-btn:hover {
        background: var(--color-brand-600);
    }
    .table-wrap {
        overflow-x: auto;
        border: 1px solid var(--border);
        border-radius: 10px;
    }
    table {
        width: 100%;
        border-collapse: collapse;
        font-size: 13px;
    }
    th {
        text-align: left;
        padding: 10px 12px;
        background: var(--bg-elevated);
        color: var(--text-muted);
        font-size: 11px;
        font-weight: 600;
        text-transform: uppercase;
        letter-spacing: 0.06em;
        border-bottom: 1px solid var(--border);
    }
    td {
        padding: 8px 12px;
        border-bottom: 1px solid var(--border);
        color: var(--text-secondary);
        vertical-align: top;
    }
    .name-cell strong {
        color: var(--text-primary);
    }
    .aka {
        font-size: 11px;
        color: var(--text-muted);
        font-style: italic;
    }
    .year-cell {
        font-variant-numeric: tabular-nums;
        white-space: nowrap;
    }
    .minitag {
        font-size: 11px;
        font-weight: 500;
        margin-right: 4px;
        text-transform: capitalize;
    }
    .etype {
        font-weight: 600;
        text-transform: capitalize;
        font-size: 12px;
    }
    .sum-cell {
        max-width: 300px;
    }
    .contribute {
        margin-top: 32px;
        padding-top: 24px;
        border-top: 1px solid var(--border);
    }
    h2 {
        font-size: 20px;
        font-weight: 700;
        color: var(--text-primary);
        margin-bottom: 12px;
    }
    ol {
        padding-left: 20px;
    }
    li {
        font-size: 14px;
        line-height: 1.8;
        color: var(--text-secondary);
    }
    code {
        font-family: var(--font-mono);
        font-size: 12px;
        background: var(--bg-elevated);
        padding: 2px 6px;
        border-radius: 4px;
    }
    a {
        color: var(--color-brand-500);
    }

    @media (max-width: 768px) {
        .data-page {
            padding: 20px 12px;
        }
        h1 {
            font-size: 24px;
        }
        .subtitle {
            font-size: 14px;
        }
        .toolbar {
            flex-direction: column;
            align-items: stretch;
        }
        .tabs {
            justify-content: center;
        }
        .filter-input {
            min-width: 0;
        }
        table {
            font-size: 12px;
        }
        th,
        td {
            padding: 6px 8px;
        }
        .sum-cell {
            max-width: 160px;
        }
    }
</style>
