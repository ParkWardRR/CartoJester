<script lang="ts">
    import {
        selectedNode,
        selectedNodeAlliances,
        selectedNodeId,
        selectedEdgeId,
    } from "$lib/stores";
    import {
        EDGE_COLORS,
        TAG_COLORS,
        TAG_TO_CATEGORY,
        TAG_CATEGORIES,
    } from "$lib/data/types";

    function close() {
        selectedNodeId.set(null);
        selectedEdgeId.set(null);
    }
</script>

<aside
    class="details-panel"
    class:open={$selectedNode !== null}
    aria-label="Details"
>
    {#if $selectedNode}
        <div class="panel-header">
            <div>
                <h2>{$selectedNode.name}</h2>
                {#if $selectedNode.aka.length > 0}
                    <p class="aka">aka {$selectedNode.aka.join(", ")}</p>
                {/if}
            </div>
            <button
                class="close-btn"
                on:click={close}
                aria-label="Close details"
            >
                <svg
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    ><line x1="18" y1="6" x2="6" y2="18" /><line
                        x1="6"
                        y1="6"
                        x2="18"
                        y2="18"
                    /></svg
                >
            </button>
        </div>

        <div class="bio-section">
            <div class="timeline">
                {#if $selectedNode.bornYear}
                    <span class="year">{$selectedNode.bornYear}</span>
                    <span class="timeline-bar"></span>
                    <span class="year"
                        >{$selectedNode.diedYear ?? "present"}</span
                    >
                {/if}
            </div>
            {#if $selectedNode.activeStartYear}
                <p class="active-range">
                    Active: {$selectedNode.activeStartYear}–{$selectedNode.activeEndYear ??
                        "present"}
                </p>
            {/if}
        </div>

        <div class="tags-section">
            <div class="tags-row">
                {#each $selectedNode.tags as tag}
                    {@const catId = TAG_TO_CATEGORY[tag]}
                    {@const cat = TAG_CATEGORIES.find((c) => c.id === catId)}
                    <span
                        class="detail-tag"
                        style="background: {TAG_COLORS[tag] ||
                            '#94a3b8'}20; color: {TAG_COLORS[tag] ||
                            '#94a3b8'}; border: 1px solid {TAG_COLORS[tag] ||
                            '#94a3b8'}40"
                        title="{cat?.label ?? 'Tag'}: {tag}"
                        ><span class="tag-cat-icon">{cat?.emoji ?? "🏷️"}</span
                        >{tag}</span
                    >
                {/each}
            </div>
            <span class="notability-badge" title="Notability score">
                {"★".repeat($selectedNode.notability)}{"☆".repeat(
                    5 - $selectedNode.notability,
                )}
            </span>
        </div>

        {#if $selectedNode.source === "auto"}
            <div class="auto-badge">
                <svg
                    width="14"
                    height="14"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    ><circle cx="12" cy="12" r="10" /><line
                        x1="12"
                        y1="8"
                        x2="12"
                        y2="12"
                    /><line x1="12" y1="16" x2="12.01" y2="16" /></svg
                >
                Auto-ingested — needs verification
            </div>
        {/if}

        <section class="alliances-section">
            <h3>Alliances ({$selectedNodeAlliances.length})</h3>
            <div class="alliances-list">
                {#each $selectedNodeAlliances as { edge, other }}
                    <button
                        class="alliance-item"
                        on:click={() => {
                            if (other) selectedNodeId.set(other.id);
                        }}
                    >
                        <div class="alliance-header">
                            <span
                                class="alliance-type"
                                style="background: {EDGE_COLORS[
                                    edge.type
                                ]}20; color: {EDGE_COLORS[edge.type]}"
                                >{edge.type}</span
                            >
                            <span class="alliance-weight"
                                >{"●".repeat(edge.weight)}{"○".repeat(
                                    5 - edge.weight,
                                )}</span
                            >
                        </div>
                        <div class="alliance-name">
                            {other?.name ?? "Unknown"}
                        </div>
                        <div class="alliance-summary">{edge.summary}</div>
                        {#if edge.startYear}
                            <div class="alliance-years">
                                {edge.startYear}–{edge.endYear ?? "present"}
                            </div>
                        {/if}
                    </button>
                {/each}
            </div>
        </section>

        {#if $selectedNode.links.length > 0}
            <section class="links-section">
                <h3>Links</h3>
                {#each $selectedNode.links as link}
                    <a
                        href={link.url}
                        target="_blank"
                        rel="noopener noreferrer"
                        class="ext-link"
                    >
                        <svg
                            width="12"
                            height="12"
                            viewBox="0 0 24 24"
                            fill="none"
                            stroke="currentColor"
                            stroke-width="2"
                            ><path
                                d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"
                            /><polyline points="15 3 21 3 21 9" /><line
                                x1="10"
                                y1="14"
                                x2="21"
                                y2="3"
                            /></svg
                        >
                        {link.label}
                    </a>
                {/each}
            </section>
        {/if}
    {:else}
        <div class="empty-state">
            <div class="empty-icon">🃏</div>
            <p>Click a comedian on the graph to explore their alliances</p>
            <p class="hint">Use ⌘K to search</p>
        </div>
    {/if}
</aside>

<style>
    .details-panel {
        width: 320px;
        min-width: 320px;
        height: 100%;
        overflow-y: auto;
        padding: 20px 16px;
        background: var(--bg-panel);
        border-left: 1px solid var(--border);
        display: flex;
        flex-direction: column;
        gap: 14px;
    }
    @media (max-width: 768px) {
        .details-panel {
            position: fixed;
            bottom: 0;
            left: 0;
            right: 0;
            width: 100%;
            min-width: 100%;
            max-height: 60vh;
            height: auto;
            border-left: none;
            border-top: 1px solid var(--border);
            border-radius: 16px 16px 0 0;
            z-index: 900;
            transform: translateY(100%);
            transition: transform 0.25s ease;
            box-shadow: 0 -4px 20px rgba(0, 0, 0, 0.3);
        }
        .details-panel.open {
            transform: translateY(0);
        }
    }
    .panel-header {
        display: flex;
        justify-content: space-between;
        align-items: flex-start;
        gap: 8px;
    }
    .panel-header h2 {
        font-size: 18px;
        font-weight: 800;
        letter-spacing: -0.03em;
        color: var(--text-primary);
        line-height: 1.2;
    }
    .aka {
        font-size: 12px;
        color: var(--text-muted);
        font-style: italic;
        margin-top: 2px;
    }
    .close-btn {
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
        flex-shrink: 0;
        transition: all 0.15s;
    }
    .close-btn:hover {
        background: var(--bg-elevated);
        color: var(--text-primary);
    }
    .bio-section {
        display: flex;
        flex-direction: column;
        gap: 4px;
    }
    .timeline {
        display: flex;
        align-items: center;
        gap: 8px;
        font-size: 14px;
        font-weight: 600;
        color: var(--text-primary);
    }
    .timeline-bar {
        flex: 1;
        height: 2px;
        background: linear-gradient(
            90deg,
            var(--color-brand-400),
            var(--color-accent-400)
        );
        border-radius: 1px;
    }
    .active-range {
        font-size: 12px;
        color: var(--text-secondary);
    }
    .tags-section {
        display: flex;
        justify-content: space-between;
        align-items: flex-start;
        gap: 8px;
    }
    .tags-row {
        display: flex;
        flex-wrap: wrap;
        gap: 4px;
        align-items: center;
        flex: 1;
    }
    .detail-tag {
        font-size: 10px;
        font-weight: 600;
        padding: 3px 7px;
        border-radius: 5px;
        text-transform: uppercase;
        letter-spacing: 0.5px;
        display: inline-flex;
        align-items: center;
        gap: 2px;
    }
    .tag-cat-icon {
        font-size: 9px;
        line-height: 1;
    }
    .notability-badge {
        font-size: 12px;
        color: var(--color-brand-400);
        margin-left: auto;
    }
    .auto-badge {
        display: flex;
        align-items: center;
        gap: 6px;
        font-size: 12px;
        color: #f59e0b;
        background: rgba(245, 158, 11, 0.1);
        border: 1px solid rgba(245, 158, 11, 0.2);
        border-radius: 8px;
        padding: 8px 10px;
    }
    .alliances-section,
    .links-section {
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
    .alliances-list {
        display: flex;
        flex-direction: column;
        gap: 6px;
    }
    .alliance-item {
        padding: 10px 12px;
        border-radius: 8px;
        border: 1px solid var(--border);
        background: var(--bg-card);
        text-align: left;
        cursor: pointer;
        transition: all 0.15s;
        display: flex;
        flex-direction: column;
        gap: 4px;
    }
    .alliance-item:hover {
        background: var(--bg-elevated);
        border-color: var(--border-active);
        box-shadow: var(--shadow-sm);
    }
    .alliance-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    .alliance-type {
        font-size: 10px;
        font-weight: 600;
        text-transform: uppercase;
        padding: 2px 6px;
        border-radius: 4px;
        letter-spacing: 0.5px;
    }
    .alliance-weight {
        font-size: 10px;
        color: var(--text-muted);
        letter-spacing: 1px;
    }
    .alliance-name {
        font-size: 14px;
        font-weight: 600;
        color: var(--text-primary);
    }
    .alliance-summary {
        font-size: 12px;
        color: var(--text-secondary);
        line-height: 1.4;
    }
    .alliance-years {
        font-size: 11px;
        color: var(--text-muted);
        font-weight: 500;
    }
    .ext-link {
        display: flex;
        align-items: center;
        gap: 6px;
        font-size: 13px;
        color: var(--color-brand-500);
        text-decoration: none;
        padding: 6px 0;
        transition: color 0.15s;
    }
    .ext-link:hover {
        color: var(--color-brand-400);
        text-decoration: underline;
    }
    .empty-state {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100%;
        text-align: center;
        gap: 8px;
        color: var(--text-muted);
        font-size: 14px;
        padding: 24px;
    }
    .empty-icon {
        font-size: 48px;
        margin-bottom: 8px;
    }
    .hint {
        font-size: 12px;
        opacity: 0.6;
    }
</style>
