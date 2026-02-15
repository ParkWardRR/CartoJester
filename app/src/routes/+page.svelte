<script lang="ts">
    import FilterPanel from "$lib/components/FilterPanel.svelte";
    import Graph from "$lib/components/Graph.svelte";
    import DetailsPanel from "$lib/components/DetailsPanel.svelte";
    import TimeSlider from "$lib/components/TimeSlider.svelte";
    import Legend from "$lib/components/Legend.svelte";
</script>

<svelte:head>
    <title>CartoJester — Comedian Alliances Atlas</title>
    <meta
        name="description"
        content="Explore comedian alliances across decades with an interactive network graph."
    />
</svelte:head>

<div class="explorer">
    <FilterPanel />
    <div class="center-col">
        <div class="graph-wrap">
            <Graph />
            <Legend />
        </div>
        <TimeSlider />
    </div>
    <DetailsPanel />
</div>

<style>
    .explorer {
        display: flex;
        flex: 1;
        overflow: hidden;
    }
    .center-col {
        flex: 1;
        display: flex;
        flex-direction: column;
        overflow: hidden;
        min-width: 0;
    }
    .graph-wrap {
        flex: 1;
        position: relative;
        overflow: hidden;
    }
</style>
