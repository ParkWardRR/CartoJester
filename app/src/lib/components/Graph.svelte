<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import * as d3 from "d3";
  import {
    filteredNodes,
    filteredEdges,
    selectedNodeId,
    selectedEdgeId,
    yearRange,
    clusterMode,
  } from "$lib/stores";
  import {
    EDGE_COLORS,
    TAG_COLORS,
    TAG_TO_CATEGORY,
    TAG_CATEGORIES,
  } from "$lib/data/types";
  import type { ComedianNode, AllianceEdge } from "$lib/data/types";

  let container: HTMLDivElement;
  let canvas: HTMLCanvasElement;
  let ctx: CanvasRenderingContext2D;
  let simulation: d3.Simulation<ComedianNode, undefined>;
  let transform = d3.zoomIdentity;
  let width = 800;
  let height = 600;
  let dpr = 1;
  let hoveredNode: ComedianNode | null = null;
  let tooltipX = 0;
  let tooltipY = 0;
  let simulationSettled = false;

  let currentNodes: ComedianNode[] = [];
  let currentEdges: AllianceEdge[] = [];
  let nodeMap = new Map<string, ComedianNode>();

  // Pre-computed degree map for faster lookups
  let degreeMap = new Map<string, number>();

  function computeDegrees() {
    degreeMap.clear();
    for (const e of currentEdges) {
      degreeMap.set(e.sourceId, (degreeMap.get(e.sourceId) || 0) + 1);
      degreeMap.set(e.targetId, (degreeMap.get(e.targetId) || 0) + 1);
    }
  }

  function initCanvas() {
    if (!canvas || !container) return;
    dpr = window.devicePixelRatio || 1;
    width = container.clientWidth;
    height = container.clientHeight;
    canvas.width = width * dpr;
    canvas.height = height * dpr;
    canvas.style.width = width + "px";
    canvas.style.height = height + "px";
    ctx = canvas.getContext("2d", { alpha: false })!;
    ctx.scale(dpr, dpr);
  }

  function getNodeRadius(n: ComedianNode): number {
    const degree = degreeMap.get(n.id) || 0;
    return Math.max(6, Math.min(24, 4 + n.notability * 2.5 + degree * 0.8));
  }

  function getNodeColor(n: ComedianNode): string {
    const primaryTag = n.tags[0];
    return TAG_COLORS[primaryTag] || "#94a3b8";
  }

  function draw() {
    if (!ctx) return;
    const isDark = document.documentElement.classList.contains("dark");
    const zoom_k = transform.k;

    ctx.fillStyle = isDark ? "#020617" : "#ffffff";
    ctx.fillRect(0, 0, width, height);

    ctx.save();
    ctx.translate(transform.x, transform.y);
    ctx.scale(zoom_k, zoom_k);

    // === SMART EDGE RENDERING ===
    // At low zoom, hide low-weight edges to reduce visual clutter
    const edgeWeightThreshold = zoom_k < 0.3 ? 3 : zoom_k < 0.6 ? 2 : 1;
    // Calculate visible viewport in graph coordinates
    const viewLeft = -transform.x / zoom_k;
    const viewTop = -transform.y / zoom_k;
    const viewRight = (width - transform.x) / zoom_k;
    const viewBottom = (height - transform.y) / zoom_k;

    for (const edge of currentEdges) {
      // Skip low-weight edges at low zoom
      if (edge.weight < edgeWeightThreshold) continue;

      const source = nodeMap.get(edge.sourceId);
      const target = nodeMap.get(edge.targetId);
      if (!source || !target || source.x == null || target.x == null) continue;

      // Viewport culling: skip edges completely outside view (with margin)
      const margin = 100;
      const ex1 = Math.min(source.x, target.x!);
      const ey1 = Math.min(source.y!, target.y!);
      const ex2 = Math.max(source.x, target.x!);
      const ey2 = Math.max(source.y!, target.y!);
      if (
        ex2 < viewLeft - margin ||
        ex1 > viewRight + margin ||
        ey2 < viewTop - margin ||
        ey1 > viewBottom + margin
      )
        continue;

      const isSelected = edge.id === $selectedEdgeId;
      const isConnectedToSelected =
        $selectedNodeId &&
        (edge.sourceId === $selectedNodeId ||
          edge.targetId === $selectedNodeId);

      ctx.beginPath();
      ctx.moveTo(source.x, source.y!);

      // Subtle curve for readability
      const midX = (source.x + target.x!) / 2;
      const midY = (source.y! + target.y!) / 2;
      const dx = target.x! - source.x;
      const dy = target.y! - source.y!;
      const dist = Math.sqrt(dx * dx + dy * dy);
      if (dist < 1) continue; // Skip zero-length edges
      const offset = dist * 0.05;
      ctx.quadraticCurveTo(
        midX - (dy * offset) / dist,
        midY + (dx * offset) / dist,
        target.x!,
        target.y!,
      );

      ctx.strokeStyle = EDGE_COLORS[edge.type] || "#94a3b8";

      // Density-aware line width: thinner when zoomed out
      const baseWidth = 0.5 + edge.weight * 0.3;
      ctx.lineWidth = isSelected
        ? 3
        : isConnectedToSelected
          ? 2
          : Math.max(0.3, baseWidth / Math.sqrt(zoom_k));

      // Density-aware opacity: fade edges more at low zoom
      const baseAlpha = $selectedNodeId
        ? isSelected
          ? 1
          : isConnectedToSelected
            ? 0.9
            : 0.08
        : Math.min(0.6, 0.15 + edge.weight * 0.1) * Math.min(1, zoom_k * 1.5);
      ctx.globalAlpha = baseAlpha;
      ctx.stroke();
    }

    ctx.globalAlpha = 1;

    // === SMART NODE RENDERING ===
    // At very low zoom, skip low-notability nodes that aren't connected to selection
    const notabilityThreshold = zoom_k < 0.3 ? 3 : zoom_k < 0.5 ? 2 : 0;

    for (const node of currentNodes) {
      if (node.x == null || node.y == null) continue;

      // Viewport culling
      const r = getNodeRadius(node);
      if (
        node.x + r < viewLeft ||
        node.x - r > viewRight ||
        node.y + r < viewTop ||
        node.y - r > viewBottom
      )
        continue;

      // Skip low-notability at low zoom (unless selected/connected)
      const isSelected = node.id === $selectedNodeId;
      const isConnected =
        $selectedNodeId &&
        currentEdges.some(
          (e) =>
            (e.sourceId === $selectedNodeId && e.targetId === node.id) ||
            (e.targetId === $selectedNodeId && e.sourceId === node.id),
        );
      const isHovered = hoveredNode?.id === node.id;

      if (
        !isSelected &&
        !isConnected &&
        !isHovered &&
        node.notability < notabilityThreshold
      )
        continue;

      const color = getNodeColor(node);

      // Dimming when a selection is active
      ctx.globalAlpha = isSelected || isConnected || !$selectedNodeId ? 1 : 0.2;

      // Glow for selected/hovered
      if (isSelected || isHovered) {
        ctx.shadowColor = color;
        ctx.shadowBlur = isSelected ? 20 : 12;
      }

      // Node circle
      ctx.beginPath();
      ctx.arc(node.x, node.y, r, 0, Math.PI * 2);

      // Gradient fill
      const grad = ctx.createRadialGradient(
        node.x - r * 0.3,
        node.y - r * 0.3,
        r * 0.1,
        node.x,
        node.y,
        r,
      );
      grad.addColorStop(0, lightenColor(color, 40));
      grad.addColorStop(1, color);
      ctx.fillStyle = grad;
      ctx.fill();

      // Border
      ctx.strokeStyle = isSelected
        ? "#ffffff"
        : isDark
          ? "rgba(255,255,255,0.3)"
          : "rgba(0,0,0,0.15)";
      ctx.lineWidth = isSelected ? 3 : 1;
      ctx.stroke();

      ctx.shadowColor = "transparent";
      ctx.shadowBlur = 0;

      // Auto badge
      if (node.source === "auto") {
        ctx.fillStyle = "#f59e0b";
        ctx.beginPath();
        ctx.arc(node.x + r * 0.7, node.y - r * 0.7, 4, 0, Math.PI * 2);
        ctx.fill();
      }

      // === SMART LABELS ===
      // Only show labels when zoomed in enough, or for selected/important nodes
      const degree = degreeMap.get(node.id) || 0;
      const importance = node.notability + degree * 0.3;
      const labelZoomThreshold =
        importance > 6 ? 0.4 : importance > 4 ? 0.6 : 0.8;

      const showLabel =
        isSelected || isHovered || isConnected || zoom_k > labelZoomThreshold;

      if (showLabel) {
        ctx.globalAlpha =
          isSelected || isConnected || !$selectedNodeId ? 1 : 0.3;
        const fontSize = Math.max(9, Math.min(14, 12 / Math.sqrt(zoom_k)));
        ctx.font = `${isSelected ? "bold " : ""}${fontSize}px Inter, system-ui`;
        ctx.textAlign = "center";
        ctx.textBaseline = "top";

        const text = node.name;
        const metrics = ctx.measureText(text);
        const textY = node.y + r + 4;
        ctx.fillStyle = isDark
          ? "rgba(2, 6, 23, 0.8)"
          : "rgba(255, 255, 255, 0.9)";
        ctx.fillRect(
          node.x - metrics.width / 2 - 3,
          textY - 1,
          metrics.width + 6,
          fontSize + 2,
        );

        ctx.fillStyle = isDark ? "#e2e8f0" : "#1e293b";
        ctx.fillText(text, node.x, textY);
      }
    }

    ctx.globalAlpha = 1;
    ctx.restore();
  }

  function lightenColor(hex: string, percent: number): string {
    const num = parseInt(hex.replace("#", ""), 16);
    const r = Math.min(255, (num >> 16) + percent);
    const g = Math.min(255, ((num >> 8) & 0x00ff) + percent);
    const b = Math.min(255, (num & 0x0000ff) + percent);
    return `rgb(${r},${g},${b})`;
  }

  function findNodeAt(x: number, y: number): ComedianNode | null {
    const [tx, ty] = transform.invert([x, y]);
    for (let i = currentNodes.length - 1; i >= 0; i--) {
      const n = currentNodes[i];
      if (n.x == null || n.y == null) continue;
      const r = getNodeRadius(n);
      const dx = tx - n.x;
      const dy = ty - n.y;
      if (dx * dx + dy * dy < (r + 3) * (r + 3)) return n;
    }
    return null;
  }

  function handleClick(e: MouseEvent) {
    const rect = canvas.getBoundingClientRect();
    const node = findNodeAt(e.clientX - rect.left, e.clientY - rect.top);
    if (node) {
      selectedNodeId.set(node.id === $selectedNodeId ? null : node.id);
      selectedEdgeId.set(null);
    } else {
      selectedNodeId.set(null);
      selectedEdgeId.set(null);
    }
    draw();
  }

  function handleMouseMove(e: MouseEvent) {
    const rect = canvas.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    const node = findNodeAt(x, y);
    if (node !== hoveredNode) {
      hoveredNode = node;
      tooltipX = e.clientX;
      tooltipY = e.clientY;
      canvas.style.cursor = node ? "pointer" : "grab";
      draw();
    } else if (node) {
      tooltipX = e.clientX;
      tooltipY = e.clientY;
    }
  }

  function handleDblClick(e: MouseEvent) {
    const rect = canvas.getBoundingClientRect();
    const node = findNodeAt(e.clientX - rect.left, e.clientY - rect.top);
    if (node && node.x != null && node.y != null) {
      const t = d3.zoomIdentity
        .translate(width / 2, height / 2)
        .scale(2.5)
        .translate(-node.x, -node.y);
      d3.select(canvas)
        .transition()
        .duration(750)
        .call(zoom.transform as any, t);
    }
  }

  let zoom: d3.ZoomBehavior<HTMLCanvasElement, unknown>;
  let draggedNode: ComedianNode | null = null;

  function setupInteractions() {
    zoom = d3
      .zoom<HTMLCanvasElement, unknown>()
      .scaleExtent([0.1, 8])
      .on("zoom", (event) => {
        transform = event.transform;
        draw();
      });

    const sel = d3.select(canvas);
    sel.call(zoom);

    // Drag
    sel.call(
      d3
        .drag<HTMLCanvasElement, unknown>()
        .subject((event) => {
          const node = findNodeAt(event.x, event.y);
          if (node) return { x: node.x, y: node.y, node };
          return null;
        })
        .on("start", (event) => {
          if (!event.subject) return;
          draggedNode = event.subject.node;
          if (draggedNode) {
            draggedNode.fx = draggedNode.x;
            draggedNode.fy = draggedNode.y;
            // Use a low alpha to avoid jostling the whole graph
            simulation.alphaTarget(0.1).restart();
          }
        })
        .on("drag", (event) => {
          if (!draggedNode) return;
          const [x, y] = transform.invert([event.x, event.y]);
          draggedNode.fx = x;
          draggedNode.fy = y;
        })
        .on("end", () => {
          if (draggedNode) {
            draggedNode.fx = null;
            draggedNode.fy = null;
            simulation.alphaTarget(0);
          }
          draggedNode = null;
        }),
    );
  }

  function buildSimulation() {
    if (simulation) simulation.stop();

    computeDegrees();

    const linkForce = d3
      .forceLink<ComedianNode, any>()
      .id((d: any) => d.id)
      .links(
        currentEdges.map((e) => ({
          source: e.sourceId,
          target: e.targetId,
          weight: e.weight,
        })),
      )
      .distance((d) => 100 + 40 / Math.max(1, (d as any).weight))
      .strength((d) => 0.2 + (d as any).weight * 0.05);

    simulation = d3
      .forceSimulation(currentNodes)
      .force("link", linkForce)
      // Stronger repulsion to spread dense clusters
      .force(
        "charge",
        d3
          .forceManyBody()
          .strength((d) => -120 - (d as any).notability * 25)
          .distanceMax(600),
      )
      // Gentle centering — no constant pull that causes rotation
      .force("center", d3.forceCenter(width / 2, height / 2).strength(0.01))
      // Stronger collision to prevent overlap in dense areas
      .force(
        "collision",
        d3
          .forceCollide()
          .radius((d: any) => getNodeRadius(d) + 8)
          .strength(0.8),
      )
      // Very gentle gravity — prevents drift without causing rotation
      .force("x", d3.forceX(width / 2).strength(0.008))
      .force("y", d3.forceY(height / 2).strength(0.008))
      // FAST decay so it settles quickly — the core fix for "too much auto-rotation"
      .alphaDecay(0.05)
      .velocityDecay(0.5)
      .on("tick", () => {
        draw();
      })
      .on("end", () => {
        simulationSettled = true;
      });

    // Start with a reasonable alpha so it settles within a few seconds
    simulation.alpha(0.6);
  }

  function updateData(nodes: ComedianNode[], edges: AllianceEdge[]) {
    const oldPositions = new Map(
      currentNodes.map((n) => [n.id, { x: n.x, y: n.y }]),
    );

    currentNodes = nodes.map((n) => {
      const old = oldPositions.get(n.id);
      return {
        ...n,
        x: old?.x ?? width / 2 + (Math.random() - 0.5) * 400,
        y: old?.y ?? height / 2 + (Math.random() - 0.5) * 400,
      };
    });
    currentEdges = edges;
    nodeMap = new Map(currentNodes.map((n) => [n.id, n]));
    simulationSettled = false;
    buildSimulation();
  }

  let resizeObserver: ResizeObserver;

  onMount(() => {
    initCanvas();
    setupInteractions();

    const unsub1 = filteredNodes.subscribe((n) => {
      const e = currentEdges.length ? currentEdges : [];
      filteredEdges.subscribe((edges) => {
        updateData(n, edges);
      })();
    });

    const unsub2 = filteredEdges.subscribe((e) => {
      updateData(currentNodes.length ? currentNodes : [], e);
    });

    resizeObserver = new ResizeObserver(() => {
      initCanvas();
      // DON'T restart the simulation on resize — just update center forces gently
      if (simulation) {
        simulation.force(
          "center",
          d3.forceCenter(width / 2, height / 2).strength(0.01),
        );
        simulation.force("x", d3.forceX(width / 2).strength(0.008));
        simulation.force("y", d3.forceY(height / 2).strength(0.008));
        // Only restart if not yet settled, and with very low alpha
        if (!simulationSettled) {
          simulation.alpha(0.1).restart();
        } else {
          draw();
        }
      }
    });
    resizeObserver.observe(container);

    filteredNodes.subscribe(() => {})();
    filteredEdges.subscribe(() => {})();
  });

  onDestroy(() => {
    if (simulation) simulation.stop();
    if (resizeObserver) resizeObserver.disconnect();
  });

  $: if (ctx) {
    updateData($filteredNodes, $filteredEdges);
  }

  export function zoomToNode(nodeId: string) {
    const node = nodeMap.get(nodeId);
    if (node && node.x != null && node.y != null) {
      const t = d3.zoomIdentity
        .translate(width / 2, height / 2)
        .scale(2.5)
        .translate(-node.x, -node.y);
      d3.select(canvas)
        .transition()
        .duration(750)
        .call(zoom.transform as any, t);
      selectedNodeId.set(nodeId);
    }
  }

  export function resetZoom() {
    d3.select(canvas)
      .transition()
      .duration(500)
      .call(zoom.transform as any, d3.zoomIdentity);
  }
</script>

<div
  bind:this={container}
  class="graph-container gpu-canvas"
  role="img"
  aria-label="Comedian network graph"
>
  <canvas
    bind:this={canvas}
    on:click={handleClick}
    on:mousemove={handleMouseMove}
    on:dblclick={handleDblClick}
  ></canvas>

  {#if hoveredNode}
    <div
      class="tooltip"
      style="left: {tooltipX + 12}px; top: {tooltipY - 12}px;"
    >
      <div class="tooltip-name">{hoveredNode.name}</div>
      <div class="tooltip-meta">
        {#if hoveredNode.bornYear}
          {hoveredNode.bornYear}{hoveredNode.diedYear
            ? `–${hoveredNode.diedYear}`
            : "–present"}
        {/if}
        · {degreeMap.get(hoveredNode.id) || 0} connections
      </div>
      <div class="tooltip-tags">
        {#each hoveredNode.tags as tag}
          {@const catId = TAG_TO_CATEGORY[tag]}
          {@const cat = TAG_CATEGORIES.find((c) => c.id === catId)}
          <span
            class="tooltip-tag"
            style="background: {TAG_COLORS[tag] ||
              '#94a3b8'}40; color: {TAG_COLORS[tag] || '#94a3b8'}"
            >{cat?.emoji ?? ""} {tag}</span
          >
        {/each}
      </div>
    </div>
  {/if}
</div>

<style>
  .graph-container {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
    border-radius: 12px;
  }
  canvas {
    display: block;
    width: 100%;
    height: 100%;
    cursor: grab;
  }
  canvas:active {
    cursor: grabbing;
  }
  .tooltip {
    position: fixed;
    pointer-events: none;
    z-index: 100;
    background: var(--glass-bg);
    backdrop-filter: blur(12px);
    -webkit-backdrop-filter: blur(12px);
    border: 1px solid var(--glass-border);
    border-radius: 10px;
    padding: 10px 14px;
    box-shadow: var(--shadow-lg);
    max-width: 240px;
  }
  .tooltip-name {
    font-weight: 700;
    font-size: 14px;
    color: var(--text-primary);
    margin-bottom: 2px;
  }
  .tooltip-meta {
    font-size: 12px;
    color: var(--text-secondary);
    margin-bottom: 6px;
  }
  .tooltip-tags {
    display: flex;
    gap: 4px;
    flex-wrap: wrap;
  }
  .tooltip-tag {
    font-size: 10px;
    font-weight: 600;
    padding: 2px 6px;
    border-radius: 4px;
    text-transform: uppercase;
    letter-spacing: 0.5px;
  }
</style>
