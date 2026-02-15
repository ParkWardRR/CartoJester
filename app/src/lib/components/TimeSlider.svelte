<script lang="ts">
    import { yearRange, isPlaying } from "$lib/stores";
    import { onDestroy } from "svelte";

    let interval: ReturnType<typeof setInterval> | null = null;

    $: rangeStart = $yearRange[0];
    $: rangeEnd = $yearRange[1];

    function handleStartChange(e: Event) {
        const v = parseInt((e.target as HTMLInputElement).value);
        if (v <= rangeEnd) yearRange.set([v, rangeEnd]);
    }

    function handleEndChange(e: Event) {
        const v = parseInt((e.target as HTMLInputElement).value);
        if (v >= rangeStart) yearRange.set([rangeStart, v]);
    }

    function togglePlay() {
        isPlaying.update((v) => {
            if (!v) {
                // Start
                yearRange.set([1895, 1895]);
                interval = setInterval(() => {
                    yearRange.update(([s, e]) => {
                        if (e >= 2025) {
                            isPlaying.set(false);
                            if (interval) clearInterval(interval);
                            return [1895, 2025];
                        }
                        return [s, Math.min(2025, e + 1)];
                    });
                }, 80);
                return true;
            } else {
                if (interval) clearInterval(interval);
                return false;
            }
        });
    }

    const decades = [
        1910, 1920, 1930, 1940, 1950, 1960, 1970, 1980, 1990, 2000, 2010, 2020,
    ];

    function jumpToDecade(d: number) {
        yearRange.set([d, d + 9]);
    }

    onDestroy(() => {
        if (interval) clearInterval(interval);
    });
</script>

<div class="time-control" role="region" aria-label="Time range filter">
    <div class="time-main">
        <button
            class="play-btn"
            on:click={togglePlay}
            title={$isPlaying ? "Pause" : "Play through years"}
        >
            {#if $isPlaying}
                <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    ><rect x="6" y="4" width="4" height="16" /><rect
                        x="14"
                        y="4"
                        width="4"
                        height="16"
                    /></svg
                >
            {:else}
                <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="currentColor"
                    ><polygon points="5 3 19 12 5 21 5 3" /></svg
                >
            {/if}
        </button>

        <div class="slider-container">
            <div class="slider-track">
                <input
                    type="range"
                    min="1895"
                    max="2025"
                    value={rangeStart}
                    on:input={handleStartChange}
                    class="range-input start"
                    aria-label="Start year"
                />
                <input
                    type="range"
                    min="1895"
                    max="2025"
                    value={rangeEnd}
                    on:input={handleEndChange}
                    class="range-input end"
                    aria-label="End year"
                />
            </div>
        </div>

        <div class="year-display">
            <span class="year-badge">{rangeStart}</span>
            <span class="year-sep">—</span>
            <span class="year-badge">{rangeEnd}</span>
        </div>
    </div>

    <div class="decade-jumps">
        {#each decades as d}
            <button class="decade-btn" on:click={() => jumpToDecade(d)}
                >{d}s</button
            >
        {/each}
        <button
            class="decade-btn all"
            on:click={() => yearRange.set([1895, 2025])}>All</button
        >
    </div>
</div>

<style>
    .time-control {
        background: var(--bg-panel);
        border-top: 1px solid var(--border);
        padding: 12px 20px;
        display: flex;
        flex-direction: column;
        gap: 8px;
    }
    .time-main {
        display: flex;
        align-items: center;
        gap: 14px;
    }
    .play-btn {
        width: 36px;
        height: 36px;
        border-radius: 50%;
        border: none;
        background: linear-gradient(
            135deg,
            var(--color-brand-500),
            var(--color-brand-600)
        );
        color: white;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        flex-shrink: 0;
        box-shadow: 0 2px 8px rgba(168, 85, 247, 0.3);
        transition:
            transform 0.15s,
            box-shadow 0.15s;
    }
    .play-btn:hover {
        transform: scale(1.08);
        box-shadow: 0 3px 12px rgba(168, 85, 247, 0.4);
    }
    .slider-container {
        flex: 1;
        position: relative;
        height: 24px;
    }
    .slider-track {
        position: relative;
        width: 100%;
        height: 100%;
    }
    .range-input {
        position: absolute;
        width: 100%;
        top: 0;
        height: 24px;
        -webkit-appearance: none;
        appearance: none;
        background: transparent;
        pointer-events: none;
    }
    .range-input::-webkit-slider-thumb {
        -webkit-appearance: none;
        width: 16px;
        height: 16px;
        border-radius: 50%;
        background: var(--color-brand-500);
        border: 2px solid white;
        box-shadow: 0 1px 4px rgba(0, 0, 0, 0.2);
        pointer-events: all;
        cursor: grab;
    }
    .range-input::-webkit-slider-runnable-track {
        height: 4px;
        background: var(--bg-elevated);
        border-radius: 2px;
    }
    .range-input.start::-webkit-slider-runnable-track {
        background: transparent;
    }
    .year-display {
        display: flex;
        align-items: center;
        gap: 4px;
        flex-shrink: 0;
    }
    .year-badge {
        font-size: 13px;
        font-weight: 700;
        color: var(--text-primary);
        background: var(--bg-elevated);
        padding: 4px 8px;
        border-radius: 6px;
        font-variant-numeric: tabular-nums;
        min-width: 48px;
        text-align: center;
    }
    .year-sep {
        color: var(--text-muted);
        font-size: 12px;
    }
    .decade-jumps {
        display: flex;
        gap: 2px;
        flex-wrap: wrap;
    }
    .decade-btn {
        font-size: 10px;
        font-weight: 600;
        padding: 3px 7px;
        border-radius: 4px;
        border: 1px solid var(--border);
        background: var(--bg-card);
        color: var(--text-muted);
        cursor: pointer;
        transition: all 0.15s;
    }
    .decade-btn:hover {
        background: var(--bg-elevated);
        color: var(--text-primary);
        border-color: var(--border-active);
    }
    .decade-btn.all {
        background: var(--color-brand-500);
        color: white;
        border-color: var(--color-brand-500);
    }
</style>
