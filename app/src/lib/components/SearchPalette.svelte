<script lang="ts">
    import {
        searchOpen,
        searchQuery,
        searchResults,
        selectedNodeId,
    } from "$lib/stores";
    import { onMount } from "svelte";

    let inputEl: HTMLInputElement;
    let selectedIndex = 0;

    function close() {
        searchOpen.set(false);
        searchQuery.set("");
        selectedIndex = 0;
    }
    function select(id: string) {
        selectedNodeId.set(id);
        close();
    }

    function handleKeydown(e: KeyboardEvent) {
        if (e.key === "Escape") close();
        else if (e.key === "ArrowDown") {
            e.preventDefault();
            selectedIndex = Math.min(
                selectedIndex + 1,
                $searchResults.length - 1,
            );
        } else if (e.key === "ArrowUp") {
            e.preventDefault();
            selectedIndex = Math.max(selectedIndex - 1, 0);
        } else if (e.key === "Enter" && $searchResults.length > 0)
            select($searchResults[selectedIndex].id);
    }

    $: if ($searchOpen && inputEl) setTimeout(() => inputEl?.focus(), 50);
    $: selectedIndex = 0;

    onMount(() => {
        function gk(e: KeyboardEvent) {
            if ((e.metaKey || e.ctrlKey) && e.key === "k") {
                e.preventDefault();
                searchOpen.update((v) => !v);
            }
        }
        window.addEventListener("keydown", gk);
        return () => window.removeEventListener("keydown", gk);
    });
</script>

{#if $searchOpen}
    <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
    <div
        class="search-overlay"
        on:click={close}
        on:keydown={handleKeydown}
        role="dialog"
        aria-label="Search"
        tabindex="-1"
    >
        <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <div
            class="search-modal"
            on:click|stopPropagation
            on:keydown|stopPropagation
            role="search"
        >
            <div class="sinput-row">
                <svg
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    ><circle cx="11" cy="11" r="8" /><line
                        x1="21"
                        y1="21"
                        x2="16.65"
                        y2="16.65"
                    /></svg
                >
                <input
                    bind:this={inputEl}
                    bind:value={$searchQuery}
                    on:keydown={handleKeydown}
                    placeholder="Search comedians..."
                    class="sinput"
                    type="text"
                />
                <kbd class="kbd">esc</kbd>
            </div>
            {#if $searchResults.length > 0}
                <div class="sresults">
                    {#each $searchResults as r, i}
                        <button
                            class="sr"
                            class:sel={i === selectedIndex}
                            on:click={() => select(r.id)}
                        >
                            <div class="srn">{r.name}</div>
                            <div class="srm">
                                {r.bornYear
                                    ? `${r.bornYear}–${r.diedYear ?? "present"}`
                                    : ""} · {r.tags.join(", ")}
                            </div>
                        </button>
                    {/each}
                </div>
            {:else if $searchQuery.length > 0}
                <div class="nores">No comedians found</div>
            {/if}
        </div>
    </div>
{/if}

<style>
    .search-overlay {
        position: fixed;
        inset: 0;
        z-index: 200;
        background: rgba(0, 0, 0, 0.4);
        backdrop-filter: blur(4px);
        display: flex;
        align-items: flex-start;
        justify-content: center;
        padding-top: 15vh;
    }
    .search-modal {
        width: 100%;
        max-width: 520px;
        background: var(--bg-card);
        border: 1px solid var(--border);
        border-radius: 14px;
        box-shadow: var(--shadow-xl);
        overflow: hidden;
    }
    .sinput-row {
        display: flex;
        align-items: center;
        gap: 10px;
        padding: 14px 16px;
        border-bottom: 1px solid var(--border);
        color: var(--text-secondary);
    }
    .sinput {
        flex: 1;
        font-size: 16px;
        font-weight: 500;
        background: none;
        border: none;
        outline: none;
        color: var(--text-primary);
        font-family: var(--font-sans);
    }
    .sinput::placeholder {
        color: var(--text-muted);
    }
    .kbd {
        font-size: 10px;
        padding: 2px 6px;
        border: 1px solid var(--border);
        border-radius: 4px;
        color: var(--text-muted);
        font-family: var(--font-mono);
        background: var(--bg-elevated);
    }
    .sresults {
        max-height: 320px;
        overflow-y: auto;
        padding: 4px;
    }
    .sr {
        width: 100%;
        text-align: left;
        padding: 10px 12px;
        border-radius: 8px;
        border: none;
        background: none;
        cursor: pointer;
        display: flex;
        flex-direction: column;
        gap: 2px;
        transition: background 0.1s;
        color: var(--text-primary);
    }
    .sr:hover,
    .sr.sel {
        background: var(--bg-elevated);
    }
    .srn {
        font-size: 14px;
        font-weight: 600;
    }
    .srm {
        font-size: 12px;
        color: var(--text-muted);
        text-transform: capitalize;
    }
    .nores {
        padding: 24px;
        text-align: center;
        color: var(--text-muted);
        font-size: 14px;
    }
</style>
