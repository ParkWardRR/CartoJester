<script lang="ts">
    import { darkMode, toggleDarkMode, searchOpen } from "$lib/stores";
    import { page } from "$app/stores";
    import { base } from "$app/paths";
</script>

<nav class="navbar" aria-label="Main navigation">
    <a href="{base}/" class="logo">
        <span class="logo-icon">🃏</span>
        <span class="logo-text">CartoJester</span>
    </a>

    <div class="nav-links">
        <a href="{base}/" class:active={$page.url.pathname === `${base}/`}
            >Explorer</a
        >
        <a
            href="{base}/about"
            class:active={$page.url.pathname === `${base}/about`}>About</a
        >
        <a
            href="{base}/data"
            class:active={$page.url.pathname === `${base}/data`}>Data</a
        >
    </div>

    <div class="nav-actions">
        <button
            class="nav-btn"
            on:click={() => searchOpen.set(true)}
            title="Search (⌘K)"
        >
            <svg
                width="16"
                height="16"
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
            <kbd>⌘K</kbd>
        </button>
        <button
            class="nav-btn theme-toggle"
            on:click={toggleDarkMode}
            title="Toggle theme"
        >
            {#if $darkMode}
                <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    ><circle cx="12" cy="12" r="5" /><line
                        x1="12"
                        y1="1"
                        x2="12"
                        y2="3"
                    /><line x1="12" y1="21" x2="12" y2="23" /><line
                        x1="4.22"
                        y1="4.22"
                        x2="5.64"
                        y2="5.64"
                    /><line x1="18.36" y1="18.36" x2="19.78" y2="19.78" /><line
                        x1="1"
                        y1="12"
                        x2="3"
                        y2="12"
                    /><line x1="21" y1="12" x2="23" y2="12" /><line
                        x1="4.22"
                        y1="19.78"
                        x2="5.64"
                        y2="18.36"
                    /><line x1="18.36" y1="5.64" x2="19.78" y2="4.22" /></svg
                >
            {:else}
                <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    ><path
                        d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"
                    /></svg
                >
            {/if}
        </button>
    </div>
</nav>

<style>
    .navbar {
        display: flex;
        align-items: center;
        justify-content: space-between;
        height: 52px;
        padding: 0 16px;
        background: var(--bg-panel);
        border-bottom: 1px solid var(--border);
        z-index: 50;
    }
    .logo {
        display: flex;
        align-items: center;
        gap: 8px;
        text-decoration: none;
        color: var(--text-primary);
    }
    .logo-icon {
        font-size: 22px;
    }
    .logo-text {
        font-size: 16px;
        font-weight: 800;
        letter-spacing: -0.03em;
        background: linear-gradient(
            135deg,
            var(--color-brand-500),
            var(--color-accent-500)
        );
        -webkit-background-clip: text;
        -webkit-text-fill-color: transparent;
    }
    .nav-links {
        display: flex;
        gap: 4px;
    }
    .nav-links a {
        font-size: 13px;
        font-weight: 500;
        color: var(--text-secondary);
        text-decoration: none;
        padding: 6px 12px;
        border-radius: 6px;
        transition: all 0.15s;
    }
    .nav-links a:hover {
        color: var(--text-primary);
        background: var(--bg-elevated);
    }
    .nav-links a.active {
        color: var(--text-primary);
        background: var(--bg-elevated);
        font-weight: 600;
    }
    .nav-actions {
        display: flex;
        gap: 4px;
    }
    .nav-btn {
        display: flex;
        align-items: center;
        gap: 6px;
        height: 32px;
        padding: 0 10px;
        border-radius: 6px;
        border: 1px solid var(--border);
        background: var(--bg-card);
        color: var(--text-secondary);
        cursor: pointer;
        font-size: 12px;
        transition: all 0.15s;
    }
    .nav-btn:hover {
        background: var(--bg-elevated);
        color: var(--text-primary);
    }
    .nav-btn kbd {
        font-size: 10px;
        padding: 1px 4px;
        background: var(--bg-elevated);
        border: 1px solid var(--border);
        border-radius: 3px;
        font-family: var(--font-mono);
        color: var(--text-muted);
    }
    .theme-toggle {
        width: 32px;
        padding: 0;
        justify-content: center;
    }
</style>
