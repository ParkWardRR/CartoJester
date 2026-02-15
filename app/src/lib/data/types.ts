export interface ComedianNode {
    id: string;
    name: string;
    aka: string[];
    bornYear: number | null;
    diedYear: number | null;
    activeStartYear: number | null;
    activeEndYear: number | null;
    tags: string[];
    notability: number; // 1–5
    links: { label: string; url: string }[];
    source?: 'seed' | 'auto';
    externalIds?: { wikidata?: string };
    x?: number;
    y?: number;
    fx?: number | null;
    fy?: number | null;
    vx?: number;
    vy?: number;
}

export type EdgeType = 'collaboration' | 'troupe' | 'influence' | 'mentor' | 'rivalry' | 'studio' | 'family';

export interface AllianceEdge {
    id: string;
    sourceId: string;
    targetId: string;
    type: EdgeType;
    startYear: number | null;
    endYear: number | null;
    weight: number; // 1–5
    summary: string;
    evidence: { url: string; label: string }[];
    source?: 'seed' | 'auto';
}

export interface SeedData {
    nodes: ComedianNode[];
    edges: AllianceEdge[];
}

export const EDGE_TYPES: EdgeType[] = ['collaboration', 'troupe', 'influence', 'mentor', 'rivalry', 'studio', 'family'];

export const EDGE_COLORS: Record<EdgeType, string> = {
    collaboration: '#a855f7',
    troupe: '#f59e0b',
    influence: '#3b82f6',
    mentor: '#10b981',
    rivalry: '#ef4444',
    studio: '#8b5cf6',
    family: '#ec4899'
};

export interface TagCategory {
    id: string;
    label: string;
    emoji: string;
    tags: Record<string, string>; // tag → color
}

export const TAG_CATEGORIES: TagCategory[] = [
    {
        id: 'format',
        label: 'Format',
        emoji: '🎤',
        tags: {
            standup: '#f97316',
            sketch: '#06b6d4',
            improv: '#14b8a6',
            film: '#eab308',
            tv: '#8b5cf6',
            animation: '#f43f5e',
            panel: '#e879f9',
            podcast: '#10b981',
            latenight: '#f472b6',
            snl: '#3b82f6',
        }
    },
    {
        id: 'era',
        label: 'Era',
        emoji: '🕰️',
        tags: {
            silent: '#94a3b8',
            vaudeville: '#d97706',
            keystone: '#78716c',
        }
    },
    {
        id: 'region',
        label: 'Region',
        emoji: '🌏',
        tags: {
            india: '#f59e0b',
            philippines: '#22d3ee',
            indonesia: '#06b6d4',
            malaysia: '#a78bfa',
            singapore: '#818cf8',
        }
    },
];

// Flat map for backward compatibility — all components that just need tag→color can still use this
export const TAG_COLORS: Record<string, string> = TAG_CATEGORIES.reduce(
    (acc, cat) => ({ ...acc, ...cat.tags }),
    {} as Record<string, string>
);

// Lookup: tag → category id
export const TAG_TO_CATEGORY: Record<string, string> = TAG_CATEGORIES.reduce(
    (acc, cat) => {
        for (const tag of Object.keys(cat.tags)) acc[tag] = cat.id;
        return acc;
    },
    {} as Record<string, string>
);

export const ERA_PRESETS: { label: string; start: number; end: number }[] = [
    { label: 'Silent Era', start: 1895, end: 1930 },
    { label: 'Golden Age', start: 1930, end: 1960 },
    { label: 'Counterculture', start: 1960, end: 1980 },
    { label: 'Comedy Boom', start: 1980, end: 2000 },
    { label: 'Modern', start: 2000, end: 2025 },
    { label: 'All Time', start: 1895, end: 2025 }
];
