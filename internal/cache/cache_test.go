package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
			val: []byte(`{"count":1054,"next":"https://pokeapi.co/api/v2/location-area/?offset=20&limit=20","previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"},{"name":"eterna-city-area","url":"https://pokeapi.co/api/v2/location-area/2/"},{"name":"pastoria-city-area","url":"https://pokeapi.co/api/v2/location-area/3/"},{"name":"sunyshore-city-area","url":"https://pokeapi.co/api/v2/location-area/4/"},{"name":"sinnoh-pokemon-league-area","url":"https://pokeapi.co/api/v2/location-area/5/"},{"name":"oreburgh-mine-1f","url":"https://pokeapi.co/api/v2/location-area/6/"},{"name":"oreburgh-mine-b1f","url":"https://pokeapi.co/api/v2/location-area/7/"},{"name":"valley-windworks-area","url":"https://pokeapi.co/api/v2/location-area/8/"},{"name":"eterna-forest-area","url":"https://pokeapi.co/api/v2/location-area/9/"},{"name":"fuego-ironworks-area","url":"https://pokeapi.co/api/v2/location-area/10/"},{"name":"mt-coronet-1f-route-207","url":"https://pokeapi.co/api/v2/location-area/11/"},{"name":"mt-coronet-2f","url":"https://pokeapi.co/api/v2/location-area/12/"},{"name":"mt-coronet-3f","url":"https://pokeapi.co/api/v2/location-area/13/"},{"name":"mt-coronet-exterior-snowfall","url":"https://pokeapi.co/api/v2/location-area/14/"},{"name":"mt-coronet-exterior-blizzard","url":"https://pokeapi.co/api/v2/location-area/15/"},{"name":"mt-coronet-4f","url":"https://pokeapi.co/api/v2/location-area/16/"},{"name":"mt-coronet-4f-small-room","url":"https://pokeapi.co/api/v2/location-area/17/"},{"name":"mt-coronet-5f","url":"https://pokeapi.co/api/v2/location-area/18/"},{"name":"mt-coronet-6f","url":"https://pokeapi.co/api/v2/location-area/19/"},{"name":"mt-coronet-1f-from-exterior","url":"https://pokeapi.co/api/v2/location-area/20/"}]}`),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/?offset=160&limit=20",
			val: []byte(`{"count":1054,"next":"https://pokeapi.co/api/v2/location-area/?offset=180&limit=20","previous":"https://pokeapi.co/api/v2/location-area/?offset=140&limit=20","results":[{"name":"union-cave-b2f","url":"https://pokeapi.co/api/v2/location-area/200/"},{"name":"johto-route-33-area","url":"https://pokeapi.co/api/v2/location-area/201/"},{"name":"slowpoke-well-1f","url":"https://pokeapi.co/api/v2/location-area/202/"},{"name":"slowpoke-well-b1f","url":"https://pokeapi.co/api/v2/location-area/203/"},{"name":"ilex-forest-area","url":"https://pokeapi.co/api/v2/location-area/204/"},{"name":"johto-route-34-area","url":"https://pokeapi.co/api/v2/location-area/205/"},{"name":"johto-route-35-area","url":"https://pokeapi.co/api/v2/location-area/206/"},{"name":"national-park-area","url":"https://pokeapi.co/api/v2/location-area/207/"},{"name":"unknown-all-bugs-area","url":"https://pokeapi.co/api/v2/location-area/208/"},{"name":"johto-route-36-area","url":"https://pokeapi.co/api/v2/location-area/209/"},{"name":"johto-route-37-area","url":"https://pokeapi.co/api/v2/location-area/210/"},{"name":"ecruteak-city-area","url":"https://pokeapi.co/api/v2/location-area/211/"},{"name":"burned-tower-1f","url":"https://pokeapi.co/api/v2/location-area/212/"},{"name":"burned-tower-b1f","url":"https://pokeapi.co/api/v2/location-area/213/"},{"name":"bell-tower-2f","url":"https://pokeapi.co/api/v2/location-area/214/"},{"name":"bell-tower-3f","url":"https://pokeapi.co/api/v2/location-area/215/"},{"name":"bell-tower-4f","url":"https://pokeapi.co/api/v2/location-area/216/"},{"name":"bell-tower-5f","url":"https://pokeapi.co/api/v2/location-area/217/"},{"name":"bell-tower-6f","url":"https://pokeapi.co/api/v2/location-area/218/"},{"name":"bell-tower-7f","url":"https://pokeapi.co/api/v2/location-area/219/"}]}`),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
