<script lang="ts">
    import {
        EDGE_TYPES,
        EDGE_COLORS,
        TAG_COLORS,
        ERA_PRESETS,
    } from "$lib/data/types";
    import type { EdgeType } from "$lib/data/types";
    import {
        enabledEdgeTypes,
        enabledTags,
        minWeight,
        clusterMode,
        yearRange,
        nodeStats,
        showOnlyVerified,
    } from "$lib/stores";

    const allTags = Object.keys(TAG_COLORS);

    function toggleEdgeType(t: EdgeType) {
        enabledEdgeTypes.update((s) => {
            const next = new Set(s);
            if (next.has(t)) next.delete(t);
            else next.add(t);
            return next;
        });
    }

    function toggleTag(tag: string) {
        enabledTags.update((s) => {
            const next = new Set(s);
            if (next.has(tag)) next.delete(tag);
            else next.add(tag);
            return next;
        });
    }

    function setEra(start: number, end: number) {
        yearRange.set([start, end]);
    }

    function resetFilters() {
        enabledEdgeTypes.set(new Set(EDGE_TYPES));
        enabledTags.set(new Set());
        minWeight.set(1);
        yearRange.set([1895, 2025]);
        clusterMode.set("none");
        showOnlyVerified.set(false);
    }
</script>

<aside class="filter-panel" aria-label="Filters">
    <div class="panel-header">
        <h2>Filters</h2>
        <button
            class="reset-btn"
            on:click={resetFilters}
            title="Reset all filters"
        >
            <svg
                width="14"
                height="14"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                ><path d="M1 4v6h6M23 20v-6h-6" /><path
                    d="M20.49 9A9 9 0 0 0 5.64 5.64L1 10m22 4l-4.64 4.36A9 9 0 0 1 3.51 15"
                /></svg
            >
        </button>
    </div>

    <div class="stats-bar">
        <span class="stat"
            ><strong>{$nodeStats.nodeCount}</strong> comedians</span
        >
        <span class="stat-dot">·</span>
        <span class="stat"
            ><strong>{$nodeStats.edgeCount}</strong> alliances</span
        >
    </div>

    <section>
        <h3>Alliance Types</h3>
        <div class="chip-grid">
            {#each EDGE_TYPES as t}
                <button
                    class="chip"
                    class:active={$enabledEdgeTypes.has(t)}
                    style="--chip-color: {EDGE_COLORS[t]}"
                    on:click={() => toggleEdgeType(t)}
                >
                    <span class="chip-dot" style="background: {EDGE_COLORS[t]}"
                    ></span>
                    {t}
                </button>
            {/each}
        </div>
    </section>

    <section>
        <h3>Era Presets</h3>
        <div class="chip-grid">
            {#each ERA_PRESETS as era}
                <button
                    class="era-chip"
                    on:click={() => setEra(era.start, era.end)}
                >
                    {era.label}
                </button>
            {/each}
        </div>
    </section>

    <section>
        <h3>Tags</h3>
        <div class="chip-grid">
            {#each allTags as tag}
                <button
                    class="chip"
                    class:active={$enabledTags.has(tag)}
                    style="--chip-color: {TAG_COLORS[tag]}"
                    on:click={() => toggleTag(tag)}
                >
                    <span class="chip-dot" style="background: {TAG_COLORS[tag]}"
                    ></span>
                    {tag}
                </button>
            {/each}
        </div>
    </section>

    <section>
        <h3>Min Weight</h3>
        <div class="slider-row">
            <input
                type="range"
                min="1"
                max="5"
                step="1"
                bind:value={$minWeight}
            />
            <span class="slider-val">{$minWeight}</span>
        </div>
    </section>

    <section>
        <label class="toggle-row">
            <input type="checkbox" bind:checked={$showOnlyVerified} />
            <span>Verified only</span>
        </label>
    </section>
</aside>

<style>
    .filter-panel {
        width: 260px;
        min-width: 260px;
        height: 100%;
        overflow-y: auto;
        padding: 20px 16px;
        background: var(--bg-panel);
        border-right: 1px solid var(--border);
        display: flex;
        flex-direction: column;
        gap: 16px;
    }
    .panel-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .panel-header h2 {
        font-size: 15px;
        font-weight: 700;
        letter-spacing: -0.02em;
        color: var(--text-primary);
    }
    .reset-btn {
        width: 28px;
        height: 28px;
        border-radius: 6px;
        border: 1px solid var(--border);
        background: var(--bg-card);
        color: var(--text-secondary);
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all 0.15s;
    }
    .reset-btn:hover {
        background: var(--bg-elevated);
        color: var(--text-primary);
    }
    .stats-bar {
        display: flex;
        gap: 6px;
        align-items: center;
        font-size: 12px;
        color: var(--text-secondary);
        padding: 8px 10px;
        background: var(--bg-elevated);
        border-radius: 8px;
    }
    .stat strong {
        color: var(--text-primary);
        font-weight: 700;
    }
    .stat-dot {
        opacity: 0.4;
    }
    section {
        display: flex;
        flex-direction: column;
        gap: 8px;
    }
    h3 {
        font-size: 11px;
        font-weight: 600;
        text-transform: uppercase;
        letter-spacing: 0.08em;
        color: var(--text-muted);
    }
    .chip-grid {
        display: flex;
        flex-wrap: wrap;
        gap: 4px;
    }
    .chip {
        font-size: 11px;
        font-weight: 500;
        padding: 4px 8px;
        border-radius: 6px;
        border: 1px solid var(--border);
        background: var(--bg-card);
        color: var(--text-secondary);
        cursor: pointer;
        display: flex;
        align-items: center;
        gap: 4px;
        transition: all 0.15s;
        text-transform: capitalize;
    }
    .chip:hover {
        border-color: var(--chip-color);
        color: var(--text-primary);
    }
    .chip.active {
        background: color-mix(in srgb, var(--chip-color) 12%, transparent);
        border-color: var(--chip-color);
        color: var(--chip-color);
    }
    .chip-dot {
        width: 8px;
        height: 8px;
        border-radius: 50%;
        flex-shrink: 0;
    }
    .era-chip {
        font-size: 11px;
        font-weight: 500;
        padding: 4px 10px;
        border-radius: 6px;
        border: 1px solid var(--border);
        background: var(--bg-card);
        color: var(--text-secondary);
        cursor: pointer;
        transition: all 0.15s;
    }
    .era-chip:hover {
        background: var(--bg-elevated);
        color: var(--text-primary);
        border-color: var(--border-active);
    }
    .slider-row {
        display: flex;
        align-items: center;
        gap: 10px;
    }
    .slider-row input[type="range"] {
        flex: 1;
        accent-color: var(--color-brand-500);
    }
    .slider-val {
        font-size: 13px;
        font-weight: 600;
        color: var(--text-primary);
        min-width: 20px;
        text-align: center;
    }
    .toggle-row {
        display: flex;
        align-items: center;
        gap: 8px;
        font-size: 13px;
        cursor: pointer;
        color: var(--text-secondary);
    }
    .toggle-row input {
        accent-color: var(--color-brand-500);
    }
</style>
