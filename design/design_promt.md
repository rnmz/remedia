```json
{
  "$schema": "https://design-tokens.github.io/community-group/format/",
  "meta": {
    "project": "YOMI • 読み (Manga & Comic Reading Platform)",
    "version": "2.0.0",
    "designPhilosophy": [
      "Apple Human Interface Guidelines (Apple Books / App Store tactile clarity, translucent vibrancy, spring physics)",
      "Japanese Editorial Minimalism (Ma — generous negative space, Washi paper warmth, Hanko stamp accents)",
      "Modern Bento / Metro UI (Structured modular tiles, crisp 1px borders, scannable typographic hierarchy)"
    ],
    "antiPatterns": [
      "No aggressive cyberpunk neon or harsh glowing box-shadows",
      "No pure #FFFFFF backgrounds or pure #000000 text — always use warm paper and sumi ink tones",
      "No cluttered metadata overlays on top of cover artwork"
    ]
  },
  "colorThemes": {
    "light_washi": {
      "label": "Washi Paper (Default Light)",
      "surfaces": {
        "bgCanvas": "#F6F5F2",
        "bgSurface": "#FFFFFF",
        "bgElevated": "#EFECE6",
        "bgSubtle": "#EAE6DF",
        "bgGlassHeader": "rgba(246, 245, 242, 0.78)",
        "bgGlassFloating": "rgba(255, 255, 255, 0.85)"
      },
      "typography": {
        "textPrimary": "#1C1B18",
        "textSecondary": "#6E6A63",
        "textMuted": "#9E9990",
        "textOnAccent": "#FFFFFF"
      },
      "accents": {
        "primary": "#1C1B18",
        "hankoVermilion": "#C84B31",
        "indigo": "#3B6E8C",
        "kintsugiGold": "#C59B27",
        "jade": "#2E7D5B"
      },
      "borders": {
        "hairline": "rgba(28, 27, 24, 0.08)",
        "strong": "rgba(28, 27, 24, 0.15)",
        "specularInner": "rgba(255, 255, 255, 0.65)"
      },
      "shadows": {
        "cardResting": "0 2px 10px -2px rgba(28, 27, 24, 0.04), 0 1px 3px rgba(28, 27, 24, 0.03)",
        "cardHover": "0 14px 32px -6px rgba(28, 27, 24, 0.08), 0 4px 12px -2px rgba(28, 27, 24, 0.04)",
        "modal": "0 24px 60px -12px rgba(28, 27, 24, 0.18)"
      }
    },
    "dark_onyx": {
      "label": "Warm Onyx (Dark Mode)",
      "surfaces": {
        "bgCanvas": "#121211",
        "bgSurface": "#1C1B1A",
        "bgElevated": "#262523",
        "bgSubtle": "#2E2C29",
        "bgGlassHeader": "rgba(18, 18, 17, 0.78)",
        "bgGlassFloating": "rgba(28, 27, 26, 0.86)"
      },
      "typography": {
        "textPrimary": "#F3F1EC",
        "textSecondary": "#A39E93",
        "textMuted": "#6E6A61",
        "textOnAccent": "#121211"
      },
      "accents": {
        "primary": "#F3F1EC",
        "hankoVermilion": "#E05A40",
        "indigo": "#5B9BD5",
        "kintsugiGold": "#E3B94B",
        "jade": "#4E9F7A"
      },
      "borders": {
        "hairline": "rgba(243, 241, 236, 0.08)",
        "strong": "rgba(243, 241, 236, 0.16)",
        "specularInner": "rgba(255, 255, 255, 0.06)"
      },
      "shadows": {
        "cardResting": "0 4px 16px rgba(0, 0, 0, 0.25)",
        "cardHover": "0 18px 40px rgba(0, 0, 0, 0.45)",
        "modal": "0 28px 70px rgba(0, 0, 0, 0.65)"
      }
    }
  },
  "typography": {
    "fontFamilies": {
      "sans": "'Inter', -apple-system, BlinkMacSystemFont, 'SF Pro Display', sans-serif",
      "editorialSerif": "'Newsreader', 'Shippori Mincho', Georgia, serif",
      "mono": "'JetBrains Mono', 'SF Mono', monospace"
    },
    "scale": {
      "displayEditorial": {
        "fontFamily": "editorialSerif",
        "fontSize": "clamp(1.75rem, 2.8vw, 2.5rem)",
        "lineHeight": "1.15",
        "letterSpacing": "-0.02em",
        "fontWeight": "500"
      },
      "sectionHeader": {
        "fontFamily": "sans",
        "fontSize": "1.25rem",
        "lineHeight": "1.3",
        "letterSpacing": "-0.02em",
        "fontWeight": "600"
      },
      "cardTitle": {
        "fontFamily": "sans",
        "fontSize": "0.9375rem",
        "lineHeight": "1.35",
        "letterSpacing": "-0.01em",
        "fontWeight": "600"
      },
      "bodyDefault": {
        "fontFamily": "sans",
        "fontSize": "0.875rem",
        "lineHeight": "1.6",
        "fontWeight": "400"
      },
      "microMeta": {
        "fontFamily": "mono",
        "fontSize": "0.6875rem",
        "lineHeight": "1.2",
        "letterSpacing": "0.06em",
        "textTransform": "uppercase",
        "fontWeight": "500"
      }
    }
  },
  "layoutAndGeometry": {
    "container": {
      "maxWidth": "1340px",
      "paddingX": "clamp(1.25rem, 3vw, 2.5rem)"
    },
    "bentoGrid": {
      "columnsDesktop": 12,
      "gap": "1.25rem",
      "heroTileSpan": 7,
      "sideStackSpan": 5
    },
    "borderRadius": {
      "bentoTile": "24px",
      "bookCover": "14px",
      "innerCard": "18px",
      "pillButton": "9999px",
      "hankoStamp": "5px"
    },
    "glassmorphism": {
      "backdropFilter": "blur(20px) saturate(180%)",
      "border": "1px solid var(--border-hairline)"
    }
  },
  "subscriptionAvatarSystem": {
    "description": "Jewelry-grade, non-cyberpunk avatar rings in the header & profile widgets using conic-gradients and subtle inner spacing.",
    "geometry": {
      "avatarSizeHeader": "38px",
      "ringPadding": "2.5px",
      "innerGapBorder": "2px solid var(--bg-surface)",
      "kanjiCornerSeal": {
        "size": "15px",
        "borderRadius": "4px",
        "fontSize": "9px",
        "fontWeight": "700",
        "position": "bottom: -2px; right: -2px"
      }
    },
    "tiers": {
      "free": {
        "id": "free",
        "name": "Reader",
        "kanjiSeal": "読",
        "ringBackground": "linear-gradient(135deg, #D4D1C8 0%, #9E9990 50%, #D4D1C8 100%)",
        "animation": "none",
        "pillStyle": {
          "background": "rgba(158, 153, 144, 0.14)",
          "textColor": "var(--text-secondary)"
        },
        "sealStyle": {
          "background": "#7A766E",
          "textColor": "#FFFFFF"
        }
      },
      "ink": {
        "id": "ink",
        "name": "Ink+",
        "kanjiSeal": "蒼",
        "ringBackground": "conic-gradient(from 0deg, #2B5876, #4E4376, #67B26F, #3B82A6, #2B5876)",
        "animation": "spinRing 9s linear infinite",
        "outerHalo": "0 2px 8px rgba(59, 110, 140, 0.25)",
        "pillStyle": {
          "background": "rgba(59, 110, 140, 0.12)",
          "textColor": "#3B6E8C"
        },
        "sealStyle": {
          "background": "#3B6E8C",
          "textColor": "#FFFFFF"
        }
      },
      "kintsugi": {
        "id": "kintsugi",
        "name": "Kintsugi Gold",
        "kanjiSeal": "金",
        "ringBackground": "conic-gradient(from 0deg, #B38728, #FBF5B7, #D4AF37, #8A6421, #F7EF8A, #B38728)",
        "animation": "spinRing 6s linear infinite",
        "outerHalo": "0 2px 10px rgba(197, 155, 39, 0.30)",
        "usernameTreatment": "linear-gradient(90deg, #9A7418, #D4AF37, #9A7418)",
        "pillStyle": {
          "background": "rgba(197, 155, 39, 0.15)",
          "textColor": "#A67C00"
        },
        "sealStyle": {
          "background": "linear-gradient(135deg, #D4AF37, #9A7418)",
          "textColor": "#FFFFFF"
        }
      },
      "miyabi": {
        "id": "miyabi",
        "name": "Miyabi • 雅",
        "kanjiSeal": "雅",
        "ringBackground": "conic-gradient(from 0deg, #D98A9C, #E8C574, #76B09C, #8E9AAF, #D98A9C)",
        "animation": "spinRing 4.5s linear infinite",
        "outerHalo": "0 3px 14px rgba(217, 138, 156, 0.38)",
        "usernameTreatment": "linear-gradient(90deg, #C84B31, #B38728, #3B6E8C)",
        "pillStyle": {
          "background": "linear-gradient(135deg, rgba(217, 138, 156, 0.18), rgba(232, 197, 116, 0.2), rgba(118, 176, 156, 0.18))",
          "textColor": "#B84A5A"
        },
        "sealStyle": {
          "background": "linear-gradient(135deg, #C84B31, #B38728)",
          "textColor": "#FFFFFF"
        }
      }
    }
  },
  "componentPatterns": {
    "headerBar": {
      "height": "64px",
      "behavior": "Sticky top, translucent glass, 1px bottom hairline",
      "leftZone": "Hanko red square seal (読み) + bold wordmark YOMI + segmented pill navigation",
      "rightZone": "Command+K search pill + Washi/Onyx theme toggle + Interactive User Profile Pill with Avatar Ring"
    },
    "mangaCoverCard": {
      "aspectRatio": "2 / 3",
      "borderRadius": "14px",
      "spineCreaseOverlay": "linear-gradient(90deg, rgba(255,255,255,0.28) 0%, rgba(0,0,0,0.08) 3.5%, rgba(255,255,255,0.08) 5%, transparent 12%)",
      "hoverMotion": "transform: translateY(-4px) scale(1.01); transition: 350ms cubic-bezier(0.22, 1, 0.36, 1)",
      "badges": "Top-left frosted glass pill for format (Manga/Manhwa), Top-right Kintsugi badge for Early Access"
    },
    "hankoSealBadge": {
      "description": "Traditional Japanese vermilion stamp used sparingly as an editorial mark or logo emblem",
      "background": "#C84B31",
      "color": "#FFFFFF",
      "borderRadius": "5px",
      "fontFamily": "editorialSerif"
    }
  },
  "motionAndMicroInteractions": {
    "springEasing": "cubic-bezier(0.22, 1, 0.36, 1)",
    "smoothEasing": "cubic-bezier(0.16, 1, 0.3, 1)",
    "durations": {
      "fast": "180ms",
      "normal": "320ms",
      "slow": "500ms"
    }
  }
}
```

Generate [page name] strictly according to the YOMI design system, based on this JSON: retain the Washi/Onyx palette, the physical cover-binding effect (spineCreaseOverlay), the subscription-tier ring elements (subscriptionAvatarSystem), and the Bento modular grid.