<script lang="ts">
    import { COMEDY_GROUPS, type ComedyGroup } from "$lib/data/groups";
    import { nodes, selectedNode, enabledTags, yearRange } from "$lib/stores";
    import type { ComedianNode } from "$lib/data/types";
    import {
        TAG_COLORS,
        TAG_TO_CATEGORY,
        TAG_CATEGORIES,
    } from "$lib/data/types";
    import { base } from "$app/paths";
    import { onMount } from "svelte";

    let allNodes: ComedianNode[] = [];
    let activeGroup: ComedyGroup | null = null;
    let groupMembers: ComedianNode[] = [];
    let mounted = false;

    nodes.subscribe((n) => {
        allNodes = n;
    });

    function selectGroup(g: ComedyGroup) {
        if (activeGroup?.id === g.id) {
            activeGroup = null;
            groupMembers = [];
        } else {
            activeGroup = g;
            groupMembers = allNodes.filter((n) => g.memberIds.includes(n.id));
            groupMembers.sort(
                (a, b) =>
                    (a.activeStartYear ?? a.bornYear ?? 0) -
                    (b.activeStartYear ?? b.bornYear ?? 0),
            );
        }
    }

    function getNotabilityStars(n: number): string {
        return "★".repeat(n) + "☆".repeat(5 - n);
    }

    function getTagColor(tag: string): string {
        return TAG_COLORS[tag] || "#94a3b8";
    }

    onMount(() => {
        mounted = true;
    });
</script>

<svelte:head>
    <title>Groups — CartoJester</title>
    <meta
        name="description"
        content="Browse comedy movements and vibes — from Silent Era Pioneers to the Podcast Comedy Sphere."
    />
</svelte:head>

<div class="groups-page" class:mounted>
    <header class="groups-header">
        <h1>Comedy Vibes</h1>
        <p class="subtitle">
            Surf through comedy movements, scenes, and crews. Each group
            represents a distinct vibe in comedy history.
        </p>
    </header>

    <div class="groups-layout">
        <div class="groups-grid">
            {#each COMEDY_GROUPS as group, i}
                <button
                    class="group-card"
                    class:active={activeGroup?.id === group.id}
                    style="--group-color: {group.color}; --delay: {i * 60}ms"
                    on:click={() => selectGroup(group)}
                >
                    <div class="card-glow"></div>
                    <div class="card-content">
                        <div class="card-top">
                            <span class="card-icon">{group.icon}</span>
                            <span class="card-era">{group.era}</span>
                        </div>
                        <h2 class="card-name">{group.name}</h2>
                        <p class="card-desc">{group.description}</p>
                        <div class="card-chips">
                            {#each group.characteristics.slice(0, 3) as c}
                                <span class="chip">{c}</span>
                            {/each}
                        </div>
                        <div class="card-footer">
                            <span class="member-count">
                                {group.memberIds.length} comedians
                            </span>
                            <span class="expand-hint">
                                {activeGroup?.id === group.id
                                    ? "▲ Close"
                                    : "▼ Explore"}
                            </span>
                        </div>
                    </div>
                </button>
            {/each}
        </div>

        {#if activeGroup}
            <div
                class="group-detail"
                style="--group-color: {activeGroup.color}"
            >
                <div class="detail-header">
                    <span class="detail-icon">{activeGroup.icon}</span>
                    <div>
                        <h2>{activeGroup.name}</h2>
                        <span class="detail-era">{activeGroup.era}</span>
                    </div>
                </div>

                <div class="detail-traits">
                    {#each activeGroup.characteristics as trait}
                        <span class="trait-badge">{trait}</span>
                    {/each}
                </div>

                <p class="detail-desc">{activeGroup.description}</p>

                <div class="members-grid">
                    {#each groupMembers as member, i}
                        <a
                            href="{base}/"
                            class="member-card"
                            style="--i: {i}"
                            title="View {member.name} in explorer"
                        >
                            <div
                                class="member-avatar"
                                style="background: {getTagColor(
                                    member.tags[0] || 'standup',
                                )}"
                            >
                                {member.name.charAt(0)}
                            </div>
                            <div class="member-info">
                                <span class="member-name">{member.name}</span>
                                {#if member.aka.length > 0}
                                    <span class="member-aka"
                                        >aka {member.aka[0]}</span
                                    >
                                {/if}
                                <div class="member-tags">
                                    {#each member.tags.slice(0, 3) as tag}
                                        {@const catId = TAG_TO_CATEGORY[tag]}
                                        {@const cat = TAG_CATEGORIES.find(
                                            (c) => c.id === catId,
                                        )}
                                        <span
                                            class="mini-tag"
                                            style="background: {getTagColor(
                                                tag,
                                            )}20; color: {getTagColor(tag)}"
                                            title="{cat?.label ?? 'Tag'}: {tag}"
                                            >{cat?.emoji ?? ""} {tag}</span
                                        >
                                    {/each}
                                </div>
                                <div class="member-meta">
                                    <span class="notability"
                                        >{getNotabilityStars(
                                            member.notability,
                                        )}</span
                                    >
                                    <span class="years">
                                        {member.activeStartYear ??
                                            member.bornYear ??
                                            "?"} – {member.activeEndYear ??
                                            "now"}
                                    </span>
                                </div>
                            </div>
                        </a>
                    {/each}
                </div>

                {#if groupMembers.length === 0}
                    <p class="no-members">
                        No members found in the dataset for this group yet.
                    </p>
                {/if}
            </div>
        {/if}
    </div>
</div>

<style>
    .groups-page {
        overflow-y: auto;
        padding: 32px;
        max-height: 100%;
        opacity: 0;
        transform: translateY(12px);
        transition:
            opacity 0.4s ease,
            transform 0.4s ease;
    }
    .groups-page.mounted {
        opacity: 1;
        transform: translateY(0);
    }

    .groups-header {
        text-align: center;
        margin-bottom: 40px;
    }
    .groups-header h1 {
        font-size: 36px;
        font-weight: 900;
        letter-spacing: -0.04em;
        background: linear-gradient(
            135deg,
            var(--color-brand-400),
            var(--color-accent-400),
            var(--color-brand-500)
        );
        -webkit-background-clip: text;
        background-clip: text;
        -webkit-text-fill-color: transparent;
        background-size: 200% 200%;
        animation: gradientShift 6s ease infinite;
    }
    @keyframes gradientShift {
        0%,
        100% {
            background-position: 0% 50%;
        }
        50% {
            background-position: 100% 50%;
        }
    }
    .subtitle {
        margin-top: 8px;
        color: var(--text-secondary);
        font-size: 15px;
        max-width: 560px;
        margin-inline: auto;
        line-height: 1.5;
    }

    .groups-layout {
        max-width: 1200px;
        margin: 0 auto;
    }

    .groups-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 16px;
        margin-bottom: 32px;
    }

    .group-card {
        position: relative;
        text-align: left;
        border: 1px solid var(--border);
        border-radius: 16px;
        padding: 0;
        background: var(--bg-card);
        cursor: pointer;
        overflow: hidden;
        transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        font-family: inherit;
        color: inherit;
        opacity: 0;
        transform: translateY(16px);
        animation: cardEntry 0.5s ease forwards;
        animation-delay: var(--delay);
    }
    @keyframes cardEntry {
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .group-card:hover {
        border-color: var(--group-color);
        transform: translateY(-4px);
        box-shadow:
            0 12px 32px -8px color-mix(in srgb, var(--group-color) 25%, transparent),
            0 4px 8px -2px rgba(0, 0, 0, 0.1);
    }
    .group-card.active {
        border-color: var(--group-color);
        box-shadow:
            0 0 0 2px color-mix(in srgb, var(--group-color) 30%, transparent),
            0 8px 24px -4px color-mix(in srgb, var(--group-color) 20%, transparent);
    }

    .card-glow {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        height: 4px;
        background: linear-gradient(
            90deg,
            transparent,
            var(--group-color),
            transparent
        );
        opacity: 0;
        transition: opacity 0.3s ease;
    }
    .group-card:hover .card-glow,
    .group-card.active .card-glow {
        opacity: 1;
    }

    .card-content {
        padding: 20px;
    }
    .card-top {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 12px;
    }
    .card-icon {
        font-size: 28px;
    }
    .card-era {
        font-size: 11px;
        font-weight: 600;
        color: var(--group-color);
        background: color-mix(in srgb, var(--group-color) 12%, transparent);
        padding: 3px 10px;
        border-radius: 12px;
        letter-spacing: 0.03em;
    }
    .card-name {
        font-size: 17px;
        font-weight: 700;
        margin-bottom: 8px;
        letter-spacing: -0.02em;
    }
    .card-desc {
        font-size: 12.5px;
        color: var(--text-secondary);
        line-height: 1.5;
        display: -webkit-box;
        -webkit-line-clamp: 3;
        line-clamp: 3;
        -webkit-box-orient: vertical;
        overflow: hidden;
        margin-bottom: 12px;
    }
    .card-chips {
        display: flex;
        gap: 6px;
        flex-wrap: wrap;
        margin-bottom: 14px;
    }
    .chip {
        font-size: 10px;
        padding: 2px 8px;
        border-radius: 6px;
        background: var(--bg-elevated);
        color: var(--text-secondary);
        font-weight: 500;
    }
    .card-footer {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding-top: 12px;
        border-top: 1px solid var(--border);
    }
    .member-count {
        font-size: 12px;
        font-weight: 600;
        color: var(--group-color);
    }
    .expand-hint {
        font-size: 11px;
        color: var(--text-muted);
        font-weight: 500;
    }

    /* Detail panel */
    .group-detail {
        background: var(--bg-card);
        border: 1px solid var(--border);
        border-radius: 20px;
        padding: 32px;
        margin-bottom: 32px;
        border-top: 3px solid var(--group-color);
        animation: slideUp 0.35s ease;
    }
    @keyframes slideUp {
        from {
            opacity: 0;
            transform: translateY(20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }
    .detail-header {
        display: flex;
        align-items: center;
        gap: 16px;
        margin-bottom: 20px;
    }
    .detail-icon {
        font-size: 48px;
    }
    .detail-header h2 {
        font-size: 26px;
        font-weight: 800;
        letter-spacing: -0.03em;
    }
    .detail-era {
        font-size: 13px;
        color: var(--group-color);
        font-weight: 600;
    }
    .detail-traits {
        display: flex;
        gap: 8px;
        flex-wrap: wrap;
        margin-bottom: 16px;
    }
    .trait-badge {
        font-size: 12px;
        padding: 4px 12px;
        border-radius: 20px;
        background: color-mix(in srgb, var(--group-color) 10%, transparent);
        color: var(--group-color);
        font-weight: 600;
        border: 1px solid
            color-mix(in srgb, var(--group-color) 20%, transparent);
    }
    .detail-desc {
        font-size: 14px;
        color: var(--text-secondary);
        line-height: 1.7;
        margin-bottom: 24px;
        max-width: 720px;
    }

    /* Members grid */
    .members-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
        gap: 12px;
    }
    .member-card {
        display: flex;
        gap: 12px;
        padding: 12px;
        border-radius: 12px;
        border: 1px solid var(--border);
        background: var(--bg-panel);
        text-decoration: none;
        color: inherit;
        transition: all 0.2s ease;
        opacity: 0;
        animation: memberFade 0.3s ease forwards;
        animation-delay: calc(var(--i) * 30ms);
    }
    @keyframes memberFade {
        to {
            opacity: 1;
        }
    }
    .member-card:hover {
        background: var(--bg-elevated);
        border-color: var(--group-color);
        transform: translateX(4px);
    }
    .member-avatar {
        width: 40px;
        height: 40px;
        border-radius: 10px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: white;
        font-weight: 800;
        font-size: 16px;
        flex-shrink: 0;
    }
    .member-info {
        flex: 1;
        min-width: 0;
    }
    .member-name {
        font-size: 13px;
        font-weight: 700;
        display: block;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
    }
    .member-aka {
        font-size: 11px;
        color: var(--text-muted);
        font-style: italic;
        display: block;
    }
    .member-tags {
        display: flex;
        gap: 4px;
        flex-wrap: wrap;
        margin-top: 4px;
    }
    .mini-tag {
        font-size: 9px;
        padding: 1px 6px;
        border-radius: 4px;
        font-weight: 600;
        text-transform: uppercase;
        letter-spacing: 0.04em;
    }
    .member-meta {
        display: flex;
        justify-content: space-between;
        margin-top: 4px;
    }
    .notability {
        font-size: 10px;
        color: #f59e0b;
    }
    .years {
        font-size: 10px;
        color: var(--text-muted);
        font-family: var(--font-mono);
    }

    .no-members {
        text-align: center;
        padding: 32px;
        color: var(--text-muted);
        font-style: italic;
    }

    @media (max-width: 640px) {
        .groups-page {
            padding: 16px;
        }
        .groups-grid {
            grid-template-columns: 1fr;
        }
        .members-grid {
            grid-template-columns: 1fr;
        }
        .groups-header h1 {
            font-size: 28px;
        }
        .group-detail {
            padding: 20px;
        }
    }
</style>
