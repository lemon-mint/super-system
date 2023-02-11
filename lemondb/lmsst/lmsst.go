package lmsst

const LMSSTv1_MAGIC = 0x9579aea3c5752d15

// LMSSTv1 File Format

// LMSSTv1 File Header
// *---------------------------------*---------------------------------*
// | 0x9579aea3c5752d15 (8B) - Magic |           File ID (8B)          |
// *---------------------------------*---------------------------------*
// | Flags (2B) |                   Reserved (14B)                     |
// *---------------------------------*---------------------------------*

// Entry is a single key-value pair in the LMSST file.
//
// *--------------------*
// |
//
