import { writable, derived, get } from 'svelte/store';
import type { ComedianNode, AllianceEdge, EdgeType } from '$lib/data/types';
import seedData from '$lib/data/seed.json';

// Raw data
export const nodes = writable<ComedianNode[]>(seedData.nodes as ComedianNode[]);
export const edges = writable<AllianceEdge[]>(seedData.edges as AllianceEdge[]);

// Dark mode — default OFF (day mode default)
export const darkMode = writable(false);

// Theme toggle
export function toggleDarkMode() {
    darkMode.update(v => {
        const next = !v;
        if (typeof document !== 'undefined') {
            document.documentElement.classList.toggle('dark', next);
            localStorage.setItem('cartojester-theme', next ? 'dark' : 'light');
        }
        return next;
    });
}

export function initTheme() {
    if (typeof window === 'undefined') return;
    const saved = localStorage.getItem('cartojester-theme');
    const prefersDark = saved === 'dark';
    darkMode.set(prefersDark);
    document.documentElement.classList.toggle('dark', prefersDark);
}

// Time filter
export const yearRange = writable<[number, number]>([1895, 2025]);
export const isPlaying = writable(false);

// Edge type filters
export const enabledEdgeTypes = writable<Set<EdgeType>>(new Set([
    'collaboration', 'troupe', 'influence', 'mentor', 'rivalry', 'studio', 'family'
]));

// Tag filters
export const enabledTags = writable<Set<string>>(new Set());

// Min weight filter
export const minWeight = writable(1);

// Cluster mode
export type ClusterMode = 'none' | 'era' | 'medium' | 'organization';
export const clusterMode = writable<ClusterMode>('none');

// Search
export const searchQuery = writable('');
export const searchOpen = writable(false);

// Selected node/edge
export const selectedNodeId = writable<string | null>(null);
export const selectedEdgeId = writable<string | null>(null);

// Verified only filter
export const showOnlyVerified = writable(false);

// Derived: filtered nodes
export const filteredNodes = derived(
    [nodes, yearRange, enabledTags, showOnlyVerified],
    ([$nodes, $yearRange, $enabledTags, $showOnlyVerified]) => {
        return $nodes.filter(n => {
            // Time filter
            const start = n.activeStartYear ?? n.bornYear ?? 0;
            const end = n.activeEndYear ?? n.diedYear ?? 2025;
            if (end < $yearRange[0] || start > $yearRange[1]) return false;

            // Tag filter
            if ($enabledTags.size > 0) {
                if (!n.tags.some(t => $enabledTags.has(t))) return false;
            }

            // Verified filter
            if ($showOnlyVerified && n.source === 'auto') return false;

            return true;
        });
    }
);

// Derived: filtered edges
export const filteredEdges = derived(
    [edges, filteredNodes, enabledEdgeTypes, minWeight, yearRange, showOnlyVerified],
    ([$edges, $filteredNodes, $enabledEdgeTypes, $minWeight, $yearRange, $showOnlyVerified]) => {
        const nodeIds = new Set($filteredNodes.map(n => n.id));
        return $edges.filter(e => {
            if (!nodeIds.has(e.sourceId) || !nodeIds.has(e.targetId)) return false;
            if (!$enabledEdgeTypes.has(e.type)) return false;
            if (e.weight < $minWeight) return false;

            // Time filter for edges
            const start = e.startYear ?? 0;
            const end = e.endYear ?? 2025;
            if (end < $yearRange[0] || start > $yearRange[1]) return false;

            // Verified filter
            if ($showOnlyVerified && e.source === 'auto') return false;

            return true;
        });
    }
);

// Derived: selected node data
export const selectedNode = derived(
    [nodes, selectedNodeId],
    ([$nodes, $id]) => $id ? $nodes.find(n => n.id === $id) ?? null : null
);

// Derived: selected edge data
export const selectedEdge = derived(
    [edges, selectedEdgeId],
    ([$edges, $id]) => $id ? $edges.find(e => e.id === $id) ?? null : null
);

// Derived: alliances for selected node
export const selectedNodeAlliances = derived(
    [filteredEdges, selectedNodeId, nodes],
    ([$filteredEdges, $id, $nodes]) => {
        if (!$id) return [];
        return $filteredEdges
            .filter(e => e.sourceId === $id || e.targetId === $id)
            .map(e => {
                const otherId = e.sourceId === $id ? e.targetId : e.sourceId;
                const other = $nodes.find(n => n.id === otherId);
                return { edge: e, other };
            })
            .sort((a, b) => b.edge.weight - a.edge.weight);
    }
);

// Derived: node stats
export const nodeStats = derived(
    [filteredNodes, filteredEdges],
    ([$nodes, $edges]) => ({
        nodeCount: $nodes.length,
        edgeCount: $edges.length
    })
);

// Search results
export const searchResults = derived(
    [nodes, searchQuery],
    ([$nodes, $query]) => {
        if (!$query || $query.length < 1) return [];
        const q = $query.toLowerCase();
        return $nodes
            .filter(n =>
                n.name.toLowerCase().includes(q) ||
                n.aka.some(a => a.toLowerCase().includes(q)) ||
                n.id.toLowerCase().includes(q)
            )
            .slice(0, 10);
    }
);
