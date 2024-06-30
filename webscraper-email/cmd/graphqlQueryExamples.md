# Multiple inputs

query {
  media111: Media(id: 166240, type: ANIME) {
    ...MediaFields
  }
  media222: Media(id: 152681, type: ANIME) {
    ...MediaFields
  }
}

fragment MediaFields on Media {
  id
  title {
    romaji
    english
  }
  nextAiringEpisode {
    airingAt
    episode
    timeUntilAiring
  }
}

# Search by english title

query {
  Media(search: "Demon Slayer: Kimetsu no Yaiba Hashira Training Arc", type: ANIME) {
    id
    title {
      romaji
      english
      native
    }
  }
}