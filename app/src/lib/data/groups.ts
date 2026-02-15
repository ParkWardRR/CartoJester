import type { ComedianNode } from './types';

// Comedy groups/movements definitions
export interface ComedyGroup {
    id: string;
    name: string;
    era: string;
    yearStart: number;
    yearEnd: number;
    description: string;
    characteristics: string[];
    memberIds: string[];
    color: string;
    icon: string;
}

export const COMEDY_GROUPS: ComedyGroup[] = [
    {
        id: 'silent-pioneers',
        name: 'Silent Era Pioneers',
        era: '1895–1930',
        yearStart: 1895,
        yearEnd: 1930,
        description: 'The original physical comedians who invented the language of screen comedy through slapstick, sight gags, and pantomime. They proved comedy could be a universal art form requiring no words.',
        characteristics: ['Physical comedy', 'Slapstick', 'Pantomime', 'Visual storytelling', 'Studio system'],
        memberIds: ['chaplin', 'keaton', 'lloyd', 'laurel', 'hardy', 'mack-sennett', 'arbuckle', 'langdon'],
        color: '#94a3b8',
        icon: '🎬'
    },
    {
        id: 'vaudeville-stars',
        name: 'Vaudeville & Stage Comics',
        era: '1900–1950',
        yearStart: 1900,
        yearEnd: 1950,
        description: 'Performers who honed their craft on the vaudeville circuit and brought their acts to film and early television. Masters of timing, wordplay, and audience interaction.',
        characteristics: ['Timing', 'Wordplay', 'Musical comedy', 'Audience rapport', 'Variety acts'],
        memberIds: ['groucho', 'harpo', 'chico', 'wc-fields', 'mae-west', 'burns', 'jack-benny', 'three-stooges-moe', 'three-stooges-larry', 'three-stooges-curly', 'abbott', 'costello'],
        color: '#d97706',
        icon: '🎭'
    },
    {
        id: 'golden-age-tv',
        name: 'Golden Age of Television',
        era: '1950–1970',
        yearStart: 1950,
        yearEnd: 1970,
        description: 'The first generation of TV comedians who defined the sitcom format and brought comedy into every living room. They proved television could sustain long-form comedic storytelling.',
        characteristics: ['Sitcom format', 'Family-friendly', 'Physical comedy on TV', 'Live audiences', 'Network era'],
        memberIds: ['ball', 'hope', 'skelton', 'sid-caesar', 'carl-reiner', 'dick-van-dyke', 'don-rickles', 'phyllis-diller', 'jack-benny', 'burns', 'newhart'],
        color: '#8b5cf6',
        icon: '📺'
    },
    {
        id: 'counterculture-rebels',
        name: 'Counterculture Comedy',
        era: '1960–1980',
        yearStart: 1960,
        yearEnd: 1980,
        description: 'Rebels who tore down comedy conventions. They tackled politics, race, drugs, and social hypocrisy with unflinching honesty, transforming stand-up into an art form that challenged power.',
        characteristics: ['Social commentary', 'Boundary-pushing', 'Political satire', 'Confessional style', 'Anti-establishment'],
        memberIds: ['bruce', 'pryor', 'carlin', 'joan-rivers', 'lily-tomlin', 'dick-gregory', 'mort-sahl', 'bill-hicks', 'moms-mabley'],
        color: '#ef4444',
        icon: '✊'
    },
    {
        id: 'snl-original',
        name: 'SNL & Sketch Revolution',
        era: '1975–1995',
        yearStart: 1975,
        yearEnd: 1995,
        description: 'Saturday Night Live and its satellite shows launched careers and defined a generation. These performers turned sketch comedy and improv into Hollywood gold.',
        characteristics: ['Sketch comedy', 'Improv roots', 'Character work', 'Political satire', 'Ensemble energy'],
        memberIds: ['belushi', 'aykroyd', 'murray', 'murphy', 'martin-short', 'chevy-chase', 'gilda-radner', 'mike-myers', 'dana-carvey', 'phil-hartman', 'chris-farley', 'norm-macdonald', 'david-spade', 'molly-shannon', 'will-ferrell'],
        color: '#3b82f6',
        icon: '📡'
    },
    {
        id: 'standup-boom',
        name: '80s/90s Stand-Up Boom',
        era: '1980–2000',
        yearStart: 1980,
        yearEnd: 2000,
        description: 'The comedy club explosion and HBO specials era turned stand-up into mainstream entertainment. Arena tours, comedy albums, and TV deals became the new gold standard.',
        characteristics: ['Comedy clubs', 'HBO specials', 'Observational humor', 'Arena tours', 'Late-night crossover'],
        memberIds: ['seinfeld', 'rock', 'martin', 'williams', 'goldberg', 'ellen', 'ray-romano', 'tim-allen', 'roseanne', 'sinbad', 'bernie-mac', 'cedric', 'steve-harvey', 'dl-hughley', 'wanda-sykes'],
        color: '#f97316',
        icon: '🎤'
    },
    {
        id: 'alt-comedy',
        name: 'Alternative Comedy',
        era: '1995–2015',
        yearStart: 1995,
        yearEnd: 2015,
        description: 'A movement rejecting traditional punchline-driven comedy for quirky, meta, absurdist, and anti-humor styles. Often associated with indie venues and podcasts.',
        characteristics: ['Anti-humor', 'Meta-comedy', 'Absurdism', 'Indie venues', 'Podcast culture'],
        memberIds: ['zach-galifianakis', 'maria-bamford', 'patton-oswalt', 'bo-burnham', 'reggie-watts', 'tig-notaro', 'hannibal-buress', 'eric-andre', 'demetri-martin'],
        color: '#14b8a6',
        icon: '🌀'
    },
    {
        id: 'improv-ucb',
        name: 'Improv & UCB Movement',
        era: '1990–2020',
        yearStart: 1990,
        yearEnd: 2020,
        description: 'The improv pipeline from Second City, UCB, iO, and Groundlings fed TV writers rooms and created a new comedy aesthetic built on "yes, and" collaboration.',
        characteristics: ['Yes, and', 'Long-form improv', 'Writers room pipeline', 'Ensemble focus', 'Character work'],
        memberIds: ['fey', 'poehler', 'will-ferrell', 'kate-mckinnon', 'keegan-michael-key', 'jordan-peele', 'abbi-jacobson', 'ilana-glazer', 'nick-kroll', 'jason-sudeikis'],
        color: '#06b6d4',
        icon: '🎪'
    },
    {
        id: 'kings-of-comedy',
        name: 'Kings & Queens of Comedy',
        era: '2000–2020',
        yearStart: 2000,
        yearEnd: 2020,
        description: 'The original Kings of Comedy tour and its spiritual successors brought Black comedy to arena-scale audiences and launched multiple media empires.',
        characteristics: ['Arena comedy', 'Cultural commentary', 'Media empires', 'Netflix specials', 'Tour economics'],
        memberIds: ['rock', 'hart', 'chappelle', 'bernie-mac', 'cedric', 'steve-harvey', 'dl-hughley', 'tiffany-haddish', 'wanda-sykes'],
        color: '#eab308',
        icon: '👑'
    },
    {
        id: 'netflix-era',
        name: 'Streaming & Netflix Era',
        era: '2015–2025',
        yearStart: 2015,
        yearEnd: 2025,
        description: 'Comedy specials went global via streaming. Stand-up became a worldwide phenomenon, diversifying voices and reaching massive new audiences.',
        characteristics: ['Global reach', 'Diverse voices', 'Social media crossover', 'Special-driven careers', 'International comedy'],
        memberIds: ['mulaney', 'wong', 'bo-burnham', 'hasan-minhaj', 'trevor-noah', 'hannah-gadsby', 'nate-bargatze', 'taylor-tomlinson', 'sam-morril', 'mark-normand', 'shane-gillis', 'matt-rife', 'nikki-glaser'],
        color: '#ec4899',
        icon: '📱'
    }
];
