package services

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
        "strings"

        "github.com/tidwall/gjson"
)

type HowLongToBeat struct {}

func (hltb *HowLongToBeat) Search(gameName string) HltbSearchResult {
    var searchResult HltbSearchResult
    query := getDefaultQuery()
    uri := "https://howlongtobeat.com/api/search"

    query.SearchTerms = strings.Split(gameName, " ")

    jsonBody, err := json.Marshal(query)
    if err != nil {
        //  TODO : Do something with this information
    }

    print(jsonBody)

    request, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonBody))

    request.Header = hltb.getDefaultHeaders()
    client := &http.Client {}
    response, err := client.Do(request)

    if err != nil {
        println("Error in request")
        //  TODO : Do something with this information
    }

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
        println("Error in request")
        println(err)
    }

    err = json.Unmarshal(body, &searchResult)

    if err != nil {
        println("Error unmarshaling json response")
        fmt.Printf("%s", err)
    }
    fmt.Printf("%v", searchResult)
    return searchResult
}

type GameDetails struct {
    GameId              int         `json:"game_id"`
    GameName            string      `json:"game_name"`
    GameImage           string      `json:"game_image"`
    GameType            string      `json:"game_type"`
    ProfileSummary      string      `json:"profile_summary"`
    ProfileGenre        string      `json:"profile_genre"`
    ProfilePlatform     string      `json:"profile_platform"`
    CompMain            int         `json:"comp_main"`
    CompPlus            int         `json:"comp_plus"`
    Comp100             int         `json:"comp_100"`
}

func (hltb *HowLongToBeat) extractGameDetails(gameResponse []byte) GameDetails {
    var gameDetails GameDetails

    gameString := gjson.Get(string(gameResponse), "pageProps.game.data.game.0").Raw

    print("\nResulting json:\n\n", gameString, "\n")

    println(gameString)

    json.Unmarshal([]byte(gameString), &gameDetails)

    fmt.Printf("%v", gameDetails)

    return gameDetails
}

func (hltb *HowLongToBeat) GetGame(gameId string) GameDetails {
    // var gameDetails
    uri := fmt.Sprintf("https://howlongtobeat.com/_next/data/kuLnFf9C8zfMBmeInzgAq/game/%s.json", gameId)

    request, err := http.NewRequest("GET", uri, nil)

    request.Header = hltb.getDefaultHeaders()
    client := &http.Client {}
    response, err := client.Do(request)

    if err != nil {
        println("Error in request")
    }

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
        println("Error reading body")
    }
    
    return hltb.extractGameDetails(body)
} 

func (hltb *HowLongToBeat) getDefaultHeaders() http.Header {
    return http.Header {
        "Content-Type": { "application/json" },
        "User-Agent": { "Chrome: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36" },
        "Origin": { "https://howlongtobeat.com" },
        "Referer": { "https://howlongtobeat.com/" },
        "Authority": { "howlongtobeat.com" },
    }
}

type HltbGameStats struct {
    GameId          int         `json:"game_id"`
    GameName        string      `json:"game_name"`
    GameImage       string      `json:"game_image"`
    CompMain        int         `json:"comp_main"`
    Comp100         int         `json:"comp_100"`
}

type HltbSearchResult struct {
    Count           int                 `json:"count"`
    Category        string              `json:"category"`
    PageCurrent     int                 `json:"pageCurrent"`
    PageTotal       int                 `json:"pageTotal"`
    PageSize        int                 `json:"pageSize"`
    Data            []HltbGameStats     `json:"data"`
}

type QueryUsers struct {
    SortCategory    string      `json:"sortCategory"`
}

type QueryLists struct {
    SortCategory    string      `json:"sortCategory"`
}

type QuerySearchOptionsGamesRangeTime struct {
    Min             int         `json:"min"`
    Max             int         `json:"max"`
}

type QuerySearchOptionsGamesGameplay struct {
    Perspective     string      `json:"perspective"`
    Flow            string      `json:"flow"`
    Genre           string      `json:"genre"`
}

type QuerySearchOptionsGames struct {
    UserId          int         `json:"userId"`
    Platform        string      `json:"platform"`
    SortCategory    string      `json:"sortCategory"`
    RangeCategory   string      `json:"rangeCategory"`
    RangeTime       QuerySearchOptionsGamesRangeTime `json:"rangeTime"`
    Gameplay        QuerySearchOptionsGamesGameplay  `json:"gameplay"`
    Modifier        string      `json:"modifier"`
}

type QuerySearchOptionsUsers struct {}

type QuerySearchOptions struct {
    Filter          string      `json:"filter"`
    Games           QuerySearchOptionsGames `json:"games"`
    Randomizer      int         `json:"randomizer"`
    Sort            int         `json:"sort"`
    Users           QuerySearchOptionsUsers `json:"users"`
}


type Query struct {
    SearchOptions       QuerySearchOptions      `json:"searchOptions"`
    Lists               QueryLists              `json:"lists"`
    SearchPage          int                     `json:"searchPage"`
    SearchTerms         []string                `json:"searchTerms"`
    SearchType          string                  `json:"searchType"`
    Size                int                     `json:"size"`
    Users               QueryUsers              `json:"users"`
}

func getDefaultQuery() Query {
    return Query {
        SearchOptions: QuerySearchOptions {
            Filter: "",
            Games: QuerySearchOptionsGames {
                UserId: 0,
                Platform: "",
                SortCategory: "name",
                RangeCategory: "main",
                RangeTime: QuerySearchOptionsGamesRangeTime {
                    Min: 0,
                    Max: 0,
                },
                Gameplay: QuerySearchOptionsGamesGameplay {
                    Perspective: "",
                    Flow: "",
                    Genre: "",
                },
                Modifier: "hide_dlc",
            },
            Randomizer: 0,
            Sort: 0,
            Users: QuerySearchOptionsUsers {},
        },
        Lists: QueryLists {
            SortCategory: "follows",
        },
        SearchPage: 1,
        SearchTerms: nil,
        SearchType: "games",
        Size: 20,
        Users: QueryUsers {
            SortCategory: "postcount",
        },
    }
}
