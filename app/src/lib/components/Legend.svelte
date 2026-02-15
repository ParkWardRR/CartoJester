<script lang="ts">
    import { EDGE_TYPES, EDGE_COLORS, TAG_CATEGORIES } from "$lib/data/types";
</script>

<div class="legend" role="region" aria-label="Graph legend">
    <div class="legend-group">
        <span class="legend-title">Edges</span>
        {#each EDGE_TYPES as t}
            <span class="legend-item">
                <span class="leg-line" style="background:{EDGE_COLORS[t]}"
                ></span>
                <span class="leg-label">{t}</span>
            </span>
        {/each}
    </div>
    {#each TAG_CATEGORIES as cat}
        <div class="legend-group">
            <span class="legend-title">{cat.emoji} {cat.label}</span>
            {#each Object.entries(cat.tags) as [tag, color]}
                <span class="legend-item">
                    <span class="leg-dot" style="background:{color}"></span>
                    <span class="leg-label">{tag}</span>
                </span>
            {/each}
        </div>
    {/each}
</div>

<style>
    .legend {
        position: absolute;
        bottom: 12px;
        left: 12px;
        display: flex;
        gap: 16px;
        background: var(--glass-bg);
        backdrop-filter: blur(12px);
        border: 1px solid var(--glass-border);
        border-radius: 10px;
        padding: 8px 12px;
        z-index: 10;
        box-shadow: var(--shadow-md);
        max-width: calc(100% - 24px);
        overflow-x: auto;
    }
    @media (max-width: 768px) {
        .legend {
            display: none;
        }
    }
    .legend-group {
        display: flex;
        flex-wrap: wrap;
        gap: 4px 8px;
        align-items: center;
    }
    .legend-title {
        font-size: 10px;
        font-weight: 700;
        text-transform: uppercase;
        letter-spacing: 0.08em;
        color: var(--text-muted);
        margin-right: 4px;
        white-space: nowrap;
    }
    .legend-item {
        display: flex;
        align-items: center;
        gap: 3px;
    }
    .leg-line {
        width: 14px;
        height: 3px;
        border-radius: 1px;
    }
    .leg-dot {
        width: 8px;
        height: 8px;
        border-radius: 50%;
    }
    .leg-label {
        font-size: 10px;
        color: var(--text-secondary);
        text-transform: capitalize;
    }
</style>
