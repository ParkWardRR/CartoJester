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
    let showPanel = false;

    function togglePanel() {
        showPanel = !showPanel;
    }

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

<button
    class="mobile-toggle"
    on:click={togglePanel}
    aria-label="Toggle filters"
>
    <svg
        width="20"
        height="20"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
    >
        {#if showPanel}
            <path d="M18 6L6 18M6 6l12 12" />
        {:else}
            <path d="M4 6h16M4 12h16M4 18h16" />
        {/if}
    </svg>
    <span class="toggle-badge">{$nodeStats.nodeCount}</span>
</button>

{#if showPanel}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div
        class="mobile-backdrop"
        on:click={togglePanel}
        role="presentation"
    ></div>
{/if}

<aside class="filter-panel" class:open={showPanel} aria-label="Filters">
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
    .mobile-toggle {
        display: none;
        position: fixed;
        bottom: 20px;
        left: 20px;
        z-index: 1000;
        width: 48px;
        height: 48px;
        border-radius: 50%;
        border: 1px solid var(--border);
        background: var(--bg-panel);
        color: var(--text-primary);
        cursor: pointer;
        align-items: center;
        justify-content: center;
        box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
        transition: all 0.2s;
    }
    .mobile-toggle:hover {
        transform: scale(1.05);
        box-shadow: 0 6px 24px rgba(0, 0, 0, 0.4);
    }
    .toggle-badge {
        display: none;
    }
    .mobile-backdrop {
        display: none;
    }

    @media (max-width: 768px) {
        .mobile-toggle {
            display: flex;
        }
        .toggle-badge {
            display: block;
            position: absolute;
            top: -4px;
            right: -4px;
            background: var(--color-brand-500, #8b5cf6);
            color: #fff;
            font-size: 9px;
            font-weight: 700;
            padding: 2px 5px;
            border-radius: 10px;
            min-width: 18px;
            text-align: center;
        }
        .mobile-backdrop {
            display: block;
            position: fixed;
            inset: 0;
            background: rgba(0, 0, 0, 0.5);
            z-index: 998;
        }
    }

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

    @media (max-width: 768px) {
        .filter-panel {
            position: fixed;
            top: 56px;
            left: 0;
            bottom: 0;
            z-index: 999;
            transform: translateX(-100%);
            transition: transform 0.25s ease;
            box-shadow: 4px 0 20px rgba(0, 0, 0, 0.3);
        }
        .filter-panel.open {
            transform: translateX(0);
        }
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
