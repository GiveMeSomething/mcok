package config

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type Preferences struct {
	Foods    []string
	Movies   []string
	Magazine []string
}

var CountryConfig = map[Country]Preferences{
	Vietnam: {
		Foods: []string{
			"Phở gà (Chicken Pho)",
			"Cơm chiên trứng (Egg Fried Rice)",
			"Bánh cuốn (Steamed Rice Rolls)",
			"Cháo gà (Chicken Congee)",
			"Gỏi cuốn (Fresh Spring Rolls)",
			"Bún chả",
			"Phở bò",
			"Cá kho tộ",
			"Gỏi cuốn",
			"Lẩu",
			"Cháo gà (Chicken Congee)",
			"Phở gà (Chicken Pho)",
			"Canh rau (Vegetable Soup)",
			"Cá hấp (Steamed Fish)",
			"Bún thang",
		},
		Movies: []string{
			"Doraemon: Nobita's Dinosaur",
			"Trạng Tí phiêu lưu ký (Tí the Adventurer)",
			"Cô tiên xanh (The Blue Fairy)",

			"Mắt biếc (Dreamy Eyes)",
			"Bố già (Dad, I'm Sorry)",
			"Hai Phượng (Furie)",

			"Chung một dòng sông (Same River)",
			"Em bé Hà Nội (Little Girl of Hanoi)",
			"Cánh đồng hoang (The Abandoned Field)",
		},
		Magazine: []string{
			"Nhi Đồng",
			"Khăn Quàng Đỏ",

			"Tuổi Trẻ Cuối Tuần",
			"Thế Giới Tiếp Thị",

			"Người Cao Tuổi",
			"Sức Khỏe Người Cao Tuổi",
		},
	},
	India: {
		Foods: []string{
			"Aloo Paratha (Potato Stuffed Flatbread)",
			"Idli with Sambar (Steamed Rice Cakes with Lentil Stew)",
			"Vegetable Pulao (Vegetable Rice)",
			"Paneer Tikka (Indian Cottage Cheese Skewers)",
			"Sweet Semolina Upma (Sweet Semolina Porridge)",
			"Butter Chicken",
			"Biryani",
			"Dal Makhani",
			"Rogan Josh",
			"Chole Bhature",
			"Khichdi (Rice and Lentil Porridge)",
			"Dalia (Broken Wheat Porridge)",
			"Vegetable Stew",
			"Moong Dal Cheela (Lentil Pancakes)",
			"Ragi Roti (Finger Millet Flatbread)",
		},
		Movies: []string{
			"Chhota Bheem and the Throne of Bali",
			"Koi... Mil Gaya",
			"My Friend Ganesha",

			"3 Idiots",
			"Gully Boy",
			"Andhadhun",

			"Anand",
			"Baghban",
			"Piku",
		},
		Magazine: []string{
			"Champak",
			"Chandamama",

			"India Today",
			"Outlook India",

			"Harmony - Celebrate Age",
			"Life Positive",
		},
	},
	USA: {
		Foods: []string{
			"Chicken Noodle Soup",
			"Mashed Potatoes and Gravy",
			"Baked Salmon",
			"Oatmeal",
			"Apple Sauce",
			"Macaroni and Cheese",
			"Chicken Nuggets",
			"Pizza",
			"Spaghetti and Meatballs",
			"Pancakes",
			"Steak",
			"Barbecue Ribs",
			"Seafood Gumbo",
			"Chicken Pot Pie",
			"New York-Style Pizza",
		},
		Movies: []string{
			"Toy Story",
			"The Lion King",
			"Frozen",

			"The Godfather",
			"Pulp Fiction",
			"The Shawshank Redemption",

			"Casablanca",
			"Singin' in the Rain",
			"On Golden Pond",
		},
		Magazine: []string{
			"National Geographic Kids",
			"Highlights",

			"Time",
			"The New Yorker",

			"AARP The Magazine",
			"Reader's Digest",
		},
	},
}

var AgeRangeConfig = map[AgeRange]Preferences{
	Children: {
		Foods: []string{
			"Aloo Paratha (Potato Stuffed Flatbread)",
			"Idli with Sambar (Steamed Rice Cakes with Lentil Stew)",
			"Vegetable Pulao (Vegetable Rice)",
			"Paneer Tikka (Indian Cottage Cheese Skewers)",
			"Sweet Semolina Upma (Sweet Semolina Porridge)",

			"Phở gà (Chicken Pho)",
			"Cơm chiên trứng (Egg Fried Rice)",
			"Bánh cuốn (Steamed Rice Rolls)",
			"Cháo gà (Chicken Congee)",
			"Gỏi cuốn (Fresh Spring Rolls)",

			"Macaroni and Cheese",
			"Chicken Nuggets",
			"Pizza",
			"Spaghetti and Meatballs",
			"Pancakes",
		},
		Movies: []string{
			"Doraemon: Nobita's Dinosaur",
			"Trạng Tí phiêu lưu ký (Tí the Adventurer)",
			"Cô tiên xanh (The Blue Fairy)",

			"Chhota Bheem and the Throne of Bali",
			"Koi... Mil Gaya",
			"My Friend Ganesha",

			"Toy Story",
			"The Lion King",
			"Frozen",
		},
		Magazine: []string{
			"Champak",
			"Chandamama",

			"Nhi Đồng",
			"Khăn Quàng Đỏ",

			"National Geographic Kids",
			"Highlights",
		},
	},
	Teenager: {
		Foods: []string{
			"Butter Chicken",
			"Biryani",
			"Dal Makhani",
			"Rogan Josh",
			"Chole Bhature",

			"Bún chả",
			"Phở bò",
			"Cá kho tộ",
			"Gỏi cuốn",
			"Lẩu",

			"Steak",
			"Barbecue Ribs",
			"Seafood Gumbo",
			"Chicken Pot Pie",
			"New York-Style Pizza",
		},
		Movies: []string{
			"Mắt biếc (Dreamy Eyes)",
			"Bố già (Dad, I'm Sorry)",
			"Hai Phượng (Furie)",

			"3 Idiots",
			"Gully Boy",
			"Andhadhun",

			"The Godfather",
			"Pulp Fiction",
			"The Shawshank Redemption",
		},
		Magazine: []string{
			"India Today",
			"Outlook India",

			"India Today",
			"Outlook India",

			"Tuổi Trẻ Cuối Tuần",
			"Thế Giới Tiếp Thị",
		},
	},
	Adult: {
		Foods: []string{
			"Khichdi (Rice and Lentil Porridge)",
			"Dalia (Broken Wheat Porridge)",
			"Vegetable Stew",
			"Moong Dal Cheela (Lentil Pancakes)",
			"Ragi Roti (Finger Millet Flatbread)",

			"Cháo gà (Chicken Congee)",
			"Phở gà (Chicken Pho)",
			"Canh rau (Vegetable Soup)",
			"Cá hấp (Steamed Fish)",
			"Bún thang",

			"Chicken Noodle Soup",
			"Mashed Potatoes and Gravy",
			"Baked Salmon",
			"Oatmeal",
			"Apple Sauce",
		},
		Movies: []string{
			"Casablanca",
			"Singin' in the Rain",
			"On Golden Pond",

			"Anand",
			"Baghban",
			"Piku",

			"Chung một dòng sông (Same River)",
			"Em bé Hà Nội (Little Girl of Hanoi)",
			"Cánh đồng hoang (The Abandoned Field)",
		},
		Magazine: []string{

			"Harmony - Celebrate Age",
			"Life Positive",

			"Người Cao Tuổi",
			"Sức Khỏe Người Cao Tuổi",

			"AARP The Magazine",
			"Reader's Digest",
		},
	},
}

// Read config from files then put them into cache

func GenerateCacheKey(pref UserPref, ageRange AgeRange, country Country) string {
	return fmt.Sprintf("%s:%s:%s", pref, ageRange, country)
}

type UserPref string

const (
	UPrefMovie    UserPref = "movie"
	UPrefMagazine UserPref = "magazine"
	UPrefFood     UserPref = "food"
)

var ConfigCache = make(map[string][]string, 0)

func loadConfigFromFile(category UserPref, filePath string) (map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// ageRange_country_value
		parts := strings.Split(line, "_")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		age := AgeRange(parts[0])
		if !age.IsValid() {
			return nil, fmt.Errorf("invalid age: %s", age)
		}

		country := Country(parts[1])
		if !age.IsValid() {
			return nil, fmt.Errorf("invalid country: %s", country)
		}

		cacheKey := GenerateCacheKey(category, age, country)
		if _, ok := result[cacheKey]; !ok {
			result[cacheKey] = make([]string, 0)
		}

		result[cacheKey] = append(result[cacheKey], parts[2])
	}

	return result, nil
}

const FOOD_CONFIG_PATH = "/config/food_config.txt"
const MOVIE_CONFIG_PATH = "/config/movies_config.txt"
const MAGAZINE_CONFIG_PATH = "/config/magazine_config.txt"

func LoadConfig() error {
	fmt.Println("Loading configuration from file...")
	start := time.Now()

	lock := sync.Mutex{}

	errgroup.WithContext(context.Background())
	group, _ := errgroup.WithContext(context.Background())

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	group.Go(func() error {
		result, err := loadConfigFromFile(UPrefFood, path.Join(cwd, FOOD_CONFIG_PATH))
		if err != nil {
			return err
		}

		// Accquire lock and write to cache
		lock.Lock()
		defer lock.Unlock()

		for cacheKey, value := range result {
			ConfigCache[cacheKey] = value
		}

		return nil
	})

	group.Go(func() error {
		result, err := loadConfigFromFile(UPrefMovie, path.Join(cwd, MOVIE_CONFIG_PATH))
		if err != nil {
			return err
		}

		// Accquire lock and write to cache
		lock.Lock()
		defer lock.Unlock()

		for cacheKey, value := range result {
			ConfigCache[cacheKey] = value
		}

		return nil
	})

	group.Go(func() error {
		result, err := loadConfigFromFile(UPrefMagazine, path.Join(cwd, MAGAZINE_CONFIG_PATH))
		if err != nil {
			return err
		}

		// Accquire lock and write to cache
		lock.Lock()
		defer lock.Unlock()

		for cacheKey, value := range result {
			ConfigCache[cacheKey] = value
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		return err
	}

	fmt.Printf("Config loaded: %s\n", time.Since(start))

	return nil
}
